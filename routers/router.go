package routers

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
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
	Route                      []RouteWithValidQuote
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
