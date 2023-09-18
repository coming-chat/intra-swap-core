package config

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/periphery"
)

type SwapOptions struct {
	Recipient         string
	SlippageTolerance *entities.Percent
	Deadline          int
	InputTokenPermit  *periphery.PermitOptions
}

// Config passed in to determine configurations on acceptable liquidity
// to add to a position and max iterations on the route-finding algorithm
type SwapAndAddConfig struct {
	MaxIterations       int
	RatioErrorTolerance *entities.Fraction
}

// Options for executing the swap and add.
// If provided, calldata for executing the swap and add will also be returned.
type SwapAndAddOptions struct {
	SwapOptions SwapOptions
	//AddLiquidityOptions CondensedAddLiquidityOptions
}

type AlphaRouterConfig struct {
	/**
	 * The block number to use for all on-chain data. If not provided, the router will
	 * use the latest block returned by the provider.
	 */
	BlockNumber uint64
	/**
	 * The protocols to consider when finding the optimal swap. If not provided all protocols
	 * will be used.
	 */
	Protocols []base_entities.Protocol
	/**
	 * Config for selecting which pools to consider routing via on V2.
	 */
	V2PoolSelection ProtocolPoolSelection
	/**
	 * Config for selecting which pools to consider routing via on V3.
	 */
	V3PoolSelection ProtocolPoolSelection
	/**
	 * For each route, the maximum number of hops to consider. More hops will increase latency of the algorithm.
	 */
	MaxSwapsPerPath int
	/**
	 * The maximum number of splits in the returned route. A higher maximum will increase latency of the algorithm.
	 */
	MaxSplits int
	/**
	 * The minimum number of splits in the returned route.
	 * This parameters should always be set to 1. It is only included for testing purposes.
	 */
	MinSplits int
	/**
	 * Forces the returned swap to route across all protocols.
	 * This parameter should always be false. It is only included for testing purposes.
	 */
	ForceCrossProtocol bool
	/**
	 * The minimum percentage of the input token to use for each route in a split route.
	 * All routes will have a multiple of this value. For example is distribution percentage is 5,
	 * a potential return swap would be:
	 *
	 * 5% of input => Route 1
	 * 55% of input => Route 2
	 * 40% of input => Route 3
	 */
	DistributionPercent int
}

type ProtocolPoolSelection struct {
	/**
	 * The top N pools by TVL out of all pools on the protocol.
	 */
	TopN int
	/**
	 * The top N pools by TVL of pools that consist of tokenIn and tokenOut.
	 */
	TopNDirectSwaps int
	/**
	 * The top N pools by TVL of pools where one token is tokenIn and the
	 * top N pools by TVL of pools where one token is tokenOut tokenOut.
	 */
	TopNTokenInOut int
	/**
	 * Given the topNTokenInOut pools, gets the top N pools that involve the other token.
	 * E.g. for a WETH -> USDC swap, if topNTokenInOut found WETH -> DAI and WETH -> USDT,
	 * a value of 2 would find the top 2 pools that involve DAI or USDT.
	 */
	TopNSecondHop int
	/**
	 * The top N pools for token in and token out that involve a token from a list of
	 * hardcoded 'base tokens'. These are standard tokens such as WETH, USDC, DAI, etc.
	 * This is similar to how the legacy routing algorithm used by Uniswap would select
	 * pools and is intended to make the new pool selection algorithm close to a superset
	 * of the old algorithm.
	 */
	TopNWithEachBaseToken int
	/**
	 * Given the topNWithEachBaseToken pools, takes the top N pools from the full list.
	 * E.g. for a WETH -> USDC swap, if topNWithEachBaseToken found WETH -0.05-> DAI,
	 * WETH -0.01-> DAI, WETH -0.05-> USDC, WETH -0.3-> USDC, a value of 2 would reduce
	 * this set to the top 2 pools from that full list.
	 */
	TopNWithBaseToken int
}

func DefaultRoutingConfigByChain(chainId base_entities.ChainId) AlphaRouterConfig {
	switch chainId {
	// Optimism
	case base_entities.OPTIMISM, base_entities.OPTIMISTIC_KOVAN:
		return AlphaRouterConfig{
			V2PoolSelection: ProtocolPoolSelection{
				TopN:                  3,
				TopNDirectSwaps:       1,
				TopNTokenInOut:        5,
				TopNSecondHop:         2,
				TopNWithEachBaseToken: 2,
				TopNWithBaseToken:     6,
			},
			V3PoolSelection: ProtocolPoolSelection{
				TopN:                  2,
				TopNDirectSwaps:       2,
				TopNTokenInOut:        2,
				TopNSecondHop:         1,
				TopNWithEachBaseToken: 3,
				TopNWithBaseToken:     3,
			},
			MaxSwapsPerPath:     3,
			MinSplits:           1,
			MaxSplits:           7,
			DistributionPercent: 10,
			ForceCrossProtocol:  false,
		}
		// Arbitrum calls have lower gas limits and tend to timeout more, which causes us to reduce the multicall
		// batch size and send more multicalls per quote. To reduce the amount of requests each quote sends, we
		// have to adjust the routing config so we explore fewer routes.
	case base_entities.ARBITRUM_ONE, base_entities.ARBITRUM_RINKEBY:
		return AlphaRouterConfig{
			V2PoolSelection: ProtocolPoolSelection{
				TopN:                  3,
				TopNDirectSwaps:       1,
				TopNTokenInOut:        5,
				TopNSecondHop:         2,
				TopNWithEachBaseToken: 2,
				TopNWithBaseToken:     6,
			},
			V3PoolSelection: ProtocolPoolSelection{
				TopN:                  2,
				TopNDirectSwaps:       2,
				TopNTokenInOut:        2,
				TopNSecondHop:         1,
				TopNWithEachBaseToken: 3,
				TopNWithBaseToken:     2,
			},
			MaxSwapsPerPath:     2,
			MinSplits:           1,
			MaxSplits:           7,
			DistributionPercent: 25,
			ForceCrossProtocol:  false,
		}
	default:
		return AlphaRouterConfig{
			Protocols: []base_entities.Protocol{
				base_entities.V2,
				base_entities.V3,
			},
			V2PoolSelection: ProtocolPoolSelection{
				TopN:                  3,
				TopNDirectSwaps:       1,
				TopNTokenInOut:        5,
				TopNSecondHop:         2,
				TopNWithEachBaseToken: 2,
				TopNWithBaseToken:     6,
			},
			V3PoolSelection: ProtocolPoolSelection{
				TopN:                  2,
				TopNDirectSwaps:       2,
				TopNTokenInOut:        3,
				TopNSecondHop:         1,
				TopNWithEachBaseToken: 3,
				TopNWithBaseToken:     5,
			},
			MaxSwapsPerPath:     3,
			MinSplits:           1,
			MaxSplits:           7,
			DistributionPercent: 5,
			ForceCrossProtocol:  false,
		}
	}
}
