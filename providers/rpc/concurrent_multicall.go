package rpc

import (
	"errors"
	"fmt"
	"github.com/coming-chat/intra-swap-core/base_entities"
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

func ConcurrentMultiCall[T any](
	chainId base_entities.ChainId,
	core MultiCallProviderCore,
	multiCallParams []MultiCallSingle[T],
	c *MultiCallConfig,
) error {
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
		syncGroup      = sync.WaitGroup{}
		totalSuccess   = atomic.Uint32{}
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
				_, _, errReq := GetMultiCallProvider[T](core).MultiCall(chainId, multiCallParams[index:endIndex], c.BatchParams.GasLimitPerCall, c.RequireSuccess, nil)
				if errReq != nil && try+1 < c.RetryOptions.Retries {
					continue
				}
				successResult := 0
				for _, result := range multiCallParams[index:endIndex] {
					if result.CallResult.Success {
						successResult++
						totalSuccess.Add(1)
					}
				}
				if errReq != nil && successResult == 0 {
					for j := range multiCallParams[index:endIndex] {
						multiCallParams[index+j].CallResult = MultiCallResult[T]{
							Success: false,
							Err:     errReq,
						}
					}
				}
				if float64(successResult/(endIndex-index)) < c.BatchParams.MinSuccessRate && try+1 < c.RetryOptions.Retries {
					continue
				}
			}
		}(i)
	}
	syncGroup.Wait()
	var multiErrs []error
	for _, r := range multiCallParams {
		if r.CallResult.Success {
			continue
		}
		multiErrs = append(multiErrs, r.CallResult.Err)
	}

	if totalSuccess.Load() == 0 {
		allErrs := errors.Join(multiErrs...)
		if allErrs != nil {
			return fmt.Errorf("all multicall failed with errors [%w]", allErrs)
		}
	}
	return nil
}
