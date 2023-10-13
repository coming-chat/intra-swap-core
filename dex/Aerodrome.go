package dex

import (
	"errors"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/coming-chat/intra-swap-core/contracts/omni_swap"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vaulverin/uniswapv2-sdk/router"
	"math/big"
	"time"
)

type Aerodrome struct {
	UniswapV2
}

func NewAerodrome() Aerodrome {
	return Aerodrome{UniswapV2: NewUniswapV2(contracts.IAerodromeAbi, contracts.AerodromeFactoryAbi, nil)}
}

func (a Aerodrome) GetPool(tokenAAddr, tokenBAddr, factoryAddr common.Address, isStable bool, fee *big.Int) rpc.MultiCallSingle[common.Address] {
	multiParam := a.UniswapV2.GetPool(tokenAAddr, tokenBAddr, factoryAddr, isStable, fee)
	multiParam.FunctionName = "getPool"
	multiParam.FunctionParams = append(multiParam.FunctionParams, isStable)
	return multiParam
}

func (a Aerodrome) GetQuote(route *base_entities.MRoute, index int, tradeType entities.TradeType, amount *entities.CurrencyAmount) (rpc.MultiCallSingle[QuoteResult], error) {
	if tradeType == entities.ExactOutput {
		return rpc.MultiCallSingle[QuoteResult]{}, errors.New("dex Aerodrome not support getAmountsIn")
	}
	multiParaQuote, err := a.UniswapV2.GetQuote(route, index, tradeType, amount)
	if err != nil {
		return rpc.MultiCallSingle[QuoteResult]{}, err
	}
	multiParaQuote.FunctionName = "getAmountsOut"
	multiParaQuote.FunctionParams[1] = []omni_swap.IAerodromeRoute{
		{
			From:    route.Path[index].Address,
			To:      route.Path[index+1].Address,
			Stable:  route.Pools[index].Stable(),
			Factory: route.Pools[index].FactoryAddress(),
		},
	}
	return multiParaQuote, nil
}

func (a Aerodrome) PackSwap(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	amountIn, amountOut *entities.CurrencyAmount,
	pool base_entities.Pool,
	swapConfig base_entities.SwapOptions,
) (callData []byte, err error) {
	if tradeType == entities.ExactOutput {
		return nil, errors.New("dex Aerodrome not support for ExactOutput")
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
	path := []omni_swap.IAerodromeRoute{
		{
			From:    tokenIn.Address,
			To:      tokenOut.Address,
			Stable:  pool.Stable(),
			Factory: pool.FactoryAddress(),
		},
	}
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
	callData, err = a.UniswapV2.RouterContract.Pack(methodName, args...)
	return
}
