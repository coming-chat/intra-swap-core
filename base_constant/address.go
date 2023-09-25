package base_constant

import "github.com/ethereum/go-ethereum/common"

const (
	OvmGasPriceAddress = "0x420000000000000000000000000000000000000F"
	ArbGasInfoAddress  = "0x000000000000000000000000000000000000006C"

	MultiCall3Address = "0xcA11bde05977b3631167028862bE2a173976CA11"
)

var (
	AerodromePoolFactory = common.HexToAddress("0x420DD381b31aEf6683db6B902084cB0FFECe40Da")
	AerodromeRouter      = common.HexToAddress("0xcF77a3Ba9A5CA399B7c97c74d54e5b1Beb874E43")
	AerodromeQuoter      = common.HexToAddress("0xcF77a3Ba9A5CA399B7c97c74d54e5b1Beb874E43")

	//V3
	UniswapV3Router  = common.HexToAddress("0x2626664c2603336E57B271c5C0b26F421741e481")
	UniswapV3Quoter  = common.HexToAddress("0x3d4e44Eb1374240CE5F1B871ab261CD16335B76a")
	UniswapV3Factory = common.HexToAddress("0x33128a8fC17869897dcE68Ed026d694621f6FDfD")

	SwapBasedV3Router = common.HexToAddress("0xD58f563A7D6150a2575C74065CB18f53EC2e9D07")
	SwapBasedV3Quoter = common.HexToAddress("0x64F477C68049B554113Bab6526038bcF5643136c")

	SushiswapV3Router = common.HexToAddress("0xFB7eF66a7e61224DD6FcD0D7d9C3be5C8B049b9f")
	SushiswapV3Quoter = common.HexToAddress("0xb1E835Dc2785b52265711e17fCCb0fd018226a6e")

	//v2
	AlienbaseV2Router  = common.HexToAddress("0x8c1A3cF8f83074169FE5D7aD50B978e1cD6b37c7")
	AlienbaseV2Quoter  = common.HexToAddress("0x8c1A3cF8f83074169FE5D7aD50B978e1cD6b37c7")
	AlienbaseV2Factory = common.HexToAddress("0x3E84D913803b02A4a7f027165E8cA42C14C0FdE7")

	SwapBasedV2Router = common.HexToAddress("0xaaa3b1F1bd7BCc97fD1917c18ADE665C5D31F066")
	SwapBasedV2Quoter = common.HexToAddress("0xaaa3b1F1bd7BCc97fD1917c18ADE665C5D31F066")

	BaseswapV2Router = common.HexToAddress("0x327Df1E6de05895d2ab08513aaDD9313Fe505d86")
	BaseswapV2Quoter = common.HexToAddress("0x327Df1E6de05895d2ab08513aaDD9313Fe505d86")
)
