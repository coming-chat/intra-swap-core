package models

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
	v2 "github.com/coming-chat/intra-swap-core/providers/v2"
	v3 "github.com/coming-chat/intra-swap-core/providers/v3"
	"github.com/coming-chat/intra-swap-core/routers"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"math/big"
)

type L1ToL2GasCosts struct {
	GasUsedL1           *big.Int
	GasCostL1USD        *entities.CurrencyAmount
	GasCostL1QuoteToken *entities.CurrencyAmount
}

type GasModel[T routers.RouteWithValidQuote] struct {
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

type V3GasModelFactory interface {
	BuildGasModel(
		chainId base_entities.ChainId,
		gasPriceWei *big.Int,
		poolProvider v3.PoolProvider,
		inTermsOfToken *entities.Token,
	) (GasModel[V3RouteWithValidQuote], error)
}

type V2GasModelFactory interface {
	BuildGasModel(
		chainId base_entities.ChainId,
		gasPriceWei *big.Int,
		poolProvider v2.PoolProvider,
		token *entities.Token,
	) (GasModel[V2RouteWithValidQuote], error)
}
