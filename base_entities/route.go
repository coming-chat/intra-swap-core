package base_entities

import (
	"github.com/daoleno/uniswap-sdk-core/entities"
	entitiesV3 "github.com/daoleno/uniswapv3-sdk/entities"
	entitiesV2 "github.com/vaulverin/uniswapv2-sdk/entities"
)

type (
	V3Route = entitiesV3.Route
	V2Route = entitiesV2.Route
)

type Route interface {
	*V3Route | *V2Route
}

type Trade interface {
	InputAmount() *entities.CurrencyAmount
	OutputAmount() *entities.CurrencyAmount
}

type SwapAndAddParameters struct {
	// starting balance for tokenIn which will inform the tokenIn position amount
	InitialBalanceTokenIn *entities.CurrencyAmount
	// starting balance for tokenOut which will inform the tokenOut position amount
	InitialBalanceTokenOut *entities.CurrencyAmount
	// position details needed to create a new Position with the known liquidity amounts
	PreLiquidityPosition *entitiesV3.Position
}
