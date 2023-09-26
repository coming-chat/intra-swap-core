package cmd

import (
	"context"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/providers"
	"github.com/coming-chat/intra-swap-core/providers/gas"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	v2 "github.com/coming-chat/intra-swap-core/providers/v2"
	v3 "github.com/coming-chat/intra-swap-core/providers/v3"
	"github.com/coming-chat/intra-swap-core/routers/alpha_router"
	gasModelsV2 "github.com/coming-chat/intra-swap-core/routers/alpha_router/models/v2"
	gasModelsV3 "github.com/coming-chat/intra-swap-core/routers/alpha_router/models/v3"
)

func Ready(chainId base_entities.ChainId, v2Dex, v3Dex string) (*alpha_router.AlphaRouter, error) {
	rpcProvider, err := rpc.NewBaseProvider(context.TODO(), []string{"https://base.meowrpc.com"}, 100000)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	uniswapMultiCallCore, err := rpc.NewUniswapMultiCallProviderCore(rpcProvider)
	if err != nil {
		return nil, err
	}
	wNprovider := &providers.BaseWrappedNativeCurrencyProvider{}
	router := alpha_router.NewAlphaRouter(alpha_router.AlphaRouterParams{
		ChainId:                       chainId,
		Provider:                      rpcProvider,
		MultiCall2ProviderCore:        uniswapMultiCallCore,
		V3IndexerProvider:             v3.NewGeckoTerminalProvider(v3Dex),
		V3PoolProvider:                v3.NewBasePoolProvider(chainId, map[string]string{}, uniswapMultiCallCore, nil),
		V3QuoteProvider:               v3.NewBaseQuoteProvider(context.TODO(), chainId, rpcProvider, uniswapMultiCallCore, nil, false, nil),
		V2IndexerProvider:             v2.NewGeckoTerminalProvider(v2Dex),
		V2PoolProvider:                v2.NewBasePoolProvider(chainId, map[string]string{}, uniswapMultiCallCore, nil),
		V2QuoteProvider:               v2.NewBaseQuoteProvider(context.TODO(), chainId, rpcProvider, uniswapMultiCallCore, false, nil),
		TokenProvider:                 providers.NewBaseTokenProvider(uniswapMultiCallCore),
		GasPriceProvider:              gas.NewBaseGasPriceProvider(context.TODO(), rpcProvider),
		V3GasModelFactory:             gasModelsV3.NewHeuristicGasModelFactory(rpcProvider, uniswapMultiCallCore, "", wNprovider),
		V2GasModelFactory:             gasModelsV2.NewHeuristicGasModelFactory(wNprovider),
		BlockedTokenListProvider:      nil,
		WrappedNativeCurrencyProvider: wNprovider,
	}, context.TODO())
	return router, nil
}
