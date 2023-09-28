package routers

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts/omni_swap"
	"github.com/coming-chat/intra-swap-core/routers/alpha_router/config"
	"github.com/daoleno/uniswap-sdk-core/entities"
	entitiesV3 "github.com/daoleno/uniswapv3-sdk/entities"
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
	PriceImpact                *entities.Percent
	Trade                      *base_entities.MTrade
	Route                      []RouteWithValidQuote
	BlockNumber                uint64
	MethodParameters           *MethodParameters
}

type MethodParameters struct {
	SoSoData                  omni_swap.ISoSoData
	LibSwapNormalizedSwapData []omni_swap.LibSwapNormalizedSwapData
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
	QuoteList      []*big.Int
	Route          *base_entities.MRoute
}
