package main

import (
	"fmt"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/cmd"
	v3 "github.com/coming-chat/intra-swap-core/providers/v3"
	"github.com/coming-chat/intra-swap-core/routers/alpha_router/config"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	entitiesV3 "github.com/daoleno/uniswapv3-sdk/entities"
	"math/big"
)

func main() {
	router, err := cmd.Ready(base_entities.BASE)
	if err != nil {
		return
	}

	token0str := "0x1dd2d631c92b1acdfcdd51a0f7145a50130050c4"
	token1str := "0x4200000000000000000000000000000000000006"
	tokenAccessor, err := router.TokenProvider.GetTokens([]string{token0str, token1str}, nil)
	if err != nil {
		return
	}
	token0 := tokenAccessor.GetTokenByAddress(token0str)
	token1 := tokenAccessor.GetTokenByAddress(token1str)
	token0Balance := entities.FromRawAmount(token0, big.NewInt(100))
	token1Balance := entities.FromRawAmount(token1, big.NewInt(10))

	poolAccessor, err := router.V3PoolProvider.GetPools([]v3.TokenPairs{
		{
			Token0:    token0,
			Token1:    token1,
			FeeAmount: constants.FeeHigh,
		},
	}, nil)
	if err != nil {
		return
	}
	pool, err := poolAccessor.GetPool(token0, token1, constants.FeeHigh)
	if err != nil {
		return
	}
	positionPreLiquidity, err := entitiesV3.NewPosition(
		pool.Pool,
		big.NewInt(1),
		-120,
		120,
	)
	swap, err := router.RouteToRatio(
		token0Balance,
		token1Balance,
		positionPreLiquidity,
		config.SwapAndAddConfig{
			MaxIterations: 6,
		},
		&config.SwapAndAddOptions{
			SwapOptions: config.SwapOptions{
				Deadline:          100,
				Recipient:         "0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B",
				SlippageTolerance: entities.NewPercent(big.NewInt(5), big.NewInt(10_000)),
			},
		},
		config.DefaultRoutingConfigByChain(base_entities.BASE),
	)
	if err != nil {
		return
	}
	fmt.Printf("%v", swap)
}
