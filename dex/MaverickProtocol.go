package dex

import (
	"errors"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/coming-chat/intra-swap-core/contracts/omni_swap"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/coming-chat/intra-swap-core/util"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/gkirito/go-ethereum/common"
	"math/big"
)

type SingleAmountOut struct {
	AmountOut *big.Int
}

func (s SingleAmountOut) QuoteAmount() *big.Int {
	return s.AmountOut
}

type MaverickProtocol struct {
	UniswapV3
}

// TODO Need confirm the getPool from factory & getPoolInfo from pool
func NewMaverickProtocol() MaverickProtocol {
	return MaverickProtocol{
		UniswapV3: NewUniswapV3(contracts.ISwapRouterAbi, contracts.IPoolIInformationAbi, nil, nil),
	}
}

func (m MaverickProtocol) GetQuote(route *base_entities.MRoute, index int, tradeType entities.TradeType, amount *entities.CurrencyAmount) (rpc.MultiCallSingle[QuoteResult], error) {
	path := append(route.Path[index].Address.Bytes(), route.Pools[index].PoolAddress().Bytes()...)
	path = append(path, route.Path[index+1].Address.Bytes()...)
	return rpc.MultiCallSingle[QuoteResult]{
		Contract:     m.QuoterAbi,
		FunctionName: "calculateMultihopSwap",
		FunctionParams: []any{
			path,
			amount.Quotient(),
			tradeType == entities.ExactOutput,
		},
		CallResult: rpc.MultiCallResult[QuoteResult]{
			Data: &SingleAmountOut{},
		},
	}, nil
}

func (m MaverickProtocol) PackSwap(tradeType entities.TradeType, tokenIn, tokenOut *entities.Token, amountIn, amountOut *entities.CurrencyAmount, pool base_entities.Pool, swapConfig base_entities.SwapOptions) (callData []byte, err error) {
	path, err := util.EncodeRouteToPath([]base_entities.Pool{pool}, tokenIn, tradeType == entities.ExactOutput)
	if err != nil {
		return nil, err
	}
	switch tradeType {
	case entities.ExactInput:
		callData, err = m.RouterContract.Pack("exactInput", omni_swap.ISwapRouterExactInputParams{
			Path:             path,
			Recipient:        common.Address(swapConfig.Recipient),
			AmountIn:         amountIn.Quotient(),
			AmountOutMinimum: amountOut.Quotient(),
			Deadline:         swapConfig.Deadline,
		})
	case entities.ExactOutput:
		callData, err = m.RouterContract.Pack("exactOutput", omni_swap.ISwapRouterExactOutputParams{
			Path:            path,
			Recipient:       common.Address(swapConfig.Recipient),
			AmountInMaximum: amountIn.Quotient(),
			AmountOut:       amountOut.Quotient(),
			Deadline:        swapConfig.Deadline,
		})
	default:
		return nil, errors.New("unsupported tradeType")
	}
	return
}
