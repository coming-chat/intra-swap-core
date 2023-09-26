package main

import (
	"fmt"
	"github.com/coming-chat/intra-swap-core/base_constant"
	"github.com/coming-chat/intra-swap-core/cmd"
	"github.com/coming-chat/intra-swap-core/routers/alpha_router/config"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"math/big"
	"time"
)

func main() {
	log := logrus.WithField("Cmd", "Quoter")
	router, err := cmd.Ready(base_constant.BASE, "swapbased", "")
	if err != nil {
		return
	}

	token0str := "0x4200000000000000000000000000000000000006" // weth 18
	token1str := "0xd9aAEc86B65D86f6A7B5B1b0c42FFA531710b6CA" // DAI 18
	tokenAccessor, err := router.TokenProvider.GetTokens(base_constant.BASE, []string{token0str, token1str}, nil)
	if err != nil {
		return
	}
	amountIn := new(big.Int).Mul(big.NewInt(1), big.NewInt(1).Exp(big.NewInt(10), big.NewInt(18), nil))
	token0 := tokenAccessor.GetTokenByAddress(token0str)
	token1 := tokenAccessor.GetTokenByAddress(token1str)
	token0Balance := entities.FromRawAmount(token0, amountIn)

	swap, err := router.Route(
		token0Balance,
		token1,
		entities.ExactInput,
		&config.SwapOptions{
			Recipient:         common.HexToAddress("0x8F18d6260b13B77FEddba9Ac5fe367d4a8c8cEC0"),
			SlippageTolerance: entities.NewPercent(big.NewInt(1), big.NewInt(100)),
			Deadline:          big.NewInt(time.Now().Add(10 * time.Second).Unix()),
			InputTokenPermit:  nil,
			SqrtPriceLimitX96: big.NewInt(0),
		},
		config.DefaultRoutingConfigByChain(base_constant.BASE),
	)
	if err != nil {
		log.Errorf("get Route err: %v", err)
		return
	}
	fmt.Printf("Quote: %s\n", swap.Quote.ToFixed(4))
	fmt.Printf("QuoteGasAdjusted: %s\n", swap.QuoteGasAdjusted.ToFixed(4))
	fmt.Printf("priceImapct: %s\n", swap.PriceImpact.ToFixed(4))
	fmt.Printf("%v", swap)
}
