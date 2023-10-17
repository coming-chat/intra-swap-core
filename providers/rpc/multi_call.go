package rpc

import (
	"context"
	_ "embed"
	"errors"
	"github.com/coming-chat/intra-swap-core/base_constant"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/coming-chat/intra-swap-core/providers/config"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gkirito/go-ethereum/accounts/abi"
	"github.com/gkirito/go-ethereum/accounts/abi/bind"
	selfCommon "github.com/gkirito/go-ethereum/common"
	"math/big"
)

type MultiCallSingle[T any] struct {
	Contract        *abi.ABI
	FunctionParams  []any
	ContractAddress common.Address
	FunctionName    string
	CallResult      MultiCallResult[T]
}

type MultiCallResult[T any] struct {
	Success bool
	Data    T
	Err     error
}

type MultiCallResultWithInfo[T any] struct {
	ReturnData  []MultiCallResult[T]
	BlockNumber *big.Int
	BlockHash   common.Hash
}

type MultiCallProvider[MultiCallConfig any, T any] interface {
	MultiCall(
		chainId base_entities.ChainId, multiCallData []MultiCallSingle[T], providerConfig *config.Config,
	) (
		blockNumber *big.Int,
		blockHash common.Hash,
		err error,
	)
}

type BaseMultiCallProvider[T any] struct {
	core MultiCallProviderCore
}

func GetMultiCallProvider[T any](core MultiCallProviderCore) *BaseMultiCallProvider[T] {
	return &BaseMultiCallProvider[T]{
		core: core,
	}
}

func (b *BaseMultiCallProvider[T]) MultiCall(
	chainId base_entities.ChainId,
	multiCallData []MultiCallSingle[T],
	gas uint64,
	requireSuccess bool,
	providerConfig *config.Config,
) (
	blockNumber *big.Int,
	blockHash common.Hash,
	err error,
) {
	callOpts := &bind.CallOpts{}
	if providerConfig != nil {
		callOpts.BlockNumber = big.NewInt(int64(providerConfig.BlockNumber))
	}

	var calls []contracts.Multicall3Call
	for _, param := range multiCallData {
		callData, err := param.Contract.Pack(param.FunctionName, param.FunctionParams...)
		if err != nil {
			return blockNumber, blockHash, err
		}

		calls = append(calls, contracts.Multicall3Call{
			Target:   selfCommon.Address(param.ContractAddress),
			CallData: callData,
		})
	}

	decodeResults := struct {
		BlockNumber *big.Int
		BlockHash   common.Hash
		ReturnData  []contracts.Multicall3Result
	}{}
	client, err := b.core.GetMultiCallContract(chainId)
	if err != nil {
		return blockNumber, blockHash, err
	}
	method := "tryBlockAndAggregate"

	// Pack the input, call and unpack the results
	input, err := contracts.MultiCallABi.Pack(method, requireSuccess, calls)
	if err != nil {
		return blockNumber, blockHash, err
	}
	var (
		msg = ethereum.CallMsg{
			From: common.Address{},
			To:   base_constant.MultiCallAddress(chainId),
			Data: input,
			Gas:  gas,
		}
	)
	resp, err := client.CallContract(context.Background(), msg, callOpts.BlockNumber)
	if err != nil {
		return blockNumber, blockHash, err
	}
	results := &[]any{&decodeResults}
	if len(resp) == 0 {
		// Make sure we have a contract to operate on, and bail out otherwise.
		if code, err := client.CodeAt(context.Background(), *base_constant.MultiCallAddress(chainId), callOpts.BlockNumber); err != nil {
			return blockNumber, blockHash, err
		} else if len(code) == 0 {
			return blockNumber, blockHash, errors.New("error code")
		}
	}
	res := *results
	err = contracts.MultiCallABi.UnpackIntoInterface(res[0], method, resp)
	if err != nil {
		return blockNumber, blockHash, err
	}

	blockHash = decodeResults.BlockHash
	blockNumber = decodeResults.BlockNumber

	if len(decodeResults.ReturnData) != len(multiCallData) {
		return blockNumber, blockHash, errors.New("result len not equal to request params")
	}
	for i := range multiCallData {
		if !decodeResults.ReturnData[i].Success {
			continue
		}

		multiCallData[i].CallResult.Err = multiCallData[i].Contract.UnpackIntoInterface(
			&multiCallData[i].CallResult.Data,
			multiCallData[i].FunctionName,
			decodeResults.ReturnData[i].ReturnData,
		)
		if multiCallData[i].CallResult.Err == nil {
			multiCallData[i].CallResult.Success = true
		}
	}
	return
}

type MultiCallProviderCore interface {
	GetMultiCallContract(chainId base_entities.ChainId) (*ethclient.Client, error)
}
