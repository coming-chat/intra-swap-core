package dex

import (
	"errors"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vaulverin/uniswapv2-sdk/router"
	"math/big"
	"time"
)

type Camelot struct {
	UniswapV2
}

func NewCamelot() Camelot {
	return Camelot{
		UniswapV2: NewUniswapV2(contracts.CamelotRouterAbi, nil, nil),
	}
}

func (c Camelot) GetQuote(route *base_entities.MRoute, index int, tradeType entities.TradeType, amount *entities.CurrencyAmount) (rpc.MultiCallSingle[QuoteResult], error) {
	if tradeType == entities.ExactOutput {
		return rpc.MultiCallSingle[QuoteResult]{}, errors.New("dex Camelot not support getAmountsIn")
	}
	return c.UniswapV2.GetQuote(route, index, tradeType, amount)
}

func (c Camelot) PackSwap(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	amountIn, amountOut *entities.CurrencyAmount,
	pool base_entities.Pool,
	swapConfig base_entities.SwapOptions,
) (callData []byte, err error) {
	if tradeType == entities.ExactOutput {
		return nil, errors.New("dex Camelot not support ExactOutput")
	}
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

	if amountIn.Currency.IsNative() {
		args = []any{amountOut.Quotient(), path, swapConfig.Recipient, constants.AddressZero, deadline}
		methodName = "swapExactETHForTokensSupportingFeeOnTransferTokens"
	} else if amountOut.Currency.IsNative() {
		args = []any{amountIn.Quotient(), amountOut.Quotient(), path, swapConfig.Recipient, constants.AddressZero, deadline}
		methodName = "swapExactTokensForETHSupportingFeeOnTransferTokens"
	} else {
		args = []any{amountIn.Quotient(), amountOut.Quotient(), path, swapConfig.Recipient, constants.AddressZero, deadline}
		methodName = "swapExactTokensForTokensSupportingFeeOnTransferTokens"
	}
	return c.RouterContract.Pack(methodName, args...)
}
