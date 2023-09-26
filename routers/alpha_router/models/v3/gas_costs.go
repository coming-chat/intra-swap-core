package v3

import (
	"github.com/coming-chat/intra-swap-core/base_constant"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"math/big"
)

// l2 execution fee on optimism is roughly the same as mainnet
func BASE_SWAP_COST(id base_entities.ChainId) *big.Int {
	switch id {
	case base_constant.MAINNET:
		fallthrough
	case base_constant.ROPSTEN:
		fallthrough
	case base_constant.RINKEBY:
		fallthrough
	case base_constant.GÖRLI:
		fallthrough
	case base_constant.OPTIMISM:
		fallthrough
	case base_constant.OPTIMISTIC_KOVAN:
		return big.NewInt(0)
	case base_constant.KOVAN:
		return big.NewInt(2000)
	case base_constant.ARBITRUM_ONE:
		return big.NewInt(0)
	case base_constant.ARBITRUM_RINKEBY:
		return big.NewInt(5000)
	case base_constant.POLYGON:
		return big.NewInt(0)
	case base_constant.POLYGON_MUMBAI:
		return big.NewInt(2000)
	default:
		return big.NewInt(0)
	}
}

func COST_PER_INIT_TICK(id base_entities.ChainId) *big.Int {
	switch id {
	case base_constant.MAINNET:
		fallthrough
	case base_constant.ROPSTEN:
		fallthrough
	case base_constant.RINKEBY:
		fallthrough
	case base_constant.GÖRLI:
		return big.NewInt(0)
	case base_constant.KOVAN:
		return big.NewInt(31000)
	case base_constant.OPTIMISM:
		return big.NewInt(0)
	case base_constant.OPTIMISTIC_KOVAN:
		return big.NewInt(31000)
	case base_constant.ARBITRUM_ONE:
		return big.NewInt(0)
	case base_constant.ARBITRUM_RINKEBY:
		return big.NewInt(31000)
	case base_constant.POLYGON:
		return big.NewInt(0)
	case base_constant.POLYGON_MUMBAI:
		return big.NewInt(31000)
	default:
		return big.NewInt(0)
	}
}

func COST_PER_HOP(id base_entities.ChainId) *big.Int {
	switch id {
	case base_constant.MAINNET:
		fallthrough
	case base_constant.ROPSTEN:
		fallthrough
	case base_constant.RINKEBY:
		fallthrough
	case base_constant.GÖRLI:
		fallthrough
	case base_constant.KOVAN:
		fallthrough
	case base_constant.OPTIMISM:
		return big.NewInt(0)
	case base_constant.OPTIMISTIC_KOVAN:
		return big.NewInt(80000)
	case base_constant.ARBITRUM_ONE:
		return big.NewInt(0)
	case base_constant.ARBITRUM_RINKEBY:
		return big.NewInt(80000)
	case base_constant.POLYGON:
		return big.NewInt(0)
	case base_constant.POLYGON_MUMBAI:
		return big.NewInt(80000)
	default:
		return big.NewInt(0)
	}
}
