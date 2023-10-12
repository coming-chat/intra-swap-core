package providers

import (
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/coming-chat/intra-swap-core/providers/config"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
	"sync"
)

type TokenProvider interface {
	GetTokens(chainId base_entities.ChainId, addresses []string, providerConfig *config.Config) (TokenAccessor, error)
}

type TokenAccessor interface {
	GetTokenByAddress(address string) *entities.Token
	GetTokenBySymbol(symbol string) *entities.Token
	GetAllTokens() []*entities.Token
}

type BaseTokenAccessor struct {
	getTokenByAddress func(address string) *entities.Token
	getTokenBySymbol  func(symbol string) *entities.Token
	getAllTokens      func() []*entities.Token
}

func (a BaseTokenAccessor) GetTokenByAddress(address string) *entities.Token {
	return a.getTokenByAddress(address)
}

func (a BaseTokenAccessor) GetTokenBySymbol(symbol string) *entities.Token {
	return a.getTokenBySymbol(symbol)
}

func (a BaseTokenAccessor) GetAllTokens() []*entities.Token {
	return a.getAllTokens()
}

func NewBaseTokenProvider(multiCallCore rpc.MultiCallProviderCore) *BaseTokenProvider {
	return &BaseTokenProvider{
		MultiCallProviderCore: multiCallCore,
	}
}

type BaseTokenProvider struct {
	MultiCallProviderCore rpc.MultiCallProviderCore
}

func (o *BaseTokenProvider) GetTokens(chainId base_entities.ChainId, addresses []string, providerConfig *config.Config) (TokenAccessor, error) {
	addressToToken := make(map[string]*entities.Token)
	symbolToToken := make(map[string]*entities.Token)
	var tokens []*entities.Token

	var (
		multiCallSymbol   []rpc.MultiCallSingle[string]
		multiCallDecimals []rpc.MultiCallSingle[uint8]
	)
	for _, address := range addresses {
		multiCallSymbol = append(multiCallSymbol, rpc.MultiCallSingle[string]{
			Contract:        contracts.Erc20Abi,
			FunctionName:    "symbol",
			ContractAddress: common.HexToAddress(address),
		})
		multiCallDecimals = append(multiCallDecimals, rpc.MultiCallSingle[uint8]{
			Contract:        contracts.Erc20Abi,
			FunctionName:    "decimals",
			ContractAddress: common.HexToAddress(address),
		})
	}
	syncGroup := sync.WaitGroup{}
	syncGroup.Add(2)
	var (
		errSymbol, errDecimals error
	)

	go func() {
		defer syncGroup.Done()
		_, _, errSymbol = rpc.GetMultiCallProvider[string](o.MultiCallProviderCore).MultiCall(chainId, multiCallSymbol, 0, false, providerConfig)
	}()

	go func() {
		defer syncGroup.Done()
		_, _, errDecimals = rpc.GetMultiCallProvider[uint8](o.MultiCallProviderCore).MultiCall(chainId, multiCallDecimals, 0, false, providerConfig)
	}()

	syncGroup.Wait()
	if errSymbol != nil {
		return nil, errSymbol
	}
	if errDecimals != nil {
		return nil, errDecimals
	}

	for i, address := range addresses {
		if !multiCallSymbol[i].CallResult.Success || !multiCallDecimals[i].CallResult.Success {
			continue
		}
		addressToToken[address] = entities.NewToken(
			chainId,
			common.HexToAddress(address),
			uint(multiCallDecimals[i].CallResult.Data),
			multiCallSymbol[i].CallResult.Data,
			multiCallSymbol[i].CallResult.Data,
		)
		symbolToToken[multiCallSymbol[i].CallResult.Data] = addressToToken[address]
		tokens = append(tokens, addressToToken[address])
	}
	baseTokenAccessor := BaseTokenAccessor{}
	baseTokenAccessor.getAllTokens = func() []*entities.Token {
		return tokens
	}
	baseTokenAccessor.getTokenByAddress = func(address string) *entities.Token {
		return addressToToken[address]
	}
	baseTokenAccessor.getTokenBySymbol = func(symbol string) *entities.Token {
		return symbolToToken[symbol]
	}
	return baseTokenAccessor, nil
}
