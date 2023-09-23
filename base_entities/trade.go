package base_entities

import (
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	entitiesV3 "github.com/daoleno/uniswapv3-sdk/entities"
	entitiesV2 "github.com/vaulverin/uniswapv2-sdk/entities"
	"math/big"
)

type Trade interface {
	InputAmount() *entities.CurrencyAmount
	OutputAmount() *entities.CurrencyAmount
	MinimumAmountOut(slippageTolerance *entities.Percent, amountOut *entities.CurrencyAmount) (*entities.CurrencyAmount, error)
	MaximumAmountIn(slippageTolerance *entities.Percent, amountIn *entities.CurrencyAmount) (*entities.CurrencyAmount, error)
	PriceImpact() (*entities.Percent, error)
}

type Swap struct {
	Route        *MRoute
	InputAmount  *entities.CurrencyAmount
	OutputAmount *entities.CurrencyAmount
}

func NewMTrade(routes []*Swap, tradeType entities.TradeType) (*MTrade, error) {
	inputCurrency := routes[0].InputAmount.Currency
	outputCurrency := routes[0].OutputAmount.Currency
	for _, route := range routes {
		if !inputCurrency.Wrapped().Equal(route.Route.Input.Wrapped()) {
			return nil, entitiesV3.ErrInputCurrencyMismatch
		}
		if !outputCurrency.Wrapped().Equal(route.Route.Output.Wrapped()) {
			return nil, entitiesV3.ErrOutputCurrencyMismatch
		}
	}

	var numPools int
	for _, route := range routes {
		numPools += len(route.Route.Pools)
	}

	// Temp disable this check
	//var poolAddressSet = make(map[common.Address]bool)
	//for _, route := range routes {
	//	for _, pool := range route.Route.Pools {
	//		poolAddressSet[pool.PoolAddress()] = true
	//	}
	//}

	//if numPools != len(poolAddressSet) {
	//	return nil, entitiesV3.ErrDuplicatePools
	//}
	return &MTrade{
		Swaps:     routes,
		TradeType: tradeType,
	}, nil
}

type MTrade struct {
	Swaps     []*Swap
	TradeType entities.TradeType

	inputAmount    *entities.CurrencyAmount
	outputAmount   *entities.CurrencyAmount
	executionPrice *entities.Price
	priceImpact    *entities.Percent
}

func (m *MTrade) InputAmount() *entities.CurrencyAmount {
	if m.inputAmount != nil {
		return m.inputAmount
	}
	totalInputFromRoutes := entities.FromRawAmount(m.Swaps[0].InputAmount.Currency, big.NewInt(0))
	for _, swap := range m.Swaps {
		totalInputFromRoutes = totalInputFromRoutes.Add(swap.InputAmount)
	}
	m.inputAmount = totalInputFromRoutes
	return m.inputAmount
}

func (m *MTrade) OutputAmount() *entities.CurrencyAmount {
	if m.outputAmount != nil {
		return m.outputAmount
	}
	totalOutputFromRoutes := entities.FromRawAmount(m.Swaps[0].OutputAmount.Currency, big.NewInt(0))
	for _, swap := range m.Swaps {
		totalOutputFromRoutes = totalOutputFromRoutes.Add(swap.OutputAmount)
	}
	m.outputAmount = totalOutputFromRoutes
	return m.outputAmount
}

func (m *MTrade) MinimumAmountOut(slippageTolerance *entities.Percent, amountOut *entities.CurrencyAmount) (*entities.CurrencyAmount, error) {
	if amountOut == nil {
		amountOut = m.InputAmount()
	}
	if slippageTolerance.LessThan(constants.PercentZero) {
		return nil, entitiesV2.ErrInvalidSlippageTolerance
	}

	if m.TradeType == entities.ExactOutput {
		return amountOut, nil
	}

	slippageAdjustedAmountOut := entities.NewFraction(big.NewInt(1), big.NewInt(1)).
		Add(slippageTolerance.Fraction).
		Invert().
		Multiply(amountOut.Fraction).Quotient()
	return entities.FromRawAmount(amountOut.Currency, slippageAdjustedAmountOut), nil
}

func (m *MTrade) MaximumAmountIn(slippageTolerance *entities.Percent, amountIn *entities.CurrencyAmount) (*entities.CurrencyAmount, error) {
	if amountIn == nil {
		amountIn = m.InputAmount()
	}
	if slippageTolerance.LessThan(constants.PercentZero) {
		return nil, entitiesV2.ErrInvalidSlippageTolerance
	}

	if m.TradeType == entities.ExactInput {
		return amountIn, nil
	}

	slippageAdjustedAmountIn := entities.NewFraction(big.NewInt(1), big.NewInt(1)).
		Add(slippageTolerance.Fraction).
		Multiply(amountIn.Fraction).Quotient()
	return entities.FromRawAmount(amountIn.Currency, slippageAdjustedAmountIn), nil
}

func (m *MTrade) PriceImpact() (*entities.Percent, error) {
	if m.priceImpact != nil {
		return m.priceImpact, nil
	}

	spotOutputAmount := entities.FromRawAmount(m.OutputAmount().Currency, big.NewInt(0))
	for _, swap := range m.Swaps {
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
