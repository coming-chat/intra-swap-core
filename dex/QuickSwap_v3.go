package dex

import (
	"errors"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/coming-chat/intra-swap-core/contracts/omni_swap"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/coming-chat/intra-swap-core/util"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
	selfCommon "github.com/gkirito/go-ethereum/common"
	"math/big"
)

type QuoteDataWithFees struct {
	AmountOut *big.Int
	Fees      []uint16
}

func (q QuoteDataWithFees) QuoteAmount() *big.Int {
	return q.AmountOut
}

type QuickSwapV3 struct {
	UniswapV3
}

func NewQuickSwapV3() QuickSwapV3 {
	return QuickSwapV3{
		UniswapV3: NewUniswapV3(contracts.IQuickSwapRouterAbi, contracts.IQuickQuoterAbi, contracts.QuickSwapFactoryAbi, contracts.IQuickSwapPoolAbi),
	}
}

func (q QuickSwapV3) GetPool(tokenAAddr, tokenBAddr, factoryAddr common.Address, isStable bool, fee *big.Int) rpc.MultiCallSingle[common.Address] {
	return rpc.MultiCallSingle[common.Address]{
		Contract:     q.FactoryContract,
		FunctionName: "poolByPair",
		FunctionParams: []any{
			tokenAAddr, tokenBAddr,
		},
		ContractAddress: factoryAddr,
	}
}

func (q QuickSwapV3) GetQuote(route *base_entities.MRoute, index int, tradeType entities.TradeType, amount *entities.CurrencyAmount) (rpc.MultiCallSingle[QuoteResult], error) {
	param, err := q.UniswapV3.GetQuote(route, index, tradeType, amount)
	if err != nil {
		return rpc.MultiCallSingle[QuoteResult]{}, err
	}
	param.CallResult.Data = &QuoteDataWithFees{}
	return param, nil
}

func (q QuickSwapV3) PackSwap(tradeType entities.TradeType, tokenIn, tokenOut *entities.Token, amountIn, amountOut *entities.CurrencyAmount, pool base_entities.Pool, swapConfig base_entities.SwapOptions) (callData []byte, err error) {
	path, err := util.EncodeRouteToPath([]base_entities.Pool{pool}, tokenIn, tradeType == entities.ExactOutput)
	if err != nil {
		return nil, err
	}
	switch tradeType {
	case entities.ExactInput:
		param := omni_swap.IQuickSwapRouterExactInputParams{
			Path:             path,
			Recipient:        selfCommon.Address(swapConfig.Recipient),
			AmountIn:         amountIn.Quotient(),
			AmountOutMinimum: amountOut.Quotient(),
			Deadline:         swapConfig.Deadline,
		}
		callData, err = q.RouterContract.Pack("exactInput", param)
	case entities.ExactOutput:
		param := omni_swap.IQuickSwapRouterExactOutputParams{
			Path:            path,
			Recipient:       selfCommon.Address(swapConfig.Recipient),
			AmountInMaximum: amountIn.Quotient(),
			AmountOut:       amountOut.Quotient(),
			Deadline:        swapConfig.Deadline,
		}
		callData, err = q.RouterContract.Pack("exactOutput", param)
	default:
		return nil, errors.New("unsupported tradeType")
	}
	return
}

type QuickSwapGlobalState struct {
	Price              *big.Int
	Tick               *big.Int
	Fee                uint16
	TimepointIndex     uint16
	CommunityFeeToken0 uint8
	CommunityFeeToken1 uint8
	Unlocked           bool
}

func (q QuickSwapGlobalState) GetPrice() *big.Int {
	return q.Price
}

func (q QuickSwapGlobalState) GetTick() *big.Int {
	return q.Tick
}

type QuickSwapTicks struct {
	LiquidityTotal           *big.Int
	LiquidityDelta           *big.Int
	OuterFeeGrowth0Token     *big.Int
	OuterFeeGrowth1Token     *big.Int
	OuterTickCumulative      *big.Int
	OuterSecondsPerLiquidity *big.Int
	OuterSecondsSpent        uint32
	Initialized              bool
}

func (q QuickSwapTicks) GetLiquidityNet() *big.Int {
	return q.LiquidityDelta
}

func (q QuickSwapTicks) GetLiquidityGross() *big.Int {
	return q.LiquidityTotal
}

type QuickSwapPoolInfo struct {
	UniswapV3PoolInfo
}

func (q QuickSwapPoolInfo) GetSlot0(poolAddr common.Address) rpc.MultiCallSingle[IPoolState] {
	param := q.UniswapV3PoolInfo.GetSlot0(poolAddr)
	param.FunctionName = "globalState"
	param.CallResult.Data = &QuickSwapGlobalState{}
	return param
}

func (q QuickSwapPoolInfo) GetTicks(poolAddr common.Address, tick *big.Int) rpc.MultiCallSingle[PoolTick] {
	param := q.UniswapV3PoolInfo.GetTicks(poolAddr, tick)
	param.CallResult.Data = &QuickSwapTicks{}
	return param
}

func (q QuickSwapV3) GetPoolInfo() PoolInfo {
	return QuickSwapPoolInfo{
		UniswapV3PoolInfo{
			PoolContract: q.PoolContract,
		},
	}
}
