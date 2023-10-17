package dex

import (
	"errors"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/coming-chat/intra-swap-core/contracts/omni_swap"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/coming-chat/intra-swap-core/util"
	"github.com/daoleno/uniswap-sdk-core/entities"
	entitiesV3 "github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/periphery"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gkirito/go-ethereum/accounts/abi"
	selfCommon "github.com/gkirito/go-ethereum/common"
	"math/big"
)

type BaseUniswapV3QuoteData struct {
	AmountOut                   *big.Int
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	GasEstimate                 *big.Int
}

func (b BaseUniswapV3QuoteData) QuoteAmount() *big.Int {
	return b.AmountOut
}

func NewUniswapV3(routerAbi, quoterAbi, factoryAbi, poolAbi *abi.ABI) UniswapV3 {
	if routerAbi == nil {
		routerAbi = contracts.ISwapRouter02Abi
	}
	if quoterAbi == nil {
		quoterAbi = contracts.IQuoterV2Abi
	}
	if factoryAbi == nil {
		factoryAbi = contracts.IUniswapV3FactoryAbi
	}
	if poolAbi == nil {
		poolAbi = contracts.IUniswapV3PoolAbi
	}
	return UniswapV3{
		RouterContract:  routerAbi,
		QuoterAbi:       quoterAbi,
		FactoryContract: factoryAbi,
		PoolContract:    poolAbi,
	}
}

type UniswapV3 struct {
	RouterContract  *abi.ABI
	QuoterAbi       *abi.ABI
	FactoryContract *abi.ABI
	PoolContract    *abi.ABI
}

func (u UniswapV3) GetQuote(route *base_entities.MRoute, index int, tradeType entities.TradeType, amount *entities.CurrencyAmount) (rpc.MultiCallSingle[QuoteResult], error) {
	stepRoute, err := entitiesV3.NewRoute([]*entitiesV3.Pool{route.Pools[index].(*base_entities.V3Pool).Pool}, route.Path[index], route.Path[index+1])
	if err != nil {
		return rpc.MultiCallSingle[QuoteResult]{}, nil
	}
	encodedRoute, err := periphery.EncodeRouteToPath(stepRoute, tradeType == entities.ExactOutput)
	if err != nil {
		return rpc.MultiCallSingle[QuoteResult]{}, nil
	}
	param := rpc.MultiCallSingle[QuoteResult]{
		FunctionParams: []any{
			encodedRoute,
			amount.Quotient(),
		},
		ContractAddress: route.Pools[index].QuoterAddress(),
		Contract:        u.QuoterAbi,
		CallResult: rpc.MultiCallResult[QuoteResult]{
			Data: &BaseUniswapV3QuoteData{},
		},
	}
	param.FunctionName = "quoteExactInput"
	if tradeType == entities.ExactOutput {
		param.FunctionName = "quoteExactOutput"
	}
	return param, nil
}

func (u UniswapV3) GetPool(tokenAAddr, tokenBAddr, factoryAddr common.Address, isStable bool, fee *big.Int) rpc.MultiCallSingle[common.Address] {
	return rpc.MultiCallSingle[common.Address]{
		Contract:     u.FactoryContract,
		FunctionName: "getPool",
		FunctionParams: []any{
			tokenAAddr, tokenBAddr, fee,
		},
		ContractAddress: factoryAddr,
	}
}

func (u UniswapV3) GetPoolInfo() PoolInfo {
	return UniswapV3PoolInfo{
		PoolContract: u.PoolContract,
	}
}

type UniswapV3PoolInfo struct {
	PoolContract *abi.ABI
}

func (u UniswapV3PoolInfo) GetSlot0(poolAddr common.Address) rpc.MultiCallSingle[IPoolState] {
	return rpc.MultiCallSingle[IPoolState]{
		FunctionName:    "slot0",
		Contract:        u.PoolContract,
		ContractAddress: poolAddr,
		CallResult: rpc.MultiCallResult[IPoolState]{
			Data: &ISlot0{},
		},
	}
}

func (u UniswapV3PoolInfo) GetLiquidity(poolAddr common.Address) rpc.MultiCallSingle[*big.Int] {
	return rpc.MultiCallSingle[*big.Int]{
		FunctionName:    "liquidity",
		ContractAddress: poolAddr,
		Contract:        u.PoolContract,
	}
}

func (u UniswapV3PoolInfo) GetTicks(poolAddr common.Address, tick *big.Int) rpc.MultiCallSingle[PoolTick] {
	return rpc.MultiCallSingle[PoolTick]{
		FunctionName:    "ticks",
		ContractAddress: poolAddr,
		Contract:        u.PoolContract,
		FunctionParams: []any{
			tick,
		},
		CallResult: rpc.MultiCallResult[PoolTick]{
			Data: &Tick{},
		},
	}
}

func (u UniswapV3PoolInfo) GetReserves(poolAddr common.Address) rpc.MultiCallSingle[IReserves] {
	panic("v3 pool not support v2 info")
}

func (u UniswapV3) PackSwap(tradeType entities.TradeType, tokenIn, tokenOut *entities.Token, amountIn, amountOut *entities.CurrencyAmount, pool base_entities.Pool, swapConfig base_entities.SwapOptions) (callData []byte, err error) {
	path, err := util.EncodeRouteToPath([]base_entities.Pool{pool}, tokenIn, tradeType == entities.ExactOutput)
	if err != nil {
		return nil, err
	}
	switch tradeType {
	case entities.ExactInput:
		param := omni_swap.ISwapRouter02ExactInputParams{
			Path:             path,
			Recipient:        selfCommon.Address(swapConfig.Recipient),
			AmountIn:         amountIn.Quotient(),
			AmountOutMinimum: amountOut.Quotient(),
		}
		callData, err = u.RouterContract.Pack("exactInput", param)
	case entities.ExactOutput:
		param := omni_swap.ISwapRouter02ExactOutputParams{
			Path:            path,
			Recipient:       selfCommon.Address(swapConfig.Recipient),
			AmountInMaximum: amountIn.Quotient(),
			AmountOut:       amountOut.Quotient(),
		}
		callData, err = u.RouterContract.Pack("exactOutput", param)
	default:
		return nil, errors.New("unsupported tradeType")
	}
	return
}
