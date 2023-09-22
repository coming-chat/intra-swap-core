package v2

import (
	"errors"
	"fmt"
	"github.com/coming-chat/intra-swap-core/base_entities"
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
	GetQuotesManyExactIn(amountIns []*entities.CurrencyAmount, routes []*base_entities.MRoute) ([]RouteWithQuotes, error)
	GetQuotesManyExactOut(amountOuts []*entities.CurrencyAmount, routes []*base_entities.MRoute) ([]RouteWithQuotes, error)
}

func NewBaseQuoteProvider() *BaseQuoteProvider {
	return &BaseQuoteProvider{}
}

type BaseQuoteProvider struct {
}

func (b *BaseQuoteProvider) GetQuotesManyExactIn(amountIns []*entities.CurrencyAmount, routes []*base_entities.MRoute) ([]RouteWithQuotes, error) {
	return b.getQuotes(amountIns, routes, entities.ExactInput)
}

func (b *BaseQuoteProvider) GetQuotesManyExactOut(amountOuts []*entities.CurrencyAmount, routes []*base_entities.MRoute) ([]RouteWithQuotes, error) {
	return b.getQuotes(amountOuts, routes, entities.ExactOutput)
}

func (b *BaseQuoteProvider) getQuotes(
	amounts []*entities.CurrencyAmount,
	routes []*base_entities.MRoute,
	tradeType entities.TradeType,
) (routesWithQuotes []RouteWithQuotes, err error) {
	for _, route := range routes {
		var amountQuotes []AmountQuote
		for _, amount := range amounts {
			var quoteList []*big.Int
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
					quoteList = append(quoteList, cycleAmount.Quotient())
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
					quoteList = append(quoteList, cycleAmount.Quotient())
				}
			default:
				return nil, fmt.Errorf("unsupported tradeType: %#v", tradeType)
			}
			amountQuotes = append(amountQuotes, AmountQuote{
				Amount:    amount,
				Quote:     cycleAmount.Quotient(),
				QuoteList: quoteList,
			})
		}
		routesWithQuotes = append(routesWithQuotes, RouteWithQuotes{
			Route:        route,
			AmountQuotes: amountQuotes,
		})
	}
	return
}
