package v2

import (
	"context"
	"errors"
	"fmt"
	"github.com/coming-chat/intra-swap-core/base_constant"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/dex"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"math/big"
)

type AmountQuote struct {
	Amount    *entities.CurrencyAmount
	Quote     *big.Int
	QuoteList []*big.Int
}

type RouteWithQuotes struct {
	Route        *base_entities.MRoute
	AmountQuotes []AmountQuote
}

type QuoteProvider interface {
	GetQuotesMany(amounts []*entities.CurrencyAmount, routes []*base_entities.MRoute, tradeType entities.TradeType) ([]RouteWithQuotes, error)
}

func NewBaseQuoteProvider(
	ctx context.Context,
	chainId base_entities.ChainId,
	baseProvider rpc.Provider,
	multiCallCore rpc.MultiCallProviderCore,
	offline bool,
	config *rpc.MultiCallConfig,
) *BaseQuoteProvider {
	provider := &BaseQuoteProvider{
		ChainId:            chainId,
		Provider:           baseProvider,
		MultiCall2Provider: multiCallCore,
		Offline:            offline,
		Ctx:                ctx,
	}
	if config == nil {
		config = &rpc.MultiCallConfig{
			RetryOptions: rpc.RetryOptions{
				Retries:    2,
				MinTimeout: 25,
				MaxTimeout: 250,
			},
			BatchParams: rpc.BatchParams{
				MultiCallChunk:  1000,
				GasLimitPerCall: 1_000_000,
				MinSuccessRate:  0.2,
			},
			MaxConcurrentNum: 3,
		}
	}
	provider.MultiCallConfig = config
	return provider
}

type BaseQuoteProvider struct {
	ChainId            base_entities.ChainId
	Provider           rpc.Provider
	MultiCall2Provider rpc.MultiCallProviderCore
	MultiCallConfig    *rpc.MultiCallConfig
	Offline            bool
	Ctx                context.Context
}

func (b *BaseQuoteProvider) GetQuotesMany(amounts []*entities.CurrencyAmount, routes []*base_entities.MRoute, tradeType entities.TradeType) ([]RouteWithQuotes, error) {
	if b.Offline {
		return b.getQuotesOffline(amounts, routes, entities.ExactInput)
	} else {
		return b.getQuotesOnline(amounts, routes, tradeType)
	}
}

func (b *BaseQuoteProvider) getQuotesOnline(amounts []*entities.CurrencyAmount,
	routes []*base_entities.MRoute,
	tradeType entities.TradeType,
) (routesWithQuotes []RouteWithQuotes, err error) {
	var (
		maxHop = 0
		result []RouteWithQuotes
	)
	for _, route := range routes {
		if len(route.Pools) > maxHop {
			maxHop = len(route.Pools)
		}
		result = append(result, RouteWithQuotes{
			Route: route,
		})
	}
	syncAmounts := make([][]*entities.CurrencyAmount, len(routes))
	for i := range syncAmounts {
		syncAmounts[i] = make([]*entities.CurrencyAmount, len(amounts))
		copy(syncAmounts[i], amounts)
	}
	for i := 0; i < maxHop; i++ {
		var (
			multiCallParams []rpc.MultiCallSingle[dex.QuoteResult]
			params          []struct {
				routeIndex  int
				amountIndex int
			}
		)
		for ri, route := range routes {
			if len(route.Pools) <= i {
				continue
			}
			for ai, a := range syncAmounts[ri] {
				if a.EqualTo(base_constant.ZeroFraction) {
					continue
				}
				call, err := dex.RouterAddrDexMap[route.Pools[i].RouterAddress()].GetQuote(route, i, tradeType, a)
				if err != nil {
					return nil, err
				}
				multiCallParams = append(multiCallParams, call)
				params = append(params, struct {
					routeIndex  int
					amountIndex int
				}{routeIndex: ri, amountIndex: ai})
				if i == 0 {
					result[ri].AmountQuotes = append(result[ri].AmountQuotes, AmountQuote{
						Amount: entities.FromRawAmount(a.Currency, a.Quotient()),
					})
				}
			}
		}
		if len(multiCallParams) == 0 {
			continue
		}
		err = rpc.ConcurrentMultiCall[dex.QuoteResult](b.ChainId, b.MultiCall2Provider, multiCallParams, b.MultiCallConfig)
		if err != nil {
			return nil, err
		}
		for pi, param := range params {
			if !multiCallParams[pi].CallResult.Success {
				result[param.routeIndex].AmountQuotes[param.amountIndex].Quote = big.NewInt(0)
				syncAmounts[param.routeIndex][param.amountIndex] = entities.FromRawAmount(syncAmounts[param.routeIndex][param.amountIndex].Currency, result[param.routeIndex].AmountQuotes[param.amountIndex].Quote)
				result[param.routeIndex].AmountQuotes[param.amountIndex].QuoteList = append(result[param.routeIndex].AmountQuotes[param.amountIndex].QuoteList, result[param.routeIndex].AmountQuotes[param.amountIndex].Quote)
				continue
			}
			syncAmounts[param.routeIndex][param.amountIndex] = entities.FromRawAmount(syncAmounts[param.routeIndex][param.amountIndex].Currency, multiCallParams[pi].CallResult.Data.QuoteAmount())
			result[param.routeIndex].AmountQuotes[param.amountIndex].Quote = multiCallParams[pi].CallResult.Data.QuoteAmount()
			result[param.routeIndex].AmountQuotes[param.amountIndex].QuoteList = append(result[param.routeIndex].AmountQuotes[param.amountIndex].QuoteList, multiCallParams[pi].CallResult.Data.QuoteAmount())
		}
	}

	return result, nil
}

func (b *BaseQuoteProvider) getQuotesOffline(
	amounts []*entities.CurrencyAmount,
	routes []*base_entities.MRoute,
	tradeType entities.TradeType,
) (routesWithQuotes []RouteWithQuotes, err error) {
	for _, route := range routes {
		var amountQuotes []AmountQuote
		for _, amount := range amounts {
			cycleAmount := amount.Wrapped()
			switch tradeType {
			case entities.ExactInput:
				for _, pair := range route.Pools {
					if pair.Protocol() != base_entities.V2 {
						return nil, errors.New("get pool v3 quote on v2 function")
					}
					cycleAmount, _, err = pair.(*base_entities.V2Pool).GetOutputAmount(cycleAmount)
					if err != nil {
						return nil, err
					}
				}
			case entities.ExactOutput:
				for i := len(route.Pools) - 1; i >= 0; i-- {
					if route.Pools[i].Protocol() != base_entities.V2 {
						return nil, errors.New("get pool v3 quote on v2 function")
					}
					cycleAmount, _, err = route.Pools[i].(*base_entities.V2Pool).GetInputAmount(cycleAmount)
					if err != nil {
						return nil, err
					}
				}
			default:
				return nil, fmt.Errorf("unsupported tradeType: %#v", tradeType)
			}
			amountQuotes = append(amountQuotes, AmountQuote{
				Amount: amount,
				Quote:  cycleAmount.Quotient(),
			})
		}
		routesWithQuotes = append(routesWithQuotes, RouteWithQuotes{
			Route:        route,
			AmountQuotes: amountQuotes,
		})
	}
	return
}
