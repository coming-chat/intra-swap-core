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
	stepRoute, err := entitiesV3.NewRoute([]*entitiesV3.Pool{route.Pools[index].(*base_entities.V3Pool).Pool}, route.Path[index], route.Path[index+1])
	if err != nil {
		return rpc.MultiCallSingleParam{}, nil
	}
	encodedRoute, err := periphery.EncodeRouteToPath(stepRoute, tradeType == entities.ExactOutput)
	if err != nil {
		return rpc.MultiCallSingleParam{}, nil
	}
	param := rpc.MultiCallSingleParam{
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
		base_constant.ArbitrumUniswapV3Quoter:
		param.Contract = contracts.IQuoterV2Abi
	case base_constant.BaseSwapBasedV3Quoter:
		param.Contract = contracts.IQuoterAbi
	default:
		return rpc.MultiCallSingleParam{}, errors.New("unsupported quote")
	}
	return param, nil
}
