package v2

import (
	"context"
	"errors"
	"fmt"
	"github.com/coming-chat/intra-swap-core/base_constant"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"math/big"
)

type AmountQuote struct {
	Amount *entities.CurrencyAmount
	Quote  *big.Int
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
				MultiCallChunk:  100,
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
		var multiCallParams []rpc.MultiCallSingleParam
		for ri, route := range routes {
			if len(route.Pools) <= i {
				continue
			}
			for _, a := range syncAmounts[ri] {
				if a.EqualTo(base_constant.ZeroFraction) {
					continue
				}
				call, err := QuoteMultiCall(route, i, tradeType, a)
				if err != nil {
					return nil, err
				}
				multiCallParams = append(multiCallParams, call)
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
		callResult, err := rpc.ConcurrentMultiCall[[]*big.Int](b.ChainId, b.MultiCall2Provider, multiCallParams, b.MultiCallConfig)
		if err != nil {
			return nil, err
		}
		callDataIndex := 0
		for ri, route := range routes {
			if len(route.Pools) <= i {
				continue
			}
			for ra, a := range syncAmounts[ri] {
				if a.EqualTo(base_constant.ZeroFraction) {
					continue
				}
				quoteData := callResult.ReturnData[callDataIndex]
				syncAmounts[ri][ra] = entities.FromRawAmount(syncAmounts[ri][ra].Currency, quoteData.Data[len(quoteData.Data)-1])
				result[ri].AmountQuotes[ra].Quote = quoteData.Data[len(quoteData.Data)-1]
				callDataIndex++
			}
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
