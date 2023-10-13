package dex

import (
	"errors"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vaulverin/uniswapv2-sdk/router"
	"math/big"
	"time"
)

type MuteIo struct {
	UniswapV2
}

func NewMuteIo() MuteIo {
	return MuteIo{
		UniswapV2: NewUniswapV2(contracts.IMuteRouterAbi, contracts.MuteIoFatoryAbi, nil),
	}
}

func (m MuteIo) GetPool(tokenAAddr, tokenBAddr, factoryAddr common.Address, isStable bool, fee *big.Int) rpc.MultiCallSingle[common.Address] {
	param := m.UniswapV2.GetPool(tokenAAddr, tokenBAddr, factoryAddr, isStable, fee)
	param.FunctionParams = append(param.FunctionParams, isStable)
	return param
}

func (m MuteIo) GetQuote(route *base_entities.MRoute, index int, tradeType entities.TradeType, amount *entities.CurrencyAmount) (rpc.MultiCallSingle[QuoteResult], error) {
	if tradeType == entities.ExactOutput {
		return rpc.MultiCallSingle[QuoteResult]{}, errors.New("dex Mute.io not support getAmountsIn")
	}
	param, err := m.UniswapV2.GetQuote(route, index, tradeType, amount)
	if err != nil {
		return rpc.MultiCallSingle[QuoteResult]{}, err
	}
	param.FunctionParams = append(param.FunctionParams, []bool{route.Pools[index].Stable()})
	return param, nil
}

func (m MuteIo) PackSwap(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	amountIn, amountOut *entities.CurrencyAmount,
	pool base_entities.Pool,
	swapConfig base_entities.SwapOptions,
) (callData []byte, err error) {
	if tradeType == entities.ExactOutput {
		return nil, errors.New("dex Mute.io not support ExactOutput")
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
		args = []any{amountOut.Quotient(), path, swapConfig.Recipient, deadline, []bool{pool.Stable()}}
		methodName = "swapExactETHForTokens"
		// TODO handle the FeeOnTransfer
		//methodName = "swapExactETHForTokensSupportingFeeOnTransferTokens"
	} else if amountOut.Currency.IsNative() {
		// TODO handle the FeeOnTransfer
		args = []any{amountIn.Quotient(), amountOut.Quotient(), path, swapConfig.Recipient, deadline, []bool{pool.Stable()}}
		methodName = "swapExactTokensForETH"
		//methodName = "swapExactTokensForETHSupportingFeeOnTransferTokens"
	} else {
		args = []any{amountIn.Quotient(), amountOut.Quotient(), path, swapConfig.Recipient, deadline, []bool{pool.Stable()}}
		methodName = "swapExactTokensForTokens"
		// TODO handle the FeeOnTransfer
		//methodName = "swapExactTokensForTokensSupportingFeeOnTransferTokens"
	}
	return m.RouterContract.Pack(methodName, args...)
}
