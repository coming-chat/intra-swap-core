package rpc

import (
	"errors"
	"fmt"
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
	GasLimitPerCall int64
	MinSuccessRate  float64
}

type MultiCallConfig struct {
	RetryOptions     RetryOptions
	BatchParams      BatchParams
	MaxConcurrentNum int
}

func ConcurrentMultiCall[T any](chainId base_entities.ChainId, core MultiCallProviderCore, multiCallParams []MultiCallSingleParam, c *MultiCallConfig) (rInfo MultiCallResultWithInfo[T], err error) {
	if c == nil {
		c = &MultiCallConfig{
			RetryOptions: RetryOptions{
				Retries: 1,
			},
			BatchParams: BatchParams{
				MultiCallChunk: 200,
				MinSuccessRate: 1,
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
				returnData, err := GetMultiCallProvider[T](core).MultiCall(chainId, multiCallParams[index:endIndex], nil)
				if err != nil {
					fmt.Printf("%v\n", err)
					continue
				}
				successResult := 0
				for _, result := range returnData.ReturnData {
					if result.Success {
						successResult++
					}
				}
				if float64(successResult/len(returnData.ReturnData)) < c.BatchParams.MinSuccessRate {
					continue
				}
				rwLock.Lock()
				successResults[index/c.BatchParams.MultiCallChunk] = returnData
				rwLock.Unlock()
				totalSuccess.Add(1)
			}
		}(i)
	}
	syncGroup.Wait()
	if totalSuccess.Load() == 0 {
		return rInfo, errors.New("quotes returned null")
	}
	for _, r := range successResults {
		rInfo.ReturnData = append(rInfo.ReturnData, r.ReturnData...)
		rInfo.BlockHash = r.BlockHash
		rInfo.BlockNumber = r.BlockNumber
	}

	return rInfo, nil
}
