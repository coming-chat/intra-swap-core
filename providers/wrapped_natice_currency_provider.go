package providers

import (
	"fmt"
	"github.com/coming-chat/intra-swap-core/base_constant"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/daoleno/uniswap-sdk-core/entities"
)

type WrappedNativeCurrencyProvider interface {
	GetTokenByChain(chainId base_entities.ChainId) (*entities.Token, error)
}

type BaseWrappedNativeCurrencyProvider struct {
}

func (b *BaseWrappedNativeCurrencyProvider) GetTokenByChain(chainId base_entities.ChainId) (*entities.Token, error) {
	wToken, has := base_constant.WrappedNativeCurrency[chainId]
	if !has {
		return nil, fmt.Errorf("not found wrapped native ETH on chain[%d]", chainId)
	}
	return wToken, nil
}
