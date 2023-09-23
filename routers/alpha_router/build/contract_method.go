package build

import (
	"errors"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/coming-chat/intra-swap-core/contracts/omni_swap"
	"github.com/coming-chat/intra-swap-core/routers/alpha_router/config"
	"github.com/coming-chat/intra-swap-core/util"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vaulverin/uniswapv2-sdk/router"
	"math/big"
	"time"
)

type contractMethod func(trade *base_entities.MTrade, swapIndex int, amountIn, amountOut *entities.CurrencyAmount, swapConfig config.SwapOptions) (callData []byte, err error)

func iSwapRouter02(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	pool *base_entities.V3Pool,
	amountIn, amountOut *entities.CurrencyAmount,
	swapConfig config.SwapOptions,
) (callData []byte, err error) {
	// encode permit if necessary
	//if swapConfig.InputTokenPermit != nil {
	//	if amountIn.Currency.IsToken() {
	//		return nil, periphery.ErrNonTokenPermit
	//	}
	//
	//	permit, err := periphery.EncodePermit(tokenIn, swapConfig.InputTokenPermit)
	//	if err != nil {
	//		return nil, err
	//	}
	//	calldatas = append(calldatas, permit)
	//}
	path, err := util.EncodeRouteToPath([]base_entities.Pool{pool}, tokenIn, tradeType == entities.ExactOutput)
	if err != nil {
		return nil, err
	}
	switch tradeType {
	case entities.ExactInput:
		callData, err = contracts.ISwapRouter02Abi.Pack("exactInput", omni_swap.ISwapRouter02ExactInputParams{
			Path:             path,
			Recipient:        swapConfig.Recipient,
			AmountIn:         amountIn.Quotient(),
			AmountOutMinimum: amountOut.Quotient(),
		})
	case entities.ExactOutput:
		callData, err = contracts.ISwapRouter02Abi.Pack("exactOutput", omni_swap.ISwapRouter02ExactOutputParams{
			Path:            path,
			Recipient:       swapConfig.Recipient,
			AmountInMaximum: amountIn.Quotient(),
			AmountOut:       amountOut.Quotient(),
		})
	default:
		return nil, errors.New("unsupported tradeType")
	}
	return nil, err
}

func iUniswapV2Router02(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	amountIn, amountOut *entities.CurrencyAmount,
	swapConfig config.SwapOptions,
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
		if tokenIn.IsNative() {
			args = []any{amountOut, path, swapConfig.Recipient, deadline}
			methodName = "swapExactETHForTokens"
			// TODO handle the FeeOnTransfer
			//methodName = "swapExactETHForTokensSupportingFeeOnTransferTokens"
		} else if tokenOut.IsNative() {
			// TODO handle the FeeOnTransfer
			args = []any{amountIn, amountOut, path, swapConfig.Recipient, deadline}
			methodName = "swapExactTokensForETH"
			//methodName = "swapExactTokensForETHSupportingFeeOnTransferTokens"
		} else {
			args = []any{amountIn, amountOut, path, swapConfig.Recipient, deadline}
			methodName = "swapExactTokensForTokens"
			// TODO handle the FeeOnTransfer
			//methodName = "swapExactTokensForTokensSupportingFeeOnTransferTokens"
		}
	case entities.ExactOutput:
		if tokenIn.IsNative() {
			methodName = "swapETHForExactTokens"
			args = []any{amountOut, path, swapConfig.Recipient, deadline}
		} else if tokenOut.IsNative() {
			methodName = "swapTokensForExactETH"
			args = []any{amountOut, amountIn, path, swapConfig.Recipient, deadline}
		} else {
			methodName = "swapTokensForExactTokens"
			args = []any{amountOut, amountIn, path, swapConfig.Recipient, deadline}
		}
	default:
		return nil, errors.New("unsupported tradeType")
	}
	callData, err = contracts.IUniswapV2Router02Abi.Pack(methodName, args...)
	return
}
