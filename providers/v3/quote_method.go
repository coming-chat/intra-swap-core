package v3

import (
	"errors"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts/omni_swap"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	entitiesV3 "github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/periphery"
	"github.com/ethereum/go-ethereum/common"
)

type functionName = string

const (
	QuoteExactInput  functionName = "quoteExactInput"
	QuoteExactOutput functionName = "quoteExactOutput"
)

func QuoteMultiCall(
	pool base_entities.Pool,
	name functionName,
	amount *entities.CurrencyAmount,
) (rpc.MultiCallSingleParam, error) {
	switch pool.QuoterAddress() {
	case common.HexToAddress("0x3d4e44Eb1374240CE5F1B871ab261CD16335B76a"): //uniswap v3
		abi, err := omni_swap.IQuoterV2MetaData.GetAbi()
		if err != nil {
			panic(err)
		}
		route, err := entitiesV3.NewRoute([]*entitiesV3.Pool{pool.(*base_entities.V3Pool).Pool}, pool.Token0(), pool.Token1())
		if err != nil {
			return rpc.MultiCallSingleParam{}, nil
		}
		encodedRoute, err := periphery.EncodeRouteToPath(route, name == QuoteExactOutput)
		if err != nil {
			return rpc.MultiCallSingleParam{}, nil
		}
		return rpc.MultiCallSingleParam{
			Contract: abi,
			FunctionParams: []any{
				encodedRoute,
				amount.Quotient(),
			},
			ContractAddress: pool.QuoterAddress(),
			FunctionName:    name,
		}, nil
	default:
		return rpc.MultiCallSingleParam{}, errors.New("unsupported quote")
	}
}
