package functions

import (
	"errors"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"math/big"
)

func CalculateRatioAmountIn(
	optimalRatio *entities.Fraction,
	inputTokenPrice *entities.Fraction,
	inputBalance *entities.CurrencyAmount,
	outputBalance *entities.CurrencyAmount,
) (*entities.CurrencyAmount, error) {
	// formula: amountToSwap = (inputBalance - (optimalRatio * outputBalance)) / ((optimalRatio * inputTokenPrice) + 1))
	amountToSwapRaw := inputBalance.Fraction.
		Subtract(optimalRatio.Multiply(outputBalance.Fraction)).
		Divide(
			optimalRatio.Multiply(inputTokenPrice).
				Add(entities.NewFraction(big.NewInt(1), big.NewInt(1))),
		)

	if amountToSwapRaw.LessThan(entities.NewFraction(big.NewInt(0), big.NewInt(1))) {
		return nil, errors.New("routeToRatio: insufficient input token amount")
	}
	return entities.FromRawAmount(inputBalance.Currency, amountToSwapRaw.Quotient()), nil
}
