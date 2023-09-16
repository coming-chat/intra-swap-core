package gas_price

import "math/big"

type GasPrice struct {
	GasPriceWei *big.Int
}

type GasPriceProvider interface {
	GetGasPrice() GasPrice
}

type BaseGasPriceProvider struct {
}

func (b *BaseGasPriceProvider) GetGasPrice() GasPrice {
	return GasPrice{
		GasPriceWei: big.NewInt(20),
	}
}
