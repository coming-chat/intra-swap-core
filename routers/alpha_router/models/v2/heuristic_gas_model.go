package v2

import (
	"errors"
	"fmt"
	"github.com/coming-chat/intra-swap-core/base_constant"
	"github.com/coming-chat/intra-swap-core/base_entities"
	v2 "github.com/coming-chat/intra-swap-core/providers/v2"
	"github.com/coming-chat/intra-swap-core/routers/alpha_router/models"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"math/big"
	"sort"
	"strconv"
)

var (
	// Constant cost for doing any swap regardless of pools.
	baseSwapCost = big.NewInt(115000)
	// Constant per extra hop in the route.
	costPerExtraHop = big.NewInt(20000)
)

type HeuristicGasModelFactory struct {
}

func (h *HeuristicGasModelFactory) BuildGasModel(
	chainId base_entities.ChainId,
	gasPriceWei *big.Int,
	poolProvider v2.PoolProvider,
	token *entities.Token,
) (gasModel models.GasModel[models.V2RouteWithValidQuote], err error) {
	wrappedNativeCurrency, ok := base_entities.WRAPPED_NATIVE_CURRENCY[chainId]
	if !ok {
		return gasModel, fmt.Errorf("not found wrapped native currency token by chain_id[%d]", chainId)
	}
	if token.Equal(wrappedNativeCurrency) {
		usdPool, err := h.getHighestLiquidityUSDPool(
			chainId,
			poolProvider,
		)
		if err != nil {
			return gasModel, err
		}

		gasModel.EstimateGasCostFn = func(routeWithValidQuote models.V2RouteWithValidQuote) (
			*big.Int,
			*entities.CurrencyAmount,
			*entities.CurrencyAmount,
			error,
		) {
			gasCostInEth, gasUse := h.estimateGas(
				routeWithValidQuote,
				gasPriceWei,
				chainId,
			)

			var ethTokenPrice *entities.Price
			if usdPool.Token0().Address == base_entities.WRAPPED_NATIVE_CURRENCY[chainId].Address {
				ethTokenPrice = usdPool.Token0Price()
			} else {
				ethTokenPrice = usdPool.Token1Price()
			}

			gasCostInTermsOfUSD, err := ethTokenPrice.Quote(gasCostInEth)
			if err != nil {
				return nil, nil, nil, err
			}

			return gasUse, gasCostInEth, gasCostInTermsOfUSD, nil
		}
		return gasModel, err
	}

	usdPool, err := h.getHighestLiquidityUSDPool(
		chainId,
		poolProvider,
	)
	if err != nil {
		return gasModel, err
	}

	// If the quote token is not WETH, we convert the gas cost to be in terms of the quote token.
	// We do this by getting the highest liquidity <token>/ETH pool.
	ethPool, err := h.getEthPool(
		chainId,
		token,
		poolProvider,
	)
	if err != nil {
		return gasModel, err
	}

	gasModel.EstimateGasCostFn = func(routeWithValidQuote models.V2RouteWithValidQuote) (
		gasEstimate *big.Int,
		gasCostInToken,
		gasCostInUSD *entities.CurrencyAmount,
		err error,
	) {
		var usdToken *entities.Token
		if usdPool.Token0().Address == base_entities.WRAPPED_NATIVE_CURRENCY[chainId].Address {
			usdToken = usdPool.Token1()
		} else {
			usdToken = usdPool.Token0()
		}

		gasCostInEth, gasUse := h.estimateGas(
			routeWithValidQuote,
			gasPriceWei,
			chainId,
		)

		if ethPool == nil {
			return gasUse, entities.FromRawAmount(token, big.NewInt(0)), entities.FromRawAmount(usdToken, big.NewInt(0)), nil
		}

		var ethTokenPrice *entities.Price
		if ethPool.Token0().Address == base_entities.WRAPPED_NATIVE_CURRENCY[chainId].Address {
			ethTokenPrice = ethPool.Token0Price()
		} else {
			ethTokenPrice = ethPool.Token1Price()
		}
		gasCostInTermsOfQuoteToken, err := ethTokenPrice.Quote(gasCostInEth)
		if err != nil {
			return nil, nil, nil, err
		}

		var ethTokenPriceUSDPool *entities.Price
		if usdPool.Token0().Address == base_entities.WRAPPED_NATIVE_CURRENCY[chainId].Address {
			ethTokenPriceUSDPool = usdPool.Token0Price()
		} else {
			ethTokenPriceUSDPool = usdPool.Token1Price()
		}

		gasCostInTermsOfUSD, err := ethTokenPriceUSDPool.Quote(gasCostInEth)
		if err != nil {
			return nil, nil, nil, err
		}
		return gasUse, gasCostInTermsOfQuoteToken, gasCostInTermsOfUSD, nil
	}
	return
}

func (h *HeuristicGasModelFactory) estimateGas(
	routeWithValidQuote models.V2RouteWithValidQuote,
	gasPriceWei *big.Int,
	chainId base_entities.ChainId,
) (*entities.CurrencyAmount, *big.Int) {
	gasUse := new(big.Int).Add(
		baseSwapCost,
		new(big.Int).Mul(
			costPerExtraHop,
			big.NewInt(int64(len(routeWithValidQuote.Route.Pools)-1)),
		),
	)
	totalGasCostWei := new(big.Int).Mul(gasPriceWei, gasUse)
	gasCostInEth := entities.FromRawAmount(
		base_entities.WRAPPED_NATIVE_CURRENCY[chainId],
		totalGasCostWei,
	)
	return gasCostInEth, gasUse
}

func (h *HeuristicGasModelFactory) getHighestLiquidityUSDPool(
	chainId base_entities.ChainId,
	poolProvider v2.PoolProvider,
) (*base_entities.V2Pool, error) {
	usdTokens, ok := models.UsdGasTokensByChain[chainId]

	if !ok {
		return nil, fmt.Errorf("could not find a USD token for computing gas costs on chain_id[%d]", chainId)
	}

	var usdPools []v2.TokenPairs
	for _, token := range usdTokens {
		usdPools = append(usdPools, v2.TokenPairs{
			Token0:      token,
			Token1:      base_entities.WRAPPED_NATIVE_CURRENCY[chainId], //Already checked in the above func
			PairAddress: "0xef24722d5dae32dc155d961561cffbc5f347eee7",
		})
	}
	poolAccessor, err := poolProvider.GetPools(usdPools, nil)
	if err != nil {
		return nil, err
	}
	poolsRaw := poolAccessor.GetAllPools()
	var pools []*base_entities.V2Pool
	for _, poolRaw := range poolsRaw {
		if poolRaw.Reserve0().GreaterThan(base_constant.ZeroFraction) && poolRaw.Reserve1().GreaterThan(base_constant.ZeroFraction) {
			pools = append(pools, poolRaw)
		}
	}

	if len(pools) == 0 {
		return nil, errors.New("could not find a USD/WETH pool for computing gas costs")
	}

	sort.Slice(pools, func(i, j int) bool {
		var numi, numj float64
		if pools[i].Token0().Equal(base_entities.WRAPPED_NATIVE_CURRENCY[chainId]) {
			numi, _ = strconv.ParseFloat(pools[i].Reserve0().ToSignificant(2), 64)
		} else {
			numi, _ = strconv.ParseFloat(pools[i].Reserve1().ToSignificant(2), 64)
		}
		if pools[j].Token0().Equal(base_entities.WRAPPED_NATIVE_CURRENCY[chainId]) {
			numj, _ = strconv.ParseFloat(pools[j].Reserve0().ToSignificant(2), 64)
		} else {
			numj, _ = strconv.ParseFloat(pools[j].Reserve1().ToSignificant(2), 64)
		}
		return numi < numj
	})

	return pools[0], nil
}

func (h *HeuristicGasModelFactory) getEthPool(
	chainId base_entities.ChainId,
	token *entities.Token,
	poolProvider v2.PoolProvider,
) (*base_entities.V2Pool, error) {
	weth := base_entities.WRAPPED_NATIVE_CURRENCY[chainId]
	poolAccessor, err := poolProvider.GetPools([]v2.TokenPairs{
		{
			Token0: weth,
			Token1: token,
		},
	}, nil)
	if err != nil {
		return nil, err
	}
	pool, err := poolAccessor.GetPool(weth, token)
	if err != nil {
		return nil, err
	}

	if pool.Reserve0().EqualTo(base_constant.ZeroFraction) || pool.Reserve1().EqualTo(base_constant.ZeroFraction) {
		return nil, fmt.Errorf("could not find a valid WETH pool with %s for computing gas costs", token.Symbol())
	}

	return pool, nil
}
