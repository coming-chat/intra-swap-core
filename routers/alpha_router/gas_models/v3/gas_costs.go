package v3

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
	"math/big"
)

// l2 execution fee on optimism is roughly the same as mainnet
func BASE_SWAP_COST(id base_entities.ChainId) *big.Int {
	switch id {
	case base_entities.MAINNET:
		fallthrough
	case base_entities.ROPSTEN:
		fallthrough
	case base_entities.RINKEBY:
		fallthrough
	case base_entities.GÖRLI:
		fallthrough
	case base_entities.OPTIMISM:
		fallthrough
	case base_entities.OPTIMISTIC_KOVAN:
		return big.NewInt(0)
	case base_entities.KOVAN:
		return big.NewInt(2000)
	case base_entities.ARBITRUM_ONE:
		return big.NewInt(0)
	case base_entities.ARBITRUM_RINKEBY:
		return big.NewInt(5000)
	case base_entities.POLYGON:
		return big.NewInt(0)
	case base_entities.POLYGON_MUMBAI:
		return big.NewInt(2000)
	default:
		return big.NewInt(0)
	}
}

func COST_PER_INIT_TICK(id base_entities.ChainId) *big.Int {
	switch id {
	case base_entities.MAINNET:
		fallthrough
	case base_entities.ROPSTEN:
		fallthrough
	case base_entities.RINKEBY:
		fallthrough
	case base_entities.GÖRLI:
		return big.NewInt(0)
	case base_entities.KOVAN:
		return big.NewInt(31000)
	case base_entities.OPTIMISM:
		return big.NewInt(0)
	case base_entities.OPTIMISTIC_KOVAN:
		return big.NewInt(31000)
	case base_entities.ARBITRUM_ONE:
		return big.NewInt(0)
	case base_entities.ARBITRUM_RINKEBY:
		return big.NewInt(31000)
	case base_entities.POLYGON:
		return big.NewInt(0)
	case base_entities.POLYGON_MUMBAI:
		return big.NewInt(31000)
	default:
		return big.NewInt(0)
	}
}

func COST_PER_HOP(id base_entities.ChainId) *big.Int {
	switch id {
	case base_entities.MAINNET:
		fallthrough
	case base_entities.ROPSTEN:
		fallthrough
	case base_entities.RINKEBY:
		fallthrough
	case base_entities.GÖRLI:
		fallthrough
	case base_entities.KOVAN:
		fallthrough
	case base_entities.OPTIMISM:
		return big.NewInt(0)
	case base_entities.OPTIMISTIC_KOVAN:
		return big.NewInt(80000)
	case base_entities.ARBITRUM_ONE:
		return big.NewInt(0)
	case base_entities.ARBITRUM_RINKEBY:
		return big.NewInt(80000)
	case base_entities.POLYGON:
		return big.NewInt(0)
	case base_entities.POLYGON_MUMBAI:
		return big.NewInt(80000)
	default:
		return big.NewInt(0)
	}
}
