package rpc

import (
	_ "embed"
	"github.com/coming-chat/intra-swap-core/base_constant"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/coming-chat/intra-swap-core/contracts"
	"github.com/coming-chat/intra-swap-core/providers/config"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func NewMultiCallV3(backend *ethclient.Client) (*contracts.Multicall3Raw, error) {
	mc3, err := contracts.NewMulticall3(common.HexToAddress(base_constant.MultiCall3Address), backend)
	if err != nil {
		return nil, err
	}
	return &contracts.Multicall3Raw{
		Contract: mc3,
	}, nil
}

type MultiCallSingleParam struct {
	Contract        *abi.ABI
	FunctionParams  []any
	ContractAddress common.Address
	FunctionName    string
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
		chainId base_entities.ChainId, multiCallData []MultiCallSingleParam, providerConfig *config.Config,
	) (
		result MultiCallResultWithInfo[T],
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

func (b *BaseMultiCallProvider[T]) MultiCall(chainId base_entities.ChainId, multiCallData []MultiCallSingleParam, requireSuccess bool, providerConfig *config.Config) (result MultiCallResultWithInfo[T], err error) {
	callOpts := &bind.CallOpts{}
	if providerConfig != nil {
		callOpts.BlockNumber = big.NewInt(int64(providerConfig.BlockNumber))
	}

	var calls []contracts.Multicall3Call
	for _, param := range multiCallData {
		callData, err := param.Contract.Pack(param.FunctionName, param.FunctionParams...)
		if err != nil {
			return result, err
		}

		calls = append(calls, contracts.Multicall3Call{
			Target:   param.ContractAddress,
			CallData: callData,
		})
	}

	decodeResults := struct {
		BlockNumber *big.Int
		BlockHash   common.Hash
		ReturnData  []contracts.Multicall3Result
	}{}
	multiCall, err := b.core.GetMultiCallContract(chainId)
	if err != nil {
		return result, err
	}
	err = multiCall.Call(callOpts, &[]any{&decodeResults}, "tryBlockAndAggregate", requireSuccess, calls)
	if err != nil {
		return
	}
	result.BlockHash = decodeResults.BlockHash
	result.BlockNumber = decodeResults.BlockNumber

	for i, data := range decodeResults.ReturnData {
		output := MultiCallResult[T]{}
		if data.Success {
			decodeErr := multiCallData[i].Contract.UnpackIntoInterface(
				&output.Data,
				multiCallData[i].FunctionName,
				data.ReturnData,
			)
			output.Err = decodeErr
			if output.Err == nil {
				output.Success = true
			}
		}
		result.ReturnData = append(result.ReturnData, output)
	}
	return
}

type MultiCallProviderCore interface {
	GetMultiCallContract(chainId base_entities.ChainId) (*contracts.Multicall3Raw, error)
}
