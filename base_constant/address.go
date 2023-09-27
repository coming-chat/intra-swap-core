package base_constant

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	EtherAddress       = "0x0000000000000000000000000000000000000000"
	OvmGasPriceAddress = "0x420000000000000000000000000000000000000F"
	ArbGasInfoAddress  = "0x000000000000000000000000000000000000006C"

	MultiCall3Address = "0xcA11bde05977b3631167028862bE2a173976CA11"
)

// BASE Chain
var (
	BaseAerodromePoolFactory = common.HexToAddress("0x420DD381b31aEf6683db6B902084cB0FFECe40Da")
	BaseAerodromerouter      = common.HexToAddress("0xcF77a3Ba9A5CA399B7c97c74d54e5b1Beb874E43")
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
	OptimismUniswapV3Quoter = common.HexToAddress("0xb27308f9F90D607463bb33eA1BeBb41C27CE5AB6")
)

// Arb Chain
var (
	//v3
	ArbitrumUniswapV3Router = common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45")
	ArbitrumUniswapV3Quoter = common.HexToAddress("0xb27308f9F90D607463bb33eA1BeBb41C27CE5AB6")

	//v2
	ArbitrumCamelotRouter = common.HexToAddress("0xc873fEcbd354f5A56E00E710B90EF4201db2448d")
	ArbitrumCamelotQuoter = common.HexToAddress("0xc873fEcbd354f5A56E00E710B90EF4201db2448d")

	ArbitrumSushiRouter = common.HexToAddress("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506")
	ArbitrumSushiQuoter = common.HexToAddress("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506")

	ArbitrumTraderJoeRouter = common.HexToAddress("0xb4315e873dBcf96Ffd0acd8EA43f689D8c20fB30")
	ArbitrumTraderJoeQuoter = common.HexToAddress("0xb4315e873dBcf96Ffd0acd8EA43f689D8c20fB30")
)

var (
	MultiCallAddress = common.HexToAddress(MultiCall3Address)
)
