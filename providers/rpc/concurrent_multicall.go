package rpc

import (
	"errors"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"math"
	"sync"
	"sync/atomic"
)

type RetryOptions struct {
	Retries    int
	MinTimeout int
	MaxTimeout int
}

type BatchParams struct {
	MultiCallChunk  int
	GasLimitPerCall uint64
	MinSuccessRate  float64
}

type MultiCallConfig struct {
	RetryOptions     RetryOptions
	BatchParams      BatchParams
	MaxConcurrentNum int
	RequireSuccess   bool
}

func ConcurrentMultiCall[T any](chainId base_entities.ChainId, core MultiCallProviderCore, multiCallParams []MultiCallSingleParam, c *MultiCallConfig) (rInfo MultiCallResultWithInfo[T], err error) {
	if c == nil {
		c = &MultiCallConfig{
			RetryOptions: RetryOptions{
				Retries: 1,
			},
			BatchParams: BatchParams{
				MultiCallChunk:  200,
				MinSuccessRate:  1,
				GasLimitPerCall: 1000_000,
			},
			MaxConcurrentNum: 5,
		}
	}
	var (
		split          = math.Ceil(float64(len(multiCallParams)) / float64(c.BatchParams.MultiCallChunk))
		syncGroup      = sync.WaitGroup{}
		rwLock         = sync.Mutex{}
		totalSuccess   = atomic.Uint32{}
		successResults = make([]MultiCallResultWithInfo[T], int(split))
		concurrentPool = make(chan struct{}, c.MaxConcurrentNum)
	)

	for i := 0; i < len(multiCallParams); i += c.BatchParams.MultiCallChunk {
		syncGroup.Add(1)
		go func(index int) {
			defer func() {
				<-concurrentPool
				syncGroup.Done()
			}()
			concurrentPool <- struct{}{}
			for try := 0; try < c.RetryOptions.Retries; try++ {
				endIndex := index + c.BatchParams.MultiCallChunk
				if index+c.BatchParams.MultiCallChunk > len(multiCallParams) {
					endIndex = len(multiCallParams)
				}
				returnData, errReq := GetMultiCallProvider[T](core).MultiCall(chainId, multiCallParams[index:endIndex], c.BatchParams.GasLimitPerCall, c.RequireSuccess, nil)
				if errReq != nil && try+1 < c.RetryOptions.Retries {
					continue
				}
				successResult := 0
				for _, result := range returnData.ReturnData {
					if result.Success {
						successResult++
						totalSuccess.Add(1)
					}
				}
				if errReq != nil && len(returnData.ReturnData) == 0 {
					for range multiCallParams[index:endIndex] {
						returnData.ReturnData = append(returnData.ReturnData, MultiCallResult[T]{
							Success: false,
							Err:     errReq,
						})
					}
				}
				if float64(successResult/len(returnData.ReturnData)) < c.BatchParams.MinSuccessRate && try+1 < c.RetryOptions.Retries {
					continue
				}
				rwLock.Lock()
				successResults[index/c.BatchParams.MultiCallChunk] = returnData
				rwLock.Unlock()
			}
		}(i)
	}
	syncGroup.Wait()
	if totalSuccess.Load() == 0 {
		return rInfo, errors.New("no success multicall req")
	}
	for _, r := range successResults {
		rInfo.ReturnData = append(rInfo.ReturnData, r.ReturnData...)
		rInfo.BlockHash = r.BlockHash
		rInfo.BlockNumber = r.BlockNumber
	}

	if len(multiCallParams) != len(rInfo.ReturnData) {
		return rInfo, errors.New("concurrent result length not equal to req params length")
	}

	return
}
