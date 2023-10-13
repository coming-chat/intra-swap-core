package base_constant

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/ethereum/go-ethereum/common"
)

const (
	EtherAddress       = "0x0000000000000000000000000000000000000000"
	OvmGasPriceAddress = "0x420000000000000000000000000000000000000F"
	ArbGasInfoAddress  = "0x000000000000000000000000000000000000006C"
)

// BASE Chain
var (
	BaseAerodromePoolFactory = common.HexToAddress("0x420DD381b31aEf6683db6B902084cB0FFECe40Da")
	BaseAerodromeRouter      = common.HexToAddress("0xcF77a3Ba9A5CA399B7c97c74d54e5b1Beb874E43")
	BaseAerodromeQuoter      = common.HexToAddress("0xcF77a3Ba9A5CA399B7c97c74d54e5b1Beb874E43")

	//V3
	BaseUniswapV3Router  = common.HexToAddress("0x2626664c2603336E57B271c5C0b26F421741e481")
	BaseUniswapV3Quoter  = common.HexToAddress("0x3d4e44Eb1374240CE5F1B871ab261CD16335B76a")
	BaseUniswapV3Factory = common.HexToAddress("0x33128a8fC17869897dcE68Ed026d694621f6FDfD")

	BaseSwapBasedV3Router = common.HexToAddress("0xD58f563A7D6150a2575C74065CB18f53EC2e9D07")
	BaseSwapBasedV3Quoter = common.HexToAddress("0x64F477C68049B554113Bab6526038bcF5643136c")

	BaseSushiswapV3Router = common.HexToAddress("0xFB7eF66a7e61224DD6FcD0D7d9C3be5C8B049b9f")
	BaseSushiswapV3Quoter = common.HexToAddress("0xb1E835Dc2785b52265711e17fCCb0fd018226a6e")

	//v2
	BaseAlienbaseV2Router  = common.HexToAddress("0x8c1A3cF8f83074169FE5D7aD50B978e1cD6b37c7")
	BaseAlienbaseV2Quoter  = common.HexToAddress("0x8c1A3cF8f83074169FE5D7aD50B978e1cD6b37c7")
	BaseAlienbaseV2Factory = common.HexToAddress("0x3E84D913803b02A4a7f027165E8cA42C14C0FdE7")

	BaseSwapBasedV2Router = common.HexToAddress("0xaaa3b1F1bd7BCc97fD1917c18ADE665C5D31F066")
	BaseSwapBasedV2Quoter = common.HexToAddress("0xaaa3b1F1bd7BCc97fD1917c18ADE665C5D31F066")

	BaseBaseswapV2Router = common.HexToAddress("0x327Df1E6de05895d2ab08513aaDD9313Fe505d86")
	BaseBaseswapV2Quoter = common.HexToAddress("0x327Df1E6de05895d2ab08513aaDD9313Fe505d86")
)

// OP Chain
var (
	//v2
	OptimismVelodromeV2Router = common.HexToAddress("0xa062aE8A9c5e11aaA026fc2670B0D65cCc8B2858")
	OptimismVelodromeV2Quoter = common.HexToAddress("0xa062aE8A9c5e11aaA026fc2670B0D65cCc8B2858")

	//v3
	OptimismUniswapV3Router = common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45")
	OptimismUniswapV3Quoter = common.HexToAddress("0x61fFE014bA17989E743c5F6cB21bF9697530B21e")
)

// Arb Chain
var (
	//v3
	ArbitrumUniswapV3Router = common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45")
	ArbitrumUniswapV3Quoter = common.HexToAddress("0x61fFE014bA17989E743c5F6cB21bF9697530B21e")

	//v2
	ArbitrumCamelotRouter = common.HexToAddress("0xc873fEcbd354f5A56E00E710B90EF4201db2448d")
	ArbitrumCamelotQuoter = common.HexToAddress("0xc873fEcbd354f5A56E00E710B90EF4201db2448d")

	ArbitrumSushiRouter = common.HexToAddress("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506")
	ArbitrumSushiQuoter = common.HexToAddress("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506")

	ArbitrumTraderJoeRouter = common.HexToAddress("0xb4315e873dBcf96Ffd0acd8EA43f689D8c20fB30")
	ArbitrumTraderJoeQuoter = common.HexToAddress("0xb4315e873dBcf96Ffd0acd8EA43f689D8c20fB30")
)

// zkSync
var (
	ZkSyncMuteRouter = common.HexToAddress("0x8B791913eB07C32779a16750e3868aA8495F5964")
	ZkSyncMuteQuoter = common.HexToAddress("0x8B791913eB07C32779a16750e3868aA8495F5964")

	ZkSyncMaverickRouter = common.HexToAddress("0x39E098A153Ad69834a9Dac32f0FCa92066aD03f4")
	ZkSyncMaverickQuoter = common.HexToAddress("0x57D47F505EdaA8Ae1eFD807A860A79A28bE06449")

	ZkSynciZiSwapRouter = common.HexToAddress("0x943ac2310D9BC703d6AB5e5e76876e212100f894")
	ZkSynciZiSwapQuoter = common.HexToAddress("0x30C089574551516e5F1169C32C6D429C92bf3CD7")
)

// Polygon
var (
	PolygonQuickSwapV3Router = common.HexToAddress("0xf5b509bB0909a69B1c207E495f687a596C168E12")
	PolygonQuickSwapV3Quoter = common.HexToAddress("0xa15F0D7377B2A0C0c10db057f641beD21028FC89")

	PolygonQuickSwapV2Router = common.HexToAddress("0xa5E0829CaCEd8fFDD4De3c43696c57F7D7A678ff")
	PolygonQuickSwapV2Quoter = common.HexToAddress("0xa5E0829CaCEd8fFDD4De3c43696c57F7D7A678ff")

	PolygonUniswapV3Router = common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45")
	PolygonUniswapV3Quoter = common.HexToAddress("0xa5E0829CaCEd8fFDD4De3c43696c57F7D7A678ff")

	PolygonPearlFiRouter = common.HexToAddress("0xcC25C0FD84737F44a7d38649b69491BBf0c7f083")
	PolygonPearlFiQuoter = common.HexToAddress("0xcC25C0FD84737F44a7d38649b69491BBf0c7f083")
)

var (
	defaultMultiCallAddr   = common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")
	zksyncEraMultiCallAddr = common.HexToAddress("0xb27308f9F90D607463bb33eA1BeBb41C27CE5AB6")
)

func MultiCallAddress(chainId base_entities.ChainId) *common.Address {
	switch chainId {
	case ZKSYNC_ERA:
		return &zksyncEraMultiCallAddr
	default:
		return &defaultMultiCallAddr
	}
}
