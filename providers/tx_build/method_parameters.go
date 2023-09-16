package tx_build

import (
	"errors"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/daoleno/uniswap-sdk-core/entities"
	entitiesV3 "github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/periphery"
	"github.com/daoleno/uniswapv3-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
	entitiesV2 "github.com/vaulverin/uniswapv2-sdk/entities"
	"github.com/vaulverin/uniswapv2-sdk/router"
)

func BuildTrade(tokenInCurrency *entities.Currency,
	tokenOutCurrency *entities.Currency,
	tradeType entities.TradeType,
	routeAmounts []base_entities.RouteWithValidQuote,
) []base_entities.Trade {

}

func BuildSwapMethodParameters(
	trade []base_entities.Trade,
	swapConfig base_entities.SwapOptions,
) (*utils.MethodParameters, error) {
	switch trade.(type) {
	case *entitiesV3.Trade:
		return periphery.SwapCallParameters([]*entitiesV3.Trade{trade.(*entitiesV3.Trade)}, &periphery.SwapOptions{
			Recipient:         common.HexToAddress(swapConfig.Recipient),
			SlippageTolerance: swapConfig.SlippageTolerance,
			Deadline:          swapConfig.Deadline,
			InputTokenPermit:  swapConfig.InputTokenPermit,
		})
	case *entitiesV2.Trade:
		value, data, err := router.SwapCallParametersPacked(trade.(*entitiesV2.Trade), router.TradeOptions{
			Recipient: common.HexToAddress(swapConfig.Recipient),
		})
		if err != nil {
			return nil, err
		}
		return &utils.MethodParameters{
			Value:    value,
			Calldata: data,
		}, nil
	default:
		return nil, errors.New("unsupported method parameters")
	}

}
