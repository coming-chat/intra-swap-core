package routers

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/providers"
	"github.com/coming-chat/intra-swap-core/routers/alpha_router/config"
	"github.com/daoleno/uniswap-sdk-core/entities"
	entitiesV3 "github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/utils"
	"math/big"
)

type RoutingConfig interface {
}

type SwapRoute struct {
	Quote                      *entities.CurrencyAmount
	QuoteGasAdjusted           *entities.CurrencyAmount
	EstimatedGasUsed           *big.Int
	EstimatedGasUsedQuoteToken *entities.CurrencyAmount
	EstimatedGasUsedUSD        *entities.CurrencyAmount
	GasPriceWei                *big.Int
	Trade                      base_entities.Trade
	Route                      []providers.RouteWithValidQuote
	BlockNumber                uint64
	MethodParameters           *utils.MethodParameters
}

type SwapToRatioRoute struct {
	*SwapRoute
	OptimalRatio       *entities.Fraction
	PostSwapTargetPool *entitiesV3.Pool
}

type Router[T RoutingConfig] interface {
	Route(
		amount *entities.CurrencyAmount,
		quoteCurrency *entities.Currency,
		swapType entities.TradeType,
		partialRoutingConfig T,
		swapOptions config.SwapOptions,
	) (SwapRoute, error)
}

type SwapToRatio interface {
	RouteToRatio(
		token0Balance, token1Balance *entities.CurrencyAmount,
		position *entitiesV3.Position,
		swapAndAddConfig config.SwapAndAddConfig,
		swapAndAddOptions config.SwapAndAddOptions,
		routingConfig config.AlphaRouterConfig,
	) (*SwapToRatioRoute, error)
}
