package v3

import (
	"errors"
	"github.com/coming-chat/intra-swap-core/base_constant"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	entitiesV3 "github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/periphery"
)

func QuoteMultiCall(
	route *base_entities.MRoute,
	index int,
	tradeType entities.TradeType,
	amount *entities.CurrencyAmount,
) (rpc.MultiCallSingleParam, error) {
	name := "quoteExactInput"
	if tradeType == entities.ExactOutput {
		name = "quoteExactOutput"
	}
	stepRoute, err := entitiesV3.NewRoute([]*entitiesV3.Pool{route.Pools[index].(*base_entities.V3Pool).Pool}, route.Path[index], route.Path[index+1])
	if err != nil {
		return rpc.MultiCallSingleParam{}, nil
	}
	encodedRoute, err := periphery.EncodeRouteToPath(stepRoute, tradeType == entities.ExactOutput)
	if err != nil {
		return rpc.MultiCallSingleParam{}, nil
	}
	switch route.Pools[index].QuoterAddress() {
	case base_constant.UniswapV3Quoter, base_constant.SushiswapV3Quoter:
		return rpc.MultiCallSingleParam{
			Contract: contracts.IQuoterV2Abi,
			FunctionParams: []any{
				encodedRoute,
				amount.Quotient(),
			},
			ContractAddress: route.Pools[index].QuoterAddress(),
			FunctionName:    name,
		}, nil
	case base_constant.SwapBasedV3Quoter:
		return rpc.MultiCallSingleParam{
			Contract: contracts.IQuoterAbi,
			FunctionParams: []any{
				encodedRoute,
				amount.Quotient(),
			},
			ContractAddress: base_constant.SwapBasedV3Quoter,
			FunctionName:    name,
		}, nil
	default:
		return rpc.MultiCallSingleParam{}, errors.New("unsupported quote")
	}
}
