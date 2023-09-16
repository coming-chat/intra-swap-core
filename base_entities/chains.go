package base_entities

import (
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
)

type ChainId = uint

const (
	MAINNET          ChainId = 1
	ROPSTEN          ChainId = 3
	RINKEBY          ChainId = 4
	GÖRLI            ChainId = 5
	KOVAN            ChainId = 42
	OPTIMISM         ChainId = 10
	OPTIMISTIC_KOVAN ChainId = 69
	ARBITRUM_ONE     ChainId = 42161
	ARBITRUM_RINKEBY ChainId = 421611
	POLYGON          ChainId = 137
	POLYGON_MUMBAI   ChainId = 80001
	BASE             ChainId = 8453
	BaseGoerli       ChainId = 84531
)

var (
	HasL1Fee = map[ChainId]struct{}{
		OPTIMISM:         {},
		OPTIMISTIC_KOVAN: {},
		ARBITRUM_ONE:     {},
		ARBITRUM_RINKEBY: {},
	}
	V2Supported = map[ChainId]struct{}{
		MAINNET: {},
		KOVAN:   {},
		GÖRLI:   {},
		RINKEBY: {},
		ROPSTEN: {},
		BASE:    {},
	}
	ChainName = map[ChainId]string{
		BASE: "base",
	}
	SupportDex = map[string]struct{}{
		"aerodrome-base": {},
	}
)

var WRAPPED_NATIVE_CURRENCY = map[ChainId]*entities.Token{
	MAINNET: entities.NewToken(
		MAINNET,
		common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		18,
		"WETH",
		"Wrapped Ether",
	),
	ROPSTEN: entities.NewToken(
		ROPSTEN,
		common.HexToAddress("0xc778417E063141139Fce010982780140Aa0cD5Ab"),
		18,
		"WETH",
		"Wrapped Ether",
	),
	RINKEBY: entities.NewToken(
		RINKEBY,
		common.HexToAddress("0xc778417E063141139Fce010982780140Aa0cD5Ab"),
		18,
		"WETH",
		"Wrapped Ether",
	),
	GÖRLI: entities.NewToken(
		GÖRLI,
		common.HexToAddress("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"),
		18,
		"WETH",
		"Wrapped Ether",
	),
	KOVAN: entities.NewToken(
		KOVAN,
		common.HexToAddress("0xd0A1E359811322d97991E03f863a0C30C2cF029C"),
		18,
		"WETH",
		"Wrapped Ether",
	),
	OPTIMISM: entities.NewToken(
		OPTIMISM,
		common.HexToAddress("0x4200000000000000000000000000000000000006"),
		18,
		"WETH",
		"Wrapped Ether",
	),
	OPTIMISTIC_KOVAN: entities.NewToken(
		OPTIMISTIC_KOVAN,
		common.HexToAddress("0x4200000000000000000000000000000000000006"),
		18,
		"WETH",
		"Wrapped Ether",
	),
	ARBITRUM_ONE: entities.NewToken(
		ARBITRUM_ONE,
		common.HexToAddress("0x82aF49447D8a07e3bd95BD0d56f35241523fBab1"),
		18,
		"WETH",
		"Wrapped Ether",
	),
	ARBITRUM_RINKEBY: entities.NewToken(
		ARBITRUM_RINKEBY,
		common.HexToAddress("0xB47e6A5f8b33b3F17603C83a0535A9dcD7E32681"),
		18,
		"WETH",
		"Wrapped Ether",
	),
	POLYGON: entities.NewToken(
		POLYGON,
		common.HexToAddress("0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270"),
		18,
		"WMATIC",
		"Wrapped MATIC",
	),
	POLYGON_MUMBAI: entities.NewToken(
		POLYGON_MUMBAI,
		common.HexToAddress("0x9c3C9283D3e44854697Cd22D3Faa240Cfb032889"),
		18,
		"WMATIC",
		"Wrapped MATIC",
	),
	BASE: entities.NewToken(
		BASE,
		common.HexToAddress("0x4200000000000000000000000000000000000006"),
		18,
		"WETH",
		"Wrapped Ether",
	),
}
