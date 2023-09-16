package providers

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/daoleno/uniswap-sdk-core/entities"
)

type WrappedNativeCurrencyProvider interface {
	GetTokenByChain(chainId base_entities.ChainId) (*entities.Token, error)
}
