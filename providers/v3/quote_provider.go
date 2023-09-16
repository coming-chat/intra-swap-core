package v3

import (
	"context"
	"errors"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts/omni_swap"
	"github.com/coming-chat/intra-swap-core/providers/provider"
	"github.com/coming-chat/intra-swap-core/providers/rpc"
	"github.com/daoleno/uniswap-sdk-core/entities"
	entitiesV3 "github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/periphery"
	"github.com/ethereum/go-ethereum/common"
	"math"
	"math/big"
	"sync"
	"sync/atomic"
)

type AmountQuote struct {
	Amount                      *entities.CurrencyAmount
	Quote                       *big.Int
	SqrtPriceX96AfterList       []*big.Int
	InitializedTicksCrossedList []uint32
	GasEstimate                 *big.Int
}

type RouteWithQuotes struct {
	Route       *entitiesV3.Route
	AmountQuote []AmountQuote
}

type functionName = string

const (
	QuoteExactInput  functionName = "quoteExactInput"
	QuoteExactOutput functionName = "quoteExactOutput"
)

type QuoteProvider interface {
	GetQuotesManyExactIn(
		amountIns []*entities.CurrencyAmount,
		routes []*base_entities.V3Route,
		providerConfig *provider.Config,
	) ([]RouteWithQuotes, int64, error)

	GetQuotesManyExactOut(
		amountOuts []*entities.CurrencyAmount,
		routes []*base_entities.V3Route,
		providerConfig *provider.Config,
	) ([]RouteWithQuotes, int64, error)
}

type BatchParams struct {
	MultiCallChunk      int
	GasLimitPerCall     int64
	QuoteMinSuccessRate float64
}

type BlockNumberConfig struct {
	BaseBlockOffset uint64
	Rollback        bool
}

type QuoteRetryOptions struct {
	Retries    int
	MinTimeout int
	MaxTimeout int
}

type BaseQuoteProvider struct {
	QuoterAddress      common.Address
	ChainId            base_entities.ChainId
	Provider           rpc.BaseProvider
	MultiCall2Provider rpc.MultiCallProviderCore
	BatchParams        BatchParams
	BlockNumberConfig  BlockNumberConfig
	RetryOptions       QuoteRetryOptions
	Ctx                context.Context
}

func NewBaseQuoteProvider(
	ctx context.Context,
	address common.Address,
	chainId base_entities.ChainId,
	baseProvider rpc.BaseProvider,
	multiCallCore rpc.MultiCallProviderCore,
	batchParams *BatchParams,
	blockNumberConfig *BlockNumberConfig,
	retryOptions *QuoteRetryOptions,
) *BaseQuoteProvider {
	quoteProvider := &BaseQuoteProvider{
		ChainId:            chainId,
		Ctx:                ctx,
		QuoterAddress:      address,
		Provider:           baseProvider,
		MultiCall2Provider: multiCallCore,
	}
	if batchParams == nil {
		batchParams = &BatchParams{
			MultiCallChunk:      150,
			GasLimitPerCall:     1_000_000,
			QuoteMinSuccessRate: 0.2,
		}
	}
	if blockNumberConfig == nil {
		blockNumberConfig = &BlockNumberConfig{
			BaseBlockOffset: 0,
			Rollback:        false,
		}
	}
	if retryOptions == nil {
		retryOptions = &QuoteRetryOptions{
			Retries:    2,
			MinTimeout: 25,
			MaxTimeout: 250,
		}
	}
	quoteProvider.BatchParams = *batchParams
	quoteProvider.BlockNumberConfig = *blockNumberConfig
	quoteProvider.RetryOptions = *retryOptions
	return quoteProvider
}

func (b *BaseQuoteProvider) GetQuotesManyExactIn(
	amountIns []*entities.CurrencyAmount,
	routes []*base_entities.V3Route,
	providerConfig *provider.Config,
) ([]RouteWithQuotes, int64, error) {
	return b.getQuotesManyData(amountIns, routes, QuoteExactInput, providerConfig)
}

func (b *BaseQuoteProvider) GetQuotesManyExactOut(
	amountOuts []*entities.CurrencyAmount,
	routes []*base_entities.V3Route,
	providerConfig *provider.Config,
) ([]RouteWithQuotes, int64, error) {
	return b.getQuotesManyData(amountOuts, routes, QuoteExactOutput, providerConfig)
}

func (b *BaseQuoteProvider) getQuotesManyData(
	amounts []*entities.CurrencyAmount,
	routes []*base_entities.V3Route,
	functionName functionName,
	providerConfig *provider.Config,
) ([]RouteWithQuotes, int64, error) {
	originalBlockNumber, err := b.Provider.Rpc.BlockNumber(b.Ctx)
	if err != nil {
		return nil, 0, err
	}
	if providerConfig == nil {
		providerConfig = &provider.Config{
			BlockNumber: originalBlockNumber + b.BlockNumberConfig.BaseBlockOffset,
		}
	}
	var multiCallParams []rpc.MultiCallSingleParam
	quoterContract, err := omni_swap.IQuoterV2MetaData.GetAbi()
	if err != nil {
		return nil, 0, err
	}
	for _, r := range routes {
		encodedRoute, err := periphery.EncodeRouteToPath(r, functionName == QuoteExactOutput)
		if err != nil {
			return nil, 0, err
		}
		for _, a := range amounts {
			multiCallParams = append(multiCallParams, rpc.MultiCallSingleParam{
				Contract: quoterContract,
				FunctionParams: []any{
					encodedRoute,
					a.Quotient(),
				},
				ContractAddress: b.QuoterAddress,
				FunctionName:    functionName,
			})
		}
	}

	var (
		split          = math.Ceil(float64(len(multiCallParams)) / float64(b.BatchParams.MultiCallChunk))
		syncGroup      = sync.WaitGroup{}
		rwLock         = sync.Mutex{}
		totalSuccess   = atomic.Uint32{}
		successResults = make([]rpc.MultiCallResultWithInfo[struct {
			AmountOut                   *big.Int
			AmountIn                    *big.Int
			SqrtPriceX96AfterList       []*big.Int
			InitializedTicksCrossedList []uint32
			GasEstimate                 *big.Int
		}], int(split))
	)

	for i := 0; i < len(multiCallParams); i += b.BatchParams.MultiCallChunk {
		syncGroup.Add(1)
		go func(index int) {
			defer func() {
				syncGroup.Done()
			}()
			for try := 0; try < b.RetryOptions.Retries; try++ {
				endIndex := index + b.BatchParams.MultiCallChunk
				if index+b.BatchParams.MultiCallChunk > len(multiCallParams) {
					endIndex = len(multiCallParams)
				}
				returnData, err := rpc.NewUniswapMultiCallProvider[struct {
					AmountOut                   *big.Int
					AmountIn                    *big.Int
					SqrtPriceX96AfterList       []*big.Int
					InitializedTicksCrossedList []uint32
					GasEstimate                 *big.Int
				}](b.MultiCall2Provider).MultiCall(multiCallParams[index:endIndex], nil)
				if err != nil {
					continue
				}
				successResult := 0
				for _, result := range returnData.ReturnData {
					if result.Success {
						successResult++
					}
				}
				if float64(successResult/len(returnData.ReturnData)) < b.BatchParams.QuoteMinSuccessRate {
					continue
				}
				rwLock.Lock()
				successResults[index/b.BatchParams.MultiCallChunk] = returnData
				rwLock.Unlock()
				totalSuccess.Add(1)
			}
		}(i)
	}
	syncGroup.Wait()

	if totalSuccess.Load() == 0 {
		return nil, 0, errors.New("quotes returned null")
	}
	var (
		blockNumber             = successResults[0].BlockNumber
		routesQuotes            []RouteWithQuotes
		quotes                  []AmountQuote
		amountIndex, routeIndex = 0, 0
	)
	for _, result := range successResults {
		if result.BlockNumber != blockNumber {
			return nil, 0, errors.New("quotes returned from different blocks")
		}
		for _, quote := range result.ReturnData {
			if amountIndex == len(amounts) {
				routesQuotes = append(routesQuotes, RouteWithQuotes{
					Route:       routes[routeIndex],
					AmountQuote: quotes,
				})
				quotes = make([]AmountQuote, 0)
				routeIndex++
				amountIndex = 0
			}

			if !quote.Success {
				quotes = append(quotes, AmountQuote{
					Amount:                      amounts[amountIndex],
					Quote:                       nil,
					SqrtPriceX96AfterList:       nil,
					GasEstimate:                 nil,
					InitializedTicksCrossedList: nil,
				})
			} else {
				aq := AmountQuote{
					Amount:                      amounts[amountIndex],
					Quote:                       quote.ReturnData.AmountIn,
					SqrtPriceX96AfterList:       quote.ReturnData.SqrtPriceX96AfterList,
					GasEstimate:                 quote.ReturnData.GasEstimate,
					InitializedTicksCrossedList: quote.ReturnData.InitializedTicksCrossedList,
				}
				if quote.ReturnData.AmountIn == nil {
					aq.Quote = quote.ReturnData.AmountOut
				}
				quotes = append(quotes, aq)
			}
			amountIndex++
		}
	}
	return routesQuotes, blockNumber.Int64(), nil
}
