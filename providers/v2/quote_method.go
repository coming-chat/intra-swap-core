package v2

import (
	"errors"
	"github.com/coming-chat/intra-swap-core/base_constant"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
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
	case base_constant.AlienbaseV2Quoter, base_constant.SwapBasedV2Quoter, base_constant.BaseswapV2Quoter:
		param.Contract = contracts.IUniswapV2Router02Abi
	case base_constant.AerodromeQuoter:
		param.Contract = contracts.IAerodromeAbi
	default:
		return rpc.MultiCallSingleParam{}, errors.New("unsupported quote")
	}
	return param, nil
}
