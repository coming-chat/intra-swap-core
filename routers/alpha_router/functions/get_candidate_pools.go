package functions

import (
	"github.com/coming-chat/intra-swap-core/base_constant"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/providers"
	providersConfig "github.com/coming-chat/intra-swap-core/providers/config"
	v2 "github.com/coming-chat/intra-swap-core/providers/v2"
	v3 "github.com/coming-chat/intra-swap-core/providers/v3"
	"github.com/coming-chat/intra-swap-core/routers/alpha_router/config"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/shopspring/decimal"
	"sort"
)

type PoolId string

type CandidatePoolsBySelectionCriteria struct {
	Protocol   base_entities.Protocol
	Selections struct {
		TopByBaseWithTokenIn            []PoolId
		TopByBaseWithTokenOut           []PoolId
		TopByDirectSwapPool             []PoolId
		TopByEthQuoteTokenPool          []PoolId
		TopByTVL                        []PoolId
		TopByTVLUsingTokenIn            []PoolId
		TopByTVLUsingTokenOut           []PoolId
		TopByTVLUsingTokenInSecondHops  []PoolId
		TopByTVLUsingTokenOutSecondHops []PoolId
	}
}

type V2GetCandidatePoolsParams struct {
	TokenIn                  *entities.Token
	TokenOut                 *entities.Token
	RouteType                entities.TradeType
	RoutingConfig            config.AlphaRouterConfig
	IndexerProvider          v2.IndexerProvider
	TokenProvider            providers.TokenProvider
	PoolProvider             v2.PoolProvider
	BlockedTokenListProvider providers.TokenListProvider
	WNprovider               providers.WrappedNativeCurrencyProvider
	ChainId                  base_entities.ChainId
}

type V3GetCandidatePoolsParams struct {
	TokenIn                  *entities.Token
	TokenOut                 *entities.Token
	RouteType                entities.TradeType
	RoutingConfig            config.AlphaRouterConfig
	IndexerProvider          v3.IndexerProvider
	TokenProvider            providers.TokenProvider
	PoolProvider             v3.PoolProvider
	WNProvider               providers.WrappedNativeCurrencyProvider
	BlockedTokenListProvider providers.TokenListProvider
	ChainId                  base_entities.ChainId
}

func GetV3CandidatePools(params V3GetCandidatePoolsParams) (
	poolAccessor v3.PoolAccessor,
	candidatePools CandidatePoolsBySelectionCriteria,
	err error,
) {
	allPools, err := params.IndexerProvider.GetPools(params.ChainId, params.TokenIn, params.TokenOut, &providersConfig.Config{
		BlockNumber: params.RoutingConfig.BlockNumber,
	})
	if err != nil {
		return nil, CandidatePoolsBySelectionCriteria{}, err
	}

	var filteredPools []v3.IndexerPool
	if params.BlockedTokenListProvider != nil {
		for _, pool := range allPools {
			if params.BlockedTokenListProvider != nil {
				token0InBlocklist, _ := params.BlockedTokenListProvider.GetTokenByAddress(params.ChainId, pool.Token0.Id)
				token1InBlocklist, _ := params.BlockedTokenListProvider.GetTokenByAddress(params.ChainId, pool.Token1.Id)
				if token0InBlocklist != nil || token1InBlocklist != nil {
					continue
				}
			}
			filteredPools = append(filteredPools, pool)
		}
	} else {
		filteredPools = allPools
	}

	sort.Slice(filteredPools, func(i, j int) bool {
		return filteredPools[i].TvlUSD.LessThan(filteredPools[j].TvlUSD)
	})
	poolAddressesSoFar := make(map[string]struct{})
	addToAddressSet := func(pools []v3.IndexerPool) {
		for _, p := range pools {
			poolAddressesSoFar[p.Id] = struct{}{}
		}
	}
	baseTokens := baseTokensByChain[params.ChainId]
	var (
		topByBaseWithTokenIn, topByBaseWithTokenOut []v3.IndexerPool
	)
	for _, baseToken := range baseTokens {
		var (
			baseWithTokenInPools, baseWithTokenOutPools []v3.IndexerPool
		)
		for _, pool := range filteredPools {
			switch {
			case pool.Token0.Id == baseToken.Address.String() && pool.Token1.Id == params.TokenIn.Address.String():
				fallthrough
			case pool.Token1.Id == baseToken.Address.String() && pool.Token0.Id == params.TokenIn.Address.String():
				baseWithTokenInPools = append(baseWithTokenInPools, pool)
			case pool.Token0.Id == baseToken.Address.String() && pool.Token1.Id == params.TokenOut.Address.String():
				fallthrough
			case pool.Token1.Id == baseToken.Address.String() && pool.Token0.Id == params.TokenOut.Address.String():
				baseWithTokenOutPools = append(baseWithTokenOutPools, pool)
			}
		}
		sort.Slice(baseWithTokenInPools, func(i, j int) bool {
			return baseWithTokenInPools[i].TvlUSD.LessThan(baseWithTokenInPools[j].TvlUSD)
		})
		sort.Slice(baseWithTokenOutPools, func(i, j int) bool {
			return baseWithTokenOutPools[i].TvlUSD.LessThan(baseWithTokenOutPools[j].TvlUSD)
		})
		if len(baseWithTokenInPools) > params.RoutingConfig.V3PoolSelection.TopNWithEachBaseToken {
			baseWithTokenInPools = baseWithTokenInPools[:params.RoutingConfig.V3PoolSelection.TopNWithEachBaseToken]
		}
		if len(baseWithTokenOutPools) > params.RoutingConfig.V3PoolSelection.TopNWithEachBaseToken {
			baseWithTokenOutPools = baseWithTokenOutPools[:params.RoutingConfig.V3PoolSelection.TopNWithEachBaseToken]
		}
		topByBaseWithTokenIn = append(topByBaseWithTokenIn, baseWithTokenInPools...)
		topByBaseWithTokenOut = append(topByBaseWithTokenOut, baseWithTokenOutPools...)
	}
	sort.Slice(topByBaseWithTokenIn, func(i, j int) bool {
		return topByBaseWithTokenIn[i].TvlUSD.LessThan(topByBaseWithTokenIn[j].TvlUSD)
	})
	sort.Slice(topByBaseWithTokenOut, func(i, j int) bool {
		return topByBaseWithTokenOut[i].TvlUSD.LessThan(topByBaseWithTokenOut[j].TvlUSD)
	})
	if len(topByBaseWithTokenIn) > params.RoutingConfig.V3PoolSelection.TopNWithBaseToken {
		topByBaseWithTokenIn = topByBaseWithTokenIn[:params.RoutingConfig.V3PoolSelection.TopNWithBaseToken]
	}
	if len(topByBaseWithTokenOut) > params.RoutingConfig.V3PoolSelection.TopNWithBaseToken {
		topByBaseWithTokenOut = topByBaseWithTokenOut[:params.RoutingConfig.V3PoolSelection.TopNWithBaseToken]
	}

	var top2DirectSwapPool []v3.IndexerPool
	for _, pool := range filteredPools {
		if _, ok := poolAddressesSoFar[pool.Id]; !ok &&
			((pool.Token0.Id == params.TokenIn.Address.String() && pool.Token1.Id == params.TokenOut.Address.String()) ||
				(pool.Token0.Id == params.TokenOut.Address.String() && pool.Token1.Id == params.TokenIn.Address.String())) {
			top2DirectSwapPool = append(top2DirectSwapPool, pool)
		}
	}
	if len(top2DirectSwapPool) > params.RoutingConfig.V3PoolSelection.TopNDirectSwaps {
		top2DirectSwapPool = top2DirectSwapPool[:params.RoutingConfig.V3PoolSelection.TopNDirectSwaps]
	}
	if len(top2DirectSwapPool) == 0 && params.RoutingConfig.V3PoolSelection.TopNDirectSwaps > 0 {
		feeAmounts := []constants.FeeAmount{
			constants.FeeHigh, constants.FeeMedium, constants.FeeLow, constants.FeeLowest,
		}
		//Fake pool
		for _, v := range feeAmounts {
			address, token0, token1, err := params.PoolProvider.GetPoolAddress(params.TokenIn, params.TokenOut, v)
			if err != nil {
				return nil, candidatePools, err
			}
			top2DirectSwapPool = append(top2DirectSwapPool, v3.IndexerPool{
				Id:        address,
				FeeTier:   v,
				Liquidity: "10000",
				Token0:    struct{ Id string }{Id: token0.Address.String()},
				Token1:    struct{ Id string }{Id: token1.Address.String()},
				TvlETH:    decimal.NewFromInt(10000),
				TvlUSD:    decimal.NewFromInt(10000),
			})
		}
	}
	addToAddressSet(top2DirectSwapPool)
	wrappedNativeToken, err := params.WNProvider.GetTokenByChain(params.ChainId)
	if err != nil {
		return poolAccessor, candidatePools, err
	}
	// Main reason we need this is for gas estimates, only needed if token out is not native.
	// We don't check the seen address set because if we've already added pools for getting native quotes
	// theres no need to add more.
	var top2EthQuoteTokenPool []v3.IndexerPool
	if (wrappedNativeToken.Symbol() == base_constant.WrappedNativeCurrency[base_constant.MAINNET].Symbol() &&
		params.TokenOut.Symbol() != "WETH" &&
		params.TokenOut.Symbol() != "WETH9" &&
		params.TokenOut.Symbol() != "ETH") ||
		(wrappedNativeToken.Symbol() == base_constant.WMATIC_POLYGON.Symbol() &&
			params.TokenOut.Symbol() != "MATIC" &&
			params.TokenOut.Symbol() != "WMATIC") {
		for _, pool := range filteredPools {
			switch {
			case params.RouteType == entities.ExactInput && ((pool.Token0.Id == wrappedNativeToken.Address.String() && pool.Token1.Id == params.TokenOut.Address.String()) ||
				(pool.Token1.Id == wrappedNativeToken.Address.String() && pool.Token0.Id == params.TokenOut.Address.String())):
				fallthrough
			case params.RouteType == entities.ExactOutput && ((pool.Token0.Id == wrappedNativeToken.Address.String() && pool.Token1.Id == params.TokenIn.Address.String()) ||
				(pool.Token1.Id == wrappedNativeToken.Address.String() && pool.Token0.Id == params.TokenIn.Address.String())):
				top2EthQuoteTokenPool = append(top2EthQuoteTokenPool, pool)
			default:
				continue
			}
		}
		if len(top2EthQuoteTokenPool) > 1 {
			top2EthQuoteTokenPool = top2EthQuoteTokenPool[:1]
		}
	}
	addToAddressSet(top2EthQuoteTokenPool)

	var topByTVL []v3.IndexerPool
	for _, pool := range filteredPools {
		if _, has := poolAddressesSoFar[pool.Id]; !has {
			topByTVL = append(topByTVL, pool)
		}
	}
	if len(topByTVL) > params.RoutingConfig.V3PoolSelection.TopN {
		topByTVL = topByTVL[:params.RoutingConfig.V3PoolSelection.TopN]
	}
	addToAddressSet(topByTVL)

	var topByTVLUsingTokenIn []v3.IndexerPool
	for _, pool := range filteredPools {
		if _, has := poolAddressesSoFar[pool.Id]; !has && (pool.Token0.Id == params.TokenIn.Address.String() ||
			pool.Token1.Id == params.TokenIn.Address.String()) {
			topByTVLUsingTokenIn = append(topByTVLUsingTokenIn, pool)
		}
	}
	if len(topByTVLUsingTokenIn) > params.RoutingConfig.V3PoolSelection.TopNTokenInOut {
		topByTVLUsingTokenIn = topByTVLUsingTokenIn[:params.RoutingConfig.V3PoolSelection.TopNTokenInOut]
	}
	addToAddressSet(topByTVLUsingTokenIn)

	var topByTVLUsingTokenOut []v3.IndexerPool
	for _, pool := range filteredPools {
		if _, has := poolAddressesSoFar[pool.Id]; !has && (pool.Token0.Id == params.TokenOut.Address.String() ||
			pool.Token1.Id == params.TokenOut.Address.String()) {
			topByTVLUsingTokenOut = append(topByTVLUsingTokenOut, pool)
		}
	}
	if len(topByTVLUsingTokenOut) > params.RoutingConfig.V3PoolSelection.TopNTokenInOut {
		topByTVLUsingTokenOut = topByTVLUsingTokenOut[:params.RoutingConfig.V3PoolSelection.TopNTokenInOut]
	}
	addToAddressSet(topByTVLUsingTokenOut)

	var (
		topByTVLUsingTokenInSecondHops    []v3.IndexerPool
		topByTVLUsingTokenInSecondHopsMap = make(map[string]struct{})
	)
	for _, pool := range topByTVLUsingTokenIn {
		var (
			secondHopId    string
			secondHopPools []v3.IndexerPool
		)
		if params.TokenIn.Address.String() == pool.Token0.Id {
			secondHopId = pool.Token1.Id
		} else {
			secondHopId = pool.Token0.Id
		}
		for _, filterPool := range filteredPools {
			if _, has := poolAddressesSoFar[pool.Id]; !has && (filterPool.Token0.Id == secondHopId || filterPool.Token1.Id == secondHopId) {
				secondHopPools = append(secondHopPools, filterPool)
			}
		}
		if len(secondHopPools) > params.RoutingConfig.V3PoolSelection.TopNSecondHop {
			secondHopPools = secondHopPools[:params.RoutingConfig.V3PoolSelection.TopNSecondHop]
		}
		for _, v := range secondHopPools {
			if _, has := topByTVLUsingTokenInSecondHopsMap[v.Id]; has {
				continue
			}
			topByTVLUsingTokenInSecondHopsMap[v.Id] = struct{}{}
			topByTVLUsingTokenInSecondHops = append(topByTVLUsingTokenInSecondHops, v)
		}
	}
	sort.Slice(topByTVLUsingTokenInSecondHops, func(i, j int) bool {
		return topByTVLUsingTokenInSecondHops[i].TvlUSD.LessThan(topByTVLUsingTokenInSecondHops[j].TvlUSD)
	})
	if len(topByTVLUsingTokenInSecondHops) > params.RoutingConfig.V3PoolSelection.TopNSecondHop {
		topByTVLUsingTokenInSecondHops = topByTVLUsingTokenInSecondHops[:params.RoutingConfig.V3PoolSelection.TopNSecondHop]
	}
	addToAddressSet(topByTVLUsingTokenInSecondHops)

	var (
		topByTVLUsingTokenOutSecondHops    []v3.IndexerPool
		topByTVLUsingTokenOutSecondHopsMap = make(map[string]struct{})
	)
	for _, pool := range topByTVLUsingTokenIn {
		var (
			secondHopId    string
			secondHopPools []v3.IndexerPool
		)
		if params.TokenOut.Address.String() == pool.Token0.Id {
			secondHopId = pool.Token1.Id
		} else {
			secondHopId = pool.Token0.Id
		}
		for _, filterPool := range filteredPools {
			if _, has := poolAddressesSoFar[pool.Id]; !has && (filterPool.Token0.Id == secondHopId || filterPool.Token1.Id == secondHopId) {
				secondHopPools = append(secondHopPools, filterPool)
			}
		}
		if len(secondHopPools) > params.RoutingConfig.V3PoolSelection.TopNSecondHop {
			secondHopPools = secondHopPools[:params.RoutingConfig.V3PoolSelection.TopNSecondHop]
		}
		for _, v := range secondHopPools {
			if _, has := topByTVLUsingTokenOutSecondHopsMap[v.Id]; has {
				continue
			}
			topByTVLUsingTokenOutSecondHopsMap[v.Id] = struct{}{}
			topByTVLUsingTokenOutSecondHops = append(topByTVLUsingTokenOutSecondHops, v)
		}
	}
	sort.Slice(topByTVLUsingTokenOutSecondHops, func(i, j int) bool {
		return topByTVLUsingTokenOutSecondHops[i].TvlUSD.LessThan(topByTVLUsingTokenOutSecondHops[j].TvlUSD)
	})
	if len(topByTVLUsingTokenOutSecondHops) > params.RoutingConfig.V3PoolSelection.TopNSecondHop {
		topByTVLUsingTokenOutSecondHops = topByTVLUsingTokenOutSecondHops[:params.RoutingConfig.V3PoolSelection.TopNSecondHop]
	}
	addToAddressSet(topByTVLUsingTokenOutSecondHops)

	indexerPools := append(topByBaseWithTokenIn, topByBaseWithTokenOut...)
	indexerPools = append(indexerPools, top2DirectSwapPool...)
	indexerPools = append(indexerPools, top2EthQuoteTokenPool...)
	indexerPools = append(indexerPools, topByTVL...)
	indexerPools = append(indexerPools, topByTVLUsingTokenIn...)
	indexerPools = append(indexerPools, topByTVLUsingTokenOut...)
	indexerPools = append(indexerPools, topByTVLUsingTokenInSecondHops...)
	indexerPools = append(indexerPools, topByTVLUsingTokenOutSecondHops...)

	//indexPools is not unique
	var (
		tokenAddresses    []string
		tokenAddressesMap = make(map[string]struct{})
	)
	for _, pool := range indexerPools {
		if _, has := tokenAddressesMap[pool.Token0.Id]; !has {
			tokenAddressesMap[pool.Token0.Id] = struct{}{}
			tokenAddresses = append(tokenAddresses, pool.Token0.Id)
		}
		if _, has := tokenAddressesMap[pool.Token1.Id]; !has {
			tokenAddressesMap[pool.Token1.Id] = struct{}{}
			tokenAddresses = append(tokenAddresses, pool.Token1.Id)
		}
	}

	tokenAccessor, err := params.TokenProvider.GetTokens(params.ChainId, tokenAddresses, &providersConfig.Config{
		BlockNumber: params.RoutingConfig.BlockNumber,
	})
	if err != nil {
		return nil, CandidatePoolsBySelectionCriteria{}, err
	}

	var (
		tokenPairs      []v3.TokenPairs
		indexerPoolsMap = make(map[string]struct{})
	)
	for _, pool := range indexerPools {
		if _, has := indexerPoolsMap[pool.Id]; has {
			continue
		}
		indexerPoolsMap[pool.Id] = struct{}{}
		token0 := tokenAccessor.GetTokenByAddress(pool.Token0.Id)
		token1 := tokenAccessor.GetTokenByAddress(pool.Token1.Id)
		if token0 == nil || token1 == nil {
			continue
		}
		tokenPairs = append(tokenPairs, v3.TokenPairs{
			Token0:         token0,
			Token1:         token1,
			FeeAmount:      pool.FeeTier,
			PairAddress:    pool.Id,
			RouterAddress:  pool.RouterAddress,
			QuoteAddress:   pool.QuoteAddress,
			FactoryAddress: pool.FactoryAddress,
		})
	}
	poolAccessor, err = params.PoolProvider.GetPools(tokenPairs, nil)
	return
}

func GetV2CandidatePools(params V2GetCandidatePoolsParams) (
	poolAccessor v2.PoolAccessor,
	candidatePools CandidatePoolsBySelectionCriteria,
	err error,
) {
	allPools, err := params.IndexerProvider.GetPools(params.ChainId, params.TokenIn, params.TokenOut, &providersConfig.Config{
		BlockNumber: params.RoutingConfig.BlockNumber,
	})
	if err != nil {
		return nil, CandidatePoolsBySelectionCriteria{}, err
	}

	var filteredPools []v2.IndexerPool
	if params.BlockedTokenListProvider != nil {
		for _, pool := range allPools {
			if params.BlockedTokenListProvider != nil {
				token0InBlockList, _ := params.BlockedTokenListProvider.GetTokenByAddress(params.ChainId, pool.Token0.Id)
				token1InBlockList, _ := params.BlockedTokenListProvider.GetTokenBySymbol(params.ChainId, pool.Token1.Id)
				if token0InBlockList != nil || token1InBlockList != nil {
					continue
				}
			}
			filteredPools = append(filteredPools, pool)
		}
	} else {
		filteredPools = allPools
	}
	sort.Slice(filteredPools, func(i, j int) bool {
		return filteredPools[i].Reserve.LessThan(filteredPools[j].Reserve)
	})

	poolAddressesSoFar := make(map[string]struct{})
	addToAddressSet := func(pools []v2.IndexerPool) {
		for _, p := range pools {
			poolAddressesSoFar[p.Id] = struct{}{}
		}
	}
	baseTokens := baseTokensByChain[params.ChainId]
	var (
		topByBaseWithTokenIn, topByBaseWithTokenOut []v2.IndexerPool
	)
	for _, baseToken := range baseTokens {
		var (
			baseWithTokenInPools, baseWithTokenOutPools []v2.IndexerPool
		)
		for _, pool := range filteredPools {
			switch {
			case pool.Token0.Id == baseToken.Address.String() && pool.Token1.Id == params.TokenIn.Address.String():
				fallthrough
			case pool.Token1.Id == baseToken.Address.String() && pool.Token0.Id == params.TokenIn.Address.String():
				baseWithTokenInPools = append(baseWithTokenInPools, pool)
			case pool.Token0.Id == baseToken.Address.String() && pool.Token1.Id == params.TokenOut.Address.String():
				fallthrough
			case pool.Token1.Id == baseToken.Address.String() && pool.Token0.Id == params.TokenOut.Address.String():
				baseWithTokenOutPools = append(baseWithTokenOutPools, pool)
			}
		}
		sort.Slice(baseWithTokenInPools, func(i, j int) bool {
			return baseWithTokenInPools[i].Reserve.LessThan(baseWithTokenInPools[j].Reserve)
		})
		sort.Slice(baseWithTokenOutPools, func(i, j int) bool {
			return baseWithTokenOutPools[i].Reserve.LessThan(baseWithTokenOutPools[j].Reserve)
		})
		if len(baseWithTokenInPools) > params.RoutingConfig.V3PoolSelection.TopNWithEachBaseToken {
			baseWithTokenInPools = baseWithTokenInPools[:params.RoutingConfig.V3PoolSelection.TopNWithEachBaseToken]
		}
		if len(baseWithTokenOutPools) > params.RoutingConfig.V3PoolSelection.TopNWithEachBaseToken {
			baseWithTokenOutPools = baseWithTokenOutPools[:params.RoutingConfig.V3PoolSelection.TopNWithEachBaseToken]
		}
		topByBaseWithTokenIn = append(topByBaseWithTokenIn, baseWithTokenInPools...)
		topByBaseWithTokenOut = append(topByBaseWithTokenOut, baseWithTokenOutPools...)
	}
	sort.Slice(topByBaseWithTokenIn, func(i, j int) bool {
		return topByBaseWithTokenIn[i].Reserve.LessThan(topByBaseWithTokenIn[j].Reserve)
	})
	sort.Slice(topByBaseWithTokenOut, func(i, j int) bool {
		return topByBaseWithTokenOut[i].Reserve.LessThan(topByBaseWithTokenOut[j].Reserve)
	})
	if len(topByBaseWithTokenIn) > params.RoutingConfig.V3PoolSelection.TopNWithBaseToken {
		topByBaseWithTokenIn = topByBaseWithTokenIn[:params.RoutingConfig.V3PoolSelection.TopNWithBaseToken]
	}
	if len(topByBaseWithTokenOut) > params.RoutingConfig.V3PoolSelection.TopNWithBaseToken {
		topByBaseWithTokenOut = topByBaseWithTokenOut[:params.RoutingConfig.V3PoolSelection.TopNWithBaseToken]
	}

	var topByDirectSwapPool []v2.IndexerPool
	if params.RoutingConfig.V2PoolSelection.TopNDirectSwaps != 0 {
		address, token0, token1, err := params.PoolProvider.GetPoolAddress(params.TokenIn, params.TokenOut)
		if err != nil {
			return nil, candidatePools, err

		}
		topByDirectSwapPool = []v2.IndexerPool{
			{
				Id: address,
				Token0: struct {
					Id string
				}{
					Id: token0.Address.String(),
				},
				Token1: struct {
					Id string
				}{
					Id: token1.Address.String(),
				},
				Supply:  decimal.NewFromInt(10000), // Not used. Set to arbitrary number.
				Reserve: decimal.NewFromInt(10000), // Not used. Set to arbitrary number.
				Dex:     "pretend_direct",
			},
		}
	}
	addToAddressSet(topByDirectSwapPool)

	wrappedNativeToken, err := params.WNprovider.GetTokenByChain(params.ChainId)
	if err != nil {
		return poolAccessor, candidatePools, err
	}
	var topByEthQuoteTokenPool []v2.IndexerPool
	if params.TokenOut.Symbol() != "WETH" && params.TokenOut.Symbol() != "WETH9" && params.TokenOut.Symbol() != "ETH" {
		for _, pool := range filteredPools {
			switch {
			case params.RouteType == entities.ExactInput && ((pool.Token0.Id == wrappedNativeToken.Address.String() && pool.Token1.Id == params.TokenOut.Address.String()) ||
				(pool.Token1.Id == wrappedNativeToken.Address.String() && pool.Token0.Id == params.TokenOut.Address.String())):
				fallthrough
			case params.RouteType == entities.ExactOutput && ((pool.Token0.Id == wrappedNativeToken.Address.String() && pool.Token1.Id == params.TokenIn.Address.String()) ||
				(pool.Token1.Id == wrappedNativeToken.Address.String() && pool.Token0.Id == params.TokenIn.Address.String())):
				topByEthQuoteTokenPool = append(topByEthQuoteTokenPool, pool)
			default:
				continue
			}
		}
		if len(topByEthQuoteTokenPool) > 1 {
			topByEthQuoteTokenPool = topByEthQuoteTokenPool[:1]
		}
	}
	addToAddressSet(topByEthQuoteTokenPool)

	var topByTVL []v2.IndexerPool
	for _, pool := range filteredPools {
		if _, has := poolAddressesSoFar[pool.Id]; !has {
			topByTVL = append(topByTVL, pool)
		}
	}
	if len(topByTVL) > params.RoutingConfig.V3PoolSelection.TopN {
		topByTVL = topByTVL[:params.RoutingConfig.V3PoolSelection.TopN]
	}
	addToAddressSet(topByTVL)

	var topByTVLUsingTokenIn []v2.IndexerPool
	for _, pool := range filteredPools {
		if _, has := poolAddressesSoFar[pool.Id]; !has && (pool.Token0.Id == params.TokenIn.Address.String() ||
			pool.Token1.Id == params.TokenIn.Address.String()) {
			topByTVLUsingTokenIn = append(topByTVLUsingTokenIn, pool)
		}
	}
	if len(topByTVLUsingTokenIn) > params.RoutingConfig.V3PoolSelection.TopNTokenInOut {
		topByTVLUsingTokenIn = topByTVLUsingTokenIn[:params.RoutingConfig.V3PoolSelection.TopNTokenInOut]
	}
	addToAddressSet(topByTVLUsingTokenIn)

	var topByTVLUsingTokenOut []v2.IndexerPool
	for _, pool := range filteredPools {
		if _, has := poolAddressesSoFar[pool.Id]; !has && (pool.Token0.Id == params.TokenOut.Address.String() ||
			pool.Token1.Id == params.TokenOut.Address.String()) {
			topByTVLUsingTokenOut = append(topByTVLUsingTokenOut, pool)
		}
	}
	if len(topByTVLUsingTokenOut) > params.RoutingConfig.V3PoolSelection.TopNTokenInOut {
		topByTVLUsingTokenOut = topByTVLUsingTokenOut[:params.RoutingConfig.V3PoolSelection.TopNTokenInOut]
	}
	addToAddressSet(topByTVLUsingTokenOut)

	var (
		topByTVLUsingTokenInSecondHops    []v2.IndexerPool
		topByTVLUsingTokenInSecondHopsMap = make(map[string]struct{})
	)
	for _, pool := range topByTVLUsingTokenIn {
		var (
			secondHopId    string
			secondHopPools []v2.IndexerPool
		)
		if params.TokenIn.Address.String() == pool.Token0.Id {
			secondHopId = pool.Token1.Id
		} else {
			secondHopId = pool.Token0.Id
		}
		for _, filterPool := range filteredPools {
			if _, has := poolAddressesSoFar[pool.Id]; !has && (filterPool.Token0.Id == secondHopId || filterPool.Token1.Id == secondHopId) {
				secondHopPools = append(secondHopPools, filterPool)
			}
		}
		if len(secondHopPools) > params.RoutingConfig.V3PoolSelection.TopNSecondHop {
			secondHopPools = secondHopPools[:params.RoutingConfig.V3PoolSelection.TopNSecondHop]
		}
		for _, v := range secondHopPools {
			if _, has := topByTVLUsingTokenInSecondHopsMap[v.Id]; has {
				continue
			}
			topByTVLUsingTokenInSecondHopsMap[v.Id] = struct{}{}
			topByTVLUsingTokenInSecondHops = append(topByTVLUsingTokenInSecondHops, v)
		}
	}
	sort.Slice(topByTVLUsingTokenInSecondHops, func(i, j int) bool {
		return topByTVLUsingTokenInSecondHops[i].Reserve.LessThan(topByTVLUsingTokenInSecondHops[j].Reserve)
	})
	if len(topByTVLUsingTokenInSecondHops) > params.RoutingConfig.V3PoolSelection.TopNSecondHop {
		topByTVLUsingTokenInSecondHops = topByTVLUsingTokenInSecondHops[:params.RoutingConfig.V3PoolSelection.TopNSecondHop]
	}
	addToAddressSet(topByTVLUsingTokenInSecondHops)

	var (
		topByTVLUsingTokenOutSecondHops    []v2.IndexerPool
		topByTVLUsingTokenOutSecondHopsMap = make(map[string]struct{})
	)
	for _, pool := range topByTVLUsingTokenIn {
		var (
			secondHopId    string
			secondHopPools []v2.IndexerPool
		)
		if params.TokenOut.Address.String() == pool.Token0.Id {
			secondHopId = pool.Token1.Id
		} else {
			secondHopId = pool.Token0.Id
		}
		for _, filterPool := range filteredPools {
			if _, has := poolAddressesSoFar[pool.Id]; !has && (filterPool.Token0.Id == secondHopId || filterPool.Token1.Id == secondHopId) {
				secondHopPools = append(secondHopPools, filterPool)
			}
		}
		if len(secondHopPools) > params.RoutingConfig.V3PoolSelection.TopNSecondHop {
			secondHopPools = secondHopPools[:params.RoutingConfig.V3PoolSelection.TopNSecondHop]
		}
		for _, v := range secondHopPools {
			if _, has := topByTVLUsingTokenOutSecondHopsMap[v.Id]; has {
				continue
			}
			topByTVLUsingTokenOutSecondHopsMap[v.Id] = struct{}{}
			topByTVLUsingTokenOutSecondHops = append(topByTVLUsingTokenOutSecondHops, v)
		}
	}
	sort.Slice(topByTVLUsingTokenOutSecondHops, func(i, j int) bool {
		return topByTVLUsingTokenOutSecondHops[i].Reserve.LessThan(topByTVLUsingTokenOutSecondHops[j].Reserve)
	})
	if len(topByTVLUsingTokenOutSecondHops) > params.RoutingConfig.V3PoolSelection.TopNSecondHop {
		topByTVLUsingTokenOutSecondHops = topByTVLUsingTokenOutSecondHops[:params.RoutingConfig.V3PoolSelection.TopNSecondHop]
	}
	addToAddressSet(topByTVLUsingTokenOutSecondHops)

	indexerPools := append(topByBaseWithTokenIn, topByBaseWithTokenOut...)
	indexerPools = append(indexerPools, topByDirectSwapPool...)
	indexerPools = append(indexerPools, topByEthQuoteTokenPool...)
	indexerPools = append(indexerPools, topByTVL...)
	indexerPools = append(indexerPools, topByTVLUsingTokenIn...)
	indexerPools = append(indexerPools, topByTVLUsingTokenOut...)
	indexerPools = append(indexerPools, topByTVLUsingTokenInSecondHops...)
	indexerPools = append(indexerPools, topByTVLUsingTokenOutSecondHops...)

	//indexPools is not unique
	var (
		tokenAddresses    []string
		tokenAddressesMap = make(map[string]struct{})
	)
	for _, pool := range indexerPools {
		if _, has := tokenAddressesMap[pool.Token0.Id]; !has {
			tokenAddressesMap[pool.Token0.Id] = struct{}{}
			tokenAddresses = append(tokenAddresses, pool.Token0.Id)
		}
		if _, has := tokenAddressesMap[pool.Token1.Id]; !has {
			tokenAddressesMap[pool.Token1.Id] = struct{}{}
			tokenAddresses = append(tokenAddresses, pool.Token1.Id)
		}
	}

	tokenAccessor, err := params.TokenProvider.GetTokens(params.ChainId, tokenAddresses, &providersConfig.Config{
		BlockNumber: params.RoutingConfig.BlockNumber,
	})
	if err != nil {
		return nil, CandidatePoolsBySelectionCriteria{}, err
	}

	var (
		tokenPairs      []v2.TokenPairs
		indexerPoolsMap = make(map[string]struct{})
	)
	for _, pool := range indexerPools {
		if _, has := indexerPoolsMap[pool.Id]; has {
			continue
		}
		indexerPoolsMap[pool.Id] = struct{}{}
		token0 := tokenAccessor.GetTokenByAddress(pool.Token0.Id)
		token1 := tokenAccessor.GetTokenByAddress(pool.Token1.Id)
		if token0 == nil || token1 == nil {
			continue
		}
		tokenPairs = append(tokenPairs, v2.TokenPairs{
			Token0:         token0,
			Token1:         token1,
			PairAddress:    pool.Id,
			RouterAddress:  pool.RouterAddress,
			FactoryAddress: pool.FactoryAddress,
		})
	}
	poolAccessor, err = params.PoolProvider.GetPools(tokenPairs, nil)
	return

}

var baseTokensByChain = map[base_entities.ChainId][]*entities.Token{
	base_constant.MAINNET: {
		base_constant.USDC_MAINNET,
		base_constant.USDT_MAINNET,
		base_constant.WBTC_MAINNET,
		base_constant.DAI_MAINNET,
		base_constant.WrappedNativeCurrency[base_constant.MAINNET],
		base_constant.FEI_MAINNET,
	},
	base_constant.RINKEBY: {base_constant.DAI_RINKEBY_1, base_constant.DAI_RINKEBY_2},
	base_constant.OPTIMISM: {
		base_constant.DAI_OPTIMISM,
		base_constant.USDC_OPTIMISM,
		base_constant.USDT_OPTIMISM,
		base_constant.WBTC_OPTIMISM,
	},
	base_constant.OPTIMISTIC_KOVAN: {
		base_constant.DAI_OPTIMISTIC_KOVAN,
		base_constant.USDC_OPTIMISTIC_KOVAN,
		base_constant.WBTC_OPTIMISTIC_KOVAN,
		base_constant.USDT_OPTIMISTIC_KOVAN,
	},
	base_constant.ARBITRUM_ONE: {
		base_constant.DAI_ARBITRUM,
		base_constant.USDC_ARBITRUM,
		base_constant.WBTC_ARBITRUM,
		base_constant.USDT_ARBITRUM,
	},
	base_constant.ARBITRUM_RINKEBY: {base_constant.DAI_ARBITRUM_RINKEBY, base_constant.USDT_ARBITRUM_RINKEBY},
	base_constant.POLYGON:          {base_constant.USDC_POLYGON, base_constant.WMATIC_POLYGON},
	base_constant.POLYGON_MUMBAI:   {base_constant.DAI_POLYGON_MUMBAI, base_constant.WMATIC_POLYGON_MUMBAI},
	base_constant.BASE:             {base_constant.WrappedNativeCurrency[base_constant.BASE]},
}
