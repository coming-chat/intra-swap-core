package build

import (
	"github.com/coming-chat/intra-swap-core/base_constant"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts/omni_swap"
	"github.com/coming-chat/intra-swap-core/dex"
	"github.com/coming-chat/intra-swap-core/routers"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
	selfCommon "github.com/gkirito/go-ethereum/common"
	"math/big"
)

// every routers.RouteWithValidQuote is Path split by percentage
// but the route pool/pair maybe the same dex(diff with routerAddres，factoryAddress, quoteAddress)
// so the trade should build to multi dex callData
// routers
//  10% | A -> B -> C -> D (v3)
//  30% | A -> C -> D      (v3)
//  55% | A -> D           (v2)
//  5%  | A -> E -> D      (v3)
// TODO support mix v2 & v3

func BuildTrade(
	tokenInCurrency entities.Currency,
	tokenOutCurrency entities.Currency,
	tradeType entities.TradeType,
	routeAmounts []routers.RouteWithValidQuote,
) (*base_entities.MTrade, error) {
	var swaps []*base_entities.Swap
	for _, routeAmount := range routeAmounts {
		var inputAmount, outputAmount *entities.CurrencyAmount
		switch tradeType {
		case entities.ExactInput:
			inputAmount = entities.FromFractionalAmount(
				tokenInCurrency,
				routeAmount.GetBaseRouteWithValidQuote().Amount.Numerator,
				routeAmount.GetBaseRouteWithValidQuote().Amount.Denominator,
			)
			outputAmount = entities.FromFractionalAmount(
				tokenOutCurrency,
				routeAmount.GetBaseRouteWithValidQuote().Quote.Numerator,
				routeAmount.GetBaseRouteWithValidQuote().Quote.Denominator,
			)

		case entities.ExactOutput:
			inputAmount = entities.FromFractionalAmount(
				tokenInCurrency,
				routeAmount.GetBaseRouteWithValidQuote().Quote.Numerator,
				routeAmount.GetBaseRouteWithValidQuote().Quote.Denominator,
			)

			outputAmount = entities.FromFractionalAmount(
				tokenOutCurrency,
				routeAmount.GetBaseRouteWithValidQuote().Amount.Numerator,
				routeAmount.GetBaseRouteWithValidQuote().Amount.Denominator,
			)
		}

		swaps = append(swaps, &base_entities.Swap{
			Route:        routeAmount.GetBaseRouteWithValidQuote().Route,
			InputAmount:  inputAmount,
			OutputAmount: outputAmount,
		})
	}
	return base_entities.NewMTrade(swaps, tradeType)
}

func BuildOmniSwapMethodParameters(
	chainId base_entities.ChainId,
	trade *base_entities.MTrade,
	swapConfig base_entities.SwapOptions,
) (omni_swap.ISoSoData, []omni_swap.LibSwapNormalizedSwapData, error) {
	soData := omni_swap.ISoSoData{
		Receiver:           selfCommon.Address(swapConfig.Recipient),
		SendingAssetId:     selfCommon.Address(trade.Swaps[0].Route.Input.Wrapped().Address),
		SourceChainId:      uint16(trade.Swaps[0].Route.Path[0].ChainId()),
		DestinationChainId: uint16(trade.Swaps[0].Route.Path[0].ChainId()),
		Amount:             trade.InputAmount().Quotient(),
		ReceivingAssetId:   selfCommon.Address(trade.Swaps[0].Route.Output.Wrapped().Address),
	}

	var (
		libSNSwapData []omni_swap.LibSwapNormalizedSwapData
		err           error
	)
	for _, swap := range trade.Swaps {
		for i, pool := range swap.Route.Pools {
			swapData := omni_swap.LibSwapNormalizedSwapData{
				CallTo:           pool.RouterAddress().Bytes(),
				ApproveTo:        pool.RouterAddress().Bytes(),
				SendingAssetId:   swap.Route.Path[i].Address.Bytes(),
				ReceivingAssetId: swap.Route.Path[i+1].Wrapped().Address.Bytes(),
			}
			swapData.FromAmount = big.NewInt(0)
			if i == 0 {
				swapData.FromAmount = swap.InputAmount.Quotient()
				if swap.InputAmount.Currency.IsNative() {
					swapData.SendingAssetId = common.HexToAddress(base_constant.EtherAddress).Bytes()
				}
			}
			if i == len(swap.Route.Pools)-1 && swap.OutputAmount.Currency.IsNative() {
				swapData.ReceivingAssetId = common.HexToAddress(base_constant.EtherAddress).Bytes()
			}
			swapData.CallData, err = packedCallData(chainId, trade, swap, i, swapConfig)
			if err != nil {
				return soData, nil, err
			}
			libSNSwapData = append(libSNSwapData, swapData)
		}
	}
	return soData, libSNSwapData, nil
}

func packedCallData(
	chainId base_entities.ChainId,
	trade *base_entities.MTrade,
	swap *base_entities.Swap,
	poolIndex int,
	swapConfig base_entities.SwapOptions,
) ([]byte, error) {
	amountIn := entities.FromRawAmount(swap.Route.Path[poolIndex], big.NewInt(0))
	amountOut := entities.FromRawAmount(swap.Route.Path[poolIndex+1], big.NewInt(0))
	if len(swap.Route.Pools) == poolIndex+1 {
		var err error
		switch trade.TradeType {
		case entities.ExactInput:
			amountOut, err = trade.MinimumAmountOut(swapConfig.SlippageTolerance, swap.OutputAmount)
		case entities.ExactOutput:
			amountIn, err = trade.MaximumAmountIn(swapConfig.SlippageTolerance, swap.InputAmount)
		}
		if err != nil {
			return nil, err
		}
	}
	if poolIndex == 0 {
		switch trade.TradeType {
		case entities.ExactInput:
			amountIn = swap.InputAmount
		case entities.ExactOutput:
			amountOut = swap.InputAmount
		}
	}
	return dex.RouterAddrDexMap[swap.Route.Pools[poolIndex].RouterAddress()].PackSwap(
		trade.TradeType,
		swap.Route.Path[poolIndex],
		swap.Route.Path[poolIndex+1],
		amountIn,
		amountOut,
		swap.Route.Pools[poolIndex],
		swapConfig,
	)
}
