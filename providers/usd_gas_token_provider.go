package providers

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/daoleno/uniswap-sdk-core/entities"
)

type UsdGasTokensProvider interface {
	GetTokensByChain(chainId base_entities.ChainId) ([]*entities.Token, error)
}
