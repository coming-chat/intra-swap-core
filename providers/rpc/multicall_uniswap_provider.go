package rpc

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/ethereum/go-ethereum/ethclient"
)

type UniswapMultiCallProviderCore struct {
	provider        Provider
	GasLimitPerCall int64
}

func (u *UniswapMultiCallProviderCore) GetMultiCallContract(chainId base_entities.ChainId) (*ethclient.Client, error) {
	return u.provider.GetEthRpc(chainId)
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
