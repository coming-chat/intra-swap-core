package util

import (
	"github.com/daoleno/uniswap-sdk-core/entities"
	"math/big"
)

var (
	ZeroFraction = entities.NewFraction(big.NewInt(0), big.NewInt(1))
	OneFraction  = entities.NewFraction(big.NewInt(1), big.NewInt(1))
)
