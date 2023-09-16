package providers

import (
	"errors"
	"github.com/daoleno/uniswap-sdk-core/entities"
)

var ErrTokenNotFound = errors.New("token not found")

type TokenListProvider interface {
	GetTokenBySymbol(symbol string) (*entities.Token, error)
	GetTokenByAddress(address string) (*entities.Token, error)
}
