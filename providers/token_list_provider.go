package providers

import (
	"errors"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/daoleno/uniswap-sdk-core/entities"
)

var ErrTokenNotFound = errors.New("token not found")

type TokenListProvider interface {
	GetTokenBySymbol(chainId base_entities.ChainId, symbol string) (*entities.Token, error)
	GetTokenByAddress(chaiId base_entities.ChainId, address string) (*entities.Token, error)
}
