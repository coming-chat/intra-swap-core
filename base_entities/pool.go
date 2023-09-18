package base_entities

import (
	"github.com/daoleno/uniswap-sdk-core/entities"
	entitiesV3 "github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/ethereum/go-ethereum/common"
	entitiesV2 "github.com/vaulverin/uniswapv2-sdk/entities"
)

type Pool interface {
	*V2Pool | *V3Pool
	InvolvesToken(token *entities.Token) bool
	Token0() *entities.Token
	Token1() *entities.Token
	QuoterAddress() common.Address
	RouterAddress() common.Address
	PoolAddress() common.Address
	FactoryAddress() common.Address
}

type V2Pool struct {
	*entitiesV2.Pair
	poolAddress    common.Address
	quoterAddress  common.Address
	routerAddress  common.Address
	factoryAddress common.Address
}

func (v2 *V2Pool) QuoterAddress() common.Address {
	return v2.quoterAddress
}

func (v2 *V2Pool) RouterAddress() common.Address {
	return v2.routerAddress
}

func (v2 *V2Pool) PoolAddress() common.Address {
	return v2.poolAddress
}

func (v2 *V2Pool) FactoryAddress() common.Address {
	return v2.factoryAddress
}

type V3Pool struct {
	*entitiesV3.Pool
	poolAddress    common.Address
	quoterAddress  common.Address
	routerAddress  common.Address
	factoryAddress common.Address
}

func (v3 *V3Pool) QuoterAddress() common.Address {
	return v3.quoterAddress
}

func (v3 *V3Pool) RouterAddress() common.Address {
	return v3.routerAddress
}

func (v3 *V3Pool) PoolAddress() common.Address {
	return v3.poolAddress
}

func (v3 *V3Pool) FactoryAddress() common.Address {
	return v3.factoryAddress
}

func (v3 *V3Pool) Token0() *entities.Token {
	return v3.Pool.Token0
}

func (v3 *V3Pool) Token1() *entities.Token {
	return v3.Pool.Token1
}
