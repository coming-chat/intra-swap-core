package dex

import (
	"github.com/coming-chat/intra-swap-core/base_constant"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type QuoteResult interface {
	QuoteAmount() *big.Int
}

type PackMethod func(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	amountIn, amountOut *entities.CurrencyAmount,
	pool base_entities.Pool,
	swapConfig base_entities.SwapOptions,
) (callData []byte, err error)

type ISlot0 struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}

func (I ISlot0) GetPrice() *big.Int {
	return I.SqrtPriceX96
}

func (I ISlot0) GetTick() *big.Int {
	return I.Tick
}

type IPoolState interface {
	GetPrice() *big.Int
	GetTick() *big.Int
}

type PoolTick interface {
	GetLiquidityNet() *big.Int
	GetLiquidityGross() *big.Int
}

type Tick struct {
	LiquidityGross                 *big.Int
	LiquidityNet                   *big.Int
	FeeGrowthOutside0X128          *big.Int
	FeeGrowthOutside1X128          *big.Int
	TickCumulativeOutside          *big.Int
	SecondsPerLiquidityOutsideX128 *big.Int
	SecondsOutside                 uint32
	Initialized                    bool
}

func (t Tick) GetLiquidityNet() *big.Int {
	return t.LiquidityNet
}

func (t Tick) GetLiquidityGross() *big.Int {
	return t.LiquidityGross
}

type IReserves struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}

type Dex interface {
	GetQuote(
		route *base_entities.MRoute,
		index int,
		tradeType entities.TradeType,
		amount *entities.CurrencyAmount,
	) (rpc.MultiCallSingle[QuoteResult], error)
	GetPool(tokenAAddr, tokenBAddr, factoryAddr common.Address, isStable bool, fee *big.Int) rpc.MultiCallSingle[common.Address]
	GetPoolInfo() PoolInfo
	PackSwap(tradeType entities.TradeType,
		tokenIn, tokenOut *entities.Token,
		amountIn, amountOut *entities.CurrencyAmount,
		pool base_entities.Pool,
		swapConfig base_entities.SwapOptions,
	) (callData []byte, err error)
}

type PoolInfo interface {
	GetSlot0(poolAddr common.Address) rpc.MultiCallSingle[IPoolState]
	GetLiquidity(poolAddr common.Address) rpc.MultiCallSingle[*big.Int]
	GetTicks(poolAddr common.Address, tick *big.Int) rpc.MultiCallSingle[PoolTick]
	GetReserves(poolAddr common.Address) rpc.MultiCallSingle[IReserves]
}

var RouterAddrDexMap = map[common.Address]Dex{
	base_constant.BaseAlienbaseV2Router: NewAlienbase(),
	base_constant.BaseAerodromeRouter:   NewAerodrome(),
	base_constant.BaseUniswapV3Router:   NewUniswapV3(nil, nil, nil, nil),
	base_constant.BaseSwapBasedV2Router: NewSwapBasedV2(),
	base_constant.BaseBaseswapV2Router:  NewBaseSwapV2(),

	base_constant.OptimismVelodromeV2Router: NewVelodrome(),
	base_constant.OptimismUniswapV3Router:   NewUniswapV3(nil, nil, nil, nil),

	base_constant.ArbitrumUniswapV3Router: NewUniswapV3(nil, nil, nil, nil),
	base_constant.ArbitrumSushiRouter:     NewSushiSwapV2(),

	base_constant.ZkSyncMuteRouter:     NewMuteIo(),
	base_constant.ZkSyncMaverickRouter: NewMaverickProtocol(),

	base_constant.PolygonQuickSwapV2Router: NewQuickSwapV2(),
	base_constant.PolygonUniswapV3Router:   NewUniswapV3(nil, nil, nil, nil),

	base_constant.PolygonPearlFiRouter: NewPearLFi(),

	base_constant.ArbitrumCamelotRouter: NewCamelot(),

	base_constant.BaseSwapBasedV3Router:    NewSwapBasedV3(),
	base_constant.PolygonQuickSwapV3Router: NewQuickSwapV3(),
}
