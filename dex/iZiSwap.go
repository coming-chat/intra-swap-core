package dex

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/coming-chat/intra-swap-core/contracts/omni_swap"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
	selfCommon "github.com/gkirito/go-ethereum/common"
	"math/big"
)

type IiZiSwapQuote struct {
	Cost    *big.Int
	Acquire *big.Int
}

func (i IiZiSwapQuote) QuoteAmount() *big.Int {
	return i.Acquire
}

type IiZiSwapState struct {
	SqrtPrice96             *big.Int
	CurrentPoint            *big.Int
	ObservationCurrentIndex uint16
	ObservationQueueLen     uint16
	ObservationNextQueueLen uint16
	Locked                  bool
	Liquidity               *big.Int
	LiquidityX              *big.Int
}

func (i IiZiSwapState) GetPrice() *big.Int {
	return i.SqrtPrice96
}

func (i IiZiSwapState) GetTick() *big.Int {
	return i.CurrentPoint
}

func (i IiZiSwapState) GetLiquidity() *big.Int {
	return i.Liquidity
}

type IZiSwap struct {
	UniswapV3
}

func NewIZiSwap() IZiSwap {
	return IZiSwap{
		UniswapV3: NewUniswapV3(contracts.IiZiSwapAbi, contracts.IiZiSwapAbi, contracts.IiZiSwapFactoryAbi, contracts.IiZiSwapPoolAbi),
	}
}

func (i IZiSwap) GetQuote(route *base_entities.MRoute, index int, tradeType entities.TradeType, amount *entities.CurrencyAmount) (rpc.MultiCallSingle[QuoteResult], error) {
	var path []byte
	path = append(path, route.Path[index].Address.Bytes()...)
	path = append(path, route.Path[index+1].Address.Bytes()...)
	param := rpc.MultiCallSingle[QuoteResult]{
		Contract: i.QuoterAbi,
		CallResult: rpc.MultiCallResult[QuoteResult]{
			Data: &IiZiSwapQuote{},
		},
	}
	switch tradeType {
	case entities.ExactInput:
		param.FunctionName = "swapAmount"
		param.FunctionParams = []any{
			omni_swap.IiZiSwapSwapAmountParams{
				Path:        path,
				Recipient:   selfCommon.Address(common.Address{}),
				Amount:      amount.Quotient(),
				MinAcquired: big.NewInt(0),
				Deadline:    big.NewInt(10000000),
			},
		}
	case entities.ExactOutput:
		param.FunctionName = "swapDesire"
		param.FunctionParams = []any{
			omni_swap.IiZiSwapSwapDesireParams{
				Path:      path,
				Recipient: selfCommon.Address(common.Address{}),
				Desire:    amount.Quotient(),
				MaxPayed:  amount.Quotient(),
				Deadline:  big.NewInt(10000000),
			},
		}
	}
	return param, nil
}

func (i IZiSwap) GetPool(tokenAAddr, tokenBAddr, factoryAddr common.Address, isStable bool, fee *big.Int) rpc.MultiCallSingle[common.Address] {
	return rpc.MultiCallSingle[common.Address]{
		Contract:     i.FactoryContract,
		FunctionName: "pool",
		FunctionParams: []any{
			tokenAAddr, tokenBAddr, fee,
		},
		ContractAddress: factoryAddr,
	}
}

func (i IZiSwap) PackSwap(tradeType entities.TradeType, tokenIn, tokenOut *entities.Token, amountIn, amountOut *entities.CurrencyAmount, pool base_entities.Pool, swapConfig base_entities.SwapOptions) (callData []byte, err error) {
	var path []byte
	path = append(path, tokenIn.Address.Bytes()...)
	path = append(path, tokenOut.Address.Bytes()...)
	var (
		method string
		args   []any
	)
	switch tradeType {
	case entities.ExactInput:
		method = "swapAmount"
		args = []any{
			omni_swap.IiZiSwapSwapAmountParams{
				Path:        path,
				Recipient:   selfCommon.Address(swapConfig.Recipient),
				Amount:      amountIn.Quotient(),
				MinAcquired: amountOut.Quotient(),
				Deadline:    swapConfig.Deadline,
			},
		}
	case entities.ExactOutput:
		method = "swapDesire"
		args = []any{
			omni_swap.IiZiSwapSwapDesireParams{
				Path:      path,
				Recipient: selfCommon.Address(swapConfig.Recipient),
				Desire:    amountIn.Quotient(),
				MaxPayed:  amountOut.Quotient(),
				Deadline:  swapConfig.Deadline,
			},
		}
	}

	return i.RouterContract.Pack(method, args...)
}

func (i IZiSwap) GetPoolInfo() PoolInfo {
	return IZiSwapPoolInfo{
		UniswapV3PoolInfo{PoolContract: i.PoolContract},
	}
}

type IZiSwapPoolInfo struct {
	UniswapV3PoolInfo
}

func (i IZiSwapPoolInfo) GetSlot0(poolAddr common.Address) rpc.MultiCallSingle[IPoolState] {
	return rpc.MultiCallSingle[IPoolState]{
		FunctionName:    "state",
		Contract:        i.PoolContract,
		ContractAddress: poolAddr,
		CallResult: rpc.MultiCallResult[IPoolState]{
			Data: &IiZiSwapState{},
		},
	}
}

func (i IZiSwapPoolInfo) GetLiquidity(poolAddr common.Address) rpc.MultiCallSingle[ILiquidity] {
	return rpc.MultiCallSingle[ILiquidity]{
		FunctionName:    "state",
		ContractAddress: poolAddr,
		Contract:        i.PoolContract,
		CallResult: rpc.MultiCallResult[ILiquidity]{
			Data: &IiZiSwapState{},
		},
	}
}

func (i IZiSwapPoolInfo) GetTicks(poolAddr common.Address, tick *big.Int) rpc.MultiCallSingle[PoolTick] {
	panic("iziswap not found ticks")
}
