package build

import (
	"errors"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/coming-chat/intra-swap-core/contracts/omni_swap"
	"github.com/coming-chat/intra-swap-core/routers/alpha_router/config"
	"github.com/coming-chat/intra-swap-core/util"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/vaulverin/uniswapv2-sdk/router"
	"math/big"
	"time"
)

type packMethod func(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	amountIn, amountOut *entities.CurrencyAmount,
	pool base_entities.Pool,
	swapConfig config.SwapOptions,
) (callData []byte, err error)

func baseUniswapV2(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	amountIn, amountOut *entities.CurrencyAmount,
	pool base_entities.Pool,
	swapConfig config.SwapOptions,
	withStable bool,
) (methodName string, args []any, err error) {
	if tokenIn.IsNative() && tokenOut.IsNative() {
		return "", nil, router.ErrEtherInOut
	}
	path := []any{tokenIn.Address, tokenOut.Address}
	if withStable {
		path = append(path, pool.Stable())
	}

	deadline := swapConfig.Deadline
	if swapConfig.Deadline == nil {
		deadline = big.NewInt(time.Now().Add(5 * time.Minute).Unix())
	}
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
		return "", nil, errors.New("unsupported tradeType")
	}
	return
}

func iSwapRouter02(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	amountIn, amountOut *entities.CurrencyAmount,
	pool base_entities.Pool,
	swapConfig config.SwapOptions,
) (callData []byte, err error) {
	path, err := util.EncodeRouteToPath([]base_entities.Pool{pool}, tokenIn, tradeType == entities.ExactOutput)
	if err != nil {
		return nil, err
	}
	switch tradeType {
	case entities.ExactInput:
		param := omni_swap.ISwapRouter02ExactInputParams{
			Path:             path,
			Recipient:        swapConfig.Recipient,
			AmountIn:         amountIn.Quotient(),
			AmountOutMinimum: amountOut.Quotient(),
		}
		callData, err = contracts.ISwapRouter02Abi.Pack("exactInput", param)
	case entities.ExactOutput:
		param := omni_swap.ISwapRouter02ExactOutputParams{
			Path:            path,
			Recipient:       swapConfig.Recipient,
			AmountInMaximum: amountIn.Quotient(),
			AmountOut:       amountOut.Quotient(),
		}
		callData, err = contracts.ISwapRouter02Abi.Pack("exactOutput", param)
	default:
		return nil, errors.New("unsupported tradeType")
	}
	return
}

func iSwapRouter(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	amountIn, amountOut *entities.CurrencyAmount,
	pool base_entities.Pool,
	swapConfig config.SwapOptions,
) (callData []byte, err error) {
	path, err := util.EncodeRouteToPath([]base_entities.Pool{pool}, tokenIn, tradeType == entities.ExactOutput)
	if err != nil {
		return nil, err
	}
	switch tradeType {
	case entities.ExactInput:
		callData, err = contracts.ISwapRouterAbi.Pack("exactInput", omni_swap.ISwapRouterExactInputParams{
			Path:             path,
			Recipient:        swapConfig.Recipient,
			AmountIn:         amountIn.Quotient(),
			AmountOutMinimum: amountOut.Quotient(),
			Deadline:         swapConfig.Deadline,
		})
	case entities.ExactOutput:
		callData, err = contracts.ISwapRouterAbi.Pack("exactOutput", omni_swap.ISwapRouterExactOutputParams{
			Path:            path,
			Recipient:       swapConfig.Recipient,
			AmountInMaximum: amountIn.Quotient(),
			AmountOut:       amountOut.Quotient(),
			Deadline:        swapConfig.Deadline,
		})
	default:
		return nil, errors.New("unsupported tradeType")
	}
	return
}

func iUniswapV2Router02(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	amountIn, amountOut *entities.CurrencyAmount,
	pool base_entities.Pool,
	swapConfig config.SwapOptions,
) (callData []byte, err error) {
	methodName, args, err := baseUniswapV2(tradeType, tokenIn, tokenOut, amountIn, amountOut, pool, swapConfig, false)
	if err != nil {
		return nil, err
	}
	callData, err = contracts.IUniswapV2Router02Abi.Pack(methodName, args...)
	return
}

func iAerodromeRouter(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	amountIn, amountOut *entities.CurrencyAmount,
	pool base_entities.Pool,
	swapConfig config.SwapOptions,
) (callData []byte, err error) {
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
	callData, err = contracts.IAerodromeAbi.Pack(methodName, args...)
	return
}

func iVelodromeRouter(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	amountIn, amountOut *entities.CurrencyAmount,
	pool base_entities.Pool,
	swapConfig config.SwapOptions,
) (callData []byte, err error) {
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
	callData, err = contracts.IVelodromeAbi.Pack(methodName, args...)
	return
}

func iMuteRouter(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	amountIn, amountOut *entities.CurrencyAmount,
	pool base_entities.Pool,
	swapConfig config.SwapOptions,
) (callData []byte, err error) {
	method, args, err := baseUniswapV2(tradeType, tokenIn, tokenOut, amountIn, amountOut, pool, swapConfig, true)
	if err != nil {
		return nil, err
	}
	return contracts.IMuteRouterAbi.Pack(method, args)
}

func iQuickSwapRouter(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	amountIn, amountOut *entities.CurrencyAmount,
	pool base_entities.Pool,
	swapConfig config.SwapOptions,
) (callData []byte, err error) {
	path, err := util.EncodeRouteToPath([]base_entities.Pool{pool}, tokenIn, tradeType == entities.ExactOutput)
	if err != nil {
		return nil, err
	}
	switch tradeType {
	case entities.ExactInput:
		callData, err = contracts.IQuickSwapRouterAbi.Pack("exactInput", omni_swap.IQuickSwapRouterExactInputParams{
			Path:             path,
			Recipient:        swapConfig.Recipient,
			AmountIn:         amountIn.Quotient(),
			AmountOutMinimum: amountOut.Quotient(),
			Deadline:         swapConfig.Deadline,
		})
	case entities.ExactOutput:
		callData, err = contracts.ISwapRouterAbi.Pack("exactOutput", omni_swap.IQuickSwapRouterExactOutputParams{
			Path:            path,
			Recipient:       swapConfig.Recipient,
			AmountInMaximum: amountIn.Quotient(),
			AmountOut:       amountOut.Quotient(),
			Deadline:        swapConfig.Deadline,
		})
	default:
		return nil, errors.New("unsupported tradeType")
	}
	return
}

func iPearlRouter(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	amountIn, amountOut *entities.CurrencyAmount,
	pool base_entities.Pool,
	swapConfig config.SwapOptions,
) (callData []byte, err error) {
	if tradeType == entities.ExactOutput {
		return nil, errors.New("dex contract not support exactOutput")
	}
	method, args, err := baseUniswapV2(tradeType, tokenIn, tokenOut, amountIn, amountOut, pool, swapConfig, true)
	if err != nil {
		return nil, err
	}
	return contracts.IPearlRouterAbi.Pack(method, args...)
}

//
//func iiZiSwapPool(
//	tradeType entities.TradeType,
//	tokenIn, tokenOut *entities.Token,
//	amountIn, amountOut *entities.CurrencyAmount,
//	pool base_entities.Pool,
//	swapConfig config.SwapOptions,
//) (callData []byte, err error) {
//	contracts.IiZiSwapPoolAbi.Pack()
//}
