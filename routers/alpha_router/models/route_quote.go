package models

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/routers"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"math/big"
)

type V2RouteWithValidQuote struct {
	*routers.BaseRouteWithValidQuote
	GasModel GasModel[V2RouteWithValidQuote]
}

func (v V2RouteWithValidQuote) Protocol() base_entities.Protocol {
	return base_entities.V2
}

func (v V2RouteWithValidQuote) GetBaseRouteWithValidQuote() *routers.BaseRouteWithValidQuote {
	return v.BaseRouteWithValidQuote
}

type V2RouteWithValidQuoteParams struct {
	Amount     *entities.CurrencyAmount
	RawQuote   *big.Int
	Percent    int
	Route      *base_entities.MRoute
	GasModel   GasModel[V2RouteWithValidQuote]
	QuoteToken *entities.Token
	TradeType  entities.TradeType
	QuoteList  []*big.Int
}

func NewV2RouteWithValidQuote(param V2RouteWithValidQuoteParams) (*V2RouteWithValidQuote, error) {
	v2rwvq := &V2RouteWithValidQuote{
		BaseRouteWithValidQuote: &routers.BaseRouteWithValidQuote{
			Amount:     param.Amount,
			RawQuote:   param.RawQuote,
			Quote:      entities.FromRawAmount(param.QuoteToken, param.RawQuote),
			Percent:    param.Percent,
			QuoteToken: param.QuoteToken,
			TradeType:  param.TradeType,
			Route:      param.Route,
			QuoteList:  param.QuoteList,
		},
		GasModel: param.GasModel,
	}
	//gasEstimate, gasCostInToken, gasCostInUSD, err := v2rwvq.GasModel.EstimateGasCost(*v2rwvq)
	//if err != nil {
	//	return nil, err
	//}
	//v2rwvq.GasCostInToken = gasCostInToken
	//v2rwvq.GasCostInUSD = gasCostInUSD
	//v2rwvq.GasEstimate = gasEstimate

	for _, pool := range param.Route.Pools {
		v2rwvq.BaseRouteWithValidQuote.PoolAddresses = append(v2rwvq.PoolAddresses, pool.PoolAddress().String())
	}
	if v2rwvq.TradeType == entities.ExactInput {
		v2rwvq.QuoteAdjustedForGas = v2rwvq.Quote //.Subtract(gasCostInToken)
	} else {
		v2rwvq.QuoteAdjustedForGas = v2rwvq.Quote //.Add(gasCostInToken)
	}

	return v2rwvq, nil
}

type V3RouteWithValidQuote struct {
	*routers.BaseRouteWithValidQuote
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	QuoterGasEstimate           *big.Int
	GasModel                    GasModel[V3RouteWithValidQuote]
}

func (v V3RouteWithValidQuote) Protocol() base_entities.Protocol {
	return base_entities.V3
}

func (v V3RouteWithValidQuote) GetBaseRouteWithValidQuote() *routers.BaseRouteWithValidQuote {
	return v.BaseRouteWithValidQuote
}

type V3RouteWithValidQuoteParams struct {
	Amount                      *entities.CurrencyAmount
	RawQuote                    *big.Int
	QuoteList                   []*big.Int
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	QuoterGasEstimate           *big.Int
	Percent                     int
	Route                       *base_entities.MRoute
	GasModel                    GasModel[V3RouteWithValidQuote]
	QuoteToken                  *entities.Token
	TradeType                   entities.TradeType
}

func NewV3RouteWithValidQuote(param V3RouteWithValidQuoteParams) (*V3RouteWithValidQuote, error) {
	v3rwvq := &V3RouteWithValidQuote{
		BaseRouteWithValidQuote: &routers.BaseRouteWithValidQuote{
			Amount:     param.Amount,
			RawQuote:   param.RawQuote,
			Quote:      entities.FromRawAmount(param.QuoteToken, param.RawQuote),
			Percent:    param.Percent,
			QuoteToken: param.QuoteToken,
			TradeType:  param.TradeType,
			Route:      param.Route,
			QuoteList:  param.QuoteList,
		},
		SqrtPriceX96AfterList:       param.SqrtPriceX96AfterList,
		InitializedTicksCrossedList: param.InitializedTicksCrossedList,
		QuoterGasEstimate:           param.QuoterGasEstimate,
		GasModel:                    param.GasModel,
	}

	//gasEstimate, gasCostInToken, gasCostInUSD, err := v3rwvq.GasModel.EstimateGasCost(*v3rwvq)
	//if err != nil {
	//	return nil, err
	//}
	for _, pool := range param.Route.Pools {
		v3rwvq.BaseRouteWithValidQuote.PoolAddresses = append(v3rwvq.PoolAddresses, pool.PoolAddress().String())
	}
	//v3rwvq.GasCostInToken = gasCostInToken
	//v3rwvq.GasCostInUSD = gasCostInUSD
	//v3rwvq.GasEstimate = gasEstimate

	// If its exact out, we need to request *more* of the input token to account for the gas.
	if v3rwvq.TradeType == entities.ExactInput {
		v3rwvq.QuoteAdjustedForGas = v3rwvq.Quote //.Subtract(gasCostInToken)
	} else {
		v3rwvq.QuoteAdjustedForGas = v3rwvq.Quote //.Add(gasCostInToken)
	}

	return v3rwvq, nil
}
