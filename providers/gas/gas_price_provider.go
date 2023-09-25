package gas

import (
	"context"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"math/big"
)

type GasPrice struct {
	GasPriceWei *big.Int
}

type GasPriceProvider interface {
	GetGasPrice(chainId base_entities.ChainId) (GasPrice, error)
}

func NewBaseGasPriceProvider(ctx context.Context, rpc rpc.Provider) *BaseGasPriceProvider {
	return &BaseGasPriceProvider{
		rpc: rpc,
		ctx: ctx,
	}
}

type BaseGasPriceProvider struct {
	rpc rpc.Provider
	ctx context.Context
}

func (b *BaseGasPriceProvider) GetGasPrice(chainId base_entities.ChainId) (GasPrice, error) {
	ethRpc, err := b.rpc.GetEthRpc(chainId)
	if err != nil {
		return GasPrice{}, err
	}
	price, err := ethRpc.SuggestGasPrice(b.ctx)
	if err != nil {
		return GasPrice{}, err
	}
	return GasPrice{
		GasPriceWei: price,
	}, nil
}
