package v2

import (
	"errors"
	"github.com/coming-chat/intra-swap-core/base_constant"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/coming-chat/intra-swap-core/contracts/omni_swap"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
)

func QuoteMultiCall(
	route *base_entities.MRoute,
	index int,
	tradeType entities.TradeType,
	amount *entities.CurrencyAmount,
) (rpc.MultiCallSingleParam, error) {
	name := "getAmountsOut"
	if tradeType == entities.ExactOutput {
		name = "getAmountsIn"
	}
	param := rpc.MultiCallSingleParam{
		FunctionParams: []any{
			amount.Quotient(),
			[]common.Address{
				route.Path[index].Address,
				route.Path[index+1].Address,
			},
		},
		ContractAddress: route.Pools[index].QuoterAddress(),
		FunctionName:    name,
	}
	switch route.Pools[index].QuoterAddress() {
	case base_constant.BaseAlienbaseV2Quoter,
		base_constant.BaseSwapBasedV2Quoter,
		base_constant.BaseBaseswapV2Quoter,
		base_constant.ArbitrumCamelotQuoter,
		base_constant.ArbitrumSushiQuoter:
		param.Contract = contracts.IUniswapV2Router02Abi
	case base_constant.BaseAerodromeQuoter:
		param.Contract = contracts.IAerodromeAbi
		param.FunctionName = "getAmountsOut"
		param.FunctionParams = []any{
			amount.Quotient(),
			[]omni_swap.IAerodromeRoute{
				{
					From:    route.Path[index].Address,
					To:      route.Path[index+1].Address,
					Stable:  route.Pools[index].Stable(),
					Factory: route.Pools[index].FactoryAddress(),
				},
			},
		}
	case base_constant.OptimismVelodromeV2Quoter:
		param.Contract = contracts.IVelodromeAbi
		param.FunctionName = "getAmountsOut"
		param.FunctionParams = []any{
			amount.Quotient(),
			[]omni_swap.IVelodromeRoute{
				{
					From:    route.Path[index].Address,
					To:      route.Path[index+1].Address,
					Stable:  route.Pools[index].Stable(),
					Factory: route.Pools[index].FactoryAddress(),
				},
			},
		}
	//case base_constant.ArbitrumTraderJoeQuoter:
	//	param.Contract = contracts.ILBRouterAbi
	default:
		return rpc.MultiCallSingleParam{}, errors.New("unsupported quote")
	}
	return param, nil
}
