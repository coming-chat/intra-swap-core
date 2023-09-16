package base_entities

import (
	"github.com/daoleno/uniswap-sdk-core/entities"
	entitiesV3 "github.com/daoleno/uniswapv3-sdk/entities"
	entitiesV2 "github.com/vaulverin/uniswapv2-sdk/entities"
)

type Pool interface {
	*V2Pool | *V3Pool
	InvolvesToken(token *entities.Token) bool
	Token0() *entities.Token
	Token1() *entities.Token
}

type V2Pool struct {
	*entitiesV2.Pair
}

type V3Pool struct {
	*entitiesV3.Pool
}

func (v *V3Pool) Token0() *entities.Token {
	return v.Pool.Token0
}

func (v *V3Pool) Token1() *entities.Token {
	return v.Pool.Token1
}
