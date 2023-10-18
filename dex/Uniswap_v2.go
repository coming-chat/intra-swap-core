package dex

import (
	"errors"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gkirito/go-ethereum/accounts/abi"
	"github.com/vaulverin/uniswapv2-sdk/router"
	"math/big"
	"time"
)

type BaseAmounts struct {
	Amounts []*big.Int
}

func (b BaseAmounts) QuoteAmount() *big.Int {
	return b.Amounts[len(b.Amounts)-1]
}

func NewUniswapV2(routerAbi, factoryAbi, poolAbi *abi.ABI) UniswapV2 {
	if routerAbi == nil {
		routerAbi = contracts.IUniswapV2Router02Abi
	}
	if factoryAbi == nil {
		factoryAbi = contracts.IUniswapV2FactoryAbi
	}
	if poolAbi == nil {
		poolAbi = contracts.IUniswapV2PoolAbi
	}
	return UniswapV2{
		RouterContract:  routerAbi,
		FactoryContract: factoryAbi,
		PoolContract:    poolAbi,
	}
}

type UniswapV2 struct {
	RouterContract  *abi.ABI
	FactoryContract *abi.ABI
	PoolContract    *abi.ABI
}

func (b UniswapV2) GetQuote(route *base_entities.MRoute, index int, tradeType entities.TradeType, amount *entities.CurrencyAmount) (rpc.MultiCallSingle[QuoteResult], error) {
	name := "getAmountsOut"
	if tradeType == entities.ExactOutput {
		name = "getAmountsIn"
	}
	param := rpc.MultiCallSingle[QuoteResult]{
		FunctionParams: []any{
			amount.Quotient(),
			[]common.Address{
				route.Path[index].Address,
				route.Path[index+1].Address,
			},
		},
		ContractAddress: route.Pools[index].QuoterAddress(),
		FunctionName:    name,
		CallResult: rpc.MultiCallResult[QuoteResult]{
			Data: &BaseAmounts{},
		},
		Contract: b.RouterContract,
	}
	return param, nil
}

func (b UniswapV2) GetPool(tokenAAddr, tokenBAddr, factoryAddr common.Address, isStable bool, fee *big.Int) rpc.MultiCallSingle[common.Address] {
	return rpc.MultiCallSingle[common.Address]{
		Contract:     b.FactoryContract,
		FunctionName: "getPair",
		FunctionParams: []any{
			tokenAAddr, tokenBAddr,
		},
		ContractAddress: factoryAddr,
	}
}

func (b UniswapV2) GetPoolInfo() PoolInfo {
	return UniswapV2PoolInfo{
		PoolContract: b.PoolContract,
	}
}

type UniswapV2PoolInfo struct {
	PoolContract *abi.ABI
}

func (b UniswapV2PoolInfo) GetSlot0(poolAddr common.Address) rpc.MultiCallSingle[IPoolState] {
	panic("v2 pool not support v3 info")
}

func (b UniswapV2PoolInfo) GetLiquidity(poolAddr common.Address) rpc.MultiCallSingle[ILiquidity] {
	panic("v2 pool not support v3 info")
}

func (b UniswapV2PoolInfo) GetTicks(poolAddr common.Address, tick *big.Int) rpc.MultiCallSingle[PoolTick] {
	panic("v2 pool not support v3 info")
}

func (b UniswapV2PoolInfo) GetReserves(poolAddr common.Address) rpc.MultiCallSingle[IReserves] {
	return rpc.MultiCallSingle[IReserves]{
		FunctionName:    "getReserves",
		Contract:        b.PoolContract,
		ContractAddress: poolAddr,
	}
}

func (b UniswapV2) PackSwap(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	amountIn, amountOut *entities.CurrencyAmount,
	pool base_entities.Pool,
	swapConfig base_entities.SwapOptions,
) (callData []byte, err error) {
	if tokenIn.IsNative() && tokenOut.IsNative() {
		return nil, router.ErrEtherInOut
	}
	path := []common.Address{tokenIn.Address, tokenOut.Address}

	deadline := swapConfig.Deadline
	if swapConfig.Deadline == nil {
		deadline = big.NewInt(time.Now().Add(5 * time.Minute).Unix())
	}
	var (
		methodName string
		args       []any
	)

	switch tradeType {
	case entities.ExactInput:
		if amountIn.Currency.IsNative() {
			args = []any{amountOut.Quotient(), path, swapConfig.Recipient, deadline}
			methodName = "swapExactETHForTokens"
			// TODO handle the FeeOnTransfer
			//methodName = "swapExactETHForTokensSupportingFeeOnTransferTokens"
		} else if amountOut.Currency.IsNative() {
			// TODO handle the FeeOnTransfer
			args = []any{amountIn.Quotient(), amountOut.Quotient(), path, swapConfig.Recipient, deadline}
			methodName = "swapExactTokensForETH"
			//methodName = "swapExactTokensForETHSupportingFeeOnTransferTokens"
		} else {
			args = []any{amountIn.Quotient(), amountOut.Quotient(), path, swapConfig.Recipient, deadline}
			methodName = "swapExactTokensForTokens"
			// TODO handle the FeeOnTransfer
			//methodName = "swapExactTokensForTokensSupportingFeeOnTransferTokens"
		}
	case entities.ExactOutput:
		if amountIn.Currency.IsNative() {
			methodName = "swapETHForExactTokens"
			args = []any{amountOut.Quotient(), path, swapConfig.Recipient, deadline}
		} else if amountOut.Currency.IsNative() {
			methodName = "swapTokensForExactETH"
			args = []any{amountOut.Quotient(), amountIn.Quotient(), path, swapConfig.Recipient, deadline}
		} else {
			methodName = "swapTokensForExactTokens"
			args = []any{amountOut.Quotient(), amountIn.Quotient(), path, swapConfig.Recipient, deadline}
		}
	default:
		return nil, errors.New("unsupported tradeType")
	}
	return b.RouterContract.Pack(methodName, args...)
}
