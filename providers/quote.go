package providers

import "math/big"

type QuoteResult interface {
	QuoteAmount() *big.Int
}
