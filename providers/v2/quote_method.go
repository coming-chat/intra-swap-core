package v2

import (
	"errors"
	"github.com/coming-chat/intra-swap-core/base_constant"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/coming-chat/intra-swap-core/contracts/omni_swap"
	"github.com/coming-chat/intra-swap-core/dex"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func QuoteMultiCall(
	route *base_entities.MRoute,
	index int,
	tradeType entities.TradeType,
	amount *entities.CurrencyAmount,
) (rpc.MultiCallSingle[dex.QuoteResult], error) {
	name := "getAmountsOut"
	if tradeType == entities.ExactOutput {
		name = "getAmountsIn"
	}
	param := rpc.MultiCallSingle[dex.QuoteResult]{
		FunctionParams: []any{
			amount.Quotient(),
			[]common.Address{
				route.Path[index].Address,
				route.Path[index+1].Address,
			},
		},
		ContractAddress: route.Pools[index].QuoterAddress(),
		FunctionName:    name,
		CallResult: rpc.MultiCallResult[dex.QuoteResult]{
			Data: BaseAmounts{},
		},
	}
	switch route.Pools[index].QuoterAddress() {
	case base_constant.BaseAlienbaseV2Quoter,
		base_constant.BaseSwapBasedV2Quoter,
		base_constant.BaseBaseswapV2Quoter,
		base_constant.ArbitrumCamelotQuoter,
		base_constant.ArbitrumSushiQuoter,
		base_constant.PolygonQuickSwapV2Quoter:
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
	case base_constant.ZkSyncMuteQuoter:
		param.Contract = contracts.IMuteRouterAbi
		param.FunctionName = "getAmountsOut"
		param.FunctionParams = []any{
			amount.Quotient(),
			[]common.Address{
				route.Path[index].Address,
				route.Path[index+1].Address,
			},
			[]bool{route.Pools[index].Stable()},
		}
		param.CallResult.Data = MuteQuoteAmount{}
	case base_constant.PolygonPearlFiQuoter:
		param.Contract = contracts.IPearlRouterAbi
		param.FunctionParams = []any{
			amount.Quotient(),
			[]contracts.IPearlRouterroute{
				{
					route.Path[index].Address,
					route.Path[index+1].Address,
					route.Pools[index].Stable(),
				},
			},
		}
	//case base_constant.ArbitrumTraderJoeQuoter:
	//	param.Contract = contracts.ILBRouterAbi
	default:
		return rpc.MultiCallSingle[dex.QuoteResult]{}, errors.New("unsupported quote")
	}
	return param, nil
}

type BaseAmounts struct {
	Amounts []*big.Int
}

func (b BaseAmounts) QuoteAmount() *big.Int {
	return b.Amounts[len(b.Amounts)-1]
}

type MuteQuoteAmount struct {
	Amounts []*big.Int
	Stable  []bool
	Fees    []*big.Int
}

func (m MuteQuoteAmount) QuoteAmount() *big.Int {
	return m.Amounts[len(m.Amounts)-1]
}
