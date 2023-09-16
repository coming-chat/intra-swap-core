package rpc

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
)

type UniswapMultiCallProviderCore struct {
	multiCallContract *contracts.Multicall3Raw
	ChainId           base_entities.ChainId
	GasLimitPerCall   int64
}

func (u *UniswapMultiCallProviderCore) GetMultiCallContract() *contracts.Multicall3Raw {
	return u.multiCallContract
}

func NewUniswapMultiCallProviderCore(chainId base_entities.ChainId, provider BaseProvider) (*UniswapMultiCallProviderCore, error) {
	callV3, err := NewMultiCallV3(provider.Rpc)
	if err != nil {
		return nil, err
	}
	return &UniswapMultiCallProviderCore{
		ChainId:           chainId,
		GasLimitPerCall:   1_000_000,
		multiCallContract: callV3,
	}, nil
}

type UniswapMultiCallProvider[T any] struct {
	*BaseMultiCallProvider[T]
}

func NewUniswapMultiCallProvider[T any](core MultiCallProviderCore) *UniswapMultiCallProvider[T] {
	return &UniswapMultiCallProvider[T]{
		BaseMultiCallProvider: NewBaseMultiCallProvider[T](core.GetMultiCallContract()),
	}
}
