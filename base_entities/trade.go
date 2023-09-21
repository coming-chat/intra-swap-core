package base_entities

import (
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	entitiesV2 "github.com/vaulverin/uniswapv2-sdk/entities"
	"math/big"
)

type Trade interface {
	InputAmount() *entities.CurrencyAmount
	OutputAmount() *entities.CurrencyAmount
	MinimumAmountOut(slippageTolerance *entities.Percent) (*entities.CurrencyAmount, error)
	MaximumAmountIn(slippageTolerance *entities.Percent) (*entities.CurrencyAmount, error)
	PriceImpact() (*entities.Percent, error)
}

type Swap struct {
	Route        *MRoute
	InputAmount  *entities.CurrencyAmount
	OutputAmount *entities.CurrencyAmount
}

type MTrade struct {
	SwapRoute []*Swap
	TradeType entities.TradeType

	inputAmount    *entities.CurrencyAmount
	outputAmount   *entities.CurrencyAmount
	executionPrice *entities.Price
	priceImpact    *entities.Percent
}

func (m *MTrade) InputAmount() *entities.CurrencyAmount {
	return m.inputAmount
}

func (m *MTrade) OutputAmount() *entities.CurrencyAmount {
	return m.outputAmount
}

func (m *MTrade) MinimumAmountOut(slippageTolerance *entities.Percent) (*entities.CurrencyAmount, error) {
	if slippageTolerance.LessThan(constants.PercentZero) {
		return nil, entitiesV2.ErrInvalidSlippageTolerance
	}

	if m.TradeType == entities.ExactOutput {
		return m.outputAmount, nil
	}

	slippageAdjustedAmountOut := entities.NewFraction(big.NewInt(0), big.NewInt(0)).
		Add(slippageTolerance.Fraction).
		Invert().
		Multiply(entities.NewFraction(m.outputAmount.Quotient(), big.NewInt(0))).Quotient()
	return entities.FromRawAmount(m.outputAmount.Currency, slippageAdjustedAmountOut), nil
}

func (m *MTrade) MaximumAmountIn(slippageTolerance *entities.Percent) (*entities.CurrencyAmount, error) {
	if slippageTolerance.LessThan(constants.PercentZero) {
		return nil, entitiesV2.ErrInvalidSlippageTolerance
	}

	if m.TradeType == entities.ExactInput {
		return m.inputAmount, nil
	}

	slippageAdjustedAmountIn := entities.NewFraction(big.NewInt(0), big.NewInt(0)).
		Add(slippageTolerance.Fraction).
		Multiply(entities.NewFraction(m.inputAmount.Quotient(), big.NewInt(0))).Quotient()
	return entities.FromRawAmount(m.inputAmount.Currency, slippageAdjustedAmountIn), nil
}

func (m *MTrade) PriceImpact() (*entities.Percent, error) {
	if m.priceImpact != nil {
		return m.priceImpact, nil
	}

	spotOutputAmount := entities.FromRawAmount(m.OutputAmount().Currency, big.NewInt(0))
	for _, swap := range m.SwapRoute {
		midPrice, err := swap.Route.MidPrice()
		if err != nil {
			return nil, err
		}
		quotePrice, err := midPrice.Quote(swap.InputAmount)
		if err != nil {
			return nil, err
		}
		spotOutputAmount = spotOutputAmount.Add(quotePrice)
	}

	priceImpact := spotOutputAmount.Subtract(m.OutputAmount()).Divide(spotOutputAmount.Fraction)
	m.priceImpact = entities.NewPercent(priceImpact.Numerator, priceImpact.Denominator)
	return m.priceImpact, nil
}
