package alpha_router

import (
	"context"
	"errors"
	"fmt"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/providers"
	"github.com/coming-chat/intra-swap-core/providers/gas_price"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	v2 "github.com/coming-chat/intra-swap-core/providers/v2"
	v3 "github.com/coming-chat/intra-swap-core/providers/v3"
	"github.com/coming-chat/intra-swap-core/routers"
	"github.com/coming-chat/intra-swap-core/routers/alpha_router/config"
	"github.com/coming-chat/intra-swap-core/routers/alpha_router/functions"
	"github.com/coming-chat/intra-swap-core/routers/alpha_router/gas_models"
	"github.com/coming-chat/intra-swap-core/util"
	"github.com/daoleno/uniswap-sdk-core/entities"
	entitiesV3 "github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/utils"
	"math/big"
	"sync"
)

var (
	ErrNoRouteFound = errors.New("max iterations exceeded") //SwapToRatioStatus
	ErrSwapNeeded   = errors.New("no swap needed: amountToSwap = 0")
)

type AlphaRouterParams struct {
	/**
	 * The chain id for this instance of the Alpha Router.
	 */
	ChainId base_entities.ChainId
	/**
	 * The Web3 provider for getting on-chain data.
	 */
	Provider rpc.BaseProvider
	/**
	 * The provider to use for making multicalls. Used for getting on-chain data
	 * like pools, tokens, quotes in batch.
	 */
	MultiCall2ProviderCore rpc.MultiCallProviderCore
	/**
	 * The provider for getting all pools that exist on V3 from the Subgraph. The pools
	 * from this provider are filtered during the algorithm to a set of candidate pools.
	 */
	V3IndexerProvider v3.IndexerProvider
	/**
	 * The provider for getting data about V3 pools.
	 */
	V3PoolProvider v3.PoolProvider
	/**
	 * The provider for getting V3 quotes.
	 */
	V3QuoteProvider v3.QuoteProvider
	/**
	 * The provider for getting all pools that exist on V2 from the Subgraph. The pools
	 * from this provider are filtered during the algorithm to a set of candidate pools.
	 */
	V2IndexerProvider v2.IndexerProvider
	/**
	 * The provider for getting data about V2 pools.
	 */
	V2PoolProvider v2.PoolProvider
	/**
	 * The provider for getting V3 quotes.
	 */
	V2QuoteProvider v2.QuoteProvider
	/**
	 * The provider for getting data about Tokens.
	 */
	TokenProvider providers.TokenProvider
	/**
	 * The provider for getting the current gas price to use when account for gas in the
	 * algorithm.
	 */
	GasPriceProvider gas_price.GasPriceProvider
	/**
	 * A factory for generating a gas model that is used when estimating the gas used by
	 * V3 routes.
	 */
	V3GasModelFactory gas_models.V3GasModelFactory
	/**
	 * A factory for generating a gas model that is used when estimating the gas used by
	 * V2 routes.
	 */
	V2GasModelFactory gas_models.V2GasModelFactory
	/**
	 * A token list that specifies Token that should be blocked from routing through.
	 * Defaults to Uniswap's unsupported token list.
	 */
	BlockedTokenListProvider *providers.TokenListProvider

	/**
	 * Calls lens function on SwapRouter02 to determine ERC20 approval types for
	 * LP position tokens.
	 */
	//SwapRouterProvider providers.

	/**
	 * Calls the optimism gas oracle contract to fetch constants for calculating the l1 security fee.
	 */
	//OptimismGasDataProvider IL2GasDataProvider<OptimismGasData>

	/**
	 * A token validator for detecting fee-on-transfer tokens or tokens that can't be transferred.
	 */
	//TokenValidatorProvider ITokenValidatorProvider

	/**
	 * Calls the arbitrum gas data contract to fetch constants for calculating the l1 fee.
	 */
	//arbitrumGasDataProvider?: IL2GasDataProvider<ArbitrumGasData>
}

func NewAlphaRouter(params AlphaRouterParams, ctx context.Context) *AlphaRouter {
	router := &AlphaRouter{
		ChainId:                   params.ChainId,
		Provider:                  params.Provider,
		MultiCall2Provider:        params.MultiCall2ProviderCore,
		V3IndexerProviderProvider: params.V3IndexerProvider,
		V3PoolProvider:            params.V3PoolProvider,
		V3QuoteProvider:           params.V3QuoteProvider,
		V2IndexerProvider:         params.V2IndexerProvider,
		V2PoolProvider:            params.V2PoolProvider,
		V2QuoteProvider:           params.V2QuoteProvider,
		TokenProvider:             params.TokenProvider,
		GasPriceProvider:          params.GasPriceProvider,
		V3GasModelFactory:         params.V3GasModelFactory,
		V2GasModelFactory:         params.V2GasModelFactory,
		BlockedTokenListProvider:  params.BlockedTokenListProvider,
		Ctx:                       ctx,
	}
	return router
}

type AlphaRouter struct {
	ChainId                   base_entities.ChainId
	Provider                  rpc.BaseProvider
	MultiCall2Provider        rpc.MultiCallProviderCore
	V3IndexerProviderProvider v3.IndexerProvider
	V3PoolProvider            v3.PoolProvider
	V3QuoteProvider           v3.QuoteProvider
	V2IndexerProvider         v2.IndexerProvider
	V2PoolProvider            v2.PoolProvider
	V2QuoteProvider           v2.QuoteProvider
	TokenProvider             providers.TokenProvider
	GasPriceProvider          gas_price.GasPriceProvider
	//SwapRouterProvider providers.SwapRouterProvider
	V3GasModelFactory gas_models.V3GasModelFactory
	V2GasModelFactory gas_models.V2GasModelFactory
	//TokenValidatorProvider   TokenValidatorProvider
	BlockedTokenListProvider *providers.TokenListProvider
	Ctx                      context.Context
}

func (a *AlphaRouter) Route(
	amount *entities.CurrencyAmount,
	quoteCurrency entities.Currency,
	tradeType entities.TradeType,
	swapConfig *config.SwapOptions,
	routingConfig config.AlphaRouterConfig,
) (*routers.SwapRoute, error) {
	var err error
	if routingConfig.BlockNumber == 0 {
		routingConfig.BlockNumber, err = a.Provider.Rpc.BlockNumber(a.Ctx)
		if err != nil {
			return nil, err
		}
	}
	var currencyIn, currencyOut entities.Currency
	if tradeType == entities.ExactInput {
		currencyIn = amount.Currency
		currencyOut = quoteCurrency
	} else {
		currencyIn = quoteCurrency
		currencyOut = amount.Currency
	}
	tokenIn, tokenOut := currencyIn.Wrapped(), currencyOut.Wrapped()
	percents, amounts := a.getAmountDistribution(amount, routingConfig)
	gasPrice := a.GasPriceProvider.GetGasPrice()
	quoteToken := quoteCurrency.Wrapped()
	protocolsSet := make(map[base_entities.Protocol]struct{})
	for _, v := range routingConfig.Protocols {
		protocolsSet[v] = struct{}{}
	}

	gasModel, err := a.V3GasModelFactory.BuildGasModel(
		a.ChainId,
		gasPrice.GasPriceWei,
		a.V3PoolProvider,
		quoteToken,
	)
	if err != nil {
		return nil, err
	}

	_, hasV2 := protocolsSet[base_entities.V2]
	_, hasV3 := protocolsSet[base_entities.V3]
	_, v2Supported := base_entities.V2Supported[a.ChainId]
	syncGroup := sync.WaitGroup{}
	lock := sync.Mutex{}
	var routesWithValidQuotesAndPools []routesWithValidQuotesAndPool
	if (len(protocolsSet) == 0 || (hasV2 && hasV3)) && v2Supported {
		syncGroup.Add(1)
		go func() {
			defer syncGroup.Done()
			routes, err := a.getV3Quotes(tokenIn, tokenOut, amounts, percents, quoteToken, gasModel, tradeType, routingConfig)
			if err != nil {
				fmt.Printf("getV3Quotes err: %v", err)
				return
			}
			lock.Lock()
			routesWithValidQuotesAndPools = append(routesWithValidQuotesAndPools, *routes)
			lock.Unlock()
		}()
		syncGroup.Add(1)
		go func() {
			defer syncGroup.Done()
			routes, err := a.getV2Quotes(tokenIn, tokenOut, amounts, percents, quoteToken, gasPrice.GasPriceWei, tradeType, routingConfig)
			if err != nil {
				fmt.Printf("getV2Quotes err: %v", err)
				return
			}
			lock.Lock()
			routesWithValidQuotesAndPools = append(routesWithValidQuotesAndPools, *routes)
			lock.Unlock()
		}()
	} else {
		if hasV3 || (len(protocolsSet) == 0 && !v2Supported) {
			syncGroup.Add(1)
			go func() {
				defer syncGroup.Done()
				routes, err := a.getV3Quotes(tokenIn, tokenOut, amounts, percents, quoteToken, gasModel, tradeType, routingConfig)
				if err != nil {
					fmt.Printf("getV3Quotes err: %v", err)
					return
				}
				lock.Lock()
				routesWithValidQuotesAndPools = append(routesWithValidQuotesAndPools, *routes)
				lock.Unlock()
			}()
		}
		if hasV2 {
			syncGroup.Add(1)
			go func() {
				defer syncGroup.Done()
				routes, err := a.getV2Quotes(tokenIn, tokenOut, amounts, percents, quoteToken, gasPrice.GasPriceWei, tradeType, routingConfig)
				if err != nil {
					fmt.Printf("getV2Quotes err: %v", err)
					return
				}
				lock.Lock()
				routesWithValidQuotesAndPools = append(routesWithValidQuotesAndPools, *routes)
				lock.Unlock()
			}()
		}
	}
	syncGroup.Wait()
	var (
		allRoutesWithValidQuotes []providers.RouteWithValidQuote
		allCandidatePools        []functions.CandidatePoolsBySelectionCriteria
	)
	for _, r := range routesWithValidQuotesAndPools {
		allCandidatePools = append(allCandidatePools, r.CandidatePools)
		allRoutesWithValidQuotes = append(allRoutesWithValidQuotes, r.RoutesWithValidQuotes...)
	}
	if len(allRoutesWithValidQuotes) == 0 {
		return nil, errors.New("received no valid quotes")
	}

	swapRouteRaw, err := functions.GetBestSwapRoute(
		amount,
		percents,
		allRoutesWithValidQuotes,
		tradeType,
		a.ChainId,
		routingConfig,
		&gasModel,
	)
	if err != nil {
		return nil, err
	}
	//trade := utils.BuildTrade(
	//	currencyIn,
	//	currencyOut,
	//	tradeType,
	//	swapRouteRaw.Routes,
	//)
	var methodParameters *utils.MethodParameters
	//if swapConfig != nil {
	//	methodParameters, err = util.BuildSwapMethodParameters(trade, *swapConfig)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	return &routers.SwapRoute{
		Quote:                      swapRouteRaw.Quote,
		QuoteGasAdjusted:           swapRouteRaw.QuoteGasAdjusted,
		EstimatedGasUsed:           swapRouteRaw.EstimatedGasUsed,
		EstimatedGasUsedQuoteToken: swapRouteRaw.EstimatedGasUsedQuoteToken,
		EstimatedGasUsedUSD:        swapRouteRaw.EstimatedGasUsedUSD,
		GasPriceWei:                gasPrice.GasPriceWei,
		Route:                      swapRouteRaw.Routes,
		Trade:                      nil,
		MethodParameters:           methodParameters,
		BlockNumber:                routingConfig.BlockNumber,
	}, nil
}

func (a *AlphaRouter) RouteToRatio(
	token0Balance *entities.CurrencyAmount,
	token1Balance *entities.CurrencyAmount,
	position *entitiesV3.Position,
	swapAndAddConfig config.SwapAndAddConfig,
	swapAndAddOptions *config.SwapAndAddOptions,
	routingConfig config.AlphaRouterConfig,
) (*routers.SwapToRatioRoute, error) {
	before, err := token0Balance.Currency.Wrapped().SortsBefore(token1Balance.Currency.Wrapped())
	if err != nil {
		return nil, err
	}
	if before {
		token1Balance, token0Balance = token0Balance, token1Balance
	}
	preSwapOptimalRatio, err := a.calculateOptimalRatio(position, position.Pool.SqrtRatioX96, true)
	if err != nil {
		return nil, err
	}
	var zeroForOne bool
	if position.Pool.TickCurrent > position.TickUpper {
		zeroForOne = true
	} else if position.Pool.TickCurrent < position.TickLower {
		zeroForOne = false
	} else {
		zeroForOne = entities.NewFraction(token0Balance.Quotient(), token1Balance.Quotient()).GreaterThan(preSwapOptimalRatio)
		if !zeroForOne {
			preSwapOptimalRatio = preSwapOptimalRatio.Invert()
		}
	}
	var (
		swap                                      *routers.SwapRoute
		ratioAchieved                             bool
		inputBalance, outputBalance, exchangeRate = token0Balance, token1Balance, position.Pool.Token0Price().Fraction
		optimalRatio                              = preSwapOptimalRatio
		postSwapTargetPool                        = position.Pool
		n                                         = 0
	)
	if !zeroForOne {
		inputBalance = token1Balance
		outputBalance = token0Balance
		exchangeRate = position.Pool.Token1Price().Fraction
	}
	for !ratioAchieved {
		n++
		if n > swapAndAddConfig.MaxIterations {
			return nil, ErrNoRouteFound
		}
		amountToSwap, err := functions.CalculateRatioAmountIn(
			optimalRatio,
			exchangeRate,
			inputBalance,
			outputBalance,
		)
		if err != nil {
			return nil, err
		}

		if amountToSwap.EqualTo(util.ZeroFraction) {
			return nil, ErrSwapNeeded
		}
		swap, err = a.Route(amountToSwap, outputBalance.Currency, entities.ExactInput, nil, routingConfig)
		if err != nil {
			return nil, err
		}
		inputBalanceUpdated := inputBalance.Subtract(swap.Trade.InputAmount())
		outputBalanceUpdated := outputBalance.Add(swap.Trade.OutputAmount())
		newRatio := inputBalanceUpdated.Divide(outputBalanceUpdated.Fraction)

		var targetPoolPriceUpdate *big.Int
		for _, route := range swap.Route {
			if route.Protocol() != base_entities.V3 {
				continue
			}
			v3Route := route.(providers.V3RouteWithValidQuote)
			for i, v3Pool := range v3Route.Route.Pools {
				if v3Pool.Token0.Equal(position.Pool.Token0) && v3Pool.Token1.Equal(position.Pool.Token1) && v3Pool.Fee == position.Pool.Fee {
					targetPoolPriceUpdate = v3Route.SqrtPriceX96AfterList[i]
					optimalRatio, err = a.calculateOptimalRatio(position, targetPoolPriceUpdate, zeroForOne)
					if err != nil {
						return nil, err
					}
				}
			}
		}
		if targetPoolPriceUpdate == nil {
			optimalRatio = preSwapOptimalRatio
		}
		ratioAchieved = newRatio.EqualTo(optimalRatio) || absoluteValue(newRatio.Fraction.Divide(optimalRatio).Subtract(util.OneFraction)).LessThan(swapAndAddConfig.RatioErrorTolerance)

		if ratioAchieved && targetPoolPriceUpdate != nil {
			ratio, err := utils.GetTickAtSqrtRatio(targetPoolPriceUpdate)
			if err != nil {
				return nil, err
			}
			postSwapTargetPool, err = entitiesV3.NewPool(
				position.Pool.Token0,
				position.Pool.Token1,
				position.Pool.Fee,
				targetPoolPriceUpdate,
				position.Pool.Liquidity,
				ratio,
				position.Pool.TickDataProvider,
			)
		}
		exchangeRate = swap.Trade.OutputAmount().Divide(swap.Trade.InputAmount().Fraction).Fraction
		if exchangeRate.EqualTo(util.ZeroFraction) {
			return nil, errors.New("insufficient liquidity to swap to optimal ratio")
		}
	}
	if swap == nil {
		return nil, errors.New("no route found")
	}

	result := &routers.SwapToRatioRoute{
		SwapRoute:          swap,
		OptimalRatio:       optimalRatio,
		PostSwapTargetPool: postSwapTargetPool,
	}
	if swapAndAddOptions != nil {
		result.MethodParameters, err = a.buildSwapAndAddMethodParameters(swap.Trade, swapAndAddOptions, base_entities.SwapAndAddParameters{
			InitialBalanceTokenIn:  inputBalance,
			InitialBalanceTokenOut: outputBalance,
			PreLiquidityPosition:   position,
		})
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (a *AlphaRouter) calculateOptimalRatio(
	position *entitiesV3.Position,
	sqrtRatioX96 *big.Int,
	zeroForOne bool,
) (*entities.Fraction, error) {
	upperSqrtRatioX96, err := utils.GetSqrtRatioAtTick(position.TickUpper)
	if err != nil {
		return nil, err
	}
	lowerSqrtRatioX96, err := utils.GetSqrtRatioAtTick(position.TickLower)
	if err != nil {
		return nil, err
	}
	if sqrtRatioX96.Cmp(upperSqrtRatioX96) > 0 || sqrtRatioX96.Cmp(lowerSqrtRatioX96) < 0 {
		return entities.NewFraction(big.NewInt(0), big.NewInt(1)), nil
	}

	precision := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	optimalRatio := entities.NewFraction(
		utils.GetAmount0Delta(sqrtRatioX96, upperSqrtRatioX96, precision, true),
		utils.GetAmount1Delta(sqrtRatioX96, lowerSqrtRatioX96, precision, true),
	)
	if !zeroForOne {
		optimalRatio = optimalRatio.Invert()
	}
	return optimalRatio, nil
}

func (a *AlphaRouter) getAmountDistribution(
	amount *entities.CurrencyAmount,
	routingConfig config.AlphaRouterConfig,
) ([]int, []*entities.CurrencyAmount) {
	var (
		percents []int
		amounts  []*entities.CurrencyAmount
	)
	for i := 1; i <= 100/routingConfig.DistributionPercent; i++ {
		percents = append(percents, i*routingConfig.DistributionPercent)
		amounts = append(
			amounts,
			amount.Multiply(
				entities.NewFraction(
					big.NewInt(
						int64(i*routingConfig.DistributionPercent),
					),
					big.NewInt(100),
				),
			),
		)
	}
	return percents, amounts
}

func (a *AlphaRouter) getV3Quotes(
	tokenIn *entities.Token,
	tokenOut *entities.Token,
	amounts []*entities.CurrencyAmount,
	percents []int,
	quoteToken *entities.Token,
	gasModel providers.GasModel[providers.V3RouteWithValidQuote],
	swapType entities.TradeType,
	routingConfig config.AlphaRouterConfig,
) (
	*routesWithValidQuotesAndPool,
	error,
) {
	poolAccessor, candidatePools, err := functions.GetV3CandidatePools(functions.V3GetCandidatePoolsParams{
		TokenIn:                  tokenIn,
		TokenOut:                 tokenOut,
		TokenProvider:            a.TokenProvider,
		BlockedTokenListProvider: a.BlockedTokenListProvider,
		PoolProvider:             a.V3PoolProvider,
		RouteType:                swapType,
		IndexerProvider:          a.V3IndexerProviderProvider,
		RoutingConfig:            routingConfig,
		ChainId:                  a.ChainId,
	})
	if err != nil {
		return nil, err
	}
	poolsRaw := poolAccessor.GetAllPools()
	//a.applyTokenValidatorToPools() temp ship

	routes, err := functions.ComputeAllV3Routes(
		tokenIn,
		tokenOut,
		poolsRaw,
		routingConfig.MaxSwapsPerPath,
	)
	if err != nil {
		return nil, err
	}

	if len(routes) == 0 {
		return &routesWithValidQuotesAndPool{
			CandidatePools: candidatePools,
		}, nil
	}

	var (
		routesWithQuotes []v3.RouteWithQuotes
	)

	// For all our routes, and all the fractional amounts, fetch quotes on-chain.
	if swapType == entities.ExactInput {
		routesWithQuotes, err = a.V3QuoteProvider.GetQuotesManyExactIn(amounts, routes)
	} else {
		routesWithQuotes, err = a.V3QuoteProvider.GetQuotesManyExactOut(amounts, routes)
	}
	if err != nil {
		return nil, err
	}

	var routesWithValidQuotes []providers.RouteWithValidQuote
	for _, routeWithQuote := range routesWithQuotes {
		for i := 0; i < len(routeWithQuote.AmountQuote); i++ {
			percent := percents[i]
			amountQuote := routeWithQuote.AmountQuote[i]

			if amountQuote.Quote == nil || amountQuote.SqrtPriceX96AfterList == nil || amountQuote.InitializedTicksCrossedList == nil || amountQuote.GasEstimate == nil {
				continue
			}

			routeWithValidQuote, err := providers.NewV3RouteWithValidQuote(providers.V3RouteWithValidQuoteParams{
				Route:                       routeWithQuote.Route,
				RawQuote:                    amountQuote.Quote,
				Amount:                      amountQuote.Amount,
				Percent:                     percent,
				SqrtPriceX96AfterList:       amountQuote.SqrtPriceX96AfterList,
				InitializedTicksCrossedList: amountQuote.InitializedTicksCrossedList,
				QuoterGasEstimate:           amountQuote.GasEstimate,
				GasModel:                    gasModel,
				QuoteToken:                  quoteToken,
				TradeType:                   swapType,
				PoolAccessor:                poolAccessor,
			})
			if err != nil {
				return nil, err
			}
			routesWithValidQuotes = append(routesWithValidQuotes, routeWithValidQuote)
		}
	}
	return &routesWithValidQuotesAndPool{
		RoutesWithValidQuotes: routesWithValidQuotes,
		CandidatePools:        candidatePools,
	}, nil
}

//func (a *AlphaRouter[T]) applyTokenValidatorToPools(
//	pools []T,
//	isInvalidFn func(token *entities.Currency, tokenValidation TokenValidationResult) bool) []T {
//
//}

type routesWithValidQuotesAndPool struct {
	RoutesWithValidQuotes []providers.RouteWithValidQuote
	CandidatePools        functions.CandidatePoolsBySelectionCriteria
}

func absoluteValue(fraction *entities.Fraction) *entities.Fraction {
	numeratorAbs := new(big.Int).Abs(fraction.Numerator)
	denominatorAbs := new(big.Int).Abs(fraction.Denominator)
	return entities.NewFraction(numeratorAbs, denominatorAbs)
}

func (a *AlphaRouter) buildSwapAndAddMethodParameters(
	trade base_entities.Trade,
	swapAndAddOptions *config.SwapAndAddOptions,
	swapAndAddParameters base_entities.SwapAndAddParameters,
) (*utils.MethodParameters, error) {

	//periphery.SwapCallParameters()
	return nil, nil
}

func (a *AlphaRouter) getV2Quotes(
	tokenIn *entities.Token,
	tokenOut *entities.Token,
	amounts []*entities.CurrencyAmount,
	percents []int,
	quoteToken *entities.Token,
	gasPriceWei *big.Int,
	swapType entities.TradeType,
	routingConfig config.AlphaRouterConfig,
) (*routesWithValidQuotesAndPool, error) {
	poolAccessor, candidatePools, err := functions.GetV2CandidatePools(functions.V2GetCandidatePoolsParams{
		TokenIn:                  tokenIn,
		TokenOut:                 tokenOut,
		TokenProvider:            a.TokenProvider,
		BlockedTokenListProvider: a.BlockedTokenListProvider,
		PoolProvider:             a.V2PoolProvider,
		RouteType:                swapType,
		IndexerProvider:          a.V2IndexerProvider,
		RoutingConfig:            routingConfig,
		ChainId:                  a.ChainId,
	})
	if err != nil {
		return nil, err
	}
	poolsRaw := poolAccessor.GetAllPools()

	//pools : = a.
	pools := poolsRaw
	routes, err := functions.ComputeAllV2Routes(
		tokenIn,
		tokenOut,
		pools,
		routingConfig.MaxSwapsPerPath,
	)
	if err != nil {
		return nil, err
	}
	if len(routes) == 0 {
		return &routesWithValidQuotesAndPool{
			CandidatePools: candidatePools,
		}, nil
	}

	var routesWithQuotes []v2.RouteWithQuotes
	if swapType == entities.ExactInput {
		routesWithQuotes, err = a.V2QuoteProvider.GetQuotesManyExactIn(amounts, routes)
	} else {
		routesWithQuotes, err = a.V2QuoteProvider.GetQuotesManyExactOut(amounts, routes)
	}
	if err != nil {
		return nil, err
	}

	gasModel, err := a.V2GasModelFactory.BuildGasModel(
		a.ChainId,
		gasPriceWei,
		a.V2PoolProvider,
		quoteToken,
	)
	if err != nil {
		return nil, err
	}

	var routesWithValidQuotes []providers.RouteWithValidQuote
	for _, routeWithQuote := range routesWithQuotes {
		for i := 0; i < len(routeWithQuote.AmountQuotes); i++ {
			percent := percents[i]
			amountQuote := routeWithQuote.AmountQuotes[i]
			if amountQuote.Quote == nil {
				continue
			}
			routeWithValidQuote, err := providers.NewV2RouteWithValidQuote(providers.V2RouteWithValidQuoteParams{
				Route:        routeWithQuote.Route,
				RawQuote:     amountQuote.Quote,
				Amount:       amountQuote.Amount,
				Percent:      percent,
				GasModel:     gasModel,
				QuoteToken:   quoteToken,
				TradeType:    swapType,
				PoolAccessor: poolAccessor,
			})
			if err != nil {
				return nil, err
			}
			routesWithValidQuotes = append(routesWithValidQuotes, routeWithValidQuote)
		}
	}
	return &routesWithValidQuotesAndPool{
		RoutesWithValidQuotes: routesWithValidQuotes,
		CandidatePools:        candidatePools,
	}, nil
}
