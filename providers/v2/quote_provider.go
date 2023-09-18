package v2

import (
	"fmt"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"math/big"
)

type AmountQuote struct {
	Amount *entities.CurrencyAmount
	Quote  *big.Int
}

type RouteWithQuotes struct {
	Route        *base_entities.V2Route
	AmountQuotes []AmountQuote
}

type QuoteProvider interface {
	GetQuotesManyExactIn(amountIns []*entities.CurrencyAmount, routes []*base_entities.V2Route) ([]RouteWithQuotes, error)
	GetQuotesManyExactOut(amountOuts []*entities.CurrencyAmount, routes []*base_entities.V2Route) ([]RouteWithQuotes, error)
}

func NewBaseQuoteProvider() *BaseQuoteProvider {
	return &BaseQuoteProvider{}
}

type BaseQuoteProvider struct {
}

func (b *BaseQuoteProvider) GetQuotesManyExactIn(amountIns []*entities.CurrencyAmount, routes []*base_entities.V2Route) ([]RouteWithQuotes, error) {
	return b.getQuotes(amountIns, routes, entities.ExactInput)
}

func (b *BaseQuoteProvider) GetQuotesManyExactOut(amountOuts []*entities.CurrencyAmount, routes []*base_entities.V2Route) ([]RouteWithQuotes, error) {
	return b.getQuotes(amountOuts, routes, entities.ExactOutput)
}

func (b *BaseQuoteProvider) getQuotes(
	amounts []*entities.CurrencyAmount,
	routes []*base_entities.V2Route,
	tradeType entities.TradeType,
) (routesWithQuotes []RouteWithQuotes, err error) {
	for _, route := range routes {
		var amountQuotes []AmountQuote
		for _, amount := range amounts {
			cycleAmount := amount.Wrapped()
			switch tradeType {
			case entities.ExactInput:
				for _, pair := range route.Pairs {
					cycleAmount, _, err = pair.GetOutputAmount(cycleAmount)
					if err != nil {
						return nil, err
					}
				}
			case entities.ExactOutput:
				for i := len(route.Pairs) - 1; i >= 0; i-- {
					cycleAmount, _, err = route.Pairs[i].GetInputAmount(cycleAmount)
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
