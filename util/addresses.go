package util

import "github.com/ethereum/go-ethereum/common"

const (
	OvmGasPriceAddress = "0x420000000000000000000000000000000000000F"
	ArbGasInfoAddress  = "0x000000000000000000000000000000000000006C"

	MultiCall3Address = "0xcA11bde05977b3631167028862bE2a173976CA11"
)

var (
	AerodromePoolFactory = common.HexToAddress("0x420DD381b31aEf6683db6B902084cB0FFECe40Da")

	UniswapV3Router  = common.HexToAddress("0x2626664c2603336E57B271c5C0b26F421741e481")
	UniswapV3Factory = common.HexToAddress("0x33128a8fC17869897dcE68Ed026d694621f6FDfD")
	UniswapV3Quoter  = common.HexToAddress("0x3d4e44Eb1374240CE5F1B871ab261CD16335B76a")

	AlienbaseV2Router  = common.HexToAddress("0x8c1A3cF8f83074169FE5D7aD50B978e1cD6b37c7")
	AlienbaseV2Factory = common.HexToAddress("0x3E84D913803b02A4a7f027165E8cA42C14C0FdE7")
)
