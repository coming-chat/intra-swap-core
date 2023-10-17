package v3

import (
	"errors"
	"github.com/coming-chat/intra-swap-core/base_constant"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/coming-chat/intra-swap-core/contracts/omni_swap"
	"github.com/coming-chat/intra-swap-core/dex"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	entitiesV3 "github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/periphery"
	selfCommon "github.com/gkirito/go-ethereum/common"
	"math/big"
)

func QuoteMultiCall(
	route *base_entities.MRoute,
	index int,
	tradeType entities.TradeType,
	amount *entities.CurrencyAmount,
) (rpc.MultiCallSingle[dex.QuoteResult], error) {
	stepRoute, err := entitiesV3.NewRoute([]*entitiesV3.Pool{route.Pools[index].(*base_entities.V3Pool).Pool}, route.Path[index], route.Path[index+1])
	if err != nil {
		return rpc.MultiCallSingle[dex.QuoteResult]{}, nil
	}
	encodedRoute, err := periphery.EncodeRouteToPath(stepRoute, tradeType == entities.ExactOutput)
	if err != nil {
		return rpc.MultiCallSingle[dex.QuoteResult]{}, nil
	}
	param := rpc.MultiCallSingle[dex.QuoteResult]{
		FunctionParams: []any{
			encodedRoute,
			amount.Quotient(),
		},
		ContractAddress: route.Pools[index].QuoterAddress(),
	}
	param.FunctionName = "quoteExactInput"
	if tradeType == entities.ExactOutput {
		param.FunctionName = "quoteExactOutput"
	}
	switch route.Pools[index].QuoterAddress() {
	case base_constant.BaseUniswapV3Quoter,
		base_constant.BaseSushiswapV3Quoter,
		base_constant.OptimismUniswapV3Quoter,
		base_constant.ArbitrumUniswapV3Quoter,
		base_constant.PolygonUniswapV3Quoter:
		param.Contract = contracts.IQuoterV2Abi
		param.CallResult.Data = BaseUniswapV3QuoteData{}
	case base_constant.BaseSwapBasedV3Quoter:
		param.Contract = contracts.IQuoterAbi
		param.CallResult.Data = SingleAmountOut{}
	case base_constant.ZkSyncMaverickQuoter:
		param.Contract = contracts.IPoolIInformationAbi
		param.FunctionName = "calculateMultihopSwap"
		path := append(route.Path[index].Address.Bytes(), route.Pools[index].PoolAddress().Bytes()...)
		path = append(path, route.Path[index+1].Address.Bytes()...)
		param.FunctionParams = []any{
			path,
			amount.Quotient(),
			tradeType == entities.ExactOutput,
		}
		param.CallResult.Data = SingleAmountOut{}
	case base_constant.ZkSynciZiSwapQuoter:
		param.Contract = contracts.IiZiSwapQuoterAbi
		param.FunctionName = "swapAmount"
		param.FunctionParams = []any{
			omni_swap.IiZiSwapQuoterSwapAmountParams{
				Path:        encodedRoute,
				Recipient:   selfCommon.Address{},
				Amount:      amount.Quotient(),
				MinAcquired: big.NewInt(0),
				Deadline:    big.NewInt(10000000),
			},
		}
		param.CallResult.Data = IiZiSwapQuote{}
	case base_constant.PolygonQuickSwapV3Quoter:
		param.Contract = contracts.IQuickQuoterAbi
		param.CallResult.Data = QuickSwapQuoteData{}
	default:
		return rpc.MultiCallSingle[dex.QuoteResult]{}, errors.New("unsupported quote")
	}
	return param, nil
}

type BaseUniswapV3QuoteData struct {
	AmountOut                   *big.Int
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	GasEstimate                 *big.Int
}

func (b BaseUniswapV3QuoteData) QuoteAmount() *big.Int {
	return b.AmountOut
}

type SingleAmountOut struct {
	AmountOut *big.Int
}

func (s SingleAmountOut) QuoteAmount() *big.Int {
	return s.AmountOut
}

type IiZiSwapQuote struct {
	Cost    *big.Int
	Acquire *big.Int
}

func (i IiZiSwapQuote) QuoteAmount() *big.Int {
	return i.Acquire
}

type QuickSwapQuoteData struct {
	AmountOut *big.Int
	Fees      []uint16
}

func (q QuickSwapQuoteData) QuoteAmount() *big.Int {
	return q.AmountOut
}
