package dex

import (
	"errors"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/coming-chat/intra-swap-core/contracts/omni_swap"
	"github.com/coming-chat/intra-swap-core/util"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/gkirito/go-ethereum/common"
)

type SwapBasedV3 struct {
	QuickSwapV3
}

func NewSwapBasedV3() SwapBasedV3 {
	return SwapBasedV3{
		QuickSwapV3: QuickSwapV3{
			UniswapV3: NewUniswapV3(contracts.ISwapRouterAbi, contracts.ISwapBasedQuoterAbi, contracts.QuickSwapFactoryAbi, contracts.IQuickSwapPoolAbi),
		},
	}
}

func (s SwapBasedV3) PackSwap(tradeType entities.TradeType, tokenIn, tokenOut *entities.Token, amountIn, amountOut *entities.CurrencyAmount, pool base_entities.Pool, swapConfig base_entities.SwapOptions) (callData []byte, err error) {
	path, err := util.EncodeRouteToPath([]base_entities.Pool{pool}, tokenIn, tradeType == entities.ExactOutput)
	if err != nil {
		return nil, err
	}
	switch tradeType {
	case entities.ExactInput:
		param := omni_swap.ISwapRouterExactInputParams{
			Path:             path,
			Recipient:        common.Address(swapConfig.Recipient),
			AmountIn:         amountIn.Quotient(),
			AmountOutMinimum: amountOut.Quotient(),
			Deadline:         swapConfig.Deadline,
		}
		callData, err = s.RouterContract.Pack("exactInput", param)
	case entities.ExactOutput:
		param := omni_swap.ISwapRouterExactOutputParams{
			Path:            path,
			Recipient:       common.Address(swapConfig.Recipient),
			AmountInMaximum: amountIn.Quotient(),
			AmountOut:       amountOut.Quotient(),
			Deadline:        swapConfig.Deadline,
		}
		callData, err = s.RouterContract.Pack("exactOutput", param)
	default:
		return nil, errors.New("unsupported tradeType")
	}
	return
}
