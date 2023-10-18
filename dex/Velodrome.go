package dex

import (
	"errors"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/coming-chat/intra-swap-core/contracts/omni_swap"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/gkirito/go-ethereum/common"
	"github.com/vaulverin/uniswapv2-sdk/router"
	"math/big"
	"time"
)

type Velodrome struct {
	Aerodrome
}

func NewVelodrome() Velodrome {
	return Velodrome{
		Aerodrome{
			NewUniswapV2(contracts.IVelodromeAbi, contracts.AerodromeFactoryAbi, nil),
		},
	}
}

func (v Velodrome) GetQuote(route *base_entities.MRoute, index int, tradeType entities.TradeType, amount *entities.CurrencyAmount) (rpc.MultiCallSingle[QuoteResult], error) {
	if tradeType == entities.ExactOutput {
		return rpc.MultiCallSingle[QuoteResult]{}, errors.New("dex Velodrome not support getAmountsIn")
	}
	multiParaQuote, err := v.Aerodrome.GetQuote(route, index, tradeType, amount)
	if err != nil {
		return rpc.MultiCallSingle[QuoteResult]{}, err
	}
	multiParaQuote.FunctionParams[1] = []omni_swap.IVelodromeRoute{
		{
			From:    common.Address(route.Path[index].Address),
			To:      common.Address(route.Path[index+1].Address),
			Stable:  route.Pools[index].Stable(),
			Factory: common.Address(route.Pools[index].FactoryAddress()),
		},
	}
	return multiParaQuote, nil
}

func (v Velodrome) PackSwap(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	amountIn, amountOut *entities.CurrencyAmount,
	pool base_entities.Pool,
	swapConfig base_entities.SwapOptions,
) (callData []byte, err error) {
	if tradeType == entities.ExactOutput {
		return nil, errors.New("dex Velodrome not support ExactOutput")
	}
	if tokenIn.IsNative() && tokenOut.IsNative() {
		return nil, router.ErrEtherInOut
	}

	deadline := swapConfig.Deadline
	if swapConfig.Deadline == nil {
		deadline = big.NewInt(time.Now().Add(5 * time.Minute).Unix())
	}
	var (
		methodName string
		args       []any
	)
	path := []omni_swap.IVelodromeRoute{
		{
			From:    common.Address(tokenIn.Address),
			To:      common.Address(tokenOut.Address),
			Stable:  pool.Stable(),
			Factory: common.Address(pool.FactoryAddress()),
		},
	}

	if amountIn.Currency.IsNative() {
		args = []any{amountOut.Quotient(), path, swapConfig.Recipient, deadline}
		methodName = "swapExactETHForTokens"
		if pool.Stable() {
			methodName = "swapExactETHForTokensSupportingFeeOnTransferTokens"
		}
	} else if amountOut.Currency.IsNative() {
		args = []any{amountIn.Quotient(), amountOut.Quotient(), path, swapConfig.Recipient, deadline}
		methodName = "swapExactTokensForETH"
		if pool.Stable() {
			methodName = "swapExactTokensForETHSupportingFeeOnTransferTokens"
		}
	} else {
		args = []any{amountIn.Quotient(), amountOut.Quotient(), path, swapConfig.Recipient, deadline}
		methodName = "swapExactTokensForTokens"
		if pool.Stable() {
			methodName = "swapExactTokensForTokensSupportingFeeOnTransferTokens"
		}
	}
	callData, err = v.UniswapV2.RouterContract.Pack(methodName, args...)
	return
}
