package rpc

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
)

type UniswapMultiCallProviderCore struct {
	provider        Provider
	GasLimitPerCall int64
}

func (u *UniswapMultiCallProviderCore) GetMultiCallContract(chainId base_entities.ChainId) (*contracts.Multicall3Raw, error) {
	rpc, err := u.provider.GetEthRpc(chainId)
	if err != nil {
		return nil, err
	}
	return NewMultiCallV3(rpc)
}

func NewUniswapMultiCallProviderCore(provider Provider) (*UniswapMultiCallProviderCore, error) {
	return &UniswapMultiCallProviderCore{
		provider:        provider,
		GasLimitPerCall: 1_000_000,
	}, nil
}

type UniswapMultiCallProvider[T any] struct {
	*BaseMultiCallProvider[T]
}
