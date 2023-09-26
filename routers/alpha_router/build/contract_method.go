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

type contractMethod func(trade *base_entities.MTrade, swapIndex int, amountIn, amountOut *big.Int, swapConfig config.SwapOptions) (callData []byte, err error)

func baseUniswapV2(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	amountIn, amountOut *entities.CurrencyAmount,
	swapConfig config.SwapOptions,
) (methodName string, args []any, err error) {
	if tokenIn.IsNative() && tokenOut.IsNative() {
		return "", nil, router.ErrEtherInOut
	}
	path := []common.Address{tokenIn.Address, tokenOut.Address}

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

//func baseUniswapV3[exactInput any, exactOutput any](
//	tradeType entities.TradeType,
//	tokenIn, tokenOut *entities.Token,
//	pool *base_entities.V3Pool,
//	amountIn, amountOut *big.Int,
//	swapConfig config.SwapOptions,
//) (callData []byte, err error) {
//
//}

func iSwapRouter02(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	pool *base_entities.V3Pool,
	amountIn, amountOut *entities.CurrencyAmount,
	swapConfig config.SwapOptions,
) (callData []byte, err error) {
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
	return
}

func iSwapRouter(
	tradeType entities.TradeType,
	tokenIn, tokenOut *entities.Token,
	pool *base_entities.V3Pool,
	amountIn, amountOut *entities.CurrencyAmount,
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
		})
	case entities.ExactOutput:
		callData, err = contracts.ISwapRouterAbi.Pack("exactOutput", omni_swap.ISwapRouterExactOutputParams{
			Path:            path,
			Recipient:       swapConfig.Recipient,
			AmountInMaximum: amountIn.Quotient(),
			AmountOut:       amountOut.Quotient(),
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
	swapConfig config.SwapOptions,
) (callData []byte, err error) {
	methodName, args, err := baseUniswapV2(tradeType, tokenIn, tokenOut, amountIn, amountOut, swapConfig)
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
	pool *base_entities.V3Pool,
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
			Stable:  pool.Stable,
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
	pool *base_entities.V3Pool,
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
			Stable:  pool.Stable,
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

func iLBRouter(
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
		args       []any
		methodName string
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
	callData, err = contracts.ILBRouterAbi.Pack(methodName, args...)
	return
}
