package base_constant

import "github.com/coming-chat/intra-swap-core/base_entities"

const (
	MAINNET          base_entities.ChainId = 1
	ROPSTEN          base_entities.ChainId = 3
	RINKEBY          base_entities.ChainId = 4
	GÖRLI            base_entities.ChainId = 5
	KOVAN            base_entities.ChainId = 42
	OPTIMISM         base_entities.ChainId = 10
	OPTIMISTIC_KOVAN base_entities.ChainId = 69
	ARBITRUM_ONE     base_entities.ChainId = 42161
	ARBITRUM_RINKEBY base_entities.ChainId = 421611
	POLYGON          base_entities.ChainId = 137
	POLYGON_MUMBAI   base_entities.ChainId = 80001
	BASE             base_entities.ChainId = 8453
	BASE_GOERLI      base_entities.ChainId = 84531
	ZKSYNC_ERA       base_entities.ChainId = 324
	ZKSYNC_ERA_TEST  base_entities.ChainId = 280
)

var (
	HasL1Fee = map[base_entities.ChainId]struct{}{
		OPTIMISM:         {},
		OPTIMISTIC_KOVAN: {},
		ARBITRUM_ONE:     {},
		ARBITRUM_RINKEBY: {},
	}
	V2Supported = map[base_entities.ChainId]struct{}{
		MAINNET: {},
		KOVAN:   {},
		GÖRLI:   {},
		RINKEBY: {},
		ROPSTEN: {},
		BASE:    {},
	}
	ChainName = map[base_entities.ChainId]string{
		BASE: "base",
	}
	SupportDex = map[string]struct{}{
		"aerodrome-base": {},
	}
)
