package main

import (
	"fmt"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/cmd"
	"github.com/coming-chat/intra-swap-core/routers/alpha_router/config"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"math/big"
)

func main() {
	router, err := cmd.Ready(base_entities.BASE)
	if err != nil {
		return
	}

	token0str := "0xd9aAEc86B65D86f6A7B5B1b0c42FFA531710b6CA"
	token1str := "0x4200000000000000000000000000000000000006"
	tokenAccessor, err := router.TokenProvider.GetTokens([]string{token0str, token1str}, nil)
	if err != nil {
		return
	}
	token0 := tokenAccessor.GetTokenByAddress(token0str)
	token1 := tokenAccessor.GetTokenByAddress(token1str)
	token0Balance := entities.FromRawAmount(token0, new(big.Int).Mul(big.NewInt(10000), big.NewInt(1).Exp(big.NewInt(10), big.NewInt(6), nil)))
	//token1Balance := entities.FromRawAmount(token1, big.NewInt(10))

	swap, err := router.Route(
		token0Balance,
		token1,
		entities.ExactInput,
		&config.SwapOptions{
			Recipient: "0x0",
		},
		config.DefaultRoutingConfigByChain(base_entities.BASE),
	)
	if err != nil {
		return
	}
	fmt.Printf("Quote: %s\n", swap.Route[0].GetBaseRouteWithValidQuote().Quote.ToFixed(4))
	fmt.Printf("Amount: %s\n", swap.Route[0].GetBaseRouteWithValidQuote().Amount.ToFixed(4))
	fmt.Printf("%v", swap)
}
