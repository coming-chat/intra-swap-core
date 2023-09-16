package providers

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/coming-chat/intra-swap-core/providers/provider"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
	"sync"
)

type TokenProvider interface {
	GetTokens(addresses []string, providerConfig *provider.Config) (TokenAccessor, error)
}

type TokenAccessor interface {
	GetTokenByAddress(address string) *entities.Token
	GetTokenBySymbol(symbol string) *entities.Token
	GetAllTokens() []*entities.Token
}

type BaseTokenAccessor struct {
	addressToToken map[string]*entities.Token
	symbolToToken  map[string]*entities.Token
	tokens         []*entities.Token
}

func (b BaseTokenAccessor) GetTokenByAddress(address string) *entities.Token {
	return b.addressToToken[address]
}

func (b BaseTokenAccessor) GetTokenBySymbol(symbol string) *entities.Token {
	return b.symbolToToken[symbol]
}

func (b BaseTokenAccessor) GetAllTokens() []*entities.Token {
	return b.tokens
}

func NewBaseTokenProvider(chainId base_entities.ChainId, multiCallCore rpc.MultiCallProviderCore) *BaseTokenProvider {
	return &BaseTokenProvider{
		ChainId:               chainId,
		MultiCallProviderCore: multiCallCore,
	}
}

type BaseTokenProvider struct {
	ChainId               base_entities.ChainId
	MultiCallProviderCore rpc.MultiCallProviderCore
}

func (o *BaseTokenProvider) GetTokens(addresses []string, providerConfig *provider.Config) (TokenAccessor, error) {
	addressToToken := make(map[string]*entities.Token)
	symbolToToken := make(map[string]*entities.Token)
	var tokens []*entities.Token

	erc20, err := contracts.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	var (
		multiCallSymbol   []rpc.MultiCallSingleParam
		multiCallDecimals []rpc.MultiCallSingleParam
	)
	for _, address := range addresses {
		multiCallSymbol = append(multiCallSymbol, rpc.MultiCallSingleParam{
			Contract:        erc20,
			FunctionName:    "symbol",
			ContractAddress: common.HexToAddress(address),
		})
		multiCallDecimals = append(multiCallDecimals, rpc.MultiCallSingleParam{
			Contract:        erc20,
			FunctionName:    "decimals",
			ContractAddress: common.HexToAddress(address),
		})
	}
	syncGroup := sync.WaitGroup{}
	syncGroup.Add(2)
	var (
		symbolsResults         rpc.MultiCallResultWithInfo[string]
		decimalsResults        rpc.MultiCallResultWithInfo[uint8]
		errSymbol, errDecimals error
	)

	go func() {
		defer syncGroup.Done()
		symbolsResults, errSymbol = rpc.NewUniswapMultiCallProvider[string](o.MultiCallProviderCore).MultiCall(multiCallSymbol, providerConfig)
	}()

	go func() {
		defer syncGroup.Done()
		decimalsResults, errDecimals = rpc.NewUniswapMultiCallProvider[uint8](o.MultiCallProviderCore).MultiCall(multiCallDecimals, providerConfig)
	}()

	syncGroup.Wait()
	if errSymbol != nil {
		return nil, errSymbol
	}
	if errDecimals != nil {
		return nil, errDecimals
	}

	for i, address := range addresses {
		if !symbolsResults.ReturnData[i].Success || !decimalsResults.ReturnData[i].Success {
			continue
		}
		addressToToken[address] = entities.NewToken(
			o.ChainId,
			common.HexToAddress(address),
			uint(decimalsResults.ReturnData[i].ReturnData),
			symbolsResults.ReturnData[i].ReturnData,
			symbolsResults.ReturnData[i].ReturnData,
		)
		symbolToToken[symbolsResults.ReturnData[i].ReturnData] = addressToToken[address]
		tokens = append(tokens, addressToToken[address])
	}

	return BaseTokenAccessor{
		tokens:         tokens,
		addressToToken: addressToToken,
		symbolToToken:  symbolToToken,
	}, nil
}

// Some well known tokens on each chain for seeding cache / testing.
var USDC_MAINNET = entities.NewToken(
	base_entities.MAINNET,
	common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
	6,
	"USDC",
	"USD//C",
)

var USDT_MAINNET = entities.NewToken(
	base_entities.MAINNET,
	common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
	6,
	"USDT",
	"Tether USD",
)

var WBTC_MAINNET = entities.NewToken(
	base_entities.MAINNET,
	common.HexToAddress("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"),
	8,
	"WBTC",
	"Wrapped BTC",
)

var DAI_MAINNET = entities.NewToken(
	base_entities.MAINNET,
	common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"),
	18,
	"DAI",
	"Dai Stablecoin",
)

var FEI_MAINNET = entities.NewToken(
	base_entities.MAINNET,
	common.HexToAddress("0x956F47F50A910163D8BF957Cf5846D573E7f87CA"),
	18,
	"FEI",
	"Fei USD",
)

var UNI_MAINNET = entities.NewToken(
	base_entities.MAINNET,
	common.HexToAddress("0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"),
	18,
	"UNI",
	"Uniswap",
)

var USDC_ROPSTEN = entities.NewToken(
	base_entities.ROPSTEN,
	common.HexToAddress("0x07865c6e87b9f70255377e024ace6630c1eaa37f"),
	6,
	"USDC",
	"USD//C",
)

var USDT_ROPSTEN = entities.NewToken(
	base_entities.ROPSTEN,
	common.HexToAddress("0x516de3a7a567d81737e3a46ec4ff9cfd1fcb0136"),
	6,
	"USDT",
	"Tether USD",
)

var DAI_ROPSTEN = entities.NewToken(
	base_entities.ROPSTEN,
	common.HexToAddress("0xad6d458402f60fd3bd25163575031acdce07538d"),
	18,
	"DAI",
	"Dai Stablecoin",
)

var DAI_RINKEBY_1 = entities.NewToken(
	base_entities.RINKEBY,
	common.HexToAddress("0x5592ec0cfb4dbc12d3ab100b257153436a1f0fea"),
	18,
	"DAI",
	"DAI",
)

var DAI_RINKEBY_2 = entities.NewToken(
	base_entities.RINKEBY,
	common.HexToAddress("0xc7AD46e0b8a400Bb3C915120d284AafbA8fc4735"),
	18,
	"DAI",
	"DAI",
)

var USDC_RINKEBY = entities.NewToken(
	base_entities.RINKEBY,
	common.HexToAddress("0x4DBCdF9B62e891a7cec5A2568C3F4FAF9E8Abe2b"),
	6,
	"tUSDC",
	"test USD//C",
)

var USDT_RINKEBY = entities.NewToken(
	base_entities.RINKEBY,
	common.HexToAddress("0xa689352b7c1cad82864beb1d90679356d3962f4d"),
	18,
	"USDT",
	"Tether USD",
)

var USDC_GÖRLI = entities.NewToken(
	base_entities.GÖRLI,
	common.HexToAddress("0x07865c6e87b9f70255377e024ace6630c1eaa37f"),
	6,
	"USDC",
	"USD//C",
)

var USDT_GÖRLI = entities.NewToken(
	base_entities.GÖRLI,
	common.HexToAddress("0xe583769738b6dd4e7caf8451050d1948be717679"),
	18,
	"USDT",
	"Tether USD",
)

var WBTC_GÖRLI = entities.NewToken(
	base_entities.GÖRLI,
	common.HexToAddress("0xa0a5ad2296b38bd3e3eb59aaeaf1589e8d9a29a9"),
	8,
	"WBTC",
	"Wrapped BTC",
)

var DAI_GÖRLI = entities.NewToken(
	base_entities.GÖRLI,
	common.HexToAddress("0x11fe4b6ae13d2a6055c8d9cf65c55bac32b5d844"),
	18,
	"DAI",
	"Dai Stablecoin",
)

var UNI_GÖRLI = entities.NewToken(
	base_entities.GÖRLI,
	common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"),
	18,
	"UNI",
	"Uni token",
)

var USDC_KOVAN = entities.NewToken(
	base_entities.KOVAN,
	common.HexToAddress("0x31eeb2d0f9b6fd8642914ab10f4dd473677d80df"),
	6,
	"USDC",
	"USD//C",
)

var USDT_KOVAN = entities.NewToken(
	base_entities.KOVAN,
	common.HexToAddress("0xa325f1b1ebb748715dfbbaf62e0c6677e137f45d"),
	18,
	"USDT",
	"Tether USD",
)

var WBTC_KOVAN = entities.NewToken(
	base_entities.KOVAN,
	common.HexToAddress("0xe36bc5d8b689ad6d80e78c3e736670e80d4b329d"),
	8,
	"WBTC",
	"Wrapped BTC",
)

var DAI_KOVAN = entities.NewToken(
	base_entities.KOVAN,
	common.HexToAddress("0x9dc7b33c3b63fc00ed5472fbd7813edda6a64752"),
	18,
	"DAI",
	"Dai Stablecoin",
)

var USDC_OPTIMISM = entities.NewToken(
	base_entities.OPTIMISM,
	common.HexToAddress("0x7F5c764cBc14f9669B88837ca1490cCa17c31607"),
	6,
	"USDC",
	"USD//C",
)

var USDT_OPTIMISM = entities.NewToken(
	base_entities.OPTIMISM,
	common.HexToAddress("0x94b008aA00579c1307B0EF2c499aD98a8ce58e58"),
	6,
	"USDT",
	"Tether USD",
)

var WBTC_OPTIMISM = entities.NewToken(
	base_entities.OPTIMISM,
	common.HexToAddress("0x68f180fcCe6836688e9084f035309E29Bf0A2095"),
	8,
	"WBTC",
	"Wrapped BTC",
)

var DAI_OPTIMISM = entities.NewToken(
	base_entities.OPTIMISM,
	common.HexToAddress("0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1"),
	18,
	"DAI",
	"Dai Stablecoin",
)

var USDC_OPTIMISTIC_KOVAN = entities.NewToken(
	base_entities.OPTIMISTIC_KOVAN,
	common.HexToAddress("0x3b8e53b3ab8e01fb57d0c9e893bc4d655aa67d84"),
	6,
	"USDC",
	"USD//C",
)

var USDT_OPTIMISTIC_KOVAN = entities.NewToken(
	base_entities.OPTIMISTIC_KOVAN,
	common.HexToAddress("0x7F5c764cBc14f9669B88837ca1490cCa17c31607"),
	6,
	"USDT",
	"Tether USD",
)

var WBTC_OPTIMISTIC_KOVAN = entities.NewToken(
	base_entities.OPTIMISTIC_KOVAN,
	common.HexToAddress("0x2382a8f65b9120E554d1836a504808aC864E169d"),
	8,
	"WBTC",
	"Wrapped BTC",
)

var DAI_OPTIMISTIC_KOVAN = entities.NewToken(
	base_entities.OPTIMISTIC_KOVAN,
	common.HexToAddress("0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1"),
	18,
	"DAI",
	"Dai Stablecoin",
)

var USDC_ARBITRUM = entities.NewToken(
	base_entities.ARBITRUM_ONE,
	common.HexToAddress("0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8"),
	6,
	"USDC",
	"USD//C",
)

var USDT_ARBITRUM = entities.NewToken(
	base_entities.ARBITRUM_ONE,
	common.HexToAddress("0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9"),
	6,
	"USDT",
	"Tether USD",
)

var WBTC_ARBITRUM = entities.NewToken(
	base_entities.ARBITRUM_ONE,
	common.HexToAddress("0x2f2a2543B76A4166549F7aaB2e75Bef0aefC5B0f"),
	8,
	"WBTC",
	"Wrapped BTC",
)

var DAI_ARBITRUM = entities.NewToken(
	base_entities.ARBITRUM_ONE,
	common.HexToAddress("0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1"),
	18,
	"DAI",
	"Dai Stablecoin",
)

// var DAI_ARBITRUM_RINKEBY = entities.NewToken(
//   base_entities.ARBITRUM_RINKEBY,
//   common.HexToAddress("0x2f3C1B6A51A469051A22986aA0dDF98466cc8D3c"),
//   18,
//   "DAI",
//   "Dai Stablecoin"
// );

// higher liquidity in dai-weth pool on arb-rinkeby
var DAI_ARBITRUM_RINKEBY = entities.NewToken(
	base_entities.ARBITRUM_RINKEBY,
	common.HexToAddress("0x5364dc963c402aaf150700f38a8ef52c1d7d7f14"),
	18,
	"DAI",
	"Dai Stablecoin",
)

var USDT_ARBITRUM_RINKEBY = entities.NewToken(
	base_entities.ARBITRUM_RINKEBY,
	common.HexToAddress("0x920b9301c2de92186299cd2abc7199e25b9728b3"),
	6,
	"UDST",
	"Tether USD",
)

var USDC_ARBITRUM_RINKEBY = entities.NewToken(
	base_entities.ARBITRUM_RINKEBY,
	common.HexToAddress("0x09b98f8b2395d076514037ff7d39a091a536206c"),
	6,
	"USDC",
	"USD//C",
)

var UNI_ARBITRUM_RINKEBY = entities.NewToken(
	base_entities.ARBITRUM_RINKEBY,
	common.HexToAddress("0x049251a7175071316e089d0616d8b6aacd2c93b8"),
	18,
	"UNI",
	"Uni token",
)

// polygon tokens
var WMATIC_POLYGON = entities.NewToken(
	base_entities.POLYGON,
	common.HexToAddress("0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270"),
	18,
	"WMATIC",
	"Wrapped MATIC",
)

var WETH_POLYGON = entities.NewToken(
	base_entities.POLYGON,
	common.HexToAddress("0x7ceb23fd6bc0add59e62ac25578270cff1b9f619"),
	18,
	"WETH",
	"Wrapped Ether",
)

var USDC_POLYGON = entities.NewToken(
	base_entities.POLYGON,
	common.HexToAddress("0x2791bca1f2de4661ed88a30c99a7a9449aa84174"),
	6,
	"USDC",
	"USD//C",
)

var DAI_POLYGON = entities.NewToken(
	base_entities.POLYGON,
	common.HexToAddress("0x8f3Cf7ad23Cd3CaDbD9735AFf958023239c6A063"),
	18,
	"DAI",
	"Dai Stablecoin",
)

// polygon mumbai tokens
var WMATIC_POLYGON_MUMBAI = entities.NewToken(
	base_entities.POLYGON_MUMBAI,
	common.HexToAddress("0x9c3C9283D3e44854697Cd22D3Faa240Cfb032889"),
	18,
	"WMATIC",
	"Wrapped MATIC",
)

var USDC_POLYGON_MUMBAI = entities.NewToken(
	base_entities.POLYGON_MUMBAI,
	common.HexToAddress("0xe11a86849d99f524cac3e7a0ec1241828e332c62"),
	6,
	"USDC",
	"USD//C",
)

var DAI_POLYGON_MUMBAI = entities.NewToken(
	base_entities.POLYGON_MUMBAI,
	common.HexToAddress("0x001b3b4d0f3714ca98ba10f6042daebf0b1b7b6f"),
	18,
	"DAI",
	"Dai Stablecoin",
)

var WETH_POLYGON_MUMBAI = entities.NewToken(
	base_entities.POLYGON_MUMBAI,
	common.HexToAddress("0xa6fa4fb5f76172d178d61b04b0ecd319c5d1c0aa"),
	18,
	"WETH",
	"Wrapped Ether",
)

var USDC_BASE_MAINNET = entities.NewToken(
	base_entities.BASE,
	common.HexToAddress("0xd9aAEc86B65D86f6A7B5B1b0c42FFA531710b6CA"),
	6,
	"USDbC",
	"USD base Coin",
)
