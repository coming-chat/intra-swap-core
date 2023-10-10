package providers

import (
	"fmt"
	"github.com/coming-chat/intra-swap-core/base_constant"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/daoleno/uniswap-sdk-core/entities"
)

type UsdGasTokensProvider interface {
	GetTokensByChain(chainId base_entities.ChainId) ([]*entities.Token, error)
}

type BaseUsdGasTokenProvider struct {
}

func (b *BaseUsdGasTokenProvider) GetTokensByChain(chainId base_entities.ChainId) ([]*entities.Token, error) {
	usdTokens, has := base_constant.UsdGasTokensByChain[chainId]
	if !has {
		return nil, fmt.Errorf("not found usd token on chain [%d]", chainId)
	}
	return usdTokens, nil
}
