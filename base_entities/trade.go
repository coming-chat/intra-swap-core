package base_entities

import (
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	entitiesV3 "github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/ethereum/go-ethereum/common"
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
	//multi_swap means multi_hop, DO NOT CHECK every swap has same input & output
	//we make the trade swap to not multi_route, is one route with multi_hop
	//inputCurrency := routes[0].InputAmount.Currency
	//outputCurrency := routes[0].OutputAmount.Currency
	//for _, route := range routes {
	//	if !inputCurrency.Wrapped().Equal(route.Route.Input.Wrapped()) {
	//		return nil, entitiesV3.ErrInputCurrencyMismatch
	//	}
	//	if !outputCurrency.Wrapped().Equal(route.Route.Output.Wrapped()) {
	//		return nil, entitiesV3.ErrOutputCurrencyMismatch
	//	}
	//}

	var numPools int
	for _, route := range routes {
		numPools += len(route.Route.Pools)
	}

	var poolAddressSet = make(map[common.Address]bool)
	for _, route := range routes {
		for _, pool := range route.Route.Pools {
			poolAddressSet[pool.PoolAddress()] = true
		}
	}

	if numPools != len(poolAddressSet) {
		return nil, entitiesV3.ErrDuplicatePools
	}

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
	return m.inputAmount
}

func (m *MTrade) OutputAmount() *entities.CurrencyAmount {
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

	slippageAdjustedAmountOut := entities.NewFraction(big.NewInt(0), big.NewInt(0)).
		Add(slippageTolerance.Fraction).
		Invert().
		Multiply(entities.NewFraction(amountOut.Quotient(), big.NewInt(0))).Quotient()
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

	slippageAdjustedAmountIn := entities.NewFraction(big.NewInt(0), big.NewInt(0)).
		Add(slippageTolerance.Fraction).
		Multiply(entities.NewFraction(amountIn.Quotient(), big.NewInt(0))).Quotient()
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
