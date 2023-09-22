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
	"github.com/coming-chat/intra-swap-core/util"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Ready(chainId base_entities.ChainId) (*alpha_router.AlphaRouter, error) {
	client, err := ethclient.Dial("https://base.meowrpc.com")
	if err != nil {
		return nil, err
	}
	rpcBaseProvider := rpc.BaseProvider{
		Rpc: client,
	}
	uniswapMultiCallCore, err := rpc.NewUniswapMultiCallProviderCore(chainId, rpcBaseProvider)
	if err != nil {
		return nil, err
	}
	router := alpha_router.NewAlphaRouter(alpha_router.AlphaRouterParams{
		ChainId:                  chainId,
		Provider:                 rpcBaseProvider,
		MultiCall2ProviderCore:   uniswapMultiCallCore,
		V3IndexerProvider:        v3.NewGeckoTerminalProvider(chainId),
		V3PoolProvider:           v3.NewBasePoolProvider(chainId, map[string]string{}, uniswapMultiCallCore, nil),
		V3QuoteProvider:          v3.NewBaseQuoteProvider(context.TODO(), util.UniswapV3Quoter, chainId, rpcBaseProvider, uniswapMultiCallCore, nil, nil),
		V2IndexerProvider:        v2.NewGeckoTerminalProvider(chainId),
		V2PoolProvider:           v2.NewBasePoolProvider(chainId, map[string]string{}, uniswapMultiCallCore, nil),
		V2QuoteProvider:          v2.NewBaseQuoteProvider(),
		TokenProvider:            providers.NewBaseTokenProvider(chainId, uniswapMultiCallCore),
		GasPriceProvider:         &gas.BaseGasPriceProvider{},
		V3GasModelFactory:        gasModelsV3.NewHeuristicGasModelFactory(rpcBaseProvider, uniswapMultiCallCore, ""),
		V2GasModelFactory:        &gasModelsV2.HeuristicGasModelFactory{},
		BlockedTokenListProvider: nil,
	}, context.TODO())
	return router, nil
}
