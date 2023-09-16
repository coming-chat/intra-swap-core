package gas_models

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/providers"
	v2 "github.com/coming-chat/intra-swap-core/providers/v2"
	v3 "github.com/coming-chat/intra-swap-core/providers/v3"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"math/big"
)

type V3GasModelFactory interface {
	BuildGasModel(
		chainId base_entities.ChainId,
		gasPriceWei *big.Int,
		poolProvider v3.PoolProvider,
		inTermsOfToken *entities.Token,
	) (providers.GasModel[providers.V3RouteWithValidQuote], error)
}

type V2GasModelFactory interface {
	BuildGasModel(
		chainId base_entities.ChainId,
		gasPriceWei *big.Int,
		poolProvider v2.PoolProvider,
		token *entities.Token,
	) (providers.GasModel[providers.V2RouteWithValidQuote], error)
}

var UsdGasTokensByChain = map[base_entities.ChainId][]*entities.Token{
	base_entities.MAINNET:      {providers.DAI_MAINNET, providers.USDC_MAINNET, providers.USDT_MAINNET},
	base_entities.RINKEBY:      {providers.DAI_RINKEBY_1, providers.DAI_RINKEBY_2},
	base_entities.ARBITRUM_ONE: {providers.DAI_ARBITRUM, providers.USDC_ARBITRUM, providers.USDT_ARBITRUM},
	base_entities.OPTIMISM:     {providers.DAI_OPTIMISM, providers.USDC_OPTIMISM, providers.USDT_OPTIMISM},
	base_entities.OPTIMISTIC_KOVAN: {
		providers.DAI_OPTIMISTIC_KOVAN,
		providers.USDC_OPTIMISTIC_KOVAN,
		providers.USDT_OPTIMISTIC_KOVAN,
	},
	base_entities.ARBITRUM_RINKEBY: {providers.DAI_ARBITRUM_RINKEBY, providers.USDT_ARBITRUM_RINKEBY},
	base_entities.KOVAN:            {providers.DAI_KOVAN, providers.USDC_KOVAN, providers.USDT_KOVAN},
	base_entities.GÖRLI:            {providers.DAI_GÖRLI, providers.USDC_GÖRLI, providers.USDT_GÖRLI, providers.WBTC_GÖRLI},
	base_entities.ROPSTEN:          {providers.DAI_ROPSTEN, providers.USDC_ROPSTEN, providers.USDT_ROPSTEN},
	base_entities.POLYGON:          {providers.USDC_POLYGON},
	base_entities.POLYGON_MUMBAI:   {providers.DAI_POLYGON_MUMBAI},
	base_entities.BASE:             {providers.USDC_BASE_MAINNET},
}
