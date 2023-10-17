package contracts

import (
	"github.com/coming-chat/intra-swap-core/contracts/omni_swap"
	"github.com/coming-chat/intra-swap-core/contracts/uniswap_v2"
	"github.com/coming-chat/intra-swap-core/contracts/uniswap_v3"
)

var (
	// router
	ISwapRouter02Abi, _      = omni_swap.ISwapRouter02MetaData.GetAbi()
	ISwapRouterAbi, _        = omni_swap.ISwapRouterMetaData.GetAbi()
	IUniswapV2Router02Abi, _ = omni_swap.IUniswapV2Router02MetaData.GetAbi()
	IAerodromeAbi, _         = omni_swap.IAerodromeMetaData.GetAbi()
	IVelodromeAbi, _         = omni_swap.IVelodromeMetaData.GetAbi()
	ILBRouterAbi, _          = omni_swap.ILBRouterMetaData.GetAbi()
	IMuteRouterAbi, _        = omni_swap.IMuteRouterMetaData.GetAbi()
	IiZiSwapAbi, _           = omni_swap.IiZiSwapMetaData.GetAbi()
	IQuickSwapRouterAbi, _   = omni_swap.IQuickSwapRouterMetaData.GetAbi()
	IPearlRouterAbi, _       = omni_swap.IPearlRouterMetaData.GetAbi()
	CamelotRouterAbi, _      = omni_swap.ICamelotRouterMetaData.GetAbi()

	//Quoter
	IQuoterV2Abi, _         = omni_swap.IQuoterV2MetaData.GetAbi()
	IQuoterAbi, _           = omni_swap.IQuoterMetaData.GetAbi()
	IiZiSwapQuoterAbi, _    = omni_swap.IiZiSwapQuoterMetaData.GetAbi()
	IPoolIInformationAbi, _ = omni_swap.IPoolInformationMetaData.GetAbi()
	IQuickQuoterAbi, _      = omni_swap.IQuickQuoterMetaData.GetAbi()

	//Factory
	IUniswapV3FactoryAbi, _ = uniswap_v3.FactoryMetaData.GetAbi()
	IUniswapV2FactoryAbi, _ = uniswap_v2.PairFactoryMetaData.GetAbi()
	MuteIoFatoryAbi, _      = MuteIoFactoryMetaData.GetAbi()
	AerodromeFactoryAbi, _  = AerodromeFactoryMetaData.GetAbi()

	//Pool
	IUniswapV3PoolAbi, _ = uniswap_v3.PoolMetaData.GetAbi()
	IUniswapV2PoolAbi, _ = uniswap_v2.PoolMetaData.GetAbi()

	//Erc20
	Erc20Abi, _ = Erc20MetaData.GetAbi()

	//multicall
	MultiCallABi, _ = IMulticall3MetaData.GetAbi()
)
