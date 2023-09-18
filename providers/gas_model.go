package providers

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
	v2 "github.com/coming-chat/intra-swap-core/providers/v2"
	v3 "github.com/coming-chat/intra-swap-core/providers/v3"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"math/big"
)

type L1ToL2GasCosts struct {
	GasUsedL1           *big.Int
	GasCostL1USD        *entities.CurrencyAmount
	GasCostL1QuoteToken *entities.CurrencyAmount
}

type GasModel[T RouteWithValidQuote] struct {
	EstimateGasCostFn func(routeWithValidQuote T) (
		gasEstimate *big.Int,
		gasCostInToken,
		gasCostInUSD *entities.CurrencyAmount,
		err error,
	)
	CalculateL1GasFeesFn func(routes []T) (*L1ToL2GasCosts, error)
}

func (g GasModel[T]) EstimateGasCost(routeWithValidQuote T) (
	gasEstimate *big.Int,
	gasCostInToken,
	gasCostInUSD *entities.CurrencyAmount,
	err error,
) {
	return g.EstimateGasCostFn(routeWithValidQuote)
}

func (g GasModel[T]) CalculateL1GasFees(routes []T) (*L1ToL2GasCosts, error) {
	return g.CalculateL1GasFeesFn(routes)
}

type RouteWithValidQuote interface {
	Protocol() base_entities.Protocol
	GetBaseRouteWithValidQuote() *BaseRouteWithValidQuote
}

type BaseRouteWithValidQuote struct {
	Amount  *entities.CurrencyAmount
	Percent int
	// If exact in, v3rwvq is (quote - gasCostInToken). If exact out, v3rwvq is (quote + gasCostInToken).
	QuoteAdjustedForGas *entities.CurrencyAmount
	Quote               *entities.CurrencyAmount
	GasEstimate         *big.Int
	// The gas cost in terms of the quote token.
	GasCostInToken *entities.CurrencyAmount
	GasCostInUSD   *entities.CurrencyAmount
	TradeType      entities.TradeType
	PoolAddresses  []string
	TokenPath      []*entities.Token
	QuoteToken     *entities.Token
	RawQuote       *big.Int
}

type V2RouteWithValidQuote struct {
	*BaseRouteWithValidQuote
	Route        *base_entities.V2Route
	PoolAccessor v2.PoolAccessor
	GasModel     GasModel[V2RouteWithValidQuote]
}

func (v V2RouteWithValidQuote) Protocol() base_entities.Protocol {
	return base_entities.V2
}

func (v V2RouteWithValidQuote) GetBaseRouteWithValidQuote() *BaseRouteWithValidQuote {
	return v.BaseRouteWithValidQuote
}

type V2RouteWithValidQuoteParams struct {
	Amount       *entities.CurrencyAmount
	RawQuote     *big.Int
	Percent      int
	Route        *base_entities.V2Route
	GasModel     GasModel[V2RouteWithValidQuote]
	QuoteToken   *entities.Token
	TradeType    entities.TradeType
	PoolAccessor v2.PoolAccessor
}

func NewV2RouteWithValidQuote(param V2RouteWithValidQuoteParams) (*V2RouteWithValidQuote, error) {
	v2rwvq := &V2RouteWithValidQuote{
		BaseRouteWithValidQuote: &BaseRouteWithValidQuote{
			Amount:     param.Amount,
			RawQuote:   param.RawQuote,
			Quote:      entities.FromRawAmount(param.QuoteToken, param.RawQuote),
			Percent:    param.Percent,
			QuoteToken: param.QuoteToken,
			TradeType:  param.TradeType,
		},
		Route:        param.Route,
		GasModel:     param.GasModel,
		PoolAccessor: param.PoolAccessor,
	}
	gasEstimate, gasCostInToken, gasCostInUSD, err := v2rwvq.GasModel.EstimateGasCost(*v2rwvq)
	if err != nil {
		return nil, err
	}
	v2rwvq.GasCostInToken = gasCostInToken
	v2rwvq.GasCostInUSD = gasCostInUSD
	v2rwvq.GasEstimate = gasEstimate

	if v2rwvq.TradeType == entities.ExactInput {
		v2rwvq.QuoteAdjustedForGas = v2rwvq.Quote.Subtract(gasCostInToken)
	} else {
		v2rwvq.QuoteAdjustedForGas = v2rwvq.Quote.Add(gasCostInToken)
	}

	for _, p := range param.Route.Pairs {
		basePool := param.PoolAccessor.GetBasePool(p)
		v2rwvq.PoolAddresses = append(v2rwvq.PoolAddresses, basePool.GetAddress().String())
	}
	v2rwvq.TokenPath = v2rwvq.Route.Path
	return v2rwvq, nil
}

type V3RouteWithValidQuote struct {
	*BaseRouteWithValidQuote
	Route                       *base_entities.V3Route
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	QuoterGasEstimate           *big.Int
	GasModel                    GasModel[V3RouteWithValidQuote]
	PoolAccessor                v3.PoolAccessor
}

func (v V3RouteWithValidQuote) Protocol() base_entities.Protocol {
	return base_entities.V3
}

func (v V3RouteWithValidQuote) GetBaseRouteWithValidQuote() *BaseRouteWithValidQuote {
	return v.BaseRouteWithValidQuote
}

type V3RouteWithValidQuoteParams struct {
	Amount                      *entities.CurrencyAmount
	RawQuote                    *big.Int
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	QuoterGasEstimate           *big.Int
	Percent                     int
	Route                       *base_entities.V3Route
	GasModel                    GasModel[V3RouteWithValidQuote]
	QuoteToken                  *entities.Token
	TradeType                   entities.TradeType
	PoolAccessor                v3.PoolAccessor
}

func NewV3RouteWithValidQuote(param V3RouteWithValidQuoteParams) (*V3RouteWithValidQuote, error) {
	v3rwvq := &V3RouteWithValidQuote{
		BaseRouteWithValidQuote: &BaseRouteWithValidQuote{
			Amount:     param.Amount,
			RawQuote:   param.RawQuote,
			Quote:      entities.FromRawAmount(param.QuoteToken, param.RawQuote),
			Percent:    param.Percent,
			QuoteToken: param.QuoteToken,
			TradeType:  param.TradeType,
		},
		SqrtPriceX96AfterList:       param.SqrtPriceX96AfterList,
		InitializedTicksCrossedList: param.InitializedTicksCrossedList,
		QuoterGasEstimate:           param.QuoterGasEstimate,
		Route:                       param.Route,
		GasModel:                    param.GasModel,
		PoolAccessor:                param.PoolAccessor,
	}

	gasEstimate, gasCostInToken, gasCostInUSD, err := v3rwvq.GasModel.EstimateGasCost(*v3rwvq)
	if err != nil {
		return nil, err
	}

	v3rwvq.GasCostInToken = gasCostInToken
	v3rwvq.GasCostInUSD = gasCostInUSD
	v3rwvq.GasEstimate = gasEstimate

	// If its exact out, we need to request *more* of the input token to account for the gas.
	if v3rwvq.TradeType == entities.ExactInput {
		v3rwvq.QuoteAdjustedForGas = v3rwvq.Quote.Subtract(gasCostInToken)
	} else {
		v3rwvq.QuoteAdjustedForGas = v3rwvq.Quote.Add(gasCostInToken)
	}
	for _, pool := range param.Route.Pools {
		basePool := param.PoolAccessor.GetBasePool(pool)
		v3rwvq.PoolAddresses = append(v3rwvq.PoolAddresses, basePool.PoolAddress().String())
	}

	v3rwvq.TokenPath = v3rwvq.Route.TokenPath
	return v3rwvq, nil
}
