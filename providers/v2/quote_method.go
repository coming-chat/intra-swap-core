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
	switch route.Pools[index].QuoterAddress() {
	case base_constant.AlienbaseV2Quoter:
		name := "getAmountsOut"
		if tradeType == entities.ExactOutput {
			name = "getAmountsIn"
		}
		return rpc.MultiCallSingleParam{
			Contract: contracts.IUniswapV2Router02Abi,
			FunctionParams: []any{
				amount.Quotient(),
				[]common.Address{
					route.Path[index].Address,
					route.Path[index+1].Address,
				},
			},
			ContractAddress: route.Pools[index].QuoterAddress(),
			FunctionName:    name,
		}, nil
	default:
		return rpc.MultiCallSingleParam{}, errors.New("unsupported quote")
	}
}
