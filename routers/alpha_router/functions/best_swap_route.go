package functions

import (
	"errors"
	"github.com/coming-chat/intra-swap-core/base_constant"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/providers"
	"github.com/coming-chat/intra-swap-core/routers"
	"github.com/coming-chat/intra-swap-core/routers/alpha_router/config"
	"github.com/coming-chat/intra-swap-core/routers/alpha_router/models"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"math/big"
	"sort"
)

type BestRouteResult struct {
	Quote                      *entities.CurrencyAmount
	QuoteGasAdjusted           *entities.CurrencyAmount
	EstimatedGasUsed           *big.Int
	EstimatedGasUsedUSD        *entities.CurrencyAmount
	EstimatedGasUsedQuoteToken *entities.CurrencyAmount
	Routes                     []routers.RouteWithValidQuote
}

func GetBestSwapRoute(
	amount *entities.CurrencyAmount,
	percents []int,
	routesWithValidQuotes []routers.RouteWithValidQuote,
	routeType entities.TradeType,
	chainId base_entities.ChainId,
	routingConfig config.AlphaRouterConfig,
	gasModel *models.GasModel[models.V3RouteWithValidQuote],
	usdGasTokens providers.UsdGasTokensProvider,
) (*BestRouteResult, error) {
	percentToQuotes := make(map[int][]routers.RouteWithValidQuote)
	for _, v := range routesWithValidQuotes {
		percentToQuotes[v.GetBaseRouteWithValidQuote().Percent] = append(percentToQuotes[v.GetBaseRouteWithValidQuote().Percent], v)
	}
	swapRoute, err := getBestSwapRouteBy(
		routeType,
		percentToQuotes,
		percents,
		chainId,
		func(rq routers.RouteWithValidQuote) *entities.CurrencyAmount {
			return rq.GetBaseRouteWithValidQuote().QuoteAdjustedForGas
		},
		routingConfig,
		gasModel,
		usdGasTokens,
	)
	if err != nil {
		return nil, err
	}

	totalAmount := entities.FromRawAmount(swapRoute.Routes[0].GetBaseRouteWithValidQuote().Amount.Currency, big.NewInt(0))
	for _, v := range swapRoute.Routes {
		totalAmount = totalAmount.Add(v.GetBaseRouteWithValidQuote().Amount)
	}
	missingAmount := amount.Subtract(totalAmount)
	if missingAmount.GreaterThan(base_constant.ZeroFraction) {
		swapRoute.Routes[len(swapRoute.Routes)-1].GetBaseRouteWithValidQuote().Amount =
			swapRoute.Routes[len(swapRoute.Routes)-1].GetBaseRouteWithValidQuote().Amount.Add(missingAmount)
	}
	return swapRoute, nil
}

func getBestSwapRouteBy(
	routeType entities.TradeType,
	percentToQuotes map[int][]routers.RouteWithValidQuote,
	percents []int,
	chainId base_entities.ChainId,
	by func(routeQuote routers.RouteWithValidQuote) *entities.CurrencyAmount,
	routingConfig config.AlphaRouterConfig,
	gasModel *models.GasModel[models.V3RouteWithValidQuote],
	usdGasTokens providers.UsdGasTokensProvider,
) (*BestRouteResult, error) {
	for _, v := range percentToQuotes {
		sort.Slice(v, func(i, j int) bool {
			if routeType == entities.ExactInput {
				return by(v[i]).GreaterThan(by(v[j]).Fraction)
			} else {
				return by(v[i]).LessThan(by(v[j]).Fraction)
			}
		})
	}

	var quoteCompFn func(a, b *entities.CurrencyAmount) bool
	if routeType == entities.ExactInput {
		quoteCompFn = func(a, b *entities.CurrencyAmount) bool {
			return a.GreaterThan(b.Fraction)
		}
	} else {
		quoteCompFn = func(a, b *entities.CurrencyAmount) bool {
			return a.LessThan(b.Fraction)
		}
	}
	sumFn := func(currencyAmounts []*entities.CurrencyAmount) *entities.CurrencyAmount {
		if len(currencyAmounts) == 0 {
			return entities.FromRawAmount(&entities.BaseCurrency{}, big.NewInt(0))
		}
		sum := new(entities.CurrencyAmount)
		*sum = *currencyAmounts[0]
		for _, v := range currencyAmounts[1:] {
			sum.Add(v)
		}
		return sum
	}

	var (
		bestQuote *entities.CurrencyAmount
		bestSwap  []routers.RouteWithValidQuote
	)

	type swapPer struct {
		Quote  *entities.CurrencyAmount
		Routes []routers.RouteWithValidQuote
	}
	//bestSwapsPerSplit := util.NewReverseHeap[swapPer](3, func(a, b util.Item[swapPer]) bool {
	//	return quoteCompFn(a.Value.Quote, b.Value.Quote)
	//})

	if v, ok := percentToQuotes[100]; ok && len(v) > 0 && routingConfig.MinSplits <= 1 && !routingConfig.ForceCrossProtocol {
		bestQuote = by(percentToQuotes[100][0])
		bestSwap = append(bestSwap, percentToQuotes[100][0])

		//for i, routeWithQuote := range percentToQuotes[100] {
		//	if i > 4 {
		//		break
		//	}
		//bestSwapsPerSplit.Push(swapPer{
		//	Quote:  by(routeWithQuote),
		//	Routes: []providers.RouteWithValidQuote{routeWithQuote},
		//})
		//}
	}
	type queueItem struct {
		PercentIndex     int
		CurRoutes        []routers.RouteWithValidQuote
		RemainingPercent int
		Special          bool
	}
	var queue []queueItem

	for i, percent := range percents {
		quotes, ok := percentToQuotes[percent]
		if !ok || len(quotes) == 0 {
			continue
		}
		queue = append(queue, queueItem{
			PercentIndex:     i,
			CurRoutes:        []routers.RouteWithValidQuote{quotes[0]},
			RemainingPercent: 100 - percent,
			Special:          false,
		})

		if len(quotes) == 1 {
			continue
		}
		queue = append(queue, queueItem{
			PercentIndex:     i,
			CurRoutes:        []routers.RouteWithValidQuote{quotes[1]},
			RemainingPercent: 100 - percent,
			Special:          false,
		})
	}

	for splits := 1; len(queue) > 0; {
		//heap.Init(bestSwapsPerSplit)

		// Size of the queue at this point is the number of potential routes we are investigating for the given number of splits.
		layer := len(queue)
		splits++

		// If we didn't improve our quote by adding another split, very unlikely to improve it by splitting more after that.
		if splits >= 3 && len(bestSwap) > 0 && len(bestSwap) < splits-1 {
			break
		}

		if splits > routingConfig.MaxSplits {
			break
		}

		for layer > 0 {
			layer--
			if len(queue) == 0 {
				break
			}
			item := queue[0]
			queue = queue[1:]

			for i := item.PercentIndex; i >= 0; i-- {
				percentA := percents[i]
				if percentA > item.RemainingPercent {
					continue
				}

				// At some point the amount * percentage is so small that the quoter is unable to get
				// a quote. In this case there could be no quotes for that percentage.
				candidateRoutesA, ok := percentToQuotes[percentA]
				if !ok || len(candidateRoutesA) == 0 {
					continue
				}

				routeWithQuoteA, err := findFirstRouteNotUsingUsedPools(item.CurRoutes, candidateRoutesA, routingConfig.ForceCrossProtocol)
				if err != nil {
					continue
				}

				remainingPercentNew := item.RemainingPercent - percentA
				curRoutesNew := append(item.CurRoutes, routeWithQuoteA)

				// If we've found a route combination that uses all 100%, and it has at least minSplits, update our best route.
				if remainingPercentNew == 0 && splits >= routingConfig.MinSplits {
					var quotesNew []*entities.CurrencyAmount
					for _, cRN := range curRoutesNew {
						quotesNew = append(quotesNew, by(cRN))
					}
					quoteNew := sumFn(quotesNew)

					gasCostL1QuoteToken := entities.FromRawAmount(quoteNew.Currency, big.NewInt(0))

					//if _, hasL1Fee := base_constant.HasL1Fee[chainId]; hasL1Fee {
					//	var curRoutesNewV3 []*models.V3RouteWithValidQuote
					//	for _, cr := range curRoutesNew {
					//		if cr.Protocol() == base_entities.V2 {
					//			break
					//		}
					//		curRoutesNewV3 = append(curRoutesNewV3, cr.(*models.V3RouteWithValidQuote))
					//	}
					//
					//	if gasModel == nil || len(curRoutesNewV3) != len(curRoutesNew) {
					//		return nil, errors.New("can't compute L1 gas fees")
					//	}
					//
					//	gasCostL1, err := gasModel.CalculateL1GasFees(curRoutesNewV3)
					//	if err != nil {
					//		return nil, err
					//	}
					//	gasCostL1QuoteToken = gasCostL1.GasCostL1QuoteToken
					//
					//}

					var quoteAfterL1Adjust *entities.CurrencyAmount
					if routeType == entities.ExactInput {
						quoteAfterL1Adjust = quoteNew.Subtract(gasCostL1QuoteToken)
					} else {
						quoteAfterL1Adjust = quoteNew.Add(gasCostL1QuoteToken)
					}

					//bestSwapsPerSplit.Push(swapPer{
					//	Quote:  quoteAfterL1Adjust,
					//	Routes: curRoutesNew,
					//})

					if bestQuote == nil || quoteCompFn(quoteAfterL1Adjust, bestQuote) {
						bestQuote = quoteAfterL1Adjust
						bestSwap = curRoutesNew
						if item.Special {
							// test
						}
					}
				} else {
					queue = append(queue, queueItem{
						CurRoutes:        curRoutesNew,
						RemainingPercent: remainingPercentNew,
						PercentIndex:     i,
						Special:          item.Special,
					})
				}

			}

		}
	}

	if len(bestSwap) == 0 {
		return nil, errors.New("could not find a valid swap")
	}

	var (
		bestSwapAllGas   []*entities.CurrencyAmount
		estimatedGasUsed = big.NewInt(0)
	)
	for _, bs := range bestSwap {
		bestSwapAllGas = append(bestSwapAllGas, bs.GetBaseRouteWithValidQuote().QuoteAdjustedForGas)
		//estimatedGasUsed = estimatedGasUsed.Add(estimatedGasUsed, bs.GetBaseRouteWithValidQuote().GasEstimate)
	}
	quoteGasAdjusted := sumFn(bestSwapAllGas)

	// this calculates the base gas used
	// if on L1, its the estimated gas used based on hops and ticks across all the routes
	// if on L2, its the gas used on the L2 based on hops and ticks across all the routes
	token, err := usdGasTokens.GetTokensByChain(chainId)
	if err != nil {
		// Each route can use a different stablecoin to account its gas costs.
		// They should all be pegged, and this is just an estimate, so we do a merge
		// to an arbitrary stable.
		return nil, err
	}
	usdToken := token[0]
	//usdTokenDecimals := usdToken.Decimals()

	// if on L2, calculate the L1 security fee
	gasCostsL1ToL2 := &models.L1ToL2GasCosts{
		GasUsedL1:    big.NewInt(0),
		GasCostL1USD: entities.FromRawAmount(usdToken, big.NewInt(0)),
		GasCostL1QuoteToken: entities.FromRawAmount(
			bestSwap[0].GetBaseRouteWithValidQuote().QuoteToken,
			big.NewInt(0),
		),
	}

	// If swapping on an L2 that includes a L1 security fee, calculate the fee and include it in the gas adjusted quotes
	//if _, hasL1Fee := base_constant.HasL1Fee[chainId]; hasL1Fee {
	//	// ensure the gasModel exists and that the swap route is a v3 only route
	//	var bsV3 []models.V3RouteWithValidQuote
	//	for _, bs := range bestSwap {
	//		if bs.Protocol() == base_entities.V2 {
	//			break
	//		}
	//		bsV3 = append(bsV3, bs.(models.V3RouteWithValidQuote))
	//	}
	//
	//	if gasModel == nil || len(bsV3) != len(bestSwap) {
	//		return nil, errors.New("can't compute L1 gas fees")
	//	}
	//	//var err error
	//	//gasCostsL1ToL2, err = gasModel.CalculateL1GasFees(bsV3)
	//	//if err != nil {
	//	//	return nil, err
	//	//}
	//
	//}

	var (
		quotes []*entities.CurrencyAmount //estimatedGasUsedUSDs, gasCostInToken,
	)
	for _, v := range bestSwap {
		//gasCostInToken = append(gasCostInToken, v.GetBaseRouteWithValidQuote().GasCostInToken)
		quotes = append(quotes, v.GetBaseRouteWithValidQuote().Quote)
		//decimalsDiff := usdTokenDecimals - v.GetBaseRouteWithValidQuote().GasCostInUSD.Currency.Decimals()
		//if decimalsDiff == 0 {
		//	estimatedGasUsedUSDs = append(estimatedGasUsedUSDs, entities.FromRawAmount(usdToken, v.GetBaseRouteWithValidQuote().GasCostInUSD.Quotient()))
		//	continue
		//}
		//estimatedGasUsedUSDs = append(estimatedGasUsedUSDs, entities.FromRawAmount(
		//	usdToken,
		//	new(big.Int).Mul(v.GetBaseRouteWithValidQuote().GasCostInUSD.Quotient(), new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimalsDiff)), nil)),
		//))
	}
	//estimatedGasUsedUSD := sumFn(estimatedGasUsedUSDs)
	//
	//// if they are different usd pools, convert to the usdToken
	//if estimatedGasUsedUSD.Currency != gasCostsL1ToL2.GasCostL1USD.Currency {
	//	decimalsDiff := usdTokenDecimals - gasCostsL1ToL2.GasCostL1USD.Currency.Decimals()
	//	estimatedGasUsedUSD = estimatedGasUsedUSD.Add(
	//		entities.FromRawAmount(
	//			usdToken,
	//			new(big.Int).Mul(gasCostsL1ToL2.GasCostL1USD.Quotient(), new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimalsDiff)), nil)),
	//		),
	//	)
	//} else {
	//	estimatedGasUsedUSD = estimatedGasUsedUSD.Add(gasCostsL1ToL2.GasCostL1USD)
	//}
	//
	//estimatedGasUsedQuoteToken := sumFn(gasCostInToken).Add(gasCostsL1ToL2.GasCostL1QuoteToken)
	quote := sumFn(quotes)

	// Adjust the quoteGasAdjusted for the l1 fee
	if routeType == entities.ExactInput {
		quoteGasAdjusted = quoteGasAdjusted.Subtract(gasCostsL1ToL2.GasCostL1QuoteToken)
	} else {
		quoteGasAdjusted = quoteGasAdjusted.Add(gasCostsL1ToL2.GasCostL1QuoteToken)
	}

	sort.Slice(bestSwap, func(i, j int) bool {
		return bestSwap[j].GetBaseRouteWithValidQuote().Amount.GreaterThan(bestSwap[i].GetBaseRouteWithValidQuote().Amount.Fraction)
	})

	return &BestRouteResult{
		Quote:            quote,
		QuoteGasAdjusted: quoteGasAdjusted,
		EstimatedGasUsed: estimatedGasUsed,
		//EstimatedGasUsedQuoteToken: estimatedGasUsedUSD,
		//EstimatedGasUsedUSD:        estimatedGasUsedQuoteToken,
		Routes: bestSwap,
	}, nil
}

func findFirstRouteNotUsingUsedPools(
	usedRoutes []routers.RouteWithValidQuote,
	candidateRouteQuotes []routers.RouteWithValidQuote,
	forceCrossProtocol bool,
) (routers.RouteWithValidQuote, error) {
	var (
		poolAddressSet = make(map[string]struct{})
		protocolsSet   = make(map[base_entities.Protocol]struct{})
	)
	for _, v := range usedRoutes {
		for _, pa := range v.GetBaseRouteWithValidQuote().PoolAddresses {
			poolAddressSet[pa] = struct{}{}
			protocolsSet[v.Protocol()] = struct{}{}
		}
	}
	for _, c := range candidateRouteQuotes {
		for _, padd := range c.GetBaseRouteWithValidQuote().PoolAddresses {
			if _, ok := poolAddressSet[padd]; ok {
				goto next
			}
		}

		if forceCrossProtocol && len(poolAddressSet) == 1 {
			if _, ok := protocolsSet[c.Protocol()]; ok {
				continue
			}
		}
		return c, nil
	next:
	}
	return models.V2RouteWithValidQuote{}, errors.New("not found the RouteQuotes")
}
