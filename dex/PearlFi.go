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

type PearlFi struct {
	MuteIo
}

func NewPearLFi() PearlFi {
	return PearlFi{
		MuteIo{
			UniswapV2: NewUniswapV2(contracts.IPearlRouterAbi, contracts.MuteIoFatoryAbi, nil),
		},
	}
}

func (f PearlFi) GetQuote(route *base_entities.MRoute, index int, tradeType entities.TradeType, amount *entities.CurrencyAmount) (rpc.MultiCallSingle[QuoteResult], error) {
	if tradeType == entities.ExactOutput {
		return rpc.MultiCallSingle[QuoteResult]{}, errors.New("dex PearlFi not support getAmountsIn")
	}
	param, err := f.UniswapV2.GetQuote(route, index, tradeType, amount)
	if err != nil {
		return rpc.MultiCallSingle[QuoteResult]{}, err
	}
	param.FunctionParams[1] = []omni_swap.IPearlRouterroute{
		{
			From:   common.Address(route.Path[index].Address),
			To:     common.Address(route.Path[index+1].Address),
			Stable: route.Pools[index].Stable(),
		},
	}
	return param, nil
}

func (f PearlFi) PackSwap(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	amountIn, amountOut *entities.CurrencyAmount,
	pool base_entities.Pool,
	swapConfig base_entities.SwapOptions,
) (callData []byte, err error) {
	if tradeType == entities.ExactOutput {
		return nil, errors.New("dex PearlFi not support ExactOutput")
	}
	if tokenIn.IsNative() && tokenOut.IsNative() {
		return nil, router.ErrEtherInOut
	}
	route := []omni_swap.IPearlRouterroute{
		{
			From:   common.Address(tokenIn.Address),
			To:     common.Address(tokenOut.Address),
			Stable: pool.Stable(),
		},
	}

	deadline := swapConfig.Deadline
	if swapConfig.Deadline == nil {
		deadline = big.NewInt(time.Now().Add(5 * time.Minute).Unix())
	}
	var (
		methodName string
		args       []any
	)

	if amountIn.Currency.IsNative() {
		args = []any{amountOut.Quotient(), route, swapConfig.Recipient, deadline}
		methodName = "swapExactETHForTokens"
	} else if amountOut.Currency.IsNative() {
		args = []any{amountIn.Quotient(), amountOut.Quotient(), route, swapConfig.Recipient, deadline}
		methodName = "swapExactTokensForETH"
	} else {
		args = []any{amountIn.Quotient(), amountOut.Quotient(), route, swapConfig.Recipient, deadline}
		methodName = "swapExactTokensForTokens"
	}
	return f.RouterContract.Pack(methodName, args...)
}
