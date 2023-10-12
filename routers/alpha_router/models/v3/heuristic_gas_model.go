package v3

import (
	"encoding/hex"
	"fmt"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/providers"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	v3 "github.com/coming-chat/intra-swap-core/providers/v3"
	"github.com/coming-chat/intra-swap-core/routers/alpha_router/config"
	"github.com/coming-chat/intra-swap-core/routers/alpha_router/models"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"math/big"
	"regexp"
	"sort"
	"strconv"
)

// Cost for crossing an uninitialized tick.
var costPerUninitTick = big.NewInt(0)

func NewHeuristicGasModelFactory(
	baseProvider rpc.Provider,
	multiCallCore rpc.MultiCallProviderCore,
	gasPriceAddress string,
	poolIndexer v3.IndexerProvider,
	tokenProvider providers.TokenProvider,
	wNProvider providers.WrappedNativeCurrencyProvider,
	usdGasTokens providers.UsdGasTokensProvider,
) *HeuristicGasModelFactory {
	return &HeuristicGasModelFactory{
		provider:          baseProvider,
		multiCallProvider: multiCallCore,
		gasPriceAddress:   gasPriceAddress,
		wNProvider:        wNProvider,
		poolIndexer:       poolIndexer,
		tokenProvider:     tokenProvider,
		usdGasTokens:      usdGasTokens,
	}
}

type HeuristicGasModelFactory struct {
	provider          rpc.Provider
	multiCallProvider rpc.MultiCallProviderCore
	gasPriceAddress   string
	wNProvider        providers.WrappedNativeCurrencyProvider
	usdGasTokens      providers.UsdGasTokensProvider
	poolIndexer       v3.IndexerProvider
	tokenProvider     providers.TokenProvider
}

func (h *HeuristicGasModelFactory) BuildGasModel(
	chainId base_entities.ChainId,
	gasPriceWei *big.Int,
	poolProvider v3.PoolProvider,
	token *entities.Token,
) (gasModel models.GasModel[models.V3RouteWithValidQuote], err error) {
	//usdPool, err := h.getHighestLiquidityUSDPool(
	//	chainId,
	//	poolProvider,
	//)
	//if err != nil {
	//	return gasModel, err
	//}
	//
	//nativeCurrency, err := h.wNProvider.GetTokenByChain(chainId)
	//if err != nil {
	//	return gasModel, err
	//}

	gasModel.CalculateL1GasFeesFn = func(routes []models.V3RouteWithValidQuote) (*models.L1ToL2GasCosts, error) {
		//swapOptions := config.SwapOptions{
		//	Recipient:         common.HexToAddress("0x0000000000000000000000000000000000000001"),
		//	Deadline:          big.NewInt(100),
		//	SlippageTolerance: entities.NewPercent(big.NewInt(5), big.NewInt(10_000)),
		//}
		//var l1Used, l1FeeInWei *big.Int
		//switch chainId {
		//case base_constant.OPTIMISM, base_constant.OPTIMISTIC_KOVAN:
		//	gasData, err := v3.GetOptimismL2GasData(h.multiCallProvider, h.gasPriceAddress)
		//	if err != nil {
		//		return nil, err
		//	}
		//	l1Used, l1FeeInWei, err = h.calculateOptimismToL1SecurityFee(routes, swapOptions, gasData)
		//case base_constant.ARBITRUM_ONE, base_constant.ARBITRUM_RINKEBY:
		//	gasData, err := v3.GetArbitrumL2GasData(h.multiCallProvider, h.gasPriceAddress)
		//	if err != nil {
		//		return nil, err
		//	}
		//	l1Used, l1FeeInWei, err = h.calculateArbitrumToL1SecurityFee(routes, swapOptions, gasData)
		//default:
		//	l1Used, l1FeeInWei = big.NewInt(0), big.NewInt(0)
		//}
		//if err != nil {
		//	return nil, err
		//}
		//// wrap fee to native currency
		//costNativeCurrency := entities.FromRawAmount(nativeCurrency, l1FeeInWei)
		//
		//// convert fee into usd
		//var nativeTokenPrice *entities.Price
		//if usdPool.Token0().Address == nativeCurrency.Address {
		//	nativeTokenPrice = usdPool.Token0Price()
		//} else {
		//	nativeTokenPrice = usdPool.Token1Price()
		//}
		//
		//gasCostL1USD, err := nativeTokenPrice.Quote(costNativeCurrency)
		//if err != nil {
		//	return nil, err
		//}
		//gasCostL1QuoteToken := costNativeCurrency
		//if !token.Equal(nativeCurrency) {
		//	nativePool, err := h.getHighestLiquidityNativePool(chainId, token, poolProvider)
		//	if err != nil {
		//		gasCostL1QuoteToken = entities.FromRawAmount(token, big.NewInt(0))
		//	}
		//	var nativeTokenPrice *entities.Price
		//	if nativePool.Token0().Address == nativeCurrency.Address {
		//		nativeTokenPrice = nativePool.Token0Price()
		//	} else {
		//		nativeTokenPrice = nativePool.Token1Price()
		//	}
		//	gasCostL1QuoteToken, err = nativeTokenPrice.Quote(costNativeCurrency)
		//	if err != nil {
		//		return nil, err
		//	}
		//}
		//return &models.L1ToL2GasCosts{
		//	GasUsedL1:           l1Used,
		//	GasCostL1USD:        gasCostL1USD,
		//	GasCostL1QuoteToken: gasCostL1QuoteToken,
		//}, nil
		return nil, nil
	}

	//if token.Equal(nativeCurrency) {
	//	gasModel.EstimateGasCostFn = func(routeWithValidQuote models.V3RouteWithValidQuote) (
	//		gasEstimate *big.Int,
	//		gasCostInToken, gasCostInUSD *entities.CurrencyAmount,
	//		err error,
	//	) {
	//		totalGasCostNativeCurrency, _, baseGasUse, err := h.estimateGas(routeWithValidQuote, gasPriceWei, chainId)
	//		if err != nil {
	//			return nil, nil, nil, err
	//		}
	//		var nativeTokenPrice *entities.Price
	//		if usdPool.Token0().Address == nativeCurrency.Address {
	//			nativeTokenPrice = usdPool.Token0Price()
	//		} else {
	//			nativeTokenPrice = usdPool.Token1Price()
	//		}
	//		gasCostInTermsOfUSD, err := nativeTokenPrice.Quote(totalGasCostNativeCurrency)
	//		if err != nil {
	//			return nil, nil, nil, err
	//		}
	//		return baseGasUse, totalGasCostNativeCurrency, gasCostInTermsOfUSD, nil
	//	}
	//	return gasModel, nil
	//}
	// If the quote token is not in the native currency, we convert the gas cost to be in terms of the quote token.
	// We do this by getting the highest liquidity <quoteToken>/<nativeCurrency> pool. eg. <quoteToken>/ETH pool.

	//var usdToken *entities.Token
	//if usdPool.Token0().Address == nativeCurrency.Address {
	//	usdToken = usdPool.Token1()
	//} else {
	//	usdToken = usdPool.Token0()
	//}
	//
	//nativePool, _ := h.getHighestLiquidityNativePool(chainId, token, poolProvider)

	gasModel.EstimateGasCostFn = func(routeWithValidQuote models.V3RouteWithValidQuote) (gasEstimate *big.Int, gasCostInToken, gasCostInUSD *entities.CurrencyAmount, err error) {
		//totalGasCostNativeCurrency, _, baseGasUse, err := h.estimateGas(routeWithValidQuote, gasPriceWei, chainId)
		//if err != nil {
		//	return nil, nil, nil, err
		//}
		//if nativePool == nil {
		//	return baseGasUse, entities.FromRawAmount(token, big.NewInt(0)), entities.FromRawAmount(usdToken, big.NewInt(0)), nil
		//}
		//
		//var nativeTokenPrice *entities.Price
		//if nativePool.Token0().Address == nativeCurrency.Address {
		//	nativeTokenPrice = nativePool.Token0Price()
		//} else {
		//	nativeTokenPrice = nativePool.Token1Price()
		//}
		//gasCostInTermsOfQuoteToken, err := nativeTokenPrice.Quote(totalGasCostNativeCurrency)
		//if err != nil {
		//	return nil, nil, nil, err
		//}
		//
		//var nativeTokenPriceUSDPool *entities.Price
		//if usdPool.Token0().Address == nativeCurrency.Address {
		//	nativeTokenPriceUSDPool = usdPool.Token0Price()
		//} else {
		//	nativeTokenPriceUSDPool = usdPool.Token1Price()
		//}
		//gasCostInTermsOfUSD, err := nativeTokenPriceUSDPool.Quote(totalGasCostNativeCurrency)
		//if err != nil {
		//	return nil, nil, nil, err
		//}
		//return baseGasUse, gasCostInTermsOfQuoteToken, gasCostInTermsOfUSD, nil
		return nil, nil, nil, nil
	}
	return
}

func (h *HeuristicGasModelFactory) getHighestLiquidityUSDPool(
	chainId base_entities.ChainId,
	poolProvider v3.PoolProvider,
) (*base_entities.V3Pool, error) {
	usdTokens, err := h.usdGasTokens.GetTokensByChain(chainId)
	if err != nil {
		return nil, err
	}
	wrappedCurrency, err := h.wNProvider.GetTokenByChain(chainId)
	if err != nil {
		return nil, err
	}

	indexPools, err := h.poolIndexer.GetPools(chainId, usdTokens[0], wrappedCurrency, nil)
	if err != nil {
		return nil, err
	}
	var (
		tokenAddresses    []string
		tokenAddressesMap = make(map[string]struct{})
	)
	for _, pool := range indexPools {
		if _, has := tokenAddressesMap[pool.Token0.Id]; !has {
			tokenAddressesMap[pool.Token0.Id] = struct{}{}
			tokenAddresses = append(tokenAddresses, pool.Token0.Id)
		}
		if _, has := tokenAddressesMap[pool.Token1.Id]; !has {
			tokenAddressesMap[pool.Token1.Id] = struct{}{}
			tokenAddresses = append(tokenAddresses, pool.Token1.Id)
		}
	}

	var tokenPairs []v3.TokenPairs
	tokenAccessor, err := h.tokenProvider.GetTokens(chainId, tokenAddresses, nil)
	if err != nil {
		return nil, err
	}
	for _, pool := range indexPools {
		token0 := tokenAccessor.GetTokenByAddress(pool.Token0.Id)
		token1 := tokenAccessor.GetTokenByAddress(pool.Token1.Id)
		if token0 == nil || token1 == nil {
			continue
		}
		tokenPairs = append(tokenPairs, v3.TokenPairs{
			Token0:         token0,
			Token1:         token1,
			FeeAmount:      pool.FeeTier,
			PairAddress:    pool.Id,
			RouterAddress:  pool.RouterAddress,
			QuoteAddress:   pool.QuoteAddress,
			FactoryAddress: pool.FactoryAddress,
		})
	}
	poolAccessor, err := poolProvider.GetPools(tokenPairs, nil)
	if err != nil {
		return nil, err
	}
	pools := poolAccessor.GetAllPools()

	//var tokenPairs []v3.TokenPairs
	//feeAmounts := []constants.FeeAmount{
	//	constants.FeeLowest,
	//	constants.FeeLow,
	//	constants.FeeMedium,
	//	constants.FeeHigh,
	//}
	//for _, f := range feeAmounts {
	//	for _, u := range usdTokens {
	//		tokenPairs = append(
	//			tokenPairs,
	//			v3.TokenPairs{
	//				Token0:    wrappedCurrency,
	//				Token1:    u,
	//				FeeAmount: f,
	//			},
	//		)
	//	}
	//}
	//poolAccessor, err := poolProvider.GetPools(tokenPairs, nil)
	//if err != nil {
	//	return nil, err
	//}
	//var pools []*base_entities.V3Pool
	//for _, f := range feeAmounts {
	//	for _, u := range usdTokens {
	//		pool, err := poolAccessor.GetPool(wrappedCurrency, u, f)
	//		if err != nil {
	//			return nil, err
	//		}
	//		pools = append(pools, pool)
	//	}
	//}
	if len(pools) == 0 {
		return nil, fmt.Errorf("can not found USD/%s pool for computing gas costs", wrappedCurrency.Symbol())
	}
	sort.Slice(pools, func(i, j int) bool {
		return pools[i].Liquidity.Cmp(pools[j].Liquidity) < 0
	})
	return pools[0], nil
}

func (h *HeuristicGasModelFactory) calculateOptimismToL1SecurityFee(
	routes []models.V3RouteWithValidQuote,
	swapConfig config.SwapOptions,
	gasData v3.OptimismGasData,
) (*big.Int, *big.Int, error) {
	//route := routes[0]
	//var inputToken, outputToken entities.Currency
	//if route.TradeType == entities.ExactInput {
	//	inputToken, outputToken = route.Amount.Currency, route.Quote.Currency
	//} else {
	//	inputToken, outputToken = route.Quote.Currency, route.Amount.Currency
	//}

	// build trade for swap calldata
	//trade := util.BuildTrade(inputToken, outputToken, route.TradeType, routes)
	//data, err := util.BuildSwapMethodParameters(trade, swapConfig)
	//if err != nil {
	//	return nil, nil, err
	//}
	// data.Calldata
	l1GasUsed := h.getL2ToL1GasUsed("0x"+hex.EncodeToString([]byte{}), gasData.Overhead)
	// l1BaseFee is L1 Gas Price on etherscan
	l1Fee := new(big.Int).Mul(l1GasUsed, gasData.L1BaseFee)
	unscaled := new(big.Int).Mul(l1Fee, gasData.Scalar)
	// scaled = unscaled / (10 ** decimals)
	scaledConversion := new(big.Int).Exp(big.NewInt(10), gasData.Decimals, nil)
	scaled := new(big.Int).Div(unscaled, scaledConversion)
	return l1GasUsed, scaled, nil
}

func (h *HeuristicGasModelFactory) calculateArbitrumToL1SecurityFee(
	routes []models.V3RouteWithValidQuote,
	swapConfig config.SwapOptions,
	gasData v3.ArbitrumGasData,
) (*big.Int, *big.Int, error) {
	//var (
	//	route                   = routes[0]
	//	inputToken, outputToken entities.Currency
	//)
	//if route.TradeType == entities.ExactInput {
	//	inputToken = route.Amount.Currency
	//	outputToken = route.Quote.Currency
	//} else {
	//	inputToken = route.Quote.Currency
	//	outputToken = route.Quote.Currency
	//}
	// build trade for swap calldata
	//trade := util.BuildTrade(inputToken, outputToken, route.TradeType, routes)
	//data, err := util.BuildSwapMethodParameters(trade, swapConfig)
	//if err != nil {
	//	return nil, nil, err
	//}
	// calculates gas amounts based on bytes of calldata, use 0 as overhead.
	l1GasUsed := h.getL2ToL1GasUsed(hex.EncodeToString([]byte{}), big.NewInt(0))
	// multiply by the fee per calldata and add the flat l2 fee
	l1Fee := new(big.Int).Mul(l1GasUsed, gasData.PerL1CalldataFee)
	l1Fee = l1Fee.Add(l1Fee, gasData.PerL2TxFee)
	return l1GasUsed, l1Fee, nil
}

//func (h *HeuristicGasModelFactory) estimateGas(
//	routeWithValidQuote models.V3RouteWithValidQuote,
//	gasPriceWei *big.Int,
//	chainId base_entities.ChainId,
//) (*entities.CurrencyAmount, *big.Int, *big.Int, error) {
//	initializedTicksCrossedSum := new(big.Int)
//	for _, v := range routeWithValidQuote.InitializedTicksCrossedList {
//		initializedTicksCrossedSum.Add(initializedTicksCrossedSum, big.NewInt(int64(v)))
//	}
//	totalInitializedTicksCrossed := math.BigMax(big.NewInt(1), initializedTicksCrossedSum)
//	totalHops := big.NewInt(int64(len(routeWithValidQuote.Route.Pools)))
//
//	hopsGasUse := new(big.Int).Mul(COST_PER_HOP(chainId), totalHops)
//	tickGasUse := new(big.Int).Mul(COST_PER_INIT_TICK(chainId), totalInitializedTicksCrossed)
//	uninitializedTickGasUse := new(big.Int).Mul(costPerUninitTick, big.NewInt(0))
//
//	// base estimate gas used based on chainId estimates for hops and ticks gas useage
//	baseGasUse := new(big.Int).Add(BASE_SWAP_COST(chainId), hopsGasUse)
//	baseGasUse = baseGasUse.Add(baseGasUse, tickGasUse).Add(baseGasUse, uninitializedTickGasUse)
//
//	baseGasCostWei := new(big.Int).Mul(gasPriceWei, baseGasUse)
//
//	wrappedCurrency, err := h.wNProvider.GetTokenByChain(chainId)
//	if err != nil {
//		return nil, nil, nil, err
//	}
//
//	totalGasCostNativeCurrency := entities.FromRawAmount(wrappedCurrency, baseGasCostWei)
//	return totalGasCostNativeCurrency, totalInitializedTicksCrossed, baseGasUse, nil
//}

func (h *HeuristicGasModelFactory) getHighestLiquidityNativePool(
	chainId base_entities.ChainId,
	token *entities.Token,
	poolProvider v3.PoolProvider,
) (*base_entities.V3Pool, error) {
	nativeCurrency, err := h.wNProvider.GetTokenByChain(chainId)
	if err != nil {
		return nil, err
	}

	indexPools, err := h.poolIndexer.GetPools(chainId, token, nativeCurrency, nil)
	if err != nil {
		return nil, err
	}
	var (
		tokenAddresses    []string
		tokenAddressesMap = make(map[string]struct{})
	)
	for _, pool := range indexPools {
		if _, has := tokenAddressesMap[pool.Token0.Id]; !has {
			tokenAddressesMap[pool.Token0.Id] = struct{}{}
			tokenAddresses = append(tokenAddresses, pool.Token0.Id)
		}
		if _, has := tokenAddressesMap[pool.Token1.Id]; !has {
			tokenAddressesMap[pool.Token1.Id] = struct{}{}
			tokenAddresses = append(tokenAddresses, pool.Token1.Id)
		}
	}

	var tokenPairs []v3.TokenPairs
	tokenAccessor, err := h.tokenProvider.GetTokens(chainId, tokenAddresses, nil)
	if err != nil {
		return nil, err
	}
	for _, pool := range indexPools {
		token0 := tokenAccessor.GetTokenByAddress(pool.Token0.Id)
		token1 := tokenAccessor.GetTokenByAddress(pool.Token1.Id)
		if token0 == nil || token1 == nil {
			continue
		}
		tokenPairs = append(tokenPairs, v3.TokenPairs{
			Token0:         token0,
			Token1:         token1,
			FeeAmount:      pool.FeeTier,
			PairAddress:    pool.Id,
			QuoteAddress:   pool.QuoteAddress,
			RouterAddress:  pool.RouterAddress,
			FactoryAddress: pool.FactoryAddress,
		})
	}
	poolAccessor, err := poolProvider.GetPools(tokenPairs, nil)
	if err != nil {
		return nil, err
	}
	pools := poolAccessor.GetAllPools()

	//var (
	//	nativePools []v3.TokenPairs
	//	feeAmounts  = []constants.FeeAmount{
	//		constants.FeeHigh,
	//		constants.FeeMedium,
	//		constants.FeeLow,
	//	}
	//)
	//for _, v := range feeAmounts {
	//	nativePools = append(nativePools, v3.TokenPairs{
	//		Token0:    nativeCurrency,
	//		Token1:    token,
	//		FeeAmount: v,
	//	})
	//}
	//nativePools = []v3.TokenPairs{
	//	{
	//		Token0: common.HexToAddress(""),
	//	},
	//}

	//poolAccessor, err := poolProvider.GetPools(nativePools, nil)
	//if err != nil {
	//	return nil, err
	//}
	//var pools []*base_entities.V3Pool
	//for _, f := range feeAmounts {
	//	pool, err := poolAccessor.GetPool(nativeCurrency, token, f)
	//	if err != nil {
	//		return nil, err
	//	}
	//	pools = append(pools, pool)
	//}
	if len(pools) == 0 {
		return nil, fmt.Errorf("could not find a %s pool with %s for computing gas costs", nativeCurrency.Symbol(), token.Symbol())
	}
	sort.Slice(pools, func(i, j int) bool {
		return pools[i].Liquidity.Cmp(pools[j].Liquidity) < 0
	})
	return pools[0], nil
}

// based on the code from the optimism OVM_GasPriceOracle contract
func (h *HeuristicGasModelFactory) getL2ToL1GasUsed(data string, overhead *big.Int) *big.Int {
	// data is hex encoded
	reg, _ := regexp.Compile("/.{1,2}/g")
	dataArr := reg.FindStringSubmatch(data[2:])
	count := 0
	for i := 0; i < len(dataArr); i++ {
		b, _ := strconv.ParseInt(dataArr[i], 16, 64)
		if b == 0 {
			count += 4
		} else {
			count += 16
		}
	}
	unsigned := new(big.Int).Add(overhead, big.NewInt(int64(count)))
	signedConversion := 68 * 16
	return new(big.Int).Add(unsigned, big.NewInt(int64(signedConversion)))
}
