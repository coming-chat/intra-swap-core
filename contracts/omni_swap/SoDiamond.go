// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package omni_swap

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/gkirito/go-ethereum"
	"github.com/gkirito/go-ethereum/accounts/abi"
	"github.com/gkirito/go-ethereum/accounts/abi/bind"
	"github.com/gkirito/go-ethereum/common"
	"github.com/gkirito/go-ethereum/core/types"
	"github.com/gkirito/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BoolFacetBoolSwapData is an auto generated low-level Go binding around an user-defined struct.
type BoolFacetBoolSwapData struct {
	SrcBoolSwapPoolId  uint32
	DstBoolSwapChainId uint32
	DstSoDiamond       [32]byte
}

// CelerFacetCelerData is an auto generated low-level Go binding around an user-defined struct.
type CelerFacetCelerData struct {
	Sender                         common.Address
	MaxSlippage                    uint32
	DstCelerChainId                uint64
	BridgeToken                    common.Address
	DstMaxGasPriceInWeiForExecutor *big.Int
	EstimateCost                   *big.Int
	DstSoDiamond                   common.Address
}

// IDiamondCutFacetCut is an auto generated low-level Go binding around an user-defined struct.
type IDiamondCutFacetCut struct {
	FacetAddress      common.Address
	Action            uint8
	FunctionSelectors [][4]byte
}

// IDiamondLoupeFacet is an auto generated low-level Go binding around an user-defined struct.
type IDiamondLoupeFacet struct {
	FacetAddress      common.Address
	FunctionSelectors [][4]byte
}

// ISoNormalizedSoData is an auto generated low-level Go binding around an user-defined struct.
type ISoNormalizedSoData struct {
	TransactionId      []byte
	Receiver           []byte
	SourceChainId      uint16
	SendingAssetId     []byte
	DestinationChainId uint16
	ReceivingAssetId   []byte
	Amount             *big.Int
}

// ISoSoData is an auto generated low-level Go binding around an user-defined struct.
type ISoSoData struct {
	TransactionId      [32]byte
	Receiver           common.Address
	SourceChainId      uint16
	SendingAssetId     common.Address
	DestinationChainId uint16
	ReceivingAssetId   common.Address
	Amount             *big.Int
}

// LibSwapNormalizedSwapData is an auto generated low-level Go binding around an user-defined struct.
type LibSwapNormalizedSwapData struct {
	CallTo           []byte
	ApproveTo        []byte
	SendingAssetId   []byte
	ReceivingAssetId []byte
	FromAmount       *big.Int
	CallData         []byte
}

// LibSwapSwapData is an auto generated low-level Go binding around an user-defined struct.
type LibSwapSwapData struct {
	CallTo           common.Address
	ApproveTo        common.Address
	SendingAssetId   common.Address
	ReceivingAssetId common.Address
	FromAmount       *big.Int
	CallData         []byte
}

// MultiChainFacetMultiChainData is an auto generated low-level Go binding around an user-defined struct.
type MultiChainFacetMultiChainData struct {
	DstChainId   uint64
	BridgeToken  common.Address
	DstSoDiamond string
}

// StargateFacetStargateData is an auto generated low-level Go binding around an user-defined struct.
type StargateFacetStargateData struct {
	SrcStargatePoolId  *big.Int
	DstStargateChainId uint16
	DstStargatePoolId  *big.Int
	MinAmount          *big.Int
	DstGasForSgReceive *big.Int
	DstSoDiamond       common.Address
}

// WormholeFacetNormalizedWormholeData is an auto generated low-level Go binding around an user-defined struct.
type WormholeFacetNormalizedWormholeData struct {
	DstWormholeChainId            uint16
	DstMaxGasPriceInWeiForRelayer *big.Int
	WormholeFee                   *big.Int
	DstSoDiamond                  []byte
}

// SoDiamondMetaData contains all meta data concerning the SoDiamond contract.
var SoDiamondMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"facetCut\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"init\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"name\":\"diamondCut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"functionSelector\",\"type\":\"bytes4\"}],\"name\":\"facetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addresses\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facetAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"facet\",\"type\":\"address\"}],\"name\":\"facetFunctionSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"facetFunctionSelectors_\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondLoupe.Facet[]\",\"name\":\"allFacets\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"correctSwap\",\"type\":\"address\"}],\"name\":\"CorrectSwapRouterSelectorsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dexAddress\",\"type\":\"address\"}],\"name\":\"DexAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dexAddress\",\"type\":\"address\"}],\"name\":\"DexRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"functionSignature\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"FunctionSignatureApprovalChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gatewayAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"soFeeAddress\",\"type\":\"address\"}],\"name\":\"GatewaySoFeeSelectorsChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"correctSwap\",\"type\":\"address\"}],\"name\":\"addCorrectSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dex\",\"type\":\"address\"}],\"name\":\"addDex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"soFee\",\"type\":\"address\"}],\"name\":\"addFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approvedDexs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"dexs\",\"type\":\"address[]\"}],\"name\":\"batchAddDex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"dexs\",\"type\":\"address[]\"}],\"name\":\"batchRemoveDex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"signatures\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool\",\"name\":\"approval\",\"type\":\"bool\"}],\"name\":\"batchSetFunctionApprovalBySignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signature\",\"type\":\"bytes32\"}],\"name\":\"isFunctionApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dex\",\"type\":\"address\"}],\"name\":\"removeDex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signature\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"approval\",\"type\":\"bool\"}],\"name\":\"setFunctionApprovalBySignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CannotBridgeToSameNetwork\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ContractCallNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeAssetTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSwapDataProvided\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSwapFromZeroBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTransferToNullAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSupportedSwapRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullAddrIsNotAValidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullAddrIsNotAnERC20Token\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"name\":\"SetAllowedList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receiveAmount\",\"type\":\"uint256\"}],\"name\":\"SoTransferCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"revertReason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"otherReason\",\"type\":\"bytes\"}],\"name\":\"SoTransferFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"SoTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stargate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"StargateInitialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INTERDELIMITER\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"stargatePayload\",\"type\":\"bytes\"}],\"name\":\"decodeStargatePayload\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDst\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDst\",\"type\":\"tuple[]\"}],\"name\":\"encodeStargatePayload\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcStargatePoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"dstStargateChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"dstStargatePoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstGasForSgReceive\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"dstSoDiamond\",\"type\":\"address\"}],\"internalType\":\"structStargateFacet.StargateData\",\"name\":\"stargateData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"estimateStargateFinalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.SoData\",\"name\":\"soData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"callTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"approveTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.SwapData[]\",\"name\":\"swapData\",\"type\":\"tuple[]\"}],\"name\":\"executeAndCheckSwaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getAmountBeforeSoFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soDataNo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDstNo\",\"type\":\"tuple[]\"}],\"name\":\"getSgReceiveForGasPayload\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soDataNo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcStargatePoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"dstStargateChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"dstStargatePoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstGasForSgReceive\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"dstSoDiamond\",\"type\":\"address\"}],\"internalType\":\"structStargateFacet.StargateData\",\"name\":\"stargateData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDstNo\",\"type\":\"tuple[]\"}],\"name\":\"getStargateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getStargateSoFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransferGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stargate\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"}],\"name\":\"initStargate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.SoData\",\"name\":\"soData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"callTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"approveTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.SwapData[]\",\"name\":\"swapDataDst\",\"type\":\"tuple[]\"}],\"name\":\"remoteSoSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"name\":\"setAllowedAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"sgReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soDataNo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"dstStargatePoolId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDstNo\",\"type\":\"tuple[]\"}],\"name\":\"sgReceiveForGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soDataNo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataSrcNo\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcStargatePoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"dstStargateChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"dstStargatePoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstGasForSgReceive\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"dstSoDiamond\",\"type\":\"address\"}],\"internalType\":\"structStargateFacet.StargateData\",\"name\":\"stargateData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDstNo\",\"type\":\"tuple[]\"}],\"name\":\"soSwapViaStargate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"NotEnoughBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogWithdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NewOwnerMustNotBeSelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoNullOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPendingOwnershipTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotPendingOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"cancelOnwershipTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmOwnershipTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"contractOwner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ContractCallNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeAssetTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSwapDataProvided\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSwapFromZeroBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTransferToNullAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSupportedSwapRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullAddrIsNotAValidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullAddrIsNotAnERC20Token\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromAssetId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toAssetId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"}],\"name\":\"SoSwappedGeneric\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receiveAmount\",\"type\":\"uint256\"}],\"name\":\"SoTransferCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"revertReason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"otherReason\",\"type\":\"bytes\"}],\"name\":\"SoTransferFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"SoTransferStarted\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.SoData\",\"name\":\"soData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"callTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"approveTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.SwapData[]\",\"name\":\"swapData\",\"type\":\"tuple[]\"}],\"name\":\"executeAndCheckSwaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soDataNo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataNo\",\"type\":\"tuple[]\"}],\"name\":\"swapTokensGeneric\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ContractCallNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeAssetTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSwapDataProvided\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSwapFromZeroBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTransferToNullAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSupportedSwapRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullAddrIsNotAValidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullAddrIsNotAnERC20Token\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenBridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"srcWormholeChainId\",\"type\":\"uint16\"}],\"name\":\"InitWormholeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receiveAmount\",\"type\":\"uint256\"}],\"name\":\"SoTransferCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"revertReason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"otherReason\",\"type\":\"bytes\"}],\"name\":\"SoTransferFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"SoTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"srcWormholeChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"dstWormholeChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"TransferFromWormhole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"dstWormholeChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPerBytes\",\"type\":\"uint256\"}],\"name\":\"UpdateWormholeGas\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualReserve\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"estimateReserve\",\"type\":\"uint256\"}],\"name\":\"UpdateWormholeReserve\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"dstWormholeChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"dstMaxGasPriceInWeiForRelayer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wormholeFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"dstSoDiamond\",\"type\":\"bytes\"}],\"internalType\":\"structWormholeFacet.NormalizedWormholeData\",\"name\":\"wormholeData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDst\",\"type\":\"tuple[]\"}],\"name\":\"checkRelayerFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodeVm\",\"type\":\"bytes\"}],\"name\":\"completeSoSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodeVm\",\"type\":\"bytes\"}],\"name\":\"completeSoSwapByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"wormholeData\",\"type\":\"bytes\"}],\"name\":\"decodeNormalizedWormholeData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"dstWormholeChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"dstMaxGasPriceInWeiForRelayer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wormholeFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"dstSoDiamond\",\"type\":\"bytes\"}],\"internalType\":\"structWormholeFacet.NormalizedWormholeData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"wormholeData\",\"type\":\"bytes\"}],\"name\":\"decodeWormholePayload\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"dstWormholeChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"dstMaxGasPriceInWeiForRelayer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wormholeFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"dstSoDiamond\",\"type\":\"bytes\"}],\"internalType\":\"structWormholeFacet.NormalizedWormholeData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"encodeNormalizedWormholeData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstMaxGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstMaxGas\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDst\",\"type\":\"tuple[]\"}],\"name\":\"encodeWormholePayload\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"dstWormholeChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"dstMaxGasPriceInWeiForRelayer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wormholeFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"dstSoDiamond\",\"type\":\"bytes\"}],\"internalType\":\"structWormholeFacet.NormalizedWormholeData\",\"name\":\"wormholeData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDst\",\"type\":\"tuple[]\"}],\"name\":\"estimateCompleteSoSwapGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"dstWormholeChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"dstMaxGasPriceInWeiForRelayer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wormholeFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"dstSoDiamond\",\"type\":\"bytes\"}],\"internalType\":\"structWormholeFacet.NormalizedWormholeData\",\"name\":\"wormholeData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDst\",\"type\":\"tuple[]\"}],\"name\":\"estimateRelayerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.SoData\",\"name\":\"soData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"callTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"approveTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.SwapData[]\",\"name\":\"swapData\",\"type\":\"tuple[]\"}],\"name\":\"executeAndCheckSwaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transferToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getWormholeFinalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWormholeMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getWormholeSoFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenBridge\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"wormholeChainId\",\"type\":\"uint16\"}],\"name\":\"initWormhole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"dstWormholeChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerBytes\",\"type\":\"uint256\"}],\"name\":\"setWormholeGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"estimateReserve\",\"type\":\"uint256\"}],\"name\":\"setWormholeReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soDataNo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataSrcNo\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"dstWormholeChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"dstMaxGasPriceInWeiForRelayer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wormholeFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"dstSoDiamond\",\"type\":\"bytes\"}],\"internalType\":\"structWormholeFacet.NormalizedWormholeData\",\"name\":\"wormholeDataNo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDstNo\",\"type\":\"tuple[]\"}],\"name\":\"soSwapViaWormhole\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"soData\",\"type\":\"bytes\"}],\"name\":\"decodeNormalizedSoData\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"swapData\",\"type\":\"bytes\"}],\"name\":\"decodeNormalizedSwapData\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"denormalizeSoData\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.SoData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"data\",\"type\":\"tuple[]\"}],\"name\":\"denormalizeSwapData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"callTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"approveTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.SwapData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"denormalizeU256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"encodeNormalizedSoData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"data\",\"type\":\"tuple[]\"}],\"name\":\"encodeNormalizedSwapData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.SoData\",\"name\":\"soData\",\"type\":\"tuple\"}],\"name\":\"normalizeSoData\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"callTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"approveTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.SwapData[]\",\"name\":\"swapData\",\"type\":\"tuple[]\"}],\"name\":\"normalizeSwapData\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"normalizeU256\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CannotBridgeToSameNetwork\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ContractCallNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeAssetTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSwapDataProvided\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSwapFromZeroBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTransferToNullAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSupportedSwapRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullAddrIsNotAValidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullAddrIsNotAnERC20Token\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"messageBus\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"CelerInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTxId\",\"type\":\"bytes32\"}],\"name\":\"RefundCelerToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"messageBus\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"name\":\"SetAllowedList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"messageBus\",\"type\":\"address\"}],\"name\":\"SetMessageBus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receiveAmount\",\"type\":\"uint256\"}],\"name\":\"SoTransferCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"revertReason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"otherReason\",\"type\":\"bytes\"}],\"name\":\"SoTransferFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"SoTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"celerTransferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcCelerChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"dstCelerChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bridgeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"TransferFromCeler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualReserve\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"estimateReserve\",\"type\":\"uint256\"}],\"name\":\"UpdateCelerReserve\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GasPerByte\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxSlippage\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"dstCelerChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"bridgeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstMaxGasPriceInWeiForExecutor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"estimateCost\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"dstSoDiamond\",\"type\":\"address\"}],\"internalType\":\"structCelerFacet.CelerData\",\"name\":\"celerData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDst\",\"type\":\"tuple[]\"}],\"name\":\"checkExecutorFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"celerPayload\",\"type\":\"bytes\"}],\"name\":\"decodeCelerPayload\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soDataNo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDstNo\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soDataNo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDstNo\",\"type\":\"tuple[]\"}],\"name\":\"encodeCelerPayload\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstCelerChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"dstMaxGasPriceInWeiForExecutor\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soDataNo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDstNo\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"is_actual\",\"type\":\"bool\"}],\"name\":\"estCelerMessageFeeAndExecutorFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.SoData\",\"name\":\"soData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"callTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"approveTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.SwapData[]\",\"name\":\"swapData\",\"type\":\"tuple[]\"}],\"name\":\"executeAndCheckSwaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_sender\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"enumICelerMessageReceiver.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"enumICelerMessageReceiver.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[{\"internalType\":\"enumICelerMessageReceiver.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessageWithTransferFallback\",\"outputs\":[{\"internalType\":\"enumICelerMessageReceiver.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[{\"internalType\":\"enumICelerMessageReceiver.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"}],\"name\":\"getBaseGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"messageBus\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"getCelerMessageFee1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"getCelerMessageFee2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getCelerSoFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExecutorFeeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"messageBus\",\"type\":\"address\"}],\"name\":\"getNativeWrap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"messageBus\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"initCeler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"messageBus\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"name\":\"setAllowedAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"dstChainIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint256\",\"name\":\"dstBaseGas\",\"type\":\"uint256\"}],\"name\":\"setBaseGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"estimateReserve\",\"type\":\"uint256\"}],\"name\":\"setCelerReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeTo\",\"type\":\"address\"}],\"name\":\"setExecutorFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"setNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soDataNo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataSrcNo\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxSlippage\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"dstCelerChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"bridgeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstMaxGasPriceInWeiForExecutor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"estimateCost\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"dstSoDiamond\",\"type\":\"address\"}],\"internalType\":\"structCelerFacet.CelerData\",\"name\":\"celerData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDstNo\",\"type\":\"tuple[]\"}],\"name\":\"soSwapViaCeler\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CannotBridgeToSameNetwork\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ContractCallNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeAssetTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSwapDataProvided\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSwapFromZeroBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTransferToNullAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSupportedSwapRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullAddrIsNotAValidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullAddrIsNotAnERC20Token\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"anyTokens\",\"type\":\"address[]\"}],\"name\":\"AnyMappingUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fastRouter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"MultiChainInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"name\":\"SetAllowedList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receiveAmount\",\"type\":\"uint256\"}],\"name\":\"SoTransferCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"revertReason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"otherReason\",\"type\":\"bytes\"}],\"name\":\"SoTransferFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"SoTransferStarted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"multiChainPayload\",\"type\":\"bytes\"}],\"name\":\"decodeMultiChainPayload\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soDataNo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDstNo\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soDataNo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDstNo\",\"type\":\"tuple[]\"}],\"name\":\"encodeMultiChainPayload\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"exec\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.SoData\",\"name\":\"soData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"callTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"approveTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.SwapData[]\",\"name\":\"swapData\",\"type\":\"tuple[]\"}],\"name\":\"executeAndCheckSwaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeToken\",\"type\":\"address\"}],\"name\":\"getAnyToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFastRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getMultiChainSoFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fastRouter\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"initMultiChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isValidMultiChainConfig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"name\":\"setAllowedAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soDataNo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataSrcNo\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"bridgeToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"dstSoDiamond\",\"type\":\"string\"}],\"internalType\":\"structMultiChainFacet.MultiChainData\",\"name\":\"multiChainData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDstNo\",\"type\":\"tuple[]\"}],\"name\":\"soSwapViaMultiChain\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"anyTokens\",\"type\":\"address[]\"}],\"name\":\"updateAddressMappings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ContractCallNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeAssetTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSwapDataProvided\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSwapFromZeroBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTransferToNullAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSupportedSwapRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullAddrIsNotAValidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullAddrIsNotAnERC20Token\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"boolSwapRouter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"boolSwapFactory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"}],\"name\":\"BoolSwapInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"consumerChainId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"boolChainId\",\"type\":\"uint32\"}],\"name\":\"PathPairUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"boolSwapPool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"name\":\"SetAllowedList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receiveAmount\",\"type\":\"uint256\"}],\"name\":\"SoTransferCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"revertReason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"otherReason\",\"type\":\"bytes\"}],\"name\":\"SoTransferFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"SoTransferStarted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"boolSwapPools\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"isAllowed\",\"type\":\"bool[]\"}],\"name\":\"batchSetBoolAllowedAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"stargatePayload\",\"type\":\"bytes\"}],\"name\":\"decodeBoolSwapPayload\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDst\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDst\",\"type\":\"tuple[]\"}],\"name\":\"encodeBoolSwapPayload\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soDataNo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcBoolSwapPoolId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"dstBoolSwapChainId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"dstSoDiamond\",\"type\":\"bytes32\"}],\"internalType\":\"structBoolFacet.BoolSwapData\",\"name\":\"boolSwapData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"bridgeAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDstNo\",\"type\":\"tuple[]\"}],\"name\":\"estimateBoolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.SoData\",\"name\":\"soData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"callTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"approveTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.SwapData[]\",\"name\":\"swapData\",\"type\":\"tuple[]\"}],\"name\":\"executeAndCheckSwaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"boolChainId\",\"type\":\"uint32\"}],\"name\":\"fromBoolToConsumer\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"consumerChainId\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"consumerChainId\",\"type\":\"uint32\"}],\"name\":\"fromConsumerToBool\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"boolChainId\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBoolBasicBeneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBoolBasicFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getBoolSoFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBoolTransferGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"}],\"name\":\"initBoolSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"bridgeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bridgeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"receiveFromBoolSwap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.SoData\",\"name\":\"soData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"callTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"approveTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.SwapData[]\",\"name\":\"swapDataDst\",\"type\":\"tuple[]\"}],\"name\":\"remoteBoolSoSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactionId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"sourceChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"destinationChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISo.NormalizedSoData\",\"name\":\"soDataNo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataSrcNo\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcBoolSwapPoolId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"dstBoolSwapChainId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"dstSoDiamond\",\"type\":\"bytes32\"}],\"internalType\":\"structBoolFacet.BoolSwapData\",\"name\":\"boolSwapData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"callTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approveTo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sendingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receivingAssetId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structLibSwap.NormalizedSwapData[]\",\"name\":\"swapDataDstNo\",\"type\":\"tuple[]\"}],\"name\":\"soSwapViaBool\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"consumerChainIds\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"boolChainIds\",\"type\":\"uint32[]\"}],\"name\":\"updatePathPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NoSwapFromZeroBalance\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dex\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromAssetId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toAssetId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"AssetSwapped\",\"type\":\"event\"}]",
}

// SoDiamondABI is the input ABI used to generate the binding from.
// Deprecated: Use SoDiamondMetaData.ABI instead.
var SoDiamondABI = SoDiamondMetaData.ABI

// SoDiamond is an auto generated Go binding around an Ethereum contract.
type SoDiamond struct {
	SoDiamondCaller     // Read-only binding to the contract
	SoDiamondTransactor // Write-only binding to the contract
	SoDiamondFilterer   // Log filterer for contract events
}

// SoDiamondCaller is an auto generated read-only Go binding around an Ethereum contract.
type SoDiamondCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SoDiamondTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SoDiamondTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SoDiamondFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SoDiamondFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SoDiamondSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SoDiamondSession struct {
	Contract     *SoDiamond        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SoDiamondCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SoDiamondCallerSession struct {
	Contract *SoDiamondCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SoDiamondTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SoDiamondTransactorSession struct {
	Contract     *SoDiamondTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SoDiamondRaw is an auto generated low-level Go binding around an Ethereum contract.
type SoDiamondRaw struct {
	Contract *SoDiamond // Generic contract binding to access the raw methods on
}

// SoDiamondCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SoDiamondCallerRaw struct {
	Contract *SoDiamondCaller // Generic read-only contract binding to access the raw methods on
}

// SoDiamondTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SoDiamondTransactorRaw struct {
	Contract *SoDiamondTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSoDiamond creates a new instance of SoDiamond, bound to a specific deployed contract.
func NewSoDiamond(address common.Address, backend bind.ContractBackend) (*SoDiamond, error) {
	contract, err := bindSoDiamond(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SoDiamond{SoDiamondCaller: SoDiamondCaller{contract: contract}, SoDiamondTransactor: SoDiamondTransactor{contract: contract}, SoDiamondFilterer: SoDiamondFilterer{contract: contract}}, nil
}

// NewSoDiamondCaller creates a new read-only instance of SoDiamond, bound to a specific deployed contract.
func NewSoDiamondCaller(address common.Address, caller bind.ContractCaller) (*SoDiamondCaller, error) {
	contract, err := bindSoDiamond(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SoDiamondCaller{contract: contract}, nil
}

// NewSoDiamondTransactor creates a new write-only instance of SoDiamond, bound to a specific deployed contract.
func NewSoDiamondTransactor(address common.Address, transactor bind.ContractTransactor) (*SoDiamondTransactor, error) {
	contract, err := bindSoDiamond(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SoDiamondTransactor{contract: contract}, nil
}

// NewSoDiamondFilterer creates a new log filterer instance of SoDiamond, bound to a specific deployed contract.
func NewSoDiamondFilterer(address common.Address, filterer bind.ContractFilterer) (*SoDiamondFilterer, error) {
	contract, err := bindSoDiamond(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SoDiamondFilterer{contract: contract}, nil
}

// bindSoDiamond binds a generic wrapper to an already deployed contract.
func bindSoDiamond(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SoDiamondMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SoDiamond *SoDiamondRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SoDiamond.Contract.SoDiamondCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SoDiamond *SoDiamondRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SoDiamond.Contract.SoDiamondTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SoDiamond *SoDiamondRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SoDiamond.Contract.SoDiamondTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SoDiamond *SoDiamondCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SoDiamond.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SoDiamond *SoDiamondTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SoDiamond.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SoDiamond *SoDiamondTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SoDiamond.Contract.contract.Transact(opts, method, params...)
}

// GasPerByte is a free data retrieval call binding the contract method 0x932ca683.
//
// Solidity: function GasPerByte() view returns(uint256)
func (_SoDiamond *SoDiamondCaller) GasPerByte(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "GasPerByte")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasPerByte is a free data retrieval call binding the contract method 0x932ca683.
//
// Solidity: function GasPerByte() view returns(uint256)
func (_SoDiamond *SoDiamondSession) GasPerByte() (*big.Int, error) {
	return _SoDiamond.Contract.GasPerByte(&_SoDiamond.CallOpts)
}

// GasPerByte is a free data retrieval call binding the contract method 0x932ca683.
//
// Solidity: function GasPerByte() view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) GasPerByte() (*big.Int, error) {
	return _SoDiamond.Contract.GasPerByte(&_SoDiamond.CallOpts)
}

// INTERDELIMITER is a free data retrieval call binding the contract method 0x80d0d656.
//
// Solidity: function INTERDELIMITER() view returns(uint8)
func (_SoDiamond *SoDiamondCaller) INTERDELIMITER(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "INTERDELIMITER")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INTERDELIMITER is a free data retrieval call binding the contract method 0x80d0d656.
//
// Solidity: function INTERDELIMITER() view returns(uint8)
func (_SoDiamond *SoDiamondSession) INTERDELIMITER() (uint8, error) {
	return _SoDiamond.Contract.INTERDELIMITER(&_SoDiamond.CallOpts)
}

// INTERDELIMITER is a free data retrieval call binding the contract method 0x80d0d656.
//
// Solidity: function INTERDELIMITER() view returns(uint8)
func (_SoDiamond *SoDiamondCallerSession) INTERDELIMITER() (uint8, error) {
	return _SoDiamond.Contract.INTERDELIMITER(&_SoDiamond.CallOpts)
}

// RAY is a free data retrieval call binding the contract method 0x552033c4.
//
// Solidity: function RAY() view returns(uint256)
func (_SoDiamond *SoDiamondCaller) RAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RAY is a free data retrieval call binding the contract method 0x552033c4.
//
// Solidity: function RAY() view returns(uint256)
func (_SoDiamond *SoDiamondSession) RAY() (*big.Int, error) {
	return _SoDiamond.Contract.RAY(&_SoDiamond.CallOpts)
}

// RAY is a free data retrieval call binding the contract method 0x552033c4.
//
// Solidity: function RAY() view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) RAY() (*big.Int, error) {
	return _SoDiamond.Contract.RAY(&_SoDiamond.CallOpts)
}

// RAY0 is a free data retrieval call binding the contract method 0x552033c4.
//
// Solidity: function RAY() view returns(uint256)
func (_SoDiamond *SoDiamondCaller) RAY0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "RAY0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RAY0 is a free data retrieval call binding the contract method 0x552033c4.
//
// Solidity: function RAY() view returns(uint256)
func (_SoDiamond *SoDiamondSession) RAY0() (*big.Int, error) {
	return _SoDiamond.Contract.RAY0(&_SoDiamond.CallOpts)
}

// RAY0 is a free data retrieval call binding the contract method 0x552033c4.
//
// Solidity: function RAY() view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) RAY0() (*big.Int, error) {
	return _SoDiamond.Contract.RAY0(&_SoDiamond.CallOpts)
}

// ApprovedDexs is a free data retrieval call binding the contract method 0xfbb2d381.
//
// Solidity: function approvedDexs() view returns(address[] addresses)
func (_SoDiamond *SoDiamondCaller) ApprovedDexs(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "approvedDexs")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ApprovedDexs is a free data retrieval call binding the contract method 0xfbb2d381.
//
// Solidity: function approvedDexs() view returns(address[] addresses)
func (_SoDiamond *SoDiamondSession) ApprovedDexs() ([]common.Address, error) {
	return _SoDiamond.Contract.ApprovedDexs(&_SoDiamond.CallOpts)
}

// ApprovedDexs is a free data retrieval call binding the contract method 0xfbb2d381.
//
// Solidity: function approvedDexs() view returns(address[] addresses)
func (_SoDiamond *SoDiamondCallerSession) ApprovedDexs() ([]common.Address, error) {
	return _SoDiamond.Contract.ApprovedDexs(&_SoDiamond.CallOpts)
}

// DecodeBoolSwapPayload is a free data retrieval call binding the contract method 0xc736fbd8.
//
// Solidity: function decodeBoolSwapPayload(bytes stargatePayload) pure returns((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst)
func (_SoDiamond *SoDiamondCaller) DecodeBoolSwapPayload(opts *bind.CallOpts, stargatePayload []byte) (struct {
	SoData      ISoNormalizedSoData
	SwapDataDst []LibSwapNormalizedSwapData
}, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "decodeBoolSwapPayload", stargatePayload)

	outstruct := new(struct {
		SoData      ISoNormalizedSoData
		SwapDataDst []LibSwapNormalizedSwapData
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SoData = *abi.ConvertType(out[0], new(ISoNormalizedSoData)).(*ISoNormalizedSoData)
	outstruct.SwapDataDst = *abi.ConvertType(out[1], new([]LibSwapNormalizedSwapData)).(*[]LibSwapNormalizedSwapData)

	return *outstruct, err

}

// DecodeBoolSwapPayload is a free data retrieval call binding the contract method 0xc736fbd8.
//
// Solidity: function decodeBoolSwapPayload(bytes stargatePayload) pure returns((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst)
func (_SoDiamond *SoDiamondSession) DecodeBoolSwapPayload(stargatePayload []byte) (struct {
	SoData      ISoNormalizedSoData
	SwapDataDst []LibSwapNormalizedSwapData
}, error) {
	return _SoDiamond.Contract.DecodeBoolSwapPayload(&_SoDiamond.CallOpts, stargatePayload)
}

// DecodeBoolSwapPayload is a free data retrieval call binding the contract method 0xc736fbd8.
//
// Solidity: function decodeBoolSwapPayload(bytes stargatePayload) pure returns((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst)
func (_SoDiamond *SoDiamondCallerSession) DecodeBoolSwapPayload(stargatePayload []byte) (struct {
	SoData      ISoNormalizedSoData
	SwapDataDst []LibSwapNormalizedSwapData
}, error) {
	return _SoDiamond.Contract.DecodeBoolSwapPayload(&_SoDiamond.CallOpts, stargatePayload)
}

// DecodeCelerPayload is a free data retrieval call binding the contract method 0x68341d44.
//
// Solidity: function decodeCelerPayload(bytes celerPayload) pure returns(address, (bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo)
func (_SoDiamond *SoDiamondCaller) DecodeCelerPayload(opts *bind.CallOpts, celerPayload []byte) (common.Address, ISoNormalizedSoData, []LibSwapNormalizedSwapData, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "decodeCelerPayload", celerPayload)

	if err != nil {
		return *new(common.Address), *new(ISoNormalizedSoData), *new([]LibSwapNormalizedSwapData), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(ISoNormalizedSoData)).(*ISoNormalizedSoData)
	out2 := *abi.ConvertType(out[2], new([]LibSwapNormalizedSwapData)).(*[]LibSwapNormalizedSwapData)

	return out0, out1, out2, err

}

// DecodeCelerPayload is a free data retrieval call binding the contract method 0x68341d44.
//
// Solidity: function decodeCelerPayload(bytes celerPayload) pure returns(address, (bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo)
func (_SoDiamond *SoDiamondSession) DecodeCelerPayload(celerPayload []byte) (common.Address, ISoNormalizedSoData, []LibSwapNormalizedSwapData, error) {
	return _SoDiamond.Contract.DecodeCelerPayload(&_SoDiamond.CallOpts, celerPayload)
}

// DecodeCelerPayload is a free data retrieval call binding the contract method 0x68341d44.
//
// Solidity: function decodeCelerPayload(bytes celerPayload) pure returns(address, (bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo)
func (_SoDiamond *SoDiamondCallerSession) DecodeCelerPayload(celerPayload []byte) (common.Address, ISoNormalizedSoData, []LibSwapNormalizedSwapData, error) {
	return _SoDiamond.Contract.DecodeCelerPayload(&_SoDiamond.CallOpts, celerPayload)
}

// DecodeMultiChainPayload is a free data retrieval call binding the contract method 0x124e8a54.
//
// Solidity: function decodeMultiChainPayload(bytes multiChainPayload) pure returns((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo)
func (_SoDiamond *SoDiamondCaller) DecodeMultiChainPayload(opts *bind.CallOpts, multiChainPayload []byte) (struct {
	SoDataNo      ISoNormalizedSoData
	SwapDataDstNo []LibSwapNormalizedSwapData
}, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "decodeMultiChainPayload", multiChainPayload)

	outstruct := new(struct {
		SoDataNo      ISoNormalizedSoData
		SwapDataDstNo []LibSwapNormalizedSwapData
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SoDataNo = *abi.ConvertType(out[0], new(ISoNormalizedSoData)).(*ISoNormalizedSoData)
	outstruct.SwapDataDstNo = *abi.ConvertType(out[1], new([]LibSwapNormalizedSwapData)).(*[]LibSwapNormalizedSwapData)

	return *outstruct, err

}

// DecodeMultiChainPayload is a free data retrieval call binding the contract method 0x124e8a54.
//
// Solidity: function decodeMultiChainPayload(bytes multiChainPayload) pure returns((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo)
func (_SoDiamond *SoDiamondSession) DecodeMultiChainPayload(multiChainPayload []byte) (struct {
	SoDataNo      ISoNormalizedSoData
	SwapDataDstNo []LibSwapNormalizedSwapData
}, error) {
	return _SoDiamond.Contract.DecodeMultiChainPayload(&_SoDiamond.CallOpts, multiChainPayload)
}

// DecodeMultiChainPayload is a free data retrieval call binding the contract method 0x124e8a54.
//
// Solidity: function decodeMultiChainPayload(bytes multiChainPayload) pure returns((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo)
func (_SoDiamond *SoDiamondCallerSession) DecodeMultiChainPayload(multiChainPayload []byte) (struct {
	SoDataNo      ISoNormalizedSoData
	SwapDataDstNo []LibSwapNormalizedSwapData
}, error) {
	return _SoDiamond.Contract.DecodeMultiChainPayload(&_SoDiamond.CallOpts, multiChainPayload)
}

// DecodeNormalizedSoData is a free data retrieval call binding the contract method 0x7f75f81f.
//
// Solidity: function decodeNormalizedSoData(bytes soData) pure returns((bytes,bytes,uint16,bytes,uint16,bytes,uint256))
func (_SoDiamond *SoDiamondCaller) DecodeNormalizedSoData(opts *bind.CallOpts, soData []byte) (ISoNormalizedSoData, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "decodeNormalizedSoData", soData)

	if err != nil {
		return *new(ISoNormalizedSoData), err
	}

	out0 := *abi.ConvertType(out[0], new(ISoNormalizedSoData)).(*ISoNormalizedSoData)

	return out0, err

}

// DecodeNormalizedSoData is a free data retrieval call binding the contract method 0x7f75f81f.
//
// Solidity: function decodeNormalizedSoData(bytes soData) pure returns((bytes,bytes,uint16,bytes,uint16,bytes,uint256))
func (_SoDiamond *SoDiamondSession) DecodeNormalizedSoData(soData []byte) (ISoNormalizedSoData, error) {
	return _SoDiamond.Contract.DecodeNormalizedSoData(&_SoDiamond.CallOpts, soData)
}

// DecodeNormalizedSoData is a free data retrieval call binding the contract method 0x7f75f81f.
//
// Solidity: function decodeNormalizedSoData(bytes soData) pure returns((bytes,bytes,uint16,bytes,uint16,bytes,uint256))
func (_SoDiamond *SoDiamondCallerSession) DecodeNormalizedSoData(soData []byte) (ISoNormalizedSoData, error) {
	return _SoDiamond.Contract.DecodeNormalizedSoData(&_SoDiamond.CallOpts, soData)
}

// DecodeNormalizedSwapData is a free data retrieval call binding the contract method 0x741eab15.
//
// Solidity: function decodeNormalizedSwapData(bytes swapData) pure returns((bytes,bytes,bytes,bytes,uint256,bytes)[])
func (_SoDiamond *SoDiamondCaller) DecodeNormalizedSwapData(opts *bind.CallOpts, swapData []byte) ([]LibSwapNormalizedSwapData, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "decodeNormalizedSwapData", swapData)

	if err != nil {
		return *new([]LibSwapNormalizedSwapData), err
	}

	out0 := *abi.ConvertType(out[0], new([]LibSwapNormalizedSwapData)).(*[]LibSwapNormalizedSwapData)

	return out0, err

}

// DecodeNormalizedSwapData is a free data retrieval call binding the contract method 0x741eab15.
//
// Solidity: function decodeNormalizedSwapData(bytes swapData) pure returns((bytes,bytes,bytes,bytes,uint256,bytes)[])
func (_SoDiamond *SoDiamondSession) DecodeNormalizedSwapData(swapData []byte) ([]LibSwapNormalizedSwapData, error) {
	return _SoDiamond.Contract.DecodeNormalizedSwapData(&_SoDiamond.CallOpts, swapData)
}

// DecodeNormalizedSwapData is a free data retrieval call binding the contract method 0x741eab15.
//
// Solidity: function decodeNormalizedSwapData(bytes swapData) pure returns((bytes,bytes,bytes,bytes,uint256,bytes)[])
func (_SoDiamond *SoDiamondCallerSession) DecodeNormalizedSwapData(swapData []byte) ([]LibSwapNormalizedSwapData, error) {
	return _SoDiamond.Contract.DecodeNormalizedSwapData(&_SoDiamond.CallOpts, swapData)
}

// DecodeNormalizedWormholeData is a free data retrieval call binding the contract method 0x8f66f150.
//
// Solidity: function decodeNormalizedWormholeData(bytes wormholeData) pure returns((uint16,uint256,uint256,bytes))
func (_SoDiamond *SoDiamondCaller) DecodeNormalizedWormholeData(opts *bind.CallOpts, wormholeData []byte) (WormholeFacetNormalizedWormholeData, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "decodeNormalizedWormholeData", wormholeData)

	if err != nil {
		return *new(WormholeFacetNormalizedWormholeData), err
	}

	out0 := *abi.ConvertType(out[0], new(WormholeFacetNormalizedWormholeData)).(*WormholeFacetNormalizedWormholeData)

	return out0, err

}

// DecodeNormalizedWormholeData is a free data retrieval call binding the contract method 0x8f66f150.
//
// Solidity: function decodeNormalizedWormholeData(bytes wormholeData) pure returns((uint16,uint256,uint256,bytes))
func (_SoDiamond *SoDiamondSession) DecodeNormalizedWormholeData(wormholeData []byte) (WormholeFacetNormalizedWormholeData, error) {
	return _SoDiamond.Contract.DecodeNormalizedWormholeData(&_SoDiamond.CallOpts, wormholeData)
}

// DecodeNormalizedWormholeData is a free data retrieval call binding the contract method 0x8f66f150.
//
// Solidity: function decodeNormalizedWormholeData(bytes wormholeData) pure returns((uint16,uint256,uint256,bytes))
func (_SoDiamond *SoDiamondCallerSession) DecodeNormalizedWormholeData(wormholeData []byte) (WormholeFacetNormalizedWormholeData, error) {
	return _SoDiamond.Contract.DecodeNormalizedWormholeData(&_SoDiamond.CallOpts, wormholeData)
}

// DecodeStargatePayload is a free data retrieval call binding the contract method 0x762a9962.
//
// Solidity: function decodeStargatePayload(bytes stargatePayload) pure returns((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst)
func (_SoDiamond *SoDiamondCaller) DecodeStargatePayload(opts *bind.CallOpts, stargatePayload []byte) (struct {
	SoData      ISoNormalizedSoData
	SwapDataDst []LibSwapNormalizedSwapData
}, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "decodeStargatePayload", stargatePayload)

	outstruct := new(struct {
		SoData      ISoNormalizedSoData
		SwapDataDst []LibSwapNormalizedSwapData
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SoData = *abi.ConvertType(out[0], new(ISoNormalizedSoData)).(*ISoNormalizedSoData)
	outstruct.SwapDataDst = *abi.ConvertType(out[1], new([]LibSwapNormalizedSwapData)).(*[]LibSwapNormalizedSwapData)

	return *outstruct, err

}

// DecodeStargatePayload is a free data retrieval call binding the contract method 0x762a9962.
//
// Solidity: function decodeStargatePayload(bytes stargatePayload) pure returns((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst)
func (_SoDiamond *SoDiamondSession) DecodeStargatePayload(stargatePayload []byte) (struct {
	SoData      ISoNormalizedSoData
	SwapDataDst []LibSwapNormalizedSwapData
}, error) {
	return _SoDiamond.Contract.DecodeStargatePayload(&_SoDiamond.CallOpts, stargatePayload)
}

// DecodeStargatePayload is a free data retrieval call binding the contract method 0x762a9962.
//
// Solidity: function decodeStargatePayload(bytes stargatePayload) pure returns((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst)
func (_SoDiamond *SoDiamondCallerSession) DecodeStargatePayload(stargatePayload []byte) (struct {
	SoData      ISoNormalizedSoData
	SwapDataDst []LibSwapNormalizedSwapData
}, error) {
	return _SoDiamond.Contract.DecodeStargatePayload(&_SoDiamond.CallOpts, stargatePayload)
}

// DecodeWormholePayload is a free data retrieval call binding the contract method 0xb91dfdd6.
//
// Solidity: function decodeWormholePayload(bytes wormholeData) pure returns(uint256, uint256, (bytes,bytes,uint16,bytes,uint16,bytes,uint256), (bytes,bytes,bytes,bytes,uint256,bytes)[])
func (_SoDiamond *SoDiamondCaller) DecodeWormholePayload(opts *bind.CallOpts, wormholeData []byte) (*big.Int, *big.Int, ISoNormalizedSoData, []LibSwapNormalizedSwapData, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "decodeWormholePayload", wormholeData)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(ISoNormalizedSoData), *new([]LibSwapNormalizedSwapData), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(ISoNormalizedSoData)).(*ISoNormalizedSoData)
	out3 := *abi.ConvertType(out[3], new([]LibSwapNormalizedSwapData)).(*[]LibSwapNormalizedSwapData)

	return out0, out1, out2, out3, err

}

// DecodeWormholePayload is a free data retrieval call binding the contract method 0xb91dfdd6.
//
// Solidity: function decodeWormholePayload(bytes wormholeData) pure returns(uint256, uint256, (bytes,bytes,uint16,bytes,uint16,bytes,uint256), (bytes,bytes,bytes,bytes,uint256,bytes)[])
func (_SoDiamond *SoDiamondSession) DecodeWormholePayload(wormholeData []byte) (*big.Int, *big.Int, ISoNormalizedSoData, []LibSwapNormalizedSwapData, error) {
	return _SoDiamond.Contract.DecodeWormholePayload(&_SoDiamond.CallOpts, wormholeData)
}

// DecodeWormholePayload is a free data retrieval call binding the contract method 0xb91dfdd6.
//
// Solidity: function decodeWormholePayload(bytes wormholeData) pure returns(uint256, uint256, (bytes,bytes,uint16,bytes,uint16,bytes,uint256), (bytes,bytes,bytes,bytes,uint256,bytes)[])
func (_SoDiamond *SoDiamondCallerSession) DecodeWormholePayload(wormholeData []byte) (*big.Int, *big.Int, ISoNormalizedSoData, []LibSwapNormalizedSwapData, error) {
	return _SoDiamond.Contract.DecodeWormholePayload(&_SoDiamond.CallOpts, wormholeData)
}

// DenormalizeSoData is a free data retrieval call binding the contract method 0x896a8b50.
//
// Solidity: function denormalizeSoData((bytes,bytes,uint16,bytes,uint16,bytes,uint256) data) pure returns((bytes32,address,uint16,address,uint16,address,uint256))
func (_SoDiamond *SoDiamondCaller) DenormalizeSoData(opts *bind.CallOpts, data ISoNormalizedSoData) (ISoSoData, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "denormalizeSoData", data)

	if err != nil {
		return *new(ISoSoData), err
	}

	out0 := *abi.ConvertType(out[0], new(ISoSoData)).(*ISoSoData)

	return out0, err

}

// DenormalizeSoData is a free data retrieval call binding the contract method 0x896a8b50.
//
// Solidity: function denormalizeSoData((bytes,bytes,uint16,bytes,uint16,bytes,uint256) data) pure returns((bytes32,address,uint16,address,uint16,address,uint256))
func (_SoDiamond *SoDiamondSession) DenormalizeSoData(data ISoNormalizedSoData) (ISoSoData, error) {
	return _SoDiamond.Contract.DenormalizeSoData(&_SoDiamond.CallOpts, data)
}

// DenormalizeSoData is a free data retrieval call binding the contract method 0x896a8b50.
//
// Solidity: function denormalizeSoData((bytes,bytes,uint16,bytes,uint16,bytes,uint256) data) pure returns((bytes32,address,uint16,address,uint16,address,uint256))
func (_SoDiamond *SoDiamondCallerSession) DenormalizeSoData(data ISoNormalizedSoData) (ISoSoData, error) {
	return _SoDiamond.Contract.DenormalizeSoData(&_SoDiamond.CallOpts, data)
}

// DenormalizeSwapData is a free data retrieval call binding the contract method 0x1154d06d.
//
// Solidity: function denormalizeSwapData((bytes,bytes,bytes,bytes,uint256,bytes)[] data) pure returns((address,address,address,address,uint256,bytes)[])
func (_SoDiamond *SoDiamondCaller) DenormalizeSwapData(opts *bind.CallOpts, data []LibSwapNormalizedSwapData) ([]LibSwapSwapData, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "denormalizeSwapData", data)

	if err != nil {
		return *new([]LibSwapSwapData), err
	}

	out0 := *abi.ConvertType(out[0], new([]LibSwapSwapData)).(*[]LibSwapSwapData)

	return out0, err

}

// DenormalizeSwapData is a free data retrieval call binding the contract method 0x1154d06d.
//
// Solidity: function denormalizeSwapData((bytes,bytes,bytes,bytes,uint256,bytes)[] data) pure returns((address,address,address,address,uint256,bytes)[])
func (_SoDiamond *SoDiamondSession) DenormalizeSwapData(data []LibSwapNormalizedSwapData) ([]LibSwapSwapData, error) {
	return _SoDiamond.Contract.DenormalizeSwapData(&_SoDiamond.CallOpts, data)
}

// DenormalizeSwapData is a free data retrieval call binding the contract method 0x1154d06d.
//
// Solidity: function denormalizeSwapData((bytes,bytes,bytes,bytes,uint256,bytes)[] data) pure returns((address,address,address,address,uint256,bytes)[])
func (_SoDiamond *SoDiamondCallerSession) DenormalizeSwapData(data []LibSwapNormalizedSwapData) ([]LibSwapSwapData, error) {
	return _SoDiamond.Contract.DenormalizeSwapData(&_SoDiamond.CallOpts, data)
}

// DenormalizeU256 is a free data retrieval call binding the contract method 0xfbc128a3.
//
// Solidity: function denormalizeU256(bytes data) pure returns(uint256)
func (_SoDiamond *SoDiamondCaller) DenormalizeU256(opts *bind.CallOpts, data []byte) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "denormalizeU256", data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DenormalizeU256 is a free data retrieval call binding the contract method 0xfbc128a3.
//
// Solidity: function denormalizeU256(bytes data) pure returns(uint256)
func (_SoDiamond *SoDiamondSession) DenormalizeU256(data []byte) (*big.Int, error) {
	return _SoDiamond.Contract.DenormalizeU256(&_SoDiamond.CallOpts, data)
}

// DenormalizeU256 is a free data retrieval call binding the contract method 0xfbc128a3.
//
// Solidity: function denormalizeU256(bytes data) pure returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) DenormalizeU256(data []byte) (*big.Int, error) {
	return _SoDiamond.Contract.DenormalizeU256(&_SoDiamond.CallOpts, data)
}

// EncodeBoolSwapPayload is a free data retrieval call binding the contract method 0x085f9441.
//
// Solidity: function encodeBoolSwapPayload((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst) pure returns(bytes)
func (_SoDiamond *SoDiamondCaller) EncodeBoolSwapPayload(opts *bind.CallOpts, soData ISoNormalizedSoData, swapDataDst []LibSwapNormalizedSwapData) ([]byte, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "encodeBoolSwapPayload", soData, swapDataDst)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeBoolSwapPayload is a free data retrieval call binding the contract method 0x085f9441.
//
// Solidity: function encodeBoolSwapPayload((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst) pure returns(bytes)
func (_SoDiamond *SoDiamondSession) EncodeBoolSwapPayload(soData ISoNormalizedSoData, swapDataDst []LibSwapNormalizedSwapData) ([]byte, error) {
	return _SoDiamond.Contract.EncodeBoolSwapPayload(&_SoDiamond.CallOpts, soData, swapDataDst)
}

// EncodeBoolSwapPayload is a free data retrieval call binding the contract method 0x085f9441.
//
// Solidity: function encodeBoolSwapPayload((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst) pure returns(bytes)
func (_SoDiamond *SoDiamondCallerSession) EncodeBoolSwapPayload(soData ISoNormalizedSoData, swapDataDst []LibSwapNormalizedSwapData) ([]byte, error) {
	return _SoDiamond.Contract.EncodeBoolSwapPayload(&_SoDiamond.CallOpts, soData, swapDataDst)
}

// EncodeCelerPayload is a free data retrieval call binding the contract method 0x61bc39e6.
//
// Solidity: function encodeCelerPayload(address sender, (bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) pure returns(bytes)
func (_SoDiamond *SoDiamondCaller) EncodeCelerPayload(opts *bind.CallOpts, sender common.Address, soDataNo ISoNormalizedSoData, swapDataDstNo []LibSwapNormalizedSwapData) ([]byte, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "encodeCelerPayload", sender, soDataNo, swapDataDstNo)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeCelerPayload is a free data retrieval call binding the contract method 0x61bc39e6.
//
// Solidity: function encodeCelerPayload(address sender, (bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) pure returns(bytes)
func (_SoDiamond *SoDiamondSession) EncodeCelerPayload(sender common.Address, soDataNo ISoNormalizedSoData, swapDataDstNo []LibSwapNormalizedSwapData) ([]byte, error) {
	return _SoDiamond.Contract.EncodeCelerPayload(&_SoDiamond.CallOpts, sender, soDataNo, swapDataDstNo)
}

// EncodeCelerPayload is a free data retrieval call binding the contract method 0x61bc39e6.
//
// Solidity: function encodeCelerPayload(address sender, (bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) pure returns(bytes)
func (_SoDiamond *SoDiamondCallerSession) EncodeCelerPayload(sender common.Address, soDataNo ISoNormalizedSoData, swapDataDstNo []LibSwapNormalizedSwapData) ([]byte, error) {
	return _SoDiamond.Contract.EncodeCelerPayload(&_SoDiamond.CallOpts, sender, soDataNo, swapDataDstNo)
}

// EncodeMultiChainPayload is a free data retrieval call binding the contract method 0x88afbc4a.
//
// Solidity: function encodeMultiChainPayload((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) pure returns(bytes)
func (_SoDiamond *SoDiamondCaller) EncodeMultiChainPayload(opts *bind.CallOpts, soDataNo ISoNormalizedSoData, swapDataDstNo []LibSwapNormalizedSwapData) ([]byte, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "encodeMultiChainPayload", soDataNo, swapDataDstNo)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeMultiChainPayload is a free data retrieval call binding the contract method 0x88afbc4a.
//
// Solidity: function encodeMultiChainPayload((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) pure returns(bytes)
func (_SoDiamond *SoDiamondSession) EncodeMultiChainPayload(soDataNo ISoNormalizedSoData, swapDataDstNo []LibSwapNormalizedSwapData) ([]byte, error) {
	return _SoDiamond.Contract.EncodeMultiChainPayload(&_SoDiamond.CallOpts, soDataNo, swapDataDstNo)
}

// EncodeMultiChainPayload is a free data retrieval call binding the contract method 0x88afbc4a.
//
// Solidity: function encodeMultiChainPayload((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) pure returns(bytes)
func (_SoDiamond *SoDiamondCallerSession) EncodeMultiChainPayload(soDataNo ISoNormalizedSoData, swapDataDstNo []LibSwapNormalizedSwapData) ([]byte, error) {
	return _SoDiamond.Contract.EncodeMultiChainPayload(&_SoDiamond.CallOpts, soDataNo, swapDataDstNo)
}

// EncodeNormalizedSoData is a free data retrieval call binding the contract method 0x9943d550.
//
// Solidity: function encodeNormalizedSoData((bytes,bytes,uint16,bytes,uint16,bytes,uint256) data) pure returns(bytes)
func (_SoDiamond *SoDiamondCaller) EncodeNormalizedSoData(opts *bind.CallOpts, data ISoNormalizedSoData) ([]byte, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "encodeNormalizedSoData", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeNormalizedSoData is a free data retrieval call binding the contract method 0x9943d550.
//
// Solidity: function encodeNormalizedSoData((bytes,bytes,uint16,bytes,uint16,bytes,uint256) data) pure returns(bytes)
func (_SoDiamond *SoDiamondSession) EncodeNormalizedSoData(data ISoNormalizedSoData) ([]byte, error) {
	return _SoDiamond.Contract.EncodeNormalizedSoData(&_SoDiamond.CallOpts, data)
}

// EncodeNormalizedSoData is a free data retrieval call binding the contract method 0x9943d550.
//
// Solidity: function encodeNormalizedSoData((bytes,bytes,uint16,bytes,uint16,bytes,uint256) data) pure returns(bytes)
func (_SoDiamond *SoDiamondCallerSession) EncodeNormalizedSoData(data ISoNormalizedSoData) ([]byte, error) {
	return _SoDiamond.Contract.EncodeNormalizedSoData(&_SoDiamond.CallOpts, data)
}

// EncodeNormalizedSwapData is a free data retrieval call binding the contract method 0x05cabca4.
//
// Solidity: function encodeNormalizedSwapData((bytes,bytes,bytes,bytes,uint256,bytes)[] data) pure returns(bytes)
func (_SoDiamond *SoDiamondCaller) EncodeNormalizedSwapData(opts *bind.CallOpts, data []LibSwapNormalizedSwapData) ([]byte, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "encodeNormalizedSwapData", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeNormalizedSwapData is a free data retrieval call binding the contract method 0x05cabca4.
//
// Solidity: function encodeNormalizedSwapData((bytes,bytes,bytes,bytes,uint256,bytes)[] data) pure returns(bytes)
func (_SoDiamond *SoDiamondSession) EncodeNormalizedSwapData(data []LibSwapNormalizedSwapData) ([]byte, error) {
	return _SoDiamond.Contract.EncodeNormalizedSwapData(&_SoDiamond.CallOpts, data)
}

// EncodeNormalizedSwapData is a free data retrieval call binding the contract method 0x05cabca4.
//
// Solidity: function encodeNormalizedSwapData((bytes,bytes,bytes,bytes,uint256,bytes)[] data) pure returns(bytes)
func (_SoDiamond *SoDiamondCallerSession) EncodeNormalizedSwapData(data []LibSwapNormalizedSwapData) ([]byte, error) {
	return _SoDiamond.Contract.EncodeNormalizedSwapData(&_SoDiamond.CallOpts, data)
}

// EncodeNormalizedWormholeData is a free data retrieval call binding the contract method 0x740dc2b6.
//
// Solidity: function encodeNormalizedWormholeData((uint16,uint256,uint256,bytes) data) pure returns(bytes)
func (_SoDiamond *SoDiamondCaller) EncodeNormalizedWormholeData(opts *bind.CallOpts, data WormholeFacetNormalizedWormholeData) ([]byte, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "encodeNormalizedWormholeData", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeNormalizedWormholeData is a free data retrieval call binding the contract method 0x740dc2b6.
//
// Solidity: function encodeNormalizedWormholeData((uint16,uint256,uint256,bytes) data) pure returns(bytes)
func (_SoDiamond *SoDiamondSession) EncodeNormalizedWormholeData(data WormholeFacetNormalizedWormholeData) ([]byte, error) {
	return _SoDiamond.Contract.EncodeNormalizedWormholeData(&_SoDiamond.CallOpts, data)
}

// EncodeNormalizedWormholeData is a free data retrieval call binding the contract method 0x740dc2b6.
//
// Solidity: function encodeNormalizedWormholeData((uint16,uint256,uint256,bytes) data) pure returns(bytes)
func (_SoDiamond *SoDiamondCallerSession) EncodeNormalizedWormholeData(data WormholeFacetNormalizedWormholeData) ([]byte, error) {
	return _SoDiamond.Contract.EncodeNormalizedWormholeData(&_SoDiamond.CallOpts, data)
}

// EncodeStargatePayload is a free data retrieval call binding the contract method 0xdf3c5ac3.
//
// Solidity: function encodeStargatePayload((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst) pure returns(bytes)
func (_SoDiamond *SoDiamondCaller) EncodeStargatePayload(opts *bind.CallOpts, soData ISoNormalizedSoData, swapDataDst []LibSwapNormalizedSwapData) ([]byte, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "encodeStargatePayload", soData, swapDataDst)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeStargatePayload is a free data retrieval call binding the contract method 0xdf3c5ac3.
//
// Solidity: function encodeStargatePayload((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst) pure returns(bytes)
func (_SoDiamond *SoDiamondSession) EncodeStargatePayload(soData ISoNormalizedSoData, swapDataDst []LibSwapNormalizedSwapData) ([]byte, error) {
	return _SoDiamond.Contract.EncodeStargatePayload(&_SoDiamond.CallOpts, soData, swapDataDst)
}

// EncodeStargatePayload is a free data retrieval call binding the contract method 0xdf3c5ac3.
//
// Solidity: function encodeStargatePayload((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst) pure returns(bytes)
func (_SoDiamond *SoDiamondCallerSession) EncodeStargatePayload(soData ISoNormalizedSoData, swapDataDst []LibSwapNormalizedSwapData) ([]byte, error) {
	return _SoDiamond.Contract.EncodeStargatePayload(&_SoDiamond.CallOpts, soData, swapDataDst)
}

// EncodeWormholePayload is a free data retrieval call binding the contract method 0xd8b886f3.
//
// Solidity: function encodeWormholePayload(uint256 dstMaxGasPrice, uint256 dstMaxGas, (bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst) pure returns(bytes)
func (_SoDiamond *SoDiamondCaller) EncodeWormholePayload(opts *bind.CallOpts, dstMaxGasPrice *big.Int, dstMaxGas *big.Int, soData ISoNormalizedSoData, swapDataDst []LibSwapNormalizedSwapData) ([]byte, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "encodeWormholePayload", dstMaxGasPrice, dstMaxGas, soData, swapDataDst)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeWormholePayload is a free data retrieval call binding the contract method 0xd8b886f3.
//
// Solidity: function encodeWormholePayload(uint256 dstMaxGasPrice, uint256 dstMaxGas, (bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst) pure returns(bytes)
func (_SoDiamond *SoDiamondSession) EncodeWormholePayload(dstMaxGasPrice *big.Int, dstMaxGas *big.Int, soData ISoNormalizedSoData, swapDataDst []LibSwapNormalizedSwapData) ([]byte, error) {
	return _SoDiamond.Contract.EncodeWormholePayload(&_SoDiamond.CallOpts, dstMaxGasPrice, dstMaxGas, soData, swapDataDst)
}

// EncodeWormholePayload is a free data retrieval call binding the contract method 0xd8b886f3.
//
// Solidity: function encodeWormholePayload(uint256 dstMaxGasPrice, uint256 dstMaxGas, (bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst) pure returns(bytes)
func (_SoDiamond *SoDiamondCallerSession) EncodeWormholePayload(dstMaxGasPrice *big.Int, dstMaxGas *big.Int, soData ISoNormalizedSoData, swapDataDst []LibSwapNormalizedSwapData) ([]byte, error) {
	return _SoDiamond.Contract.EncodeWormholePayload(&_SoDiamond.CallOpts, dstMaxGasPrice, dstMaxGas, soData, swapDataDst)
}

// EstCelerMessageFeeAndExecutorFee is a free data retrieval call binding the contract method 0x9f84807a.
//
// Solidity: function estCelerMessageFeeAndExecutorFee(uint64 dstCelerChainId, uint256 dstMaxGasPriceInWeiForExecutor, (bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo, bool is_actual) view returns(uint256, uint256, uint256)
func (_SoDiamond *SoDiamondCaller) EstCelerMessageFeeAndExecutorFee(opts *bind.CallOpts, dstCelerChainId uint64, dstMaxGasPriceInWeiForExecutor *big.Int, soDataNo ISoNormalizedSoData, swapDataDstNo []LibSwapNormalizedSwapData, is_actual bool) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "estCelerMessageFeeAndExecutorFee", dstCelerChainId, dstMaxGasPriceInWeiForExecutor, soDataNo, swapDataDstNo, is_actual)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// EstCelerMessageFeeAndExecutorFee is a free data retrieval call binding the contract method 0x9f84807a.
//
// Solidity: function estCelerMessageFeeAndExecutorFee(uint64 dstCelerChainId, uint256 dstMaxGasPriceInWeiForExecutor, (bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo, bool is_actual) view returns(uint256, uint256, uint256)
func (_SoDiamond *SoDiamondSession) EstCelerMessageFeeAndExecutorFee(dstCelerChainId uint64, dstMaxGasPriceInWeiForExecutor *big.Int, soDataNo ISoNormalizedSoData, swapDataDstNo []LibSwapNormalizedSwapData, is_actual bool) (*big.Int, *big.Int, *big.Int, error) {
	return _SoDiamond.Contract.EstCelerMessageFeeAndExecutorFee(&_SoDiamond.CallOpts, dstCelerChainId, dstMaxGasPriceInWeiForExecutor, soDataNo, swapDataDstNo, is_actual)
}

// EstCelerMessageFeeAndExecutorFee is a free data retrieval call binding the contract method 0x9f84807a.
//
// Solidity: function estCelerMessageFeeAndExecutorFee(uint64 dstCelerChainId, uint256 dstMaxGasPriceInWeiForExecutor, (bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo, bool is_actual) view returns(uint256, uint256, uint256)
func (_SoDiamond *SoDiamondCallerSession) EstCelerMessageFeeAndExecutorFee(dstCelerChainId uint64, dstMaxGasPriceInWeiForExecutor *big.Int, soDataNo ISoNormalizedSoData, swapDataDstNo []LibSwapNormalizedSwapData, is_actual bool) (*big.Int, *big.Int, *big.Int, error) {
	return _SoDiamond.Contract.EstCelerMessageFeeAndExecutorFee(&_SoDiamond.CallOpts, dstCelerChainId, dstMaxGasPriceInWeiForExecutor, soDataNo, swapDataDstNo, is_actual)
}

// EstimateBoolFee is a free data retrieval call binding the contract method 0x097ea22c.
//
// Solidity: function estimateBoolFee((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (uint32,uint32,bytes32) boolSwapData, uint256 bridgeAmount, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) view returns(uint256)
func (_SoDiamond *SoDiamondCaller) EstimateBoolFee(opts *bind.CallOpts, soDataNo ISoNormalizedSoData, boolSwapData BoolFacetBoolSwapData, bridgeAmount *big.Int, swapDataDstNo []LibSwapNormalizedSwapData) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "estimateBoolFee", soDataNo, boolSwapData, bridgeAmount, swapDataDstNo)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateBoolFee is a free data retrieval call binding the contract method 0x097ea22c.
//
// Solidity: function estimateBoolFee((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (uint32,uint32,bytes32) boolSwapData, uint256 bridgeAmount, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) view returns(uint256)
func (_SoDiamond *SoDiamondSession) EstimateBoolFee(soDataNo ISoNormalizedSoData, boolSwapData BoolFacetBoolSwapData, bridgeAmount *big.Int, swapDataDstNo []LibSwapNormalizedSwapData) (*big.Int, error) {
	return _SoDiamond.Contract.EstimateBoolFee(&_SoDiamond.CallOpts, soDataNo, boolSwapData, bridgeAmount, swapDataDstNo)
}

// EstimateBoolFee is a free data retrieval call binding the contract method 0x097ea22c.
//
// Solidity: function estimateBoolFee((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (uint32,uint32,bytes32) boolSwapData, uint256 bridgeAmount, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) EstimateBoolFee(soDataNo ISoNormalizedSoData, boolSwapData BoolFacetBoolSwapData, bridgeAmount *big.Int, swapDataDstNo []LibSwapNormalizedSwapData) (*big.Int, error) {
	return _SoDiamond.Contract.EstimateBoolFee(&_SoDiamond.CallOpts, soDataNo, boolSwapData, bridgeAmount, swapDataDstNo)
}

// EstimateCompleteSoSwapGas is a free data retrieval call binding the contract method 0x272838b8.
//
// Solidity: function estimateCompleteSoSwapGas((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (uint16,uint256,uint256,bytes) wormholeData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst) view returns(uint256)
func (_SoDiamond *SoDiamondCaller) EstimateCompleteSoSwapGas(opts *bind.CallOpts, soData ISoNormalizedSoData, wormholeData WormholeFacetNormalizedWormholeData, swapDataDst []LibSwapNormalizedSwapData) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "estimateCompleteSoSwapGas", soData, wormholeData, swapDataDst)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateCompleteSoSwapGas is a free data retrieval call binding the contract method 0x272838b8.
//
// Solidity: function estimateCompleteSoSwapGas((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (uint16,uint256,uint256,bytes) wormholeData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst) view returns(uint256)
func (_SoDiamond *SoDiamondSession) EstimateCompleteSoSwapGas(soData ISoNormalizedSoData, wormholeData WormholeFacetNormalizedWormholeData, swapDataDst []LibSwapNormalizedSwapData) (*big.Int, error) {
	return _SoDiamond.Contract.EstimateCompleteSoSwapGas(&_SoDiamond.CallOpts, soData, wormholeData, swapDataDst)
}

// EstimateCompleteSoSwapGas is a free data retrieval call binding the contract method 0x272838b8.
//
// Solidity: function estimateCompleteSoSwapGas((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (uint16,uint256,uint256,bytes) wormholeData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst) view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) EstimateCompleteSoSwapGas(soData ISoNormalizedSoData, wormholeData WormholeFacetNormalizedWormholeData, swapDataDst []LibSwapNormalizedSwapData) (*big.Int, error) {
	return _SoDiamond.Contract.EstimateCompleteSoSwapGas(&_SoDiamond.CallOpts, soData, wormholeData, swapDataDst)
}

// EstimateRelayerFee is a free data retrieval call binding the contract method 0x23e7af5d.
//
// Solidity: function estimateRelayerFee((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (uint16,uint256,uint256,bytes) wormholeData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst) view returns(uint256)
func (_SoDiamond *SoDiamondCaller) EstimateRelayerFee(opts *bind.CallOpts, soData ISoNormalizedSoData, wormholeData WormholeFacetNormalizedWormholeData, swapDataDst []LibSwapNormalizedSwapData) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "estimateRelayerFee", soData, wormholeData, swapDataDst)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateRelayerFee is a free data retrieval call binding the contract method 0x23e7af5d.
//
// Solidity: function estimateRelayerFee((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (uint16,uint256,uint256,bytes) wormholeData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst) view returns(uint256)
func (_SoDiamond *SoDiamondSession) EstimateRelayerFee(soData ISoNormalizedSoData, wormholeData WormholeFacetNormalizedWormholeData, swapDataDst []LibSwapNormalizedSwapData) (*big.Int, error) {
	return _SoDiamond.Contract.EstimateRelayerFee(&_SoDiamond.CallOpts, soData, wormholeData, swapDataDst)
}

// EstimateRelayerFee is a free data retrieval call binding the contract method 0x23e7af5d.
//
// Solidity: function estimateRelayerFee((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (uint16,uint256,uint256,bytes) wormholeData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst) view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) EstimateRelayerFee(soData ISoNormalizedSoData, wormholeData WormholeFacetNormalizedWormholeData, swapDataDst []LibSwapNormalizedSwapData) (*big.Int, error) {
	return _SoDiamond.Contract.EstimateRelayerFee(&_SoDiamond.CallOpts, soData, wormholeData, swapDataDst)
}

// EstimateStargateFinalAmount is a free data retrieval call binding the contract method 0xcd96eb9a.
//
// Solidity: function estimateStargateFinalAmount((uint256,uint16,uint256,uint256,uint256,address) stargateData, uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondCaller) EstimateStargateFinalAmount(opts *bind.CallOpts, stargateData StargateFacetStargateData, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "estimateStargateFinalAmount", stargateData, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateStargateFinalAmount is a free data retrieval call binding the contract method 0xcd96eb9a.
//
// Solidity: function estimateStargateFinalAmount((uint256,uint16,uint256,uint256,uint256,address) stargateData, uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondSession) EstimateStargateFinalAmount(stargateData StargateFacetStargateData, amount *big.Int) (*big.Int, error) {
	return _SoDiamond.Contract.EstimateStargateFinalAmount(&_SoDiamond.CallOpts, stargateData, amount)
}

// EstimateStargateFinalAmount is a free data retrieval call binding the contract method 0xcd96eb9a.
//
// Solidity: function estimateStargateFinalAmount((uint256,uint16,uint256,uint256,uint256,address) stargateData, uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) EstimateStargateFinalAmount(stargateData StargateFacetStargateData, amount *big.Int) (*big.Int, error) {
	return _SoDiamond.Contract.EstimateStargateFinalAmount(&_SoDiamond.CallOpts, stargateData, amount)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 functionSelector) view returns(address addresses)
func (_SoDiamond *SoDiamondCaller) FacetAddress(opts *bind.CallOpts, functionSelector [4]byte) (common.Address, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "facetAddress", functionSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 functionSelector) view returns(address addresses)
func (_SoDiamond *SoDiamondSession) FacetAddress(functionSelector [4]byte) (common.Address, error) {
	return _SoDiamond.Contract.FacetAddress(&_SoDiamond.CallOpts, functionSelector)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 functionSelector) view returns(address addresses)
func (_SoDiamond *SoDiamondCallerSession) FacetAddress(functionSelector [4]byte) (common.Address, error) {
	return _SoDiamond.Contract.FacetAddress(&_SoDiamond.CallOpts, functionSelector)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] addresses)
func (_SoDiamond *SoDiamondCaller) FacetAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "facetAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] addresses)
func (_SoDiamond *SoDiamondSession) FacetAddresses() ([]common.Address, error) {
	return _SoDiamond.Contract.FacetAddresses(&_SoDiamond.CallOpts)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] addresses)
func (_SoDiamond *SoDiamondCallerSession) FacetAddresses() ([]common.Address, error) {
	return _SoDiamond.Contract.FacetAddresses(&_SoDiamond.CallOpts)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address facet) view returns(bytes4[] facetFunctionSelectors_)
func (_SoDiamond *SoDiamondCaller) FacetFunctionSelectors(opts *bind.CallOpts, facet common.Address) ([][4]byte, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "facetFunctionSelectors", facet)

	if err != nil {
		return *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)

	return out0, err

}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address facet) view returns(bytes4[] facetFunctionSelectors_)
func (_SoDiamond *SoDiamondSession) FacetFunctionSelectors(facet common.Address) ([][4]byte, error) {
	return _SoDiamond.Contract.FacetFunctionSelectors(&_SoDiamond.CallOpts, facet)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address facet) view returns(bytes4[] facetFunctionSelectors_)
func (_SoDiamond *SoDiamondCallerSession) FacetFunctionSelectors(facet common.Address) ([][4]byte, error) {
	return _SoDiamond.Contract.FacetFunctionSelectors(&_SoDiamond.CallOpts, facet)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] allFacets)
func (_SoDiamond *SoDiamondCaller) Facets(opts *bind.CallOpts) ([]IDiamondLoupeFacet, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "facets")

	if err != nil {
		return *new([]IDiamondLoupeFacet), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDiamondLoupeFacet)).(*[]IDiamondLoupeFacet)

	return out0, err

}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] allFacets)
func (_SoDiamond *SoDiamondSession) Facets() ([]IDiamondLoupeFacet, error) {
	return _SoDiamond.Contract.Facets(&_SoDiamond.CallOpts)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] allFacets)
func (_SoDiamond *SoDiamondCallerSession) Facets() ([]IDiamondLoupeFacet, error) {
	return _SoDiamond.Contract.Facets(&_SoDiamond.CallOpts)
}

// FromBoolToConsumer is a free data retrieval call binding the contract method 0xa8750b9e.
//
// Solidity: function fromBoolToConsumer(uint32 boolChainId) view returns(uint32 consumerChainId)
func (_SoDiamond *SoDiamondCaller) FromBoolToConsumer(opts *bind.CallOpts, boolChainId uint32) (uint32, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "fromBoolToConsumer", boolChainId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FromBoolToConsumer is a free data retrieval call binding the contract method 0xa8750b9e.
//
// Solidity: function fromBoolToConsumer(uint32 boolChainId) view returns(uint32 consumerChainId)
func (_SoDiamond *SoDiamondSession) FromBoolToConsumer(boolChainId uint32) (uint32, error) {
	return _SoDiamond.Contract.FromBoolToConsumer(&_SoDiamond.CallOpts, boolChainId)
}

// FromBoolToConsumer is a free data retrieval call binding the contract method 0xa8750b9e.
//
// Solidity: function fromBoolToConsumer(uint32 boolChainId) view returns(uint32 consumerChainId)
func (_SoDiamond *SoDiamondCallerSession) FromBoolToConsumer(boolChainId uint32) (uint32, error) {
	return _SoDiamond.Contract.FromBoolToConsumer(&_SoDiamond.CallOpts, boolChainId)
}

// FromConsumerToBool is a free data retrieval call binding the contract method 0x72112c61.
//
// Solidity: function fromConsumerToBool(uint32 consumerChainId) view returns(uint32 boolChainId)
func (_SoDiamond *SoDiamondCaller) FromConsumerToBool(opts *bind.CallOpts, consumerChainId uint32) (uint32, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "fromConsumerToBool", consumerChainId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FromConsumerToBool is a free data retrieval call binding the contract method 0x72112c61.
//
// Solidity: function fromConsumerToBool(uint32 consumerChainId) view returns(uint32 boolChainId)
func (_SoDiamond *SoDiamondSession) FromConsumerToBool(consumerChainId uint32) (uint32, error) {
	return _SoDiamond.Contract.FromConsumerToBool(&_SoDiamond.CallOpts, consumerChainId)
}

// FromConsumerToBool is a free data retrieval call binding the contract method 0x72112c61.
//
// Solidity: function fromConsumerToBool(uint32 consumerChainId) view returns(uint32 boolChainId)
func (_SoDiamond *SoDiamondCallerSession) FromConsumerToBool(consumerChainId uint32) (uint32, error) {
	return _SoDiamond.Contract.FromConsumerToBool(&_SoDiamond.CallOpts, consumerChainId)
}

// GetAmountBeforeSoFee is a free data retrieval call binding the contract method 0x7e3c358b.
//
// Solidity: function getAmountBeforeSoFee(uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondCaller) GetAmountBeforeSoFee(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getAmountBeforeSoFee", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountBeforeSoFee is a free data retrieval call binding the contract method 0x7e3c358b.
//
// Solidity: function getAmountBeforeSoFee(uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondSession) GetAmountBeforeSoFee(amount *big.Int) (*big.Int, error) {
	return _SoDiamond.Contract.GetAmountBeforeSoFee(&_SoDiamond.CallOpts, amount)
}

// GetAmountBeforeSoFee is a free data retrieval call binding the contract method 0x7e3c358b.
//
// Solidity: function getAmountBeforeSoFee(uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) GetAmountBeforeSoFee(amount *big.Int) (*big.Int, error) {
	return _SoDiamond.Contract.GetAmountBeforeSoFee(&_SoDiamond.CallOpts, amount)
}

// GetAnyToken is a free data retrieval call binding the contract method 0x2342cb97.
//
// Solidity: function getAnyToken(address bridgeToken) view returns(address)
func (_SoDiamond *SoDiamondCaller) GetAnyToken(opts *bind.CallOpts, bridgeToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getAnyToken", bridgeToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAnyToken is a free data retrieval call binding the contract method 0x2342cb97.
//
// Solidity: function getAnyToken(address bridgeToken) view returns(address)
func (_SoDiamond *SoDiamondSession) GetAnyToken(bridgeToken common.Address) (common.Address, error) {
	return _SoDiamond.Contract.GetAnyToken(&_SoDiamond.CallOpts, bridgeToken)
}

// GetAnyToken is a free data retrieval call binding the contract method 0x2342cb97.
//
// Solidity: function getAnyToken(address bridgeToken) view returns(address)
func (_SoDiamond *SoDiamondCallerSession) GetAnyToken(bridgeToken common.Address) (common.Address, error) {
	return _SoDiamond.Contract.GetAnyToken(&_SoDiamond.CallOpts, bridgeToken)
}

// GetBaseGas is a free data retrieval call binding the contract method 0xd111cf0f.
//
// Solidity: function getBaseGas(uint64 dstChainId) view returns(uint256)
func (_SoDiamond *SoDiamondCaller) GetBaseGas(opts *bind.CallOpts, dstChainId uint64) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getBaseGas", dstChainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBaseGas is a free data retrieval call binding the contract method 0xd111cf0f.
//
// Solidity: function getBaseGas(uint64 dstChainId) view returns(uint256)
func (_SoDiamond *SoDiamondSession) GetBaseGas(dstChainId uint64) (*big.Int, error) {
	return _SoDiamond.Contract.GetBaseGas(&_SoDiamond.CallOpts, dstChainId)
}

// GetBaseGas is a free data retrieval call binding the contract method 0xd111cf0f.
//
// Solidity: function getBaseGas(uint64 dstChainId) view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) GetBaseGas(dstChainId uint64) (*big.Int, error) {
	return _SoDiamond.Contract.GetBaseGas(&_SoDiamond.CallOpts, dstChainId)
}

// GetBoolBasicBeneficiary is a free data retrieval call binding the contract method 0xda263328.
//
// Solidity: function getBoolBasicBeneficiary() view returns(address)
func (_SoDiamond *SoDiamondCaller) GetBoolBasicBeneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getBoolBasicBeneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBoolBasicBeneficiary is a free data retrieval call binding the contract method 0xda263328.
//
// Solidity: function getBoolBasicBeneficiary() view returns(address)
func (_SoDiamond *SoDiamondSession) GetBoolBasicBeneficiary() (common.Address, error) {
	return _SoDiamond.Contract.GetBoolBasicBeneficiary(&_SoDiamond.CallOpts)
}

// GetBoolBasicBeneficiary is a free data retrieval call binding the contract method 0xda263328.
//
// Solidity: function getBoolBasicBeneficiary() view returns(address)
func (_SoDiamond *SoDiamondCallerSession) GetBoolBasicBeneficiary() (common.Address, error) {
	return _SoDiamond.Contract.GetBoolBasicBeneficiary(&_SoDiamond.CallOpts)
}

// GetBoolBasicFee is a free data retrieval call binding the contract method 0x7ed5d087.
//
// Solidity: function getBoolBasicFee() view returns(uint256)
func (_SoDiamond *SoDiamondCaller) GetBoolBasicFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getBoolBasicFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBoolBasicFee is a free data retrieval call binding the contract method 0x7ed5d087.
//
// Solidity: function getBoolBasicFee() view returns(uint256)
func (_SoDiamond *SoDiamondSession) GetBoolBasicFee() (*big.Int, error) {
	return _SoDiamond.Contract.GetBoolBasicFee(&_SoDiamond.CallOpts)
}

// GetBoolBasicFee is a free data retrieval call binding the contract method 0x7ed5d087.
//
// Solidity: function getBoolBasicFee() view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) GetBoolBasicFee() (*big.Int, error) {
	return _SoDiamond.Contract.GetBoolBasicFee(&_SoDiamond.CallOpts)
}

// GetBoolSoFee is a free data retrieval call binding the contract method 0x912789f0.
//
// Solidity: function getBoolSoFee(uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondCaller) GetBoolSoFee(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getBoolSoFee", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBoolSoFee is a free data retrieval call binding the contract method 0x912789f0.
//
// Solidity: function getBoolSoFee(uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondSession) GetBoolSoFee(amount *big.Int) (*big.Int, error) {
	return _SoDiamond.Contract.GetBoolSoFee(&_SoDiamond.CallOpts, amount)
}

// GetBoolSoFee is a free data retrieval call binding the contract method 0x912789f0.
//
// Solidity: function getBoolSoFee(uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) GetBoolSoFee(amount *big.Int) (*big.Int, error) {
	return _SoDiamond.Contract.GetBoolSoFee(&_SoDiamond.CallOpts, amount)
}

// GetBoolTransferGas is a free data retrieval call binding the contract method 0x426117bd.
//
// Solidity: function getBoolTransferGas() view returns(uint256)
func (_SoDiamond *SoDiamondCaller) GetBoolTransferGas(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getBoolTransferGas")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBoolTransferGas is a free data retrieval call binding the contract method 0x426117bd.
//
// Solidity: function getBoolTransferGas() view returns(uint256)
func (_SoDiamond *SoDiamondSession) GetBoolTransferGas() (*big.Int, error) {
	return _SoDiamond.Contract.GetBoolTransferGas(&_SoDiamond.CallOpts)
}

// GetBoolTransferGas is a free data retrieval call binding the contract method 0x426117bd.
//
// Solidity: function getBoolTransferGas() view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) GetBoolTransferGas() (*big.Int, error) {
	return _SoDiamond.Contract.GetBoolTransferGas(&_SoDiamond.CallOpts)
}

// GetCelerMessageFee1 is a free data retrieval call binding the contract method 0x491ebe76.
//
// Solidity: function getCelerMessageFee1(address messageBus, bytes message) view returns(uint256)
func (_SoDiamond *SoDiamondCaller) GetCelerMessageFee1(opts *bind.CallOpts, messageBus common.Address, message []byte) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getCelerMessageFee1", messageBus, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCelerMessageFee1 is a free data retrieval call binding the contract method 0x491ebe76.
//
// Solidity: function getCelerMessageFee1(address messageBus, bytes message) view returns(uint256)
func (_SoDiamond *SoDiamondSession) GetCelerMessageFee1(messageBus common.Address, message []byte) (*big.Int, error) {
	return _SoDiamond.Contract.GetCelerMessageFee1(&_SoDiamond.CallOpts, messageBus, message)
}

// GetCelerMessageFee1 is a free data retrieval call binding the contract method 0x491ebe76.
//
// Solidity: function getCelerMessageFee1(address messageBus, bytes message) view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) GetCelerMessageFee1(messageBus common.Address, message []byte) (*big.Int, error) {
	return _SoDiamond.Contract.GetCelerMessageFee1(&_SoDiamond.CallOpts, messageBus, message)
}

// GetCelerMessageFee2 is a free data retrieval call binding the contract method 0xf64fa004.
//
// Solidity: function getCelerMessageFee2(bytes message) view returns(uint256)
func (_SoDiamond *SoDiamondCaller) GetCelerMessageFee2(opts *bind.CallOpts, message []byte) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getCelerMessageFee2", message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCelerMessageFee2 is a free data retrieval call binding the contract method 0xf64fa004.
//
// Solidity: function getCelerMessageFee2(bytes message) view returns(uint256)
func (_SoDiamond *SoDiamondSession) GetCelerMessageFee2(message []byte) (*big.Int, error) {
	return _SoDiamond.Contract.GetCelerMessageFee2(&_SoDiamond.CallOpts, message)
}

// GetCelerMessageFee2 is a free data retrieval call binding the contract method 0xf64fa004.
//
// Solidity: function getCelerMessageFee2(bytes message) view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) GetCelerMessageFee2(message []byte) (*big.Int, error) {
	return _SoDiamond.Contract.GetCelerMessageFee2(&_SoDiamond.CallOpts, message)
}

// GetCelerSoFee is a free data retrieval call binding the contract method 0x9bab85ad.
//
// Solidity: function getCelerSoFee(uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondCaller) GetCelerSoFee(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getCelerSoFee", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCelerSoFee is a free data retrieval call binding the contract method 0x9bab85ad.
//
// Solidity: function getCelerSoFee(uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondSession) GetCelerSoFee(amount *big.Int) (*big.Int, error) {
	return _SoDiamond.Contract.GetCelerSoFee(&_SoDiamond.CallOpts, amount)
}

// GetCelerSoFee is a free data retrieval call binding the contract method 0x9bab85ad.
//
// Solidity: function getCelerSoFee(uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) GetCelerSoFee(amount *big.Int) (*big.Int, error) {
	return _SoDiamond.Contract.GetCelerSoFee(&_SoDiamond.CallOpts, amount)
}

// GetExecutorFeeTo is a free data retrieval call binding the contract method 0x07db0be3.
//
// Solidity: function getExecutorFeeTo() view returns(address)
func (_SoDiamond *SoDiamondCaller) GetExecutorFeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getExecutorFeeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExecutorFeeTo is a free data retrieval call binding the contract method 0x07db0be3.
//
// Solidity: function getExecutorFeeTo() view returns(address)
func (_SoDiamond *SoDiamondSession) GetExecutorFeeTo() (common.Address, error) {
	return _SoDiamond.Contract.GetExecutorFeeTo(&_SoDiamond.CallOpts)
}

// GetExecutorFeeTo is a free data retrieval call binding the contract method 0x07db0be3.
//
// Solidity: function getExecutorFeeTo() view returns(address)
func (_SoDiamond *SoDiamondCallerSession) GetExecutorFeeTo() (common.Address, error) {
	return _SoDiamond.Contract.GetExecutorFeeTo(&_SoDiamond.CallOpts)
}

// GetFastRouter is a free data retrieval call binding the contract method 0x036fe0e5.
//
// Solidity: function getFastRouter() view returns(address)
func (_SoDiamond *SoDiamondCaller) GetFastRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getFastRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFastRouter is a free data retrieval call binding the contract method 0x036fe0e5.
//
// Solidity: function getFastRouter() view returns(address)
func (_SoDiamond *SoDiamondSession) GetFastRouter() (common.Address, error) {
	return _SoDiamond.Contract.GetFastRouter(&_SoDiamond.CallOpts)
}

// GetFastRouter is a free data retrieval call binding the contract method 0x036fe0e5.
//
// Solidity: function getFastRouter() view returns(address)
func (_SoDiamond *SoDiamondCallerSession) GetFastRouter() (common.Address, error) {
	return _SoDiamond.Contract.GetFastRouter(&_SoDiamond.CallOpts)
}

// GetMultiChainSoFee is a free data retrieval call binding the contract method 0x09790796.
//
// Solidity: function getMultiChainSoFee(uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondCaller) GetMultiChainSoFee(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getMultiChainSoFee", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMultiChainSoFee is a free data retrieval call binding the contract method 0x09790796.
//
// Solidity: function getMultiChainSoFee(uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondSession) GetMultiChainSoFee(amount *big.Int) (*big.Int, error) {
	return _SoDiamond.Contract.GetMultiChainSoFee(&_SoDiamond.CallOpts, amount)
}

// GetMultiChainSoFee is a free data retrieval call binding the contract method 0x09790796.
//
// Solidity: function getMultiChainSoFee(uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) GetMultiChainSoFee(amount *big.Int) (*big.Int, error) {
	return _SoDiamond.Contract.GetMultiChainSoFee(&_SoDiamond.CallOpts, amount)
}

// GetNativeWrap is a free data retrieval call binding the contract method 0x755ad5da.
//
// Solidity: function getNativeWrap(address messageBus) view returns(address)
func (_SoDiamond *SoDiamondCaller) GetNativeWrap(opts *bind.CallOpts, messageBus common.Address) (common.Address, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getNativeWrap", messageBus)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNativeWrap is a free data retrieval call binding the contract method 0x755ad5da.
//
// Solidity: function getNativeWrap(address messageBus) view returns(address)
func (_SoDiamond *SoDiamondSession) GetNativeWrap(messageBus common.Address) (common.Address, error) {
	return _SoDiamond.Contract.GetNativeWrap(&_SoDiamond.CallOpts, messageBus)
}

// GetNativeWrap is a free data retrieval call binding the contract method 0x755ad5da.
//
// Solidity: function getNativeWrap(address messageBus) view returns(address)
func (_SoDiamond *SoDiamondCallerSession) GetNativeWrap(messageBus common.Address) (common.Address, error) {
	return _SoDiamond.Contract.GetNativeWrap(&_SoDiamond.CallOpts, messageBus)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint64)
func (_SoDiamond *SoDiamondCaller) GetNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint64)
func (_SoDiamond *SoDiamondSession) GetNonce() (uint64, error) {
	return _SoDiamond.Contract.GetNonce(&_SoDiamond.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint64)
func (_SoDiamond *SoDiamondCallerSession) GetNonce() (uint64, error) {
	return _SoDiamond.Contract.GetNonce(&_SoDiamond.CallOpts)
}

// GetSgReceiveForGasPayload is a free data retrieval call binding the contract method 0x3136c560.
//
// Solidity: function getSgReceiveForGasPayload((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) pure returns(bytes)
func (_SoDiamond *SoDiamondCaller) GetSgReceiveForGasPayload(opts *bind.CallOpts, soDataNo ISoNormalizedSoData, swapDataDstNo []LibSwapNormalizedSwapData) ([]byte, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getSgReceiveForGasPayload", soDataNo, swapDataDstNo)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetSgReceiveForGasPayload is a free data retrieval call binding the contract method 0x3136c560.
//
// Solidity: function getSgReceiveForGasPayload((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) pure returns(bytes)
func (_SoDiamond *SoDiamondSession) GetSgReceiveForGasPayload(soDataNo ISoNormalizedSoData, swapDataDstNo []LibSwapNormalizedSwapData) ([]byte, error) {
	return _SoDiamond.Contract.GetSgReceiveForGasPayload(&_SoDiamond.CallOpts, soDataNo, swapDataDstNo)
}

// GetSgReceiveForGasPayload is a free data retrieval call binding the contract method 0x3136c560.
//
// Solidity: function getSgReceiveForGasPayload((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) pure returns(bytes)
func (_SoDiamond *SoDiamondCallerSession) GetSgReceiveForGasPayload(soDataNo ISoNormalizedSoData, swapDataDstNo []LibSwapNormalizedSwapData) ([]byte, error) {
	return _SoDiamond.Contract.GetSgReceiveForGasPayload(&_SoDiamond.CallOpts, soDataNo, swapDataDstNo)
}

// GetStargateFee is a free data retrieval call binding the contract method 0xe31dc6be.
//
// Solidity: function getStargateFee((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (uint256,uint16,uint256,uint256,uint256,address) stargateData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) view returns(uint256)
func (_SoDiamond *SoDiamondCaller) GetStargateFee(opts *bind.CallOpts, soDataNo ISoNormalizedSoData, stargateData StargateFacetStargateData, swapDataDstNo []LibSwapNormalizedSwapData) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getStargateFee", soDataNo, stargateData, swapDataDstNo)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStargateFee is a free data retrieval call binding the contract method 0xe31dc6be.
//
// Solidity: function getStargateFee((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (uint256,uint16,uint256,uint256,uint256,address) stargateData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) view returns(uint256)
func (_SoDiamond *SoDiamondSession) GetStargateFee(soDataNo ISoNormalizedSoData, stargateData StargateFacetStargateData, swapDataDstNo []LibSwapNormalizedSwapData) (*big.Int, error) {
	return _SoDiamond.Contract.GetStargateFee(&_SoDiamond.CallOpts, soDataNo, stargateData, swapDataDstNo)
}

// GetStargateFee is a free data retrieval call binding the contract method 0xe31dc6be.
//
// Solidity: function getStargateFee((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (uint256,uint16,uint256,uint256,uint256,address) stargateData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) GetStargateFee(soDataNo ISoNormalizedSoData, stargateData StargateFacetStargateData, swapDataDstNo []LibSwapNormalizedSwapData) (*big.Int, error) {
	return _SoDiamond.Contract.GetStargateFee(&_SoDiamond.CallOpts, soDataNo, stargateData, swapDataDstNo)
}

// GetStargateSoFee is a free data retrieval call binding the contract method 0x98e038d7.
//
// Solidity: function getStargateSoFee(uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondCaller) GetStargateSoFee(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getStargateSoFee", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStargateSoFee is a free data retrieval call binding the contract method 0x98e038d7.
//
// Solidity: function getStargateSoFee(uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondSession) GetStargateSoFee(amount *big.Int) (*big.Int, error) {
	return _SoDiamond.Contract.GetStargateSoFee(&_SoDiamond.CallOpts, amount)
}

// GetStargateSoFee is a free data retrieval call binding the contract method 0x98e038d7.
//
// Solidity: function getStargateSoFee(uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) GetStargateSoFee(amount *big.Int) (*big.Int, error) {
	return _SoDiamond.Contract.GetStargateSoFee(&_SoDiamond.CallOpts, amount)
}

// GetTransferGas is a free data retrieval call binding the contract method 0x0e7e8ba2.
//
// Solidity: function getTransferGas() view returns(uint256)
func (_SoDiamond *SoDiamondCaller) GetTransferGas(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getTransferGas")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransferGas is a free data retrieval call binding the contract method 0x0e7e8ba2.
//
// Solidity: function getTransferGas() view returns(uint256)
func (_SoDiamond *SoDiamondSession) GetTransferGas() (*big.Int, error) {
	return _SoDiamond.Contract.GetTransferGas(&_SoDiamond.CallOpts)
}

// GetTransferGas is a free data retrieval call binding the contract method 0x0e7e8ba2.
//
// Solidity: function getTransferGas() view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) GetTransferGas() (*big.Int, error) {
	return _SoDiamond.Contract.GetTransferGas(&_SoDiamond.CallOpts)
}

// GetWormholeFinalAmount is a free data retrieval call binding the contract method 0x599d853d.
//
// Solidity: function getWormholeFinalAmount(address transferToken, uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondCaller) GetWormholeFinalAmount(opts *bind.CallOpts, transferToken common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getWormholeFinalAmount", transferToken, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWormholeFinalAmount is a free data retrieval call binding the contract method 0x599d853d.
//
// Solidity: function getWormholeFinalAmount(address transferToken, uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondSession) GetWormholeFinalAmount(transferToken common.Address, amount *big.Int) (*big.Int, error) {
	return _SoDiamond.Contract.GetWormholeFinalAmount(&_SoDiamond.CallOpts, transferToken, amount)
}

// GetWormholeFinalAmount is a free data retrieval call binding the contract method 0x599d853d.
//
// Solidity: function getWormholeFinalAmount(address transferToken, uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) GetWormholeFinalAmount(transferToken common.Address, amount *big.Int) (*big.Int, error) {
	return _SoDiamond.Contract.GetWormholeFinalAmount(&_SoDiamond.CallOpts, transferToken, amount)
}

// GetWormholeMessageFee is a free data retrieval call binding the contract method 0x22dfbee5.
//
// Solidity: function getWormholeMessageFee() view returns(uint256)
func (_SoDiamond *SoDiamondCaller) GetWormholeMessageFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getWormholeMessageFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWormholeMessageFee is a free data retrieval call binding the contract method 0x22dfbee5.
//
// Solidity: function getWormholeMessageFee() view returns(uint256)
func (_SoDiamond *SoDiamondSession) GetWormholeMessageFee() (*big.Int, error) {
	return _SoDiamond.Contract.GetWormholeMessageFee(&_SoDiamond.CallOpts)
}

// GetWormholeMessageFee is a free data retrieval call binding the contract method 0x22dfbee5.
//
// Solidity: function getWormholeMessageFee() view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) GetWormholeMessageFee() (*big.Int, error) {
	return _SoDiamond.Contract.GetWormholeMessageFee(&_SoDiamond.CallOpts)
}

// GetWormholeSoFee is a free data retrieval call binding the contract method 0x278126b2.
//
// Solidity: function getWormholeSoFee(uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondCaller) GetWormholeSoFee(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "getWormholeSoFee", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWormholeSoFee is a free data retrieval call binding the contract method 0x278126b2.
//
// Solidity: function getWormholeSoFee(uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondSession) GetWormholeSoFee(amount *big.Int) (*big.Int, error) {
	return _SoDiamond.Contract.GetWormholeSoFee(&_SoDiamond.CallOpts, amount)
}

// GetWormholeSoFee is a free data retrieval call binding the contract method 0x278126b2.
//
// Solidity: function getWormholeSoFee(uint256 amount) view returns(uint256)
func (_SoDiamond *SoDiamondCallerSession) GetWormholeSoFee(amount *big.Int) (*big.Int, error) {
	return _SoDiamond.Contract.GetWormholeSoFee(&_SoDiamond.CallOpts, amount)
}

// IsFunctionApproved is a free data retrieval call binding the contract method 0xba1a5bdd.
//
// Solidity: function isFunctionApproved(bytes32 signature) view returns(bool approved)
func (_SoDiamond *SoDiamondCaller) IsFunctionApproved(opts *bind.CallOpts, signature [32]byte) (bool, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "isFunctionApproved", signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFunctionApproved is a free data retrieval call binding the contract method 0xba1a5bdd.
//
// Solidity: function isFunctionApproved(bytes32 signature) view returns(bool approved)
func (_SoDiamond *SoDiamondSession) IsFunctionApproved(signature [32]byte) (bool, error) {
	return _SoDiamond.Contract.IsFunctionApproved(&_SoDiamond.CallOpts, signature)
}

// IsFunctionApproved is a free data retrieval call binding the contract method 0xba1a5bdd.
//
// Solidity: function isFunctionApproved(bytes32 signature) view returns(bool approved)
func (_SoDiamond *SoDiamondCallerSession) IsFunctionApproved(signature [32]byte) (bool, error) {
	return _SoDiamond.Contract.IsFunctionApproved(&_SoDiamond.CallOpts, signature)
}

// IsValidMultiChainConfig is a free data retrieval call binding the contract method 0x5bf358df.
//
// Solidity: function isValidMultiChainConfig() view returns(bool)
func (_SoDiamond *SoDiamondCaller) IsValidMultiChainConfig(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "isValidMultiChainConfig")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidMultiChainConfig is a free data retrieval call binding the contract method 0x5bf358df.
//
// Solidity: function isValidMultiChainConfig() view returns(bool)
func (_SoDiamond *SoDiamondSession) IsValidMultiChainConfig() (bool, error) {
	return _SoDiamond.Contract.IsValidMultiChainConfig(&_SoDiamond.CallOpts)
}

// IsValidMultiChainConfig is a free data retrieval call binding the contract method 0x5bf358df.
//
// Solidity: function isValidMultiChainConfig() view returns(bool)
func (_SoDiamond *SoDiamondCallerSession) IsValidMultiChainConfig() (bool, error) {
	return _SoDiamond.Contract.IsValidMultiChainConfig(&_SoDiamond.CallOpts)
}

// NormalizeSoData is a free data retrieval call binding the contract method 0xb3609519.
//
// Solidity: function normalizeSoData((bytes32,address,uint16,address,uint16,address,uint256) soData) pure returns((bytes,bytes,uint16,bytes,uint16,bytes,uint256))
func (_SoDiamond *SoDiamondCaller) NormalizeSoData(opts *bind.CallOpts, soData ISoSoData) (ISoNormalizedSoData, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "normalizeSoData", soData)

	if err != nil {
		return *new(ISoNormalizedSoData), err
	}

	out0 := *abi.ConvertType(out[0], new(ISoNormalizedSoData)).(*ISoNormalizedSoData)

	return out0, err

}

// NormalizeSoData is a free data retrieval call binding the contract method 0xb3609519.
//
// Solidity: function normalizeSoData((bytes32,address,uint16,address,uint16,address,uint256) soData) pure returns((bytes,bytes,uint16,bytes,uint16,bytes,uint256))
func (_SoDiamond *SoDiamondSession) NormalizeSoData(soData ISoSoData) (ISoNormalizedSoData, error) {
	return _SoDiamond.Contract.NormalizeSoData(&_SoDiamond.CallOpts, soData)
}

// NormalizeSoData is a free data retrieval call binding the contract method 0xb3609519.
//
// Solidity: function normalizeSoData((bytes32,address,uint16,address,uint16,address,uint256) soData) pure returns((bytes,bytes,uint16,bytes,uint16,bytes,uint256))
func (_SoDiamond *SoDiamondCallerSession) NormalizeSoData(soData ISoSoData) (ISoNormalizedSoData, error) {
	return _SoDiamond.Contract.NormalizeSoData(&_SoDiamond.CallOpts, soData)
}

// NormalizeSwapData is a free data retrieval call binding the contract method 0xe865edfc.
//
// Solidity: function normalizeSwapData((address,address,address,address,uint256,bytes)[] swapData) pure returns((bytes,bytes,bytes,bytes,uint256,bytes)[])
func (_SoDiamond *SoDiamondCaller) NormalizeSwapData(opts *bind.CallOpts, swapData []LibSwapSwapData) ([]LibSwapNormalizedSwapData, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "normalizeSwapData", swapData)

	if err != nil {
		return *new([]LibSwapNormalizedSwapData), err
	}

	out0 := *abi.ConvertType(out[0], new([]LibSwapNormalizedSwapData)).(*[]LibSwapNormalizedSwapData)

	return out0, err

}

// NormalizeSwapData is a free data retrieval call binding the contract method 0xe865edfc.
//
// Solidity: function normalizeSwapData((address,address,address,address,uint256,bytes)[] swapData) pure returns((bytes,bytes,bytes,bytes,uint256,bytes)[])
func (_SoDiamond *SoDiamondSession) NormalizeSwapData(swapData []LibSwapSwapData) ([]LibSwapNormalizedSwapData, error) {
	return _SoDiamond.Contract.NormalizeSwapData(&_SoDiamond.CallOpts, swapData)
}

// NormalizeSwapData is a free data retrieval call binding the contract method 0xe865edfc.
//
// Solidity: function normalizeSwapData((address,address,address,address,uint256,bytes)[] swapData) pure returns((bytes,bytes,bytes,bytes,uint256,bytes)[])
func (_SoDiamond *SoDiamondCallerSession) NormalizeSwapData(swapData []LibSwapSwapData) ([]LibSwapNormalizedSwapData, error) {
	return _SoDiamond.Contract.NormalizeSwapData(&_SoDiamond.CallOpts, swapData)
}

// NormalizeU256 is a free data retrieval call binding the contract method 0xb3dac9a5.
//
// Solidity: function normalizeU256(uint256 data) pure returns(bytes)
func (_SoDiamond *SoDiamondCaller) NormalizeU256(opts *bind.CallOpts, data *big.Int) ([]byte, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "normalizeU256", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// NormalizeU256 is a free data retrieval call binding the contract method 0xb3dac9a5.
//
// Solidity: function normalizeU256(uint256 data) pure returns(bytes)
func (_SoDiamond *SoDiamondSession) NormalizeU256(data *big.Int) ([]byte, error) {
	return _SoDiamond.Contract.NormalizeU256(&_SoDiamond.CallOpts, data)
}

// NormalizeU256 is a free data retrieval call binding the contract method 0xb3dac9a5.
//
// Solidity: function normalizeU256(uint256 data) pure returns(bytes)
func (_SoDiamond *SoDiamondCallerSession) NormalizeU256(data *big.Int) ([]byte, error) {
	return _SoDiamond.Contract.NormalizeU256(&_SoDiamond.CallOpts, data)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address contractOwner)
func (_SoDiamond *SoDiamondCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address contractOwner)
func (_SoDiamond *SoDiamondSession) Owner() (common.Address, error) {
	return _SoDiamond.Contract.Owner(&_SoDiamond.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address contractOwner)
func (_SoDiamond *SoDiamondCallerSession) Owner() (common.Address, error) {
	return _SoDiamond.Contract.Owner(&_SoDiamond.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SoDiamond *SoDiamondCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SoDiamond.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SoDiamond *SoDiamondSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SoDiamond.Contract.SupportsInterface(&_SoDiamond.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SoDiamond *SoDiamondCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SoDiamond.Contract.SupportsInterface(&_SoDiamond.CallOpts, interfaceId)
}

// AddCorrectSwap is a paid mutator transaction binding the contract method 0xd0ee8c19.
//
// Solidity: function addCorrectSwap(address correctSwap) returns()
func (_SoDiamond *SoDiamondTransactor) AddCorrectSwap(opts *bind.TransactOpts, correctSwap common.Address) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "addCorrectSwap", correctSwap)
}

// AddCorrectSwap is a paid mutator transaction binding the contract method 0xd0ee8c19.
//
// Solidity: function addCorrectSwap(address correctSwap) returns()
func (_SoDiamond *SoDiamondSession) AddCorrectSwap(correctSwap common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.AddCorrectSwap(&_SoDiamond.TransactOpts, correctSwap)
}

// AddCorrectSwap is a paid mutator transaction binding the contract method 0xd0ee8c19.
//
// Solidity: function addCorrectSwap(address correctSwap) returns()
func (_SoDiamond *SoDiamondTransactorSession) AddCorrectSwap(correctSwap common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.AddCorrectSwap(&_SoDiamond.TransactOpts, correctSwap)
}

// AddDex is a paid mutator transaction binding the contract method 0x536db266.
//
// Solidity: function addDex(address dex) returns()
func (_SoDiamond *SoDiamondTransactor) AddDex(opts *bind.TransactOpts, dex common.Address) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "addDex", dex)
}

// AddDex is a paid mutator transaction binding the contract method 0x536db266.
//
// Solidity: function addDex(address dex) returns()
func (_SoDiamond *SoDiamondSession) AddDex(dex common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.AddDex(&_SoDiamond.TransactOpts, dex)
}

// AddDex is a paid mutator transaction binding the contract method 0x536db266.
//
// Solidity: function addDex(address dex) returns()
func (_SoDiamond *SoDiamondTransactorSession) AddDex(dex common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.AddDex(&_SoDiamond.TransactOpts, dex)
}

// AddFee is a paid mutator transaction binding the contract method 0xdf41d729.
//
// Solidity: function addFee(address gateway, address soFee) returns()
func (_SoDiamond *SoDiamondTransactor) AddFee(opts *bind.TransactOpts, gateway common.Address, soFee common.Address) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "addFee", gateway, soFee)
}

// AddFee is a paid mutator transaction binding the contract method 0xdf41d729.
//
// Solidity: function addFee(address gateway, address soFee) returns()
func (_SoDiamond *SoDiamondSession) AddFee(gateway common.Address, soFee common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.AddFee(&_SoDiamond.TransactOpts, gateway, soFee)
}

// AddFee is a paid mutator transaction binding the contract method 0xdf41d729.
//
// Solidity: function addFee(address gateway, address soFee) returns()
func (_SoDiamond *SoDiamondTransactorSession) AddFee(gateway common.Address, soFee common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.AddFee(&_SoDiamond.TransactOpts, gateway, soFee)
}

// BatchAddDex is a paid mutator transaction binding the contract method 0xfcd8e49e.
//
// Solidity: function batchAddDex(address[] dexs) returns()
func (_SoDiamond *SoDiamondTransactor) BatchAddDex(opts *bind.TransactOpts, dexs []common.Address) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "batchAddDex", dexs)
}

// BatchAddDex is a paid mutator transaction binding the contract method 0xfcd8e49e.
//
// Solidity: function batchAddDex(address[] dexs) returns()
func (_SoDiamond *SoDiamondSession) BatchAddDex(dexs []common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.BatchAddDex(&_SoDiamond.TransactOpts, dexs)
}

// BatchAddDex is a paid mutator transaction binding the contract method 0xfcd8e49e.
//
// Solidity: function batchAddDex(address[] dexs) returns()
func (_SoDiamond *SoDiamondTransactorSession) BatchAddDex(dexs []common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.BatchAddDex(&_SoDiamond.TransactOpts, dexs)
}

// BatchRemoveDex is a paid mutator transaction binding the contract method 0x9afc19c7.
//
// Solidity: function batchRemoveDex(address[] dexs) returns()
func (_SoDiamond *SoDiamondTransactor) BatchRemoveDex(opts *bind.TransactOpts, dexs []common.Address) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "batchRemoveDex", dexs)
}

// BatchRemoveDex is a paid mutator transaction binding the contract method 0x9afc19c7.
//
// Solidity: function batchRemoveDex(address[] dexs) returns()
func (_SoDiamond *SoDiamondSession) BatchRemoveDex(dexs []common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.BatchRemoveDex(&_SoDiamond.TransactOpts, dexs)
}

// BatchRemoveDex is a paid mutator transaction binding the contract method 0x9afc19c7.
//
// Solidity: function batchRemoveDex(address[] dexs) returns()
func (_SoDiamond *SoDiamondTransactorSession) BatchRemoveDex(dexs []common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.BatchRemoveDex(&_SoDiamond.TransactOpts, dexs)
}

// BatchSetBoolAllowedAddresses is a paid mutator transaction binding the contract method 0xf3e4836c.
//
// Solidity: function batchSetBoolAllowedAddresses(address[] boolSwapPools, bool[] isAllowed) returns()
func (_SoDiamond *SoDiamondTransactor) BatchSetBoolAllowedAddresses(opts *bind.TransactOpts, boolSwapPools []common.Address, isAllowed []bool) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "batchSetBoolAllowedAddresses", boolSwapPools, isAllowed)
}

// BatchSetBoolAllowedAddresses is a paid mutator transaction binding the contract method 0xf3e4836c.
//
// Solidity: function batchSetBoolAllowedAddresses(address[] boolSwapPools, bool[] isAllowed) returns()
func (_SoDiamond *SoDiamondSession) BatchSetBoolAllowedAddresses(boolSwapPools []common.Address, isAllowed []bool) (*types.Transaction, error) {
	return _SoDiamond.Contract.BatchSetBoolAllowedAddresses(&_SoDiamond.TransactOpts, boolSwapPools, isAllowed)
}

// BatchSetBoolAllowedAddresses is a paid mutator transaction binding the contract method 0xf3e4836c.
//
// Solidity: function batchSetBoolAllowedAddresses(address[] boolSwapPools, bool[] isAllowed) returns()
func (_SoDiamond *SoDiamondTransactorSession) BatchSetBoolAllowedAddresses(boolSwapPools []common.Address, isAllowed []bool) (*types.Transaction, error) {
	return _SoDiamond.Contract.BatchSetBoolAllowedAddresses(&_SoDiamond.TransactOpts, boolSwapPools, isAllowed)
}

// BatchSetFunctionApprovalBySignature is a paid mutator transaction binding the contract method 0x4363b848.
//
// Solidity: function batchSetFunctionApprovalBySignature(bytes32[] signatures, bool approval) returns()
func (_SoDiamond *SoDiamondTransactor) BatchSetFunctionApprovalBySignature(opts *bind.TransactOpts, signatures [][32]byte, approval bool) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "batchSetFunctionApprovalBySignature", signatures, approval)
}

// BatchSetFunctionApprovalBySignature is a paid mutator transaction binding the contract method 0x4363b848.
//
// Solidity: function batchSetFunctionApprovalBySignature(bytes32[] signatures, bool approval) returns()
func (_SoDiamond *SoDiamondSession) BatchSetFunctionApprovalBySignature(signatures [][32]byte, approval bool) (*types.Transaction, error) {
	return _SoDiamond.Contract.BatchSetFunctionApprovalBySignature(&_SoDiamond.TransactOpts, signatures, approval)
}

// BatchSetFunctionApprovalBySignature is a paid mutator transaction binding the contract method 0x4363b848.
//
// Solidity: function batchSetFunctionApprovalBySignature(bytes32[] signatures, bool approval) returns()
func (_SoDiamond *SoDiamondTransactorSession) BatchSetFunctionApprovalBySignature(signatures [][32]byte, approval bool) (*types.Transaction, error) {
	return _SoDiamond.Contract.BatchSetFunctionApprovalBySignature(&_SoDiamond.TransactOpts, signatures, approval)
}

// CancelOnwershipTransfer is a paid mutator transaction binding the contract method 0x0f4efb90.
//
// Solidity: function cancelOnwershipTransfer() returns()
func (_SoDiamond *SoDiamondTransactor) CancelOnwershipTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "cancelOnwershipTransfer")
}

// CancelOnwershipTransfer is a paid mutator transaction binding the contract method 0x0f4efb90.
//
// Solidity: function cancelOnwershipTransfer() returns()
func (_SoDiamond *SoDiamondSession) CancelOnwershipTransfer() (*types.Transaction, error) {
	return _SoDiamond.Contract.CancelOnwershipTransfer(&_SoDiamond.TransactOpts)
}

// CancelOnwershipTransfer is a paid mutator transaction binding the contract method 0x0f4efb90.
//
// Solidity: function cancelOnwershipTransfer() returns()
func (_SoDiamond *SoDiamondTransactorSession) CancelOnwershipTransfer() (*types.Transaction, error) {
	return _SoDiamond.Contract.CancelOnwershipTransfer(&_SoDiamond.TransactOpts)
}

// CheckExecutorFee is a paid mutator transaction binding the contract method 0x6035de09.
//
// Solidity: function checkExecutorFee((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (address,uint32,uint64,address,uint256,uint256,address) celerData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst) returns(bool, uint256, uint256, uint256)
func (_SoDiamond *SoDiamondTransactor) CheckExecutorFee(opts *bind.TransactOpts, soData ISoNormalizedSoData, celerData CelerFacetCelerData, swapDataDst []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "checkExecutorFee", soData, celerData, swapDataDst)
}

// CheckExecutorFee is a paid mutator transaction binding the contract method 0x6035de09.
//
// Solidity: function checkExecutorFee((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (address,uint32,uint64,address,uint256,uint256,address) celerData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst) returns(bool, uint256, uint256, uint256)
func (_SoDiamond *SoDiamondSession) CheckExecutorFee(soData ISoNormalizedSoData, celerData CelerFacetCelerData, swapDataDst []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.CheckExecutorFee(&_SoDiamond.TransactOpts, soData, celerData, swapDataDst)
}

// CheckExecutorFee is a paid mutator transaction binding the contract method 0x6035de09.
//
// Solidity: function checkExecutorFee((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (address,uint32,uint64,address,uint256,uint256,address) celerData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst) returns(bool, uint256, uint256, uint256)
func (_SoDiamond *SoDiamondTransactorSession) CheckExecutorFee(soData ISoNormalizedSoData, celerData CelerFacetCelerData, swapDataDst []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.CheckExecutorFee(&_SoDiamond.TransactOpts, soData, celerData, swapDataDst)
}

// CheckRelayerFee is a paid mutator transaction binding the contract method 0x8f6d817a.
//
// Solidity: function checkRelayerFee((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (uint16,uint256,uint256,bytes) wormholeData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst) returns(bool, uint256, uint256, uint256)
func (_SoDiamond *SoDiamondTransactor) CheckRelayerFee(opts *bind.TransactOpts, soData ISoNormalizedSoData, wormholeData WormholeFacetNormalizedWormholeData, swapDataDst []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "checkRelayerFee", soData, wormholeData, swapDataDst)
}

// CheckRelayerFee is a paid mutator transaction binding the contract method 0x8f6d817a.
//
// Solidity: function checkRelayerFee((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (uint16,uint256,uint256,bytes) wormholeData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst) returns(bool, uint256, uint256, uint256)
func (_SoDiamond *SoDiamondSession) CheckRelayerFee(soData ISoNormalizedSoData, wormholeData WormholeFacetNormalizedWormholeData, swapDataDst []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.CheckRelayerFee(&_SoDiamond.TransactOpts, soData, wormholeData, swapDataDst)
}

// CheckRelayerFee is a paid mutator transaction binding the contract method 0x8f6d817a.
//
// Solidity: function checkRelayerFee((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soData, (uint16,uint256,uint256,bytes) wormholeData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDst) returns(bool, uint256, uint256, uint256)
func (_SoDiamond *SoDiamondTransactorSession) CheckRelayerFee(soData ISoNormalizedSoData, wormholeData WormholeFacetNormalizedWormholeData, swapDataDst []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.CheckRelayerFee(&_SoDiamond.TransactOpts, soData, wormholeData, swapDataDst)
}

// CompleteSoSwap is a paid mutator transaction binding the contract method 0xe0718403.
//
// Solidity: function completeSoSwap(bytes encodeVm) returns()
func (_SoDiamond *SoDiamondTransactor) CompleteSoSwap(opts *bind.TransactOpts, encodeVm []byte) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "completeSoSwap", encodeVm)
}

// CompleteSoSwap is a paid mutator transaction binding the contract method 0xe0718403.
//
// Solidity: function completeSoSwap(bytes encodeVm) returns()
func (_SoDiamond *SoDiamondSession) CompleteSoSwap(encodeVm []byte) (*types.Transaction, error) {
	return _SoDiamond.Contract.CompleteSoSwap(&_SoDiamond.TransactOpts, encodeVm)
}

// CompleteSoSwap is a paid mutator transaction binding the contract method 0xe0718403.
//
// Solidity: function completeSoSwap(bytes encodeVm) returns()
func (_SoDiamond *SoDiamondTransactorSession) CompleteSoSwap(encodeVm []byte) (*types.Transaction, error) {
	return _SoDiamond.Contract.CompleteSoSwap(&_SoDiamond.TransactOpts, encodeVm)
}

// CompleteSoSwapByAdmin is a paid mutator transaction binding the contract method 0xeada46dd.
//
// Solidity: function completeSoSwapByAdmin(bytes encodeVm) returns()
func (_SoDiamond *SoDiamondTransactor) CompleteSoSwapByAdmin(opts *bind.TransactOpts, encodeVm []byte) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "completeSoSwapByAdmin", encodeVm)
}

// CompleteSoSwapByAdmin is a paid mutator transaction binding the contract method 0xeada46dd.
//
// Solidity: function completeSoSwapByAdmin(bytes encodeVm) returns()
func (_SoDiamond *SoDiamondSession) CompleteSoSwapByAdmin(encodeVm []byte) (*types.Transaction, error) {
	return _SoDiamond.Contract.CompleteSoSwapByAdmin(&_SoDiamond.TransactOpts, encodeVm)
}

// CompleteSoSwapByAdmin is a paid mutator transaction binding the contract method 0xeada46dd.
//
// Solidity: function completeSoSwapByAdmin(bytes encodeVm) returns()
func (_SoDiamond *SoDiamondTransactorSession) CompleteSoSwapByAdmin(encodeVm []byte) (*types.Transaction, error) {
	return _SoDiamond.Contract.CompleteSoSwapByAdmin(&_SoDiamond.TransactOpts, encodeVm)
}

// ConfirmOwnershipTransfer is a paid mutator transaction binding the contract method 0x7200b829.
//
// Solidity: function confirmOwnershipTransfer() returns()
func (_SoDiamond *SoDiamondTransactor) ConfirmOwnershipTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "confirmOwnershipTransfer")
}

// ConfirmOwnershipTransfer is a paid mutator transaction binding the contract method 0x7200b829.
//
// Solidity: function confirmOwnershipTransfer() returns()
func (_SoDiamond *SoDiamondSession) ConfirmOwnershipTransfer() (*types.Transaction, error) {
	return _SoDiamond.Contract.ConfirmOwnershipTransfer(&_SoDiamond.TransactOpts)
}

// ConfirmOwnershipTransfer is a paid mutator transaction binding the contract method 0x7200b829.
//
// Solidity: function confirmOwnershipTransfer() returns()
func (_SoDiamond *SoDiamondTransactorSession) ConfirmOwnershipTransfer() (*types.Transaction, error) {
	return _SoDiamond.Contract.ConfirmOwnershipTransfer(&_SoDiamond.TransactOpts)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] facetCut, address init, bytes callData) returns()
func (_SoDiamond *SoDiamondTransactor) DiamondCut(opts *bind.TransactOpts, facetCut []IDiamondCutFacetCut, init common.Address, callData []byte) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "diamondCut", facetCut, init, callData)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] facetCut, address init, bytes callData) returns()
func (_SoDiamond *SoDiamondSession) DiamondCut(facetCut []IDiamondCutFacetCut, init common.Address, callData []byte) (*types.Transaction, error) {
	return _SoDiamond.Contract.DiamondCut(&_SoDiamond.TransactOpts, facetCut, init, callData)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] facetCut, address init, bytes callData) returns()
func (_SoDiamond *SoDiamondTransactorSession) DiamondCut(facetCut []IDiamondCutFacetCut, init common.Address, callData []byte) (*types.Transaction, error) {
	return _SoDiamond.Contract.DiamondCut(&_SoDiamond.TransactOpts, facetCut, init, callData)
}

// Exec is a paid mutator transaction binding the contract method 0x28b7b036.
//
// Solidity: function exec(address token, address , uint256 amount, bytes message) returns(bool success, bytes result)
func (_SoDiamond *SoDiamondTransactor) Exec(opts *bind.TransactOpts, token common.Address, arg1 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "exec", token, arg1, amount, message)
}

// Exec is a paid mutator transaction binding the contract method 0x28b7b036.
//
// Solidity: function exec(address token, address , uint256 amount, bytes message) returns(bool success, bytes result)
func (_SoDiamond *SoDiamondSession) Exec(token common.Address, arg1 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _SoDiamond.Contract.Exec(&_SoDiamond.TransactOpts, token, arg1, amount, message)
}

// Exec is a paid mutator transaction binding the contract method 0x28b7b036.
//
// Solidity: function exec(address token, address , uint256 amount, bytes message) returns(bool success, bytes result)
func (_SoDiamond *SoDiamondTransactorSession) Exec(token common.Address, arg1 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _SoDiamond.Contract.Exec(&_SoDiamond.TransactOpts, token, arg1, amount, message)
}

// ExecuteAndCheckSwaps is a paid mutator transaction binding the contract method 0x3729e48c.
//
// Solidity: function executeAndCheckSwaps((bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapData) returns(uint256)
func (_SoDiamond *SoDiamondTransactor) ExecuteAndCheckSwaps(opts *bind.TransactOpts, soData ISoSoData, swapData []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "executeAndCheckSwaps", soData, swapData)
}

// ExecuteAndCheckSwaps is a paid mutator transaction binding the contract method 0x3729e48c.
//
// Solidity: function executeAndCheckSwaps((bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapData) returns(uint256)
func (_SoDiamond *SoDiamondSession) ExecuteAndCheckSwaps(soData ISoSoData, swapData []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteAndCheckSwaps(&_SoDiamond.TransactOpts, soData, swapData)
}

// ExecuteAndCheckSwaps is a paid mutator transaction binding the contract method 0x3729e48c.
//
// Solidity: function executeAndCheckSwaps((bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapData) returns(uint256)
func (_SoDiamond *SoDiamondTransactorSession) ExecuteAndCheckSwaps(soData ISoSoData, swapData []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteAndCheckSwaps(&_SoDiamond.TransactOpts, soData, swapData)
}

// ExecuteAndCheckSwaps0 is a paid mutator transaction binding the contract method 0x3729e48c.
//
// Solidity: function executeAndCheckSwaps((bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapData) returns(uint256)
func (_SoDiamond *SoDiamondTransactor) ExecuteAndCheckSwaps0(opts *bind.TransactOpts, soData ISoSoData, swapData []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "executeAndCheckSwaps0", soData, swapData)
}

// ExecuteAndCheckSwaps0 is a paid mutator transaction binding the contract method 0x3729e48c.
//
// Solidity: function executeAndCheckSwaps((bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapData) returns(uint256)
func (_SoDiamond *SoDiamondSession) ExecuteAndCheckSwaps0(soData ISoSoData, swapData []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteAndCheckSwaps0(&_SoDiamond.TransactOpts, soData, swapData)
}

// ExecuteAndCheckSwaps0 is a paid mutator transaction binding the contract method 0x3729e48c.
//
// Solidity: function executeAndCheckSwaps((bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapData) returns(uint256)
func (_SoDiamond *SoDiamondTransactorSession) ExecuteAndCheckSwaps0(soData ISoSoData, swapData []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteAndCheckSwaps0(&_SoDiamond.TransactOpts, soData, swapData)
}

// ExecuteAndCheckSwaps1 is a paid mutator transaction binding the contract method 0x3729e48c.
//
// Solidity: function executeAndCheckSwaps((bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapData) returns(uint256)
func (_SoDiamond *SoDiamondTransactor) ExecuteAndCheckSwaps1(opts *bind.TransactOpts, soData ISoSoData, swapData []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "executeAndCheckSwaps1", soData, swapData)
}

// ExecuteAndCheckSwaps1 is a paid mutator transaction binding the contract method 0x3729e48c.
//
// Solidity: function executeAndCheckSwaps((bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapData) returns(uint256)
func (_SoDiamond *SoDiamondSession) ExecuteAndCheckSwaps1(soData ISoSoData, swapData []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteAndCheckSwaps1(&_SoDiamond.TransactOpts, soData, swapData)
}

// ExecuteAndCheckSwaps1 is a paid mutator transaction binding the contract method 0x3729e48c.
//
// Solidity: function executeAndCheckSwaps((bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapData) returns(uint256)
func (_SoDiamond *SoDiamondTransactorSession) ExecuteAndCheckSwaps1(soData ISoSoData, swapData []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteAndCheckSwaps1(&_SoDiamond.TransactOpts, soData, swapData)
}

// ExecuteAndCheckSwaps2 is a paid mutator transaction binding the contract method 0x3729e48c.
//
// Solidity: function executeAndCheckSwaps((bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapData) returns(uint256)
func (_SoDiamond *SoDiamondTransactor) ExecuteAndCheckSwaps2(opts *bind.TransactOpts, soData ISoSoData, swapData []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "executeAndCheckSwaps2", soData, swapData)
}

// ExecuteAndCheckSwaps2 is a paid mutator transaction binding the contract method 0x3729e48c.
//
// Solidity: function executeAndCheckSwaps((bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapData) returns(uint256)
func (_SoDiamond *SoDiamondSession) ExecuteAndCheckSwaps2(soData ISoSoData, swapData []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteAndCheckSwaps2(&_SoDiamond.TransactOpts, soData, swapData)
}

// ExecuteAndCheckSwaps2 is a paid mutator transaction binding the contract method 0x3729e48c.
//
// Solidity: function executeAndCheckSwaps((bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapData) returns(uint256)
func (_SoDiamond *SoDiamondTransactorSession) ExecuteAndCheckSwaps2(soData ISoSoData, swapData []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteAndCheckSwaps2(&_SoDiamond.TransactOpts, soData, swapData)
}

// ExecuteAndCheckSwaps3 is a paid mutator transaction binding the contract method 0x3729e48c.
//
// Solidity: function executeAndCheckSwaps((bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapData) returns(uint256)
func (_SoDiamond *SoDiamondTransactor) ExecuteAndCheckSwaps3(opts *bind.TransactOpts, soData ISoSoData, swapData []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "executeAndCheckSwaps3", soData, swapData)
}

// ExecuteAndCheckSwaps3 is a paid mutator transaction binding the contract method 0x3729e48c.
//
// Solidity: function executeAndCheckSwaps((bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapData) returns(uint256)
func (_SoDiamond *SoDiamondSession) ExecuteAndCheckSwaps3(soData ISoSoData, swapData []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteAndCheckSwaps3(&_SoDiamond.TransactOpts, soData, swapData)
}

// ExecuteAndCheckSwaps3 is a paid mutator transaction binding the contract method 0x3729e48c.
//
// Solidity: function executeAndCheckSwaps((bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapData) returns(uint256)
func (_SoDiamond *SoDiamondTransactorSession) ExecuteAndCheckSwaps3(soData ISoSoData, swapData []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteAndCheckSwaps3(&_SoDiamond.TransactOpts, soData, swapData)
}

// ExecuteAndCheckSwaps4 is a paid mutator transaction binding the contract method 0x3729e48c.
//
// Solidity: function executeAndCheckSwaps((bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapData) returns(uint256)
func (_SoDiamond *SoDiamondTransactor) ExecuteAndCheckSwaps4(opts *bind.TransactOpts, soData ISoSoData, swapData []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "executeAndCheckSwaps4", soData, swapData)
}

// ExecuteAndCheckSwaps4 is a paid mutator transaction binding the contract method 0x3729e48c.
//
// Solidity: function executeAndCheckSwaps((bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapData) returns(uint256)
func (_SoDiamond *SoDiamondSession) ExecuteAndCheckSwaps4(soData ISoSoData, swapData []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteAndCheckSwaps4(&_SoDiamond.TransactOpts, soData, swapData)
}

// ExecuteAndCheckSwaps4 is a paid mutator transaction binding the contract method 0x3729e48c.
//
// Solidity: function executeAndCheckSwaps((bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapData) returns(uint256)
func (_SoDiamond *SoDiamondTransactorSession) ExecuteAndCheckSwaps4(soData ISoSoData, swapData []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteAndCheckSwaps4(&_SoDiamond.TransactOpts, soData, swapData)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x063ce4e5.
//
// Solidity: function executeMessage(bytes _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_SoDiamond *SoDiamondTransactor) ExecuteMessage(opts *bind.TransactOpts, _sender []byte, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "executeMessage", _sender, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x063ce4e5.
//
// Solidity: function executeMessage(bytes _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_SoDiamond *SoDiamondSession) ExecuteMessage(_sender []byte, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteMessage(&_SoDiamond.TransactOpts, _sender, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x063ce4e5.
//
// Solidity: function executeMessage(bytes _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_SoDiamond *SoDiamondTransactorSession) ExecuteMessage(_sender []byte, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteMessage(&_SoDiamond.TransactOpts, _sender, _srcChainId, _message, _executor)
}

// ExecuteMessage0 is a paid mutator transaction binding the contract method 0x9c649fdf.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_SoDiamond *SoDiamondTransactor) ExecuteMessage0(opts *bind.TransactOpts, _sender common.Address, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "executeMessage0", _sender, _srcChainId, _message, _executor)
}

// ExecuteMessage0 is a paid mutator transaction binding the contract method 0x9c649fdf.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_SoDiamond *SoDiamondSession) ExecuteMessage0(_sender common.Address, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteMessage0(&_SoDiamond.TransactOpts, _sender, _srcChainId, _message, _executor)
}

// ExecuteMessage0 is a paid mutator transaction binding the contract method 0x9c649fdf.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_SoDiamond *SoDiamondTransactorSession) ExecuteMessage0(_sender common.Address, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteMessage0(&_SoDiamond.TransactOpts, _sender, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x7cd2bffc.
//
// Solidity: function executeMessageWithTransfer(address , address token, uint256 amount, uint64 , bytes message, address ) payable returns(uint8)
func (_SoDiamond *SoDiamondTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, arg0 common.Address, token common.Address, amount *big.Int, arg3 uint64, message []byte, arg5 common.Address) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "executeMessageWithTransfer", arg0, token, amount, arg3, message, arg5)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x7cd2bffc.
//
// Solidity: function executeMessageWithTransfer(address , address token, uint256 amount, uint64 , bytes message, address ) payable returns(uint8)
func (_SoDiamond *SoDiamondSession) ExecuteMessageWithTransfer(arg0 common.Address, token common.Address, amount *big.Int, arg3 uint64, message []byte, arg5 common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteMessageWithTransfer(&_SoDiamond.TransactOpts, arg0, token, amount, arg3, message, arg5)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x7cd2bffc.
//
// Solidity: function executeMessageWithTransfer(address , address token, uint256 amount, uint64 , bytes message, address ) payable returns(uint8)
func (_SoDiamond *SoDiamondTransactorSession) ExecuteMessageWithTransfer(arg0 common.Address, token common.Address, amount *big.Int, arg3 uint64, message []byte, arg5 common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteMessageWithTransfer(&_SoDiamond.TransactOpts, arg0, token, amount, arg3, message, arg5)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0x5ab7afc6.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_SoDiamond *SoDiamondTransactor) ExecuteMessageWithTransferFallback(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "executeMessageWithTransferFallback", _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0x5ab7afc6.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_SoDiamond *SoDiamondSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteMessageWithTransferFallback(&_SoDiamond.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0x5ab7afc6.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_SoDiamond *SoDiamondTransactorSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteMessageWithTransferFallback(&_SoDiamond.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x0bcb4982.
//
// Solidity: function executeMessageWithTransferRefund(address token, uint256 amount, bytes message, address ) payable returns(uint8)
func (_SoDiamond *SoDiamondTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, token common.Address, amount *big.Int, message []byte, arg3 common.Address) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "executeMessageWithTransferRefund", token, amount, message, arg3)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x0bcb4982.
//
// Solidity: function executeMessageWithTransferRefund(address token, uint256 amount, bytes message, address ) payable returns(uint8)
func (_SoDiamond *SoDiamondSession) ExecuteMessageWithTransferRefund(token common.Address, amount *big.Int, message []byte, arg3 common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteMessageWithTransferRefund(&_SoDiamond.TransactOpts, token, amount, message, arg3)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x0bcb4982.
//
// Solidity: function executeMessageWithTransferRefund(address token, uint256 amount, bytes message, address ) payable returns(uint8)
func (_SoDiamond *SoDiamondTransactorSession) ExecuteMessageWithTransferRefund(token common.Address, amount *big.Int, message []byte, arg3 common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.ExecuteMessageWithTransferRefund(&_SoDiamond.TransactOpts, token, amount, message, arg3)
}

// InitBoolSwap is a paid mutator transaction binding the contract method 0x359b4166.
//
// Solidity: function initBoolSwap(address router, uint32 chainId) returns()
func (_SoDiamond *SoDiamondTransactor) InitBoolSwap(opts *bind.TransactOpts, router common.Address, chainId uint32) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "initBoolSwap", router, chainId)
}

// InitBoolSwap is a paid mutator transaction binding the contract method 0x359b4166.
//
// Solidity: function initBoolSwap(address router, uint32 chainId) returns()
func (_SoDiamond *SoDiamondSession) InitBoolSwap(router common.Address, chainId uint32) (*types.Transaction, error) {
	return _SoDiamond.Contract.InitBoolSwap(&_SoDiamond.TransactOpts, router, chainId)
}

// InitBoolSwap is a paid mutator transaction binding the contract method 0x359b4166.
//
// Solidity: function initBoolSwap(address router, uint32 chainId) returns()
func (_SoDiamond *SoDiamondTransactorSession) InitBoolSwap(router common.Address, chainId uint32) (*types.Transaction, error) {
	return _SoDiamond.Contract.InitBoolSwap(&_SoDiamond.TransactOpts, router, chainId)
}

// InitCeler is a paid mutator transaction binding the contract method 0x44ba83b7.
//
// Solidity: function initCeler(address messageBus, uint64 chainId) returns()
func (_SoDiamond *SoDiamondTransactor) InitCeler(opts *bind.TransactOpts, messageBus common.Address, chainId uint64) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "initCeler", messageBus, chainId)
}

// InitCeler is a paid mutator transaction binding the contract method 0x44ba83b7.
//
// Solidity: function initCeler(address messageBus, uint64 chainId) returns()
func (_SoDiamond *SoDiamondSession) InitCeler(messageBus common.Address, chainId uint64) (*types.Transaction, error) {
	return _SoDiamond.Contract.InitCeler(&_SoDiamond.TransactOpts, messageBus, chainId)
}

// InitCeler is a paid mutator transaction binding the contract method 0x44ba83b7.
//
// Solidity: function initCeler(address messageBus, uint64 chainId) returns()
func (_SoDiamond *SoDiamondTransactorSession) InitCeler(messageBus common.Address, chainId uint64) (*types.Transaction, error) {
	return _SoDiamond.Contract.InitCeler(&_SoDiamond.TransactOpts, messageBus, chainId)
}

// InitMultiChain is a paid mutator transaction binding the contract method 0xa1c74577.
//
// Solidity: function initMultiChain(address fastRouter, uint64 chainId) returns()
func (_SoDiamond *SoDiamondTransactor) InitMultiChain(opts *bind.TransactOpts, fastRouter common.Address, chainId uint64) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "initMultiChain", fastRouter, chainId)
}

// InitMultiChain is a paid mutator transaction binding the contract method 0xa1c74577.
//
// Solidity: function initMultiChain(address fastRouter, uint64 chainId) returns()
func (_SoDiamond *SoDiamondSession) InitMultiChain(fastRouter common.Address, chainId uint64) (*types.Transaction, error) {
	return _SoDiamond.Contract.InitMultiChain(&_SoDiamond.TransactOpts, fastRouter, chainId)
}

// InitMultiChain is a paid mutator transaction binding the contract method 0xa1c74577.
//
// Solidity: function initMultiChain(address fastRouter, uint64 chainId) returns()
func (_SoDiamond *SoDiamondTransactorSession) InitMultiChain(fastRouter common.Address, chainId uint64) (*types.Transaction, error) {
	return _SoDiamond.Contract.InitMultiChain(&_SoDiamond.TransactOpts, fastRouter, chainId)
}

// InitStargate is a paid mutator transaction binding the contract method 0xaf66a6d8.
//
// Solidity: function initStargate(address stargate, uint16 chainId) returns()
func (_SoDiamond *SoDiamondTransactor) InitStargate(opts *bind.TransactOpts, stargate common.Address, chainId uint16) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "initStargate", stargate, chainId)
}

// InitStargate is a paid mutator transaction binding the contract method 0xaf66a6d8.
//
// Solidity: function initStargate(address stargate, uint16 chainId) returns()
func (_SoDiamond *SoDiamondSession) InitStargate(stargate common.Address, chainId uint16) (*types.Transaction, error) {
	return _SoDiamond.Contract.InitStargate(&_SoDiamond.TransactOpts, stargate, chainId)
}

// InitStargate is a paid mutator transaction binding the contract method 0xaf66a6d8.
//
// Solidity: function initStargate(address stargate, uint16 chainId) returns()
func (_SoDiamond *SoDiamondTransactorSession) InitStargate(stargate common.Address, chainId uint16) (*types.Transaction, error) {
	return _SoDiamond.Contract.InitStargate(&_SoDiamond.TransactOpts, stargate, chainId)
}

// InitWormhole is a paid mutator transaction binding the contract method 0xd316cefe.
//
// Solidity: function initWormhole(address tokenBridge, uint16 wormholeChainId) returns()
func (_SoDiamond *SoDiamondTransactor) InitWormhole(opts *bind.TransactOpts, tokenBridge common.Address, wormholeChainId uint16) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "initWormhole", tokenBridge, wormholeChainId)
}

// InitWormhole is a paid mutator transaction binding the contract method 0xd316cefe.
//
// Solidity: function initWormhole(address tokenBridge, uint16 wormholeChainId) returns()
func (_SoDiamond *SoDiamondSession) InitWormhole(tokenBridge common.Address, wormholeChainId uint16) (*types.Transaction, error) {
	return _SoDiamond.Contract.InitWormhole(&_SoDiamond.TransactOpts, tokenBridge, wormholeChainId)
}

// InitWormhole is a paid mutator transaction binding the contract method 0xd316cefe.
//
// Solidity: function initWormhole(address tokenBridge, uint16 wormholeChainId) returns()
func (_SoDiamond *SoDiamondTransactorSession) InitWormhole(tokenBridge common.Address, wormholeChainId uint16) (*types.Transaction, error) {
	return _SoDiamond.Contract.InitWormhole(&_SoDiamond.TransactOpts, tokenBridge, wormholeChainId)
}

// ReceiveFromBoolSwap is a paid mutator transaction binding the contract method 0x4e7fbc7b.
//
// Solidity: function receiveFromBoolSwap(uint32 , address bridgeToken, uint256 bridgeAmount, bytes payload) payable returns()
func (_SoDiamond *SoDiamondTransactor) ReceiveFromBoolSwap(opts *bind.TransactOpts, arg0 uint32, bridgeToken common.Address, bridgeAmount *big.Int, payload []byte) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "receiveFromBoolSwap", arg0, bridgeToken, bridgeAmount, payload)
}

// ReceiveFromBoolSwap is a paid mutator transaction binding the contract method 0x4e7fbc7b.
//
// Solidity: function receiveFromBoolSwap(uint32 , address bridgeToken, uint256 bridgeAmount, bytes payload) payable returns()
func (_SoDiamond *SoDiamondSession) ReceiveFromBoolSwap(arg0 uint32, bridgeToken common.Address, bridgeAmount *big.Int, payload []byte) (*types.Transaction, error) {
	return _SoDiamond.Contract.ReceiveFromBoolSwap(&_SoDiamond.TransactOpts, arg0, bridgeToken, bridgeAmount, payload)
}

// ReceiveFromBoolSwap is a paid mutator transaction binding the contract method 0x4e7fbc7b.
//
// Solidity: function receiveFromBoolSwap(uint32 , address bridgeToken, uint256 bridgeAmount, bytes payload) payable returns()
func (_SoDiamond *SoDiamondTransactorSession) ReceiveFromBoolSwap(arg0 uint32, bridgeToken common.Address, bridgeAmount *big.Int, payload []byte) (*types.Transaction, error) {
	return _SoDiamond.Contract.ReceiveFromBoolSwap(&_SoDiamond.TransactOpts, arg0, bridgeToken, bridgeAmount, payload)
}

// RemoteBoolSoSwap is a paid mutator transaction binding the contract method 0x83cb2579.
//
// Solidity: function remoteBoolSoSwap(address token, uint256 amount, (bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapDataDst) returns()
func (_SoDiamond *SoDiamondTransactor) RemoteBoolSoSwap(opts *bind.TransactOpts, token common.Address, amount *big.Int, soData ISoSoData, swapDataDst []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "remoteBoolSoSwap", token, amount, soData, swapDataDst)
}

// RemoteBoolSoSwap is a paid mutator transaction binding the contract method 0x83cb2579.
//
// Solidity: function remoteBoolSoSwap(address token, uint256 amount, (bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapDataDst) returns()
func (_SoDiamond *SoDiamondSession) RemoteBoolSoSwap(token common.Address, amount *big.Int, soData ISoSoData, swapDataDst []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.RemoteBoolSoSwap(&_SoDiamond.TransactOpts, token, amount, soData, swapDataDst)
}

// RemoteBoolSoSwap is a paid mutator transaction binding the contract method 0x83cb2579.
//
// Solidity: function remoteBoolSoSwap(address token, uint256 amount, (bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapDataDst) returns()
func (_SoDiamond *SoDiamondTransactorSession) RemoteBoolSoSwap(token common.Address, amount *big.Int, soData ISoSoData, swapDataDst []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.RemoteBoolSoSwap(&_SoDiamond.TransactOpts, token, amount, soData, swapDataDst)
}

// RemoteSoSwap is a paid mutator transaction binding the contract method 0x31b5d474.
//
// Solidity: function remoteSoSwap(address token, uint256 amount, (bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapDataDst) returns()
func (_SoDiamond *SoDiamondTransactor) RemoteSoSwap(opts *bind.TransactOpts, token common.Address, amount *big.Int, soData ISoSoData, swapDataDst []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "remoteSoSwap", token, amount, soData, swapDataDst)
}

// RemoteSoSwap is a paid mutator transaction binding the contract method 0x31b5d474.
//
// Solidity: function remoteSoSwap(address token, uint256 amount, (bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapDataDst) returns()
func (_SoDiamond *SoDiamondSession) RemoteSoSwap(token common.Address, amount *big.Int, soData ISoSoData, swapDataDst []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.RemoteSoSwap(&_SoDiamond.TransactOpts, token, amount, soData, swapDataDst)
}

// RemoteSoSwap is a paid mutator transaction binding the contract method 0x31b5d474.
//
// Solidity: function remoteSoSwap(address token, uint256 amount, (bytes32,address,uint16,address,uint16,address,uint256) soData, (address,address,address,address,uint256,bytes)[] swapDataDst) returns()
func (_SoDiamond *SoDiamondTransactorSession) RemoteSoSwap(token common.Address, amount *big.Int, soData ISoSoData, swapDataDst []LibSwapSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.RemoteSoSwap(&_SoDiamond.TransactOpts, token, amount, soData, swapDataDst)
}

// RemoveDex is a paid mutator transaction binding the contract method 0x124f1ead.
//
// Solidity: function removeDex(address dex) returns()
func (_SoDiamond *SoDiamondTransactor) RemoveDex(opts *bind.TransactOpts, dex common.Address) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "removeDex", dex)
}

// RemoveDex is a paid mutator transaction binding the contract method 0x124f1ead.
//
// Solidity: function removeDex(address dex) returns()
func (_SoDiamond *SoDiamondSession) RemoveDex(dex common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.RemoveDex(&_SoDiamond.TransactOpts, dex)
}

// RemoveDex is a paid mutator transaction binding the contract method 0x124f1ead.
//
// Solidity: function removeDex(address dex) returns()
func (_SoDiamond *SoDiamondTransactorSession) RemoveDex(dex common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.RemoveDex(&_SoDiamond.TransactOpts, dex)
}

// SetAllowedAddress is a paid mutator transaction binding the contract method 0xfae36986.
//
// Solidity: function setAllowedAddress(address router, bool isAllowed) returns()
func (_SoDiamond *SoDiamondTransactor) SetAllowedAddress(opts *bind.TransactOpts, router common.Address, isAllowed bool) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "setAllowedAddress", router, isAllowed)
}

// SetAllowedAddress is a paid mutator transaction binding the contract method 0xfae36986.
//
// Solidity: function setAllowedAddress(address router, bool isAllowed) returns()
func (_SoDiamond *SoDiamondSession) SetAllowedAddress(router common.Address, isAllowed bool) (*types.Transaction, error) {
	return _SoDiamond.Contract.SetAllowedAddress(&_SoDiamond.TransactOpts, router, isAllowed)
}

// SetAllowedAddress is a paid mutator transaction binding the contract method 0xfae36986.
//
// Solidity: function setAllowedAddress(address router, bool isAllowed) returns()
func (_SoDiamond *SoDiamondTransactorSession) SetAllowedAddress(router common.Address, isAllowed bool) (*types.Transaction, error) {
	return _SoDiamond.Contract.SetAllowedAddress(&_SoDiamond.TransactOpts, router, isAllowed)
}

// SetAllowedAddress0 is a paid mutator transaction binding the contract method 0xfae36986.
//
// Solidity: function setAllowedAddress(address messageBus, bool isAllowed) returns()
func (_SoDiamond *SoDiamondTransactor) SetAllowedAddress0(opts *bind.TransactOpts, messageBus common.Address, isAllowed bool) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "setAllowedAddress0", messageBus, isAllowed)
}

// SetAllowedAddress0 is a paid mutator transaction binding the contract method 0xfae36986.
//
// Solidity: function setAllowedAddress(address messageBus, bool isAllowed) returns()
func (_SoDiamond *SoDiamondSession) SetAllowedAddress0(messageBus common.Address, isAllowed bool) (*types.Transaction, error) {
	return _SoDiamond.Contract.SetAllowedAddress0(&_SoDiamond.TransactOpts, messageBus, isAllowed)
}

// SetAllowedAddress0 is a paid mutator transaction binding the contract method 0xfae36986.
//
// Solidity: function setAllowedAddress(address messageBus, bool isAllowed) returns()
func (_SoDiamond *SoDiamondTransactorSession) SetAllowedAddress0(messageBus common.Address, isAllowed bool) (*types.Transaction, error) {
	return _SoDiamond.Contract.SetAllowedAddress0(&_SoDiamond.TransactOpts, messageBus, isAllowed)
}

// SetAllowedAddress1 is a paid mutator transaction binding the contract method 0xfae36986.
//
// Solidity: function setAllowedAddress(address account, bool isAllowed) returns()
func (_SoDiamond *SoDiamondTransactor) SetAllowedAddress1(opts *bind.TransactOpts, account common.Address, isAllowed bool) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "setAllowedAddress1", account, isAllowed)
}

// SetAllowedAddress1 is a paid mutator transaction binding the contract method 0xfae36986.
//
// Solidity: function setAllowedAddress(address account, bool isAllowed) returns()
func (_SoDiamond *SoDiamondSession) SetAllowedAddress1(account common.Address, isAllowed bool) (*types.Transaction, error) {
	return _SoDiamond.Contract.SetAllowedAddress1(&_SoDiamond.TransactOpts, account, isAllowed)
}

// SetAllowedAddress1 is a paid mutator transaction binding the contract method 0xfae36986.
//
// Solidity: function setAllowedAddress(address account, bool isAllowed) returns()
func (_SoDiamond *SoDiamondTransactorSession) SetAllowedAddress1(account common.Address, isAllowed bool) (*types.Transaction, error) {
	return _SoDiamond.Contract.SetAllowedAddress1(&_SoDiamond.TransactOpts, account, isAllowed)
}

// SetBaseGas is a paid mutator transaction binding the contract method 0xce047adb.
//
// Solidity: function setBaseGas(uint64[] dstChainIds, uint256 dstBaseGas) returns()
func (_SoDiamond *SoDiamondTransactor) SetBaseGas(opts *bind.TransactOpts, dstChainIds []uint64, dstBaseGas *big.Int) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "setBaseGas", dstChainIds, dstBaseGas)
}

// SetBaseGas is a paid mutator transaction binding the contract method 0xce047adb.
//
// Solidity: function setBaseGas(uint64[] dstChainIds, uint256 dstBaseGas) returns()
func (_SoDiamond *SoDiamondSession) SetBaseGas(dstChainIds []uint64, dstBaseGas *big.Int) (*types.Transaction, error) {
	return _SoDiamond.Contract.SetBaseGas(&_SoDiamond.TransactOpts, dstChainIds, dstBaseGas)
}

// SetBaseGas is a paid mutator transaction binding the contract method 0xce047adb.
//
// Solidity: function setBaseGas(uint64[] dstChainIds, uint256 dstBaseGas) returns()
func (_SoDiamond *SoDiamondTransactorSession) SetBaseGas(dstChainIds []uint64, dstBaseGas *big.Int) (*types.Transaction, error) {
	return _SoDiamond.Contract.SetBaseGas(&_SoDiamond.TransactOpts, dstChainIds, dstBaseGas)
}

// SetCelerReserve is a paid mutator transaction binding the contract method 0xa452def2.
//
// Solidity: function setCelerReserve(uint256 actualReserve, uint256 estimateReserve) returns()
func (_SoDiamond *SoDiamondTransactor) SetCelerReserve(opts *bind.TransactOpts, actualReserve *big.Int, estimateReserve *big.Int) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "setCelerReserve", actualReserve, estimateReserve)
}

// SetCelerReserve is a paid mutator transaction binding the contract method 0xa452def2.
//
// Solidity: function setCelerReserve(uint256 actualReserve, uint256 estimateReserve) returns()
func (_SoDiamond *SoDiamondSession) SetCelerReserve(actualReserve *big.Int, estimateReserve *big.Int) (*types.Transaction, error) {
	return _SoDiamond.Contract.SetCelerReserve(&_SoDiamond.TransactOpts, actualReserve, estimateReserve)
}

// SetCelerReserve is a paid mutator transaction binding the contract method 0xa452def2.
//
// Solidity: function setCelerReserve(uint256 actualReserve, uint256 estimateReserve) returns()
func (_SoDiamond *SoDiamondTransactorSession) SetCelerReserve(actualReserve *big.Int, estimateReserve *big.Int) (*types.Transaction, error) {
	return _SoDiamond.Contract.SetCelerReserve(&_SoDiamond.TransactOpts, actualReserve, estimateReserve)
}

// SetExecutorFeeTo is a paid mutator transaction binding the contract method 0xdcf078d0.
//
// Solidity: function setExecutorFeeTo(address feeTo) returns()
func (_SoDiamond *SoDiamondTransactor) SetExecutorFeeTo(opts *bind.TransactOpts, feeTo common.Address) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "setExecutorFeeTo", feeTo)
}

// SetExecutorFeeTo is a paid mutator transaction binding the contract method 0xdcf078d0.
//
// Solidity: function setExecutorFeeTo(address feeTo) returns()
func (_SoDiamond *SoDiamondSession) SetExecutorFeeTo(feeTo common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.SetExecutorFeeTo(&_SoDiamond.TransactOpts, feeTo)
}

// SetExecutorFeeTo is a paid mutator transaction binding the contract method 0xdcf078d0.
//
// Solidity: function setExecutorFeeTo(address feeTo) returns()
func (_SoDiamond *SoDiamondTransactorSession) SetExecutorFeeTo(feeTo common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.SetExecutorFeeTo(&_SoDiamond.TransactOpts, feeTo)
}

// SetFunctionApprovalBySignature is a paid mutator transaction binding the contract method 0xc6bd579d.
//
// Solidity: function setFunctionApprovalBySignature(bytes32 signature, bool approval) returns()
func (_SoDiamond *SoDiamondTransactor) SetFunctionApprovalBySignature(opts *bind.TransactOpts, signature [32]byte, approval bool) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "setFunctionApprovalBySignature", signature, approval)
}

// SetFunctionApprovalBySignature is a paid mutator transaction binding the contract method 0xc6bd579d.
//
// Solidity: function setFunctionApprovalBySignature(bytes32 signature, bool approval) returns()
func (_SoDiamond *SoDiamondSession) SetFunctionApprovalBySignature(signature [32]byte, approval bool) (*types.Transaction, error) {
	return _SoDiamond.Contract.SetFunctionApprovalBySignature(&_SoDiamond.TransactOpts, signature, approval)
}

// SetFunctionApprovalBySignature is a paid mutator transaction binding the contract method 0xc6bd579d.
//
// Solidity: function setFunctionApprovalBySignature(bytes32 signature, bool approval) returns()
func (_SoDiamond *SoDiamondTransactorSession) SetFunctionApprovalBySignature(signature [32]byte, approval bool) (*types.Transaction, error) {
	return _SoDiamond.Contract.SetFunctionApprovalBySignature(&_SoDiamond.TransactOpts, signature, approval)
}

// SetNonce is a paid mutator transaction binding the contract method 0x6cd4a685.
//
// Solidity: function setNonce(uint64 nonce) returns()
func (_SoDiamond *SoDiamondTransactor) SetNonce(opts *bind.TransactOpts, nonce uint64) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "setNonce", nonce)
}

// SetNonce is a paid mutator transaction binding the contract method 0x6cd4a685.
//
// Solidity: function setNonce(uint64 nonce) returns()
func (_SoDiamond *SoDiamondSession) SetNonce(nonce uint64) (*types.Transaction, error) {
	return _SoDiamond.Contract.SetNonce(&_SoDiamond.TransactOpts, nonce)
}

// SetNonce is a paid mutator transaction binding the contract method 0x6cd4a685.
//
// Solidity: function setNonce(uint64 nonce) returns()
func (_SoDiamond *SoDiamondTransactorSession) SetNonce(nonce uint64) (*types.Transaction, error) {
	return _SoDiamond.Contract.SetNonce(&_SoDiamond.TransactOpts, nonce)
}

// SetWormholeGas is a paid mutator transaction binding the contract method 0x3620a2b8.
//
// Solidity: function setWormholeGas(uint16 dstWormholeChainId, uint256 baseGas, uint256 gasPerBytes) returns()
func (_SoDiamond *SoDiamondTransactor) SetWormholeGas(opts *bind.TransactOpts, dstWormholeChainId uint16, baseGas *big.Int, gasPerBytes *big.Int) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "setWormholeGas", dstWormholeChainId, baseGas, gasPerBytes)
}

// SetWormholeGas is a paid mutator transaction binding the contract method 0x3620a2b8.
//
// Solidity: function setWormholeGas(uint16 dstWormholeChainId, uint256 baseGas, uint256 gasPerBytes) returns()
func (_SoDiamond *SoDiamondSession) SetWormholeGas(dstWormholeChainId uint16, baseGas *big.Int, gasPerBytes *big.Int) (*types.Transaction, error) {
	return _SoDiamond.Contract.SetWormholeGas(&_SoDiamond.TransactOpts, dstWormholeChainId, baseGas, gasPerBytes)
}

// SetWormholeGas is a paid mutator transaction binding the contract method 0x3620a2b8.
//
// Solidity: function setWormholeGas(uint16 dstWormholeChainId, uint256 baseGas, uint256 gasPerBytes) returns()
func (_SoDiamond *SoDiamondTransactorSession) SetWormholeGas(dstWormholeChainId uint16, baseGas *big.Int, gasPerBytes *big.Int) (*types.Transaction, error) {
	return _SoDiamond.Contract.SetWormholeGas(&_SoDiamond.TransactOpts, dstWormholeChainId, baseGas, gasPerBytes)
}

// SetWormholeReserve is a paid mutator transaction binding the contract method 0xae6ab3e5.
//
// Solidity: function setWormholeReserve(uint256 actualReserve, uint256 estimateReserve) returns()
func (_SoDiamond *SoDiamondTransactor) SetWormholeReserve(opts *bind.TransactOpts, actualReserve *big.Int, estimateReserve *big.Int) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "setWormholeReserve", actualReserve, estimateReserve)
}

// SetWormholeReserve is a paid mutator transaction binding the contract method 0xae6ab3e5.
//
// Solidity: function setWormholeReserve(uint256 actualReserve, uint256 estimateReserve) returns()
func (_SoDiamond *SoDiamondSession) SetWormholeReserve(actualReserve *big.Int, estimateReserve *big.Int) (*types.Transaction, error) {
	return _SoDiamond.Contract.SetWormholeReserve(&_SoDiamond.TransactOpts, actualReserve, estimateReserve)
}

// SetWormholeReserve is a paid mutator transaction binding the contract method 0xae6ab3e5.
//
// Solidity: function setWormholeReserve(uint256 actualReserve, uint256 estimateReserve) returns()
func (_SoDiamond *SoDiamondTransactorSession) SetWormholeReserve(actualReserve *big.Int, estimateReserve *big.Int) (*types.Transaction, error) {
	return _SoDiamond.Contract.SetWormholeReserve(&_SoDiamond.TransactOpts, actualReserve, estimateReserve)
}

// SgReceive is a paid mutator transaction binding the contract method 0xab8236f3.
//
// Solidity: function sgReceive(uint16 , bytes , uint256 , address token, uint256 amount, bytes payload) returns()
func (_SoDiamond *SoDiamondTransactor) SgReceive(opts *bind.TransactOpts, arg0 uint16, arg1 []byte, arg2 *big.Int, token common.Address, amount *big.Int, payload []byte) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "sgReceive", arg0, arg1, arg2, token, amount, payload)
}

// SgReceive is a paid mutator transaction binding the contract method 0xab8236f3.
//
// Solidity: function sgReceive(uint16 , bytes , uint256 , address token, uint256 amount, bytes payload) returns()
func (_SoDiamond *SoDiamondSession) SgReceive(arg0 uint16, arg1 []byte, arg2 *big.Int, token common.Address, amount *big.Int, payload []byte) (*types.Transaction, error) {
	return _SoDiamond.Contract.SgReceive(&_SoDiamond.TransactOpts, arg0, arg1, arg2, token, amount, payload)
}

// SgReceive is a paid mutator transaction binding the contract method 0xab8236f3.
//
// Solidity: function sgReceive(uint16 , bytes , uint256 , address token, uint256 amount, bytes payload) returns()
func (_SoDiamond *SoDiamondTransactorSession) SgReceive(arg0 uint16, arg1 []byte, arg2 *big.Int, token common.Address, amount *big.Int, payload []byte) (*types.Transaction, error) {
	return _SoDiamond.Contract.SgReceive(&_SoDiamond.TransactOpts, arg0, arg1, arg2, token, amount, payload)
}

// SgReceiveForGas is a paid mutator transaction binding the contract method 0x958b3ab0.
//
// Solidity: function sgReceiveForGas((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, uint256 dstStargatePoolId, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) returns()
func (_SoDiamond *SoDiamondTransactor) SgReceiveForGas(opts *bind.TransactOpts, soDataNo ISoNormalizedSoData, dstStargatePoolId *big.Int, swapDataDstNo []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "sgReceiveForGas", soDataNo, dstStargatePoolId, swapDataDstNo)
}

// SgReceiveForGas is a paid mutator transaction binding the contract method 0x958b3ab0.
//
// Solidity: function sgReceiveForGas((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, uint256 dstStargatePoolId, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) returns()
func (_SoDiamond *SoDiamondSession) SgReceiveForGas(soDataNo ISoNormalizedSoData, dstStargatePoolId *big.Int, swapDataDstNo []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.SgReceiveForGas(&_SoDiamond.TransactOpts, soDataNo, dstStargatePoolId, swapDataDstNo)
}

// SgReceiveForGas is a paid mutator transaction binding the contract method 0x958b3ab0.
//
// Solidity: function sgReceiveForGas((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, uint256 dstStargatePoolId, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) returns()
func (_SoDiamond *SoDiamondTransactorSession) SgReceiveForGas(soDataNo ISoNormalizedSoData, dstStargatePoolId *big.Int, swapDataDstNo []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.SgReceiveForGas(&_SoDiamond.TransactOpts, soDataNo, dstStargatePoolId, swapDataDstNo)
}

// SoSwapViaBool is a paid mutator transaction binding the contract method 0xbe113761.
//
// Solidity: function soSwapViaBool((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataSrcNo, (uint32,uint32,bytes32) boolSwapData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) payable returns()
func (_SoDiamond *SoDiamondTransactor) SoSwapViaBool(opts *bind.TransactOpts, soDataNo ISoNormalizedSoData, swapDataSrcNo []LibSwapNormalizedSwapData, boolSwapData BoolFacetBoolSwapData, swapDataDstNo []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "soSwapViaBool", soDataNo, swapDataSrcNo, boolSwapData, swapDataDstNo)
}

// SoSwapViaBool is a paid mutator transaction binding the contract method 0xbe113761.
//
// Solidity: function soSwapViaBool((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataSrcNo, (uint32,uint32,bytes32) boolSwapData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) payable returns()
func (_SoDiamond *SoDiamondSession) SoSwapViaBool(soDataNo ISoNormalizedSoData, swapDataSrcNo []LibSwapNormalizedSwapData, boolSwapData BoolFacetBoolSwapData, swapDataDstNo []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.SoSwapViaBool(&_SoDiamond.TransactOpts, soDataNo, swapDataSrcNo, boolSwapData, swapDataDstNo)
}

// SoSwapViaBool is a paid mutator transaction binding the contract method 0xbe113761.
//
// Solidity: function soSwapViaBool((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataSrcNo, (uint32,uint32,bytes32) boolSwapData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) payable returns()
func (_SoDiamond *SoDiamondTransactorSession) SoSwapViaBool(soDataNo ISoNormalizedSoData, swapDataSrcNo []LibSwapNormalizedSwapData, boolSwapData BoolFacetBoolSwapData, swapDataDstNo []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.SoSwapViaBool(&_SoDiamond.TransactOpts, soDataNo, swapDataSrcNo, boolSwapData, swapDataDstNo)
}

// SoSwapViaCeler is a paid mutator transaction binding the contract method 0x56633e3e.
//
// Solidity: function soSwapViaCeler((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataSrcNo, (address,uint32,uint64,address,uint256,uint256,address) celerData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) payable returns()
func (_SoDiamond *SoDiamondTransactor) SoSwapViaCeler(opts *bind.TransactOpts, soDataNo ISoNormalizedSoData, swapDataSrcNo []LibSwapNormalizedSwapData, celerData CelerFacetCelerData, swapDataDstNo []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "soSwapViaCeler", soDataNo, swapDataSrcNo, celerData, swapDataDstNo)
}

// SoSwapViaCeler is a paid mutator transaction binding the contract method 0x56633e3e.
//
// Solidity: function soSwapViaCeler((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataSrcNo, (address,uint32,uint64,address,uint256,uint256,address) celerData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) payable returns()
func (_SoDiamond *SoDiamondSession) SoSwapViaCeler(soDataNo ISoNormalizedSoData, swapDataSrcNo []LibSwapNormalizedSwapData, celerData CelerFacetCelerData, swapDataDstNo []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.SoSwapViaCeler(&_SoDiamond.TransactOpts, soDataNo, swapDataSrcNo, celerData, swapDataDstNo)
}

// SoSwapViaCeler is a paid mutator transaction binding the contract method 0x56633e3e.
//
// Solidity: function soSwapViaCeler((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataSrcNo, (address,uint32,uint64,address,uint256,uint256,address) celerData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) payable returns()
func (_SoDiamond *SoDiamondTransactorSession) SoSwapViaCeler(soDataNo ISoNormalizedSoData, swapDataSrcNo []LibSwapNormalizedSwapData, celerData CelerFacetCelerData, swapDataDstNo []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.SoSwapViaCeler(&_SoDiamond.TransactOpts, soDataNo, swapDataSrcNo, celerData, swapDataDstNo)
}

// SoSwapViaMultiChain is a paid mutator transaction binding the contract method 0xac1f8185.
//
// Solidity: function soSwapViaMultiChain((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataSrcNo, (uint64,address,string) multiChainData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) payable returns()
func (_SoDiamond *SoDiamondTransactor) SoSwapViaMultiChain(opts *bind.TransactOpts, soDataNo ISoNormalizedSoData, swapDataSrcNo []LibSwapNormalizedSwapData, multiChainData MultiChainFacetMultiChainData, swapDataDstNo []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "soSwapViaMultiChain", soDataNo, swapDataSrcNo, multiChainData, swapDataDstNo)
}

// SoSwapViaMultiChain is a paid mutator transaction binding the contract method 0xac1f8185.
//
// Solidity: function soSwapViaMultiChain((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataSrcNo, (uint64,address,string) multiChainData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) payable returns()
func (_SoDiamond *SoDiamondSession) SoSwapViaMultiChain(soDataNo ISoNormalizedSoData, swapDataSrcNo []LibSwapNormalizedSwapData, multiChainData MultiChainFacetMultiChainData, swapDataDstNo []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.SoSwapViaMultiChain(&_SoDiamond.TransactOpts, soDataNo, swapDataSrcNo, multiChainData, swapDataDstNo)
}

// SoSwapViaMultiChain is a paid mutator transaction binding the contract method 0xac1f8185.
//
// Solidity: function soSwapViaMultiChain((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataSrcNo, (uint64,address,string) multiChainData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) payable returns()
func (_SoDiamond *SoDiamondTransactorSession) SoSwapViaMultiChain(soDataNo ISoNormalizedSoData, swapDataSrcNo []LibSwapNormalizedSwapData, multiChainData MultiChainFacetMultiChainData, swapDataDstNo []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.SoSwapViaMultiChain(&_SoDiamond.TransactOpts, soDataNo, swapDataSrcNo, multiChainData, swapDataDstNo)
}

// SoSwapViaStargate is a paid mutator transaction binding the contract method 0xeedfa7cd.
//
// Solidity: function soSwapViaStargate((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataSrcNo, (uint256,uint16,uint256,uint256,uint256,address) stargateData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) payable returns()
func (_SoDiamond *SoDiamondTransactor) SoSwapViaStargate(opts *bind.TransactOpts, soDataNo ISoNormalizedSoData, swapDataSrcNo []LibSwapNormalizedSwapData, stargateData StargateFacetStargateData, swapDataDstNo []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "soSwapViaStargate", soDataNo, swapDataSrcNo, stargateData, swapDataDstNo)
}

// SoSwapViaStargate is a paid mutator transaction binding the contract method 0xeedfa7cd.
//
// Solidity: function soSwapViaStargate((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataSrcNo, (uint256,uint16,uint256,uint256,uint256,address) stargateData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) payable returns()
func (_SoDiamond *SoDiamondSession) SoSwapViaStargate(soDataNo ISoNormalizedSoData, swapDataSrcNo []LibSwapNormalizedSwapData, stargateData StargateFacetStargateData, swapDataDstNo []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.SoSwapViaStargate(&_SoDiamond.TransactOpts, soDataNo, swapDataSrcNo, stargateData, swapDataDstNo)
}

// SoSwapViaStargate is a paid mutator transaction binding the contract method 0xeedfa7cd.
//
// Solidity: function soSwapViaStargate((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataSrcNo, (uint256,uint16,uint256,uint256,uint256,address) stargateData, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) payable returns()
func (_SoDiamond *SoDiamondTransactorSession) SoSwapViaStargate(soDataNo ISoNormalizedSoData, swapDataSrcNo []LibSwapNormalizedSwapData, stargateData StargateFacetStargateData, swapDataDstNo []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.SoSwapViaStargate(&_SoDiamond.TransactOpts, soDataNo, swapDataSrcNo, stargateData, swapDataDstNo)
}

// SoSwapViaWormhole is a paid mutator transaction binding the contract method 0x1c4469a9.
//
// Solidity: function soSwapViaWormhole((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataSrcNo, (uint16,uint256,uint256,bytes) wormholeDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) payable returns()
func (_SoDiamond *SoDiamondTransactor) SoSwapViaWormhole(opts *bind.TransactOpts, soDataNo ISoNormalizedSoData, swapDataSrcNo []LibSwapNormalizedSwapData, wormholeDataNo WormholeFacetNormalizedWormholeData, swapDataDstNo []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "soSwapViaWormhole", soDataNo, swapDataSrcNo, wormholeDataNo, swapDataDstNo)
}

// SoSwapViaWormhole is a paid mutator transaction binding the contract method 0x1c4469a9.
//
// Solidity: function soSwapViaWormhole((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataSrcNo, (uint16,uint256,uint256,bytes) wormholeDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) payable returns()
func (_SoDiamond *SoDiamondSession) SoSwapViaWormhole(soDataNo ISoNormalizedSoData, swapDataSrcNo []LibSwapNormalizedSwapData, wormholeDataNo WormholeFacetNormalizedWormholeData, swapDataDstNo []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.SoSwapViaWormhole(&_SoDiamond.TransactOpts, soDataNo, swapDataSrcNo, wormholeDataNo, swapDataDstNo)
}

// SoSwapViaWormhole is a paid mutator transaction binding the contract method 0x1c4469a9.
//
// Solidity: function soSwapViaWormhole((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataSrcNo, (uint16,uint256,uint256,bytes) wormholeDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataDstNo) payable returns()
func (_SoDiamond *SoDiamondTransactorSession) SoSwapViaWormhole(soDataNo ISoNormalizedSoData, swapDataSrcNo []LibSwapNormalizedSwapData, wormholeDataNo WormholeFacetNormalizedWormholeData, swapDataDstNo []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.SoSwapViaWormhole(&_SoDiamond.TransactOpts, soDataNo, swapDataSrcNo, wormholeDataNo, swapDataDstNo)
}

// SwapTokensGeneric is a paid mutator transaction binding the contract method 0xc953c76a.
//
// Solidity: function swapTokensGeneric((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataNo) payable returns()
func (_SoDiamond *SoDiamondTransactor) SwapTokensGeneric(opts *bind.TransactOpts, soDataNo ISoNormalizedSoData, swapDataNo []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "swapTokensGeneric", soDataNo, swapDataNo)
}

// SwapTokensGeneric is a paid mutator transaction binding the contract method 0xc953c76a.
//
// Solidity: function swapTokensGeneric((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataNo) payable returns()
func (_SoDiamond *SoDiamondSession) SwapTokensGeneric(soDataNo ISoNormalizedSoData, swapDataNo []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.SwapTokensGeneric(&_SoDiamond.TransactOpts, soDataNo, swapDataNo)
}

// SwapTokensGeneric is a paid mutator transaction binding the contract method 0xc953c76a.
//
// Solidity: function swapTokensGeneric((bytes,bytes,uint16,bytes,uint16,bytes,uint256) soDataNo, (bytes,bytes,bytes,bytes,uint256,bytes)[] swapDataNo) payable returns()
func (_SoDiamond *SoDiamondTransactorSession) SwapTokensGeneric(soDataNo ISoNormalizedSoData, swapDataNo []LibSwapNormalizedSwapData) (*types.Transaction, error) {
	return _SoDiamond.Contract.SwapTokensGeneric(&_SoDiamond.TransactOpts, soDataNo, swapDataNo)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SoDiamond *SoDiamondTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SoDiamond *SoDiamondSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.TransferOwnership(&_SoDiamond.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SoDiamond *SoDiamondTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.TransferOwnership(&_SoDiamond.TransactOpts, newOwner)
}

// UpdateAddressMappings is a paid mutator transaction binding the contract method 0x85c8496b.
//
// Solidity: function updateAddressMappings(address[] anyTokens) returns()
func (_SoDiamond *SoDiamondTransactor) UpdateAddressMappings(opts *bind.TransactOpts, anyTokens []common.Address) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "updateAddressMappings", anyTokens)
}

// UpdateAddressMappings is a paid mutator transaction binding the contract method 0x85c8496b.
//
// Solidity: function updateAddressMappings(address[] anyTokens) returns()
func (_SoDiamond *SoDiamondSession) UpdateAddressMappings(anyTokens []common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.UpdateAddressMappings(&_SoDiamond.TransactOpts, anyTokens)
}

// UpdateAddressMappings is a paid mutator transaction binding the contract method 0x85c8496b.
//
// Solidity: function updateAddressMappings(address[] anyTokens) returns()
func (_SoDiamond *SoDiamondTransactorSession) UpdateAddressMappings(anyTokens []common.Address) (*types.Transaction, error) {
	return _SoDiamond.Contract.UpdateAddressMappings(&_SoDiamond.TransactOpts, anyTokens)
}

// UpdatePathPair is a paid mutator transaction binding the contract method 0x4dbb0636.
//
// Solidity: function updatePathPair(uint32[] consumerChainIds, uint32[] boolChainIds) returns()
func (_SoDiamond *SoDiamondTransactor) UpdatePathPair(opts *bind.TransactOpts, consumerChainIds []uint32, boolChainIds []uint32) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "updatePathPair", consumerChainIds, boolChainIds)
}

// UpdatePathPair is a paid mutator transaction binding the contract method 0x4dbb0636.
//
// Solidity: function updatePathPair(uint32[] consumerChainIds, uint32[] boolChainIds) returns()
func (_SoDiamond *SoDiamondSession) UpdatePathPair(consumerChainIds []uint32, boolChainIds []uint32) (*types.Transaction, error) {
	return _SoDiamond.Contract.UpdatePathPair(&_SoDiamond.TransactOpts, consumerChainIds, boolChainIds)
}

// UpdatePathPair is a paid mutator transaction binding the contract method 0x4dbb0636.
//
// Solidity: function updatePathPair(uint32[] consumerChainIds, uint32[] boolChainIds) returns()
func (_SoDiamond *SoDiamondTransactorSession) UpdatePathPair(consumerChainIds []uint32, boolChainIds []uint32) (*types.Transaction, error) {
	return _SoDiamond.Contract.UpdatePathPair(&_SoDiamond.TransactOpts, consumerChainIds, boolChainIds)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address assetAddress, address to, uint256 amount) returns()
func (_SoDiamond *SoDiamondTransactor) Withdraw(opts *bind.TransactOpts, assetAddress common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SoDiamond.contract.Transact(opts, "withdraw", assetAddress, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address assetAddress, address to, uint256 amount) returns()
func (_SoDiamond *SoDiamondSession) Withdraw(assetAddress common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SoDiamond.Contract.Withdraw(&_SoDiamond.TransactOpts, assetAddress, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address assetAddress, address to, uint256 amount) returns()
func (_SoDiamond *SoDiamondTransactorSession) Withdraw(assetAddress common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SoDiamond.Contract.Withdraw(&_SoDiamond.TransactOpts, assetAddress, to, amount)
}

// SoDiamondAnyMappingUpdatedIterator is returned from FilterAnyMappingUpdated and is used to iterate over the raw logs and unpacked data for AnyMappingUpdated events raised by the SoDiamond contract.
type SoDiamondAnyMappingUpdatedIterator struct {
	Event *SoDiamondAnyMappingUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondAnyMappingUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondAnyMappingUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondAnyMappingUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondAnyMappingUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondAnyMappingUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondAnyMappingUpdated represents a AnyMappingUpdated event raised by the SoDiamond contract.
type SoDiamondAnyMappingUpdated struct {
	AnyTokens []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAnyMappingUpdated is a free log retrieval operation binding the contract event 0xa35c90a7894ca381a96dcda7476b2470bf6493955b4864f023c05c8c8118a526.
//
// Solidity: event AnyMappingUpdated(address[] anyTokens)
func (_SoDiamond *SoDiamondFilterer) FilterAnyMappingUpdated(opts *bind.FilterOpts) (*SoDiamondAnyMappingUpdatedIterator, error) {

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "AnyMappingUpdated")
	if err != nil {
		return nil, err
	}
	return &SoDiamondAnyMappingUpdatedIterator{contract: _SoDiamond.contract, event: "AnyMappingUpdated", logs: logs, sub: sub}, nil
}

// WatchAnyMappingUpdated is a free log subscription operation binding the contract event 0xa35c90a7894ca381a96dcda7476b2470bf6493955b4864f023c05c8c8118a526.
//
// Solidity: event AnyMappingUpdated(address[] anyTokens)
func (_SoDiamond *SoDiamondFilterer) WatchAnyMappingUpdated(opts *bind.WatchOpts, sink chan<- *SoDiamondAnyMappingUpdated) (event.Subscription, error) {

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "AnyMappingUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondAnyMappingUpdated)
				if err := _SoDiamond.contract.UnpackLog(event, "AnyMappingUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAnyMappingUpdated is a log parse operation binding the contract event 0xa35c90a7894ca381a96dcda7476b2470bf6493955b4864f023c05c8c8118a526.
//
// Solidity: event AnyMappingUpdated(address[] anyTokens)
func (_SoDiamond *SoDiamondFilterer) ParseAnyMappingUpdated(log types.Log) (*SoDiamondAnyMappingUpdated, error) {
	event := new(SoDiamondAnyMappingUpdated)
	if err := _SoDiamond.contract.UnpackLog(event, "AnyMappingUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondAssetSwappedIterator is returned from FilterAssetSwapped and is used to iterate over the raw logs and unpacked data for AssetSwapped events raised by the SoDiamond contract.
type SoDiamondAssetSwappedIterator struct {
	Event *SoDiamondAssetSwapped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondAssetSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondAssetSwapped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondAssetSwapped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondAssetSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondAssetSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondAssetSwapped represents a AssetSwapped event raised by the SoDiamond contract.
type SoDiamondAssetSwapped struct {
	TransactionId [32]byte
	Dex           common.Address
	FromAssetId   common.Address
	ToAssetId     common.Address
	FromAmount    *big.Int
	ToAmount      *big.Int
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAssetSwapped is a free log retrieval operation binding the contract event 0x7bfdfdb5e3a3776976e53cb0607060f54c5312701c8cba1155cc4d5394440b38.
//
// Solidity: event AssetSwapped(bytes32 transactionId, address dex, address fromAssetId, address toAssetId, uint256 fromAmount, uint256 toAmount, uint256 timestamp)
func (_SoDiamond *SoDiamondFilterer) FilterAssetSwapped(opts *bind.FilterOpts) (*SoDiamondAssetSwappedIterator, error) {

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "AssetSwapped")
	if err != nil {
		return nil, err
	}
	return &SoDiamondAssetSwappedIterator{contract: _SoDiamond.contract, event: "AssetSwapped", logs: logs, sub: sub}, nil
}

// WatchAssetSwapped is a free log subscription operation binding the contract event 0x7bfdfdb5e3a3776976e53cb0607060f54c5312701c8cba1155cc4d5394440b38.
//
// Solidity: event AssetSwapped(bytes32 transactionId, address dex, address fromAssetId, address toAssetId, uint256 fromAmount, uint256 toAmount, uint256 timestamp)
func (_SoDiamond *SoDiamondFilterer) WatchAssetSwapped(opts *bind.WatchOpts, sink chan<- *SoDiamondAssetSwapped) (event.Subscription, error) {

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "AssetSwapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondAssetSwapped)
				if err := _SoDiamond.contract.UnpackLog(event, "AssetSwapped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetSwapped is a log parse operation binding the contract event 0x7bfdfdb5e3a3776976e53cb0607060f54c5312701c8cba1155cc4d5394440b38.
//
// Solidity: event AssetSwapped(bytes32 transactionId, address dex, address fromAssetId, address toAssetId, uint256 fromAmount, uint256 toAmount, uint256 timestamp)
func (_SoDiamond *SoDiamondFilterer) ParseAssetSwapped(log types.Log) (*SoDiamondAssetSwapped, error) {
	event := new(SoDiamondAssetSwapped)
	if err := _SoDiamond.contract.UnpackLog(event, "AssetSwapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondBoolSwapInitializedIterator is returned from FilterBoolSwapInitialized and is used to iterate over the raw logs and unpacked data for BoolSwapInitialized events raised by the SoDiamond contract.
type SoDiamondBoolSwapInitializedIterator struct {
	Event *SoDiamondBoolSwapInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondBoolSwapInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondBoolSwapInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondBoolSwapInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondBoolSwapInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondBoolSwapInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondBoolSwapInitialized represents a BoolSwapInitialized event raised by the SoDiamond contract.
type SoDiamondBoolSwapInitialized struct {
	BoolSwapRouter  common.Address
	BoolSwapFactory common.Address
	ChainId         uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBoolSwapInitialized is a free log retrieval operation binding the contract event 0x9576b1b0aa28e5440d98721d3d04803ae8767d60a24ee33ab76523d7dc8f6105.
//
// Solidity: event BoolSwapInitialized(address boolSwapRouter, address boolSwapFactory, uint32 chainId)
func (_SoDiamond *SoDiamondFilterer) FilterBoolSwapInitialized(opts *bind.FilterOpts) (*SoDiamondBoolSwapInitializedIterator, error) {

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "BoolSwapInitialized")
	if err != nil {
		return nil, err
	}
	return &SoDiamondBoolSwapInitializedIterator{contract: _SoDiamond.contract, event: "BoolSwapInitialized", logs: logs, sub: sub}, nil
}

// WatchBoolSwapInitialized is a free log subscription operation binding the contract event 0x9576b1b0aa28e5440d98721d3d04803ae8767d60a24ee33ab76523d7dc8f6105.
//
// Solidity: event BoolSwapInitialized(address boolSwapRouter, address boolSwapFactory, uint32 chainId)
func (_SoDiamond *SoDiamondFilterer) WatchBoolSwapInitialized(opts *bind.WatchOpts, sink chan<- *SoDiamondBoolSwapInitialized) (event.Subscription, error) {

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "BoolSwapInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondBoolSwapInitialized)
				if err := _SoDiamond.contract.UnpackLog(event, "BoolSwapInitialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBoolSwapInitialized is a log parse operation binding the contract event 0x9576b1b0aa28e5440d98721d3d04803ae8767d60a24ee33ab76523d7dc8f6105.
//
// Solidity: event BoolSwapInitialized(address boolSwapRouter, address boolSwapFactory, uint32 chainId)
func (_SoDiamond *SoDiamondFilterer) ParseBoolSwapInitialized(log types.Log) (*SoDiamondBoolSwapInitialized, error) {
	event := new(SoDiamondBoolSwapInitialized)
	if err := _SoDiamond.contract.UnpackLog(event, "BoolSwapInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondCelerInitializedIterator is returned from FilterCelerInitialized and is used to iterate over the raw logs and unpacked data for CelerInitialized events raised by the SoDiamond contract.
type SoDiamondCelerInitializedIterator struct {
	Event *SoDiamondCelerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondCelerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondCelerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondCelerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondCelerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondCelerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondCelerInitialized represents a CelerInitialized event raised by the SoDiamond contract.
type SoDiamondCelerInitialized struct {
	MessageBus common.Address
	ChainId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCelerInitialized is a free log retrieval operation binding the contract event 0x4fdd908f8304e87ab2bf41b4e68035d967886e716f829fe3dce90632b858afbc.
//
// Solidity: event CelerInitialized(address indexed messageBus, uint256 chainId)
func (_SoDiamond *SoDiamondFilterer) FilterCelerInitialized(opts *bind.FilterOpts, messageBus []common.Address) (*SoDiamondCelerInitializedIterator, error) {

	var messageBusRule []interface{}
	for _, messageBusItem := range messageBus {
		messageBusRule = append(messageBusRule, messageBusItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "CelerInitialized", messageBusRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondCelerInitializedIterator{contract: _SoDiamond.contract, event: "CelerInitialized", logs: logs, sub: sub}, nil
}

// WatchCelerInitialized is a free log subscription operation binding the contract event 0x4fdd908f8304e87ab2bf41b4e68035d967886e716f829fe3dce90632b858afbc.
//
// Solidity: event CelerInitialized(address indexed messageBus, uint256 chainId)
func (_SoDiamond *SoDiamondFilterer) WatchCelerInitialized(opts *bind.WatchOpts, sink chan<- *SoDiamondCelerInitialized, messageBus []common.Address) (event.Subscription, error) {

	var messageBusRule []interface{}
	for _, messageBusItem := range messageBus {
		messageBusRule = append(messageBusRule, messageBusItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "CelerInitialized", messageBusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondCelerInitialized)
				if err := _SoDiamond.contract.UnpackLog(event, "CelerInitialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCelerInitialized is a log parse operation binding the contract event 0x4fdd908f8304e87ab2bf41b4e68035d967886e716f829fe3dce90632b858afbc.
//
// Solidity: event CelerInitialized(address indexed messageBus, uint256 chainId)
func (_SoDiamond *SoDiamondFilterer) ParseCelerInitialized(log types.Log) (*SoDiamondCelerInitialized, error) {
	event := new(SoDiamondCelerInitialized)
	if err := _SoDiamond.contract.UnpackLog(event, "CelerInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondCorrectSwapRouterSelectorsChangedIterator is returned from FilterCorrectSwapRouterSelectorsChanged and is used to iterate over the raw logs and unpacked data for CorrectSwapRouterSelectorsChanged events raised by the SoDiamond contract.
type SoDiamondCorrectSwapRouterSelectorsChangedIterator struct {
	Event *SoDiamondCorrectSwapRouterSelectorsChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondCorrectSwapRouterSelectorsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondCorrectSwapRouterSelectorsChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondCorrectSwapRouterSelectorsChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondCorrectSwapRouterSelectorsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondCorrectSwapRouterSelectorsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondCorrectSwapRouterSelectorsChanged represents a CorrectSwapRouterSelectorsChanged event raised by the SoDiamond contract.
type SoDiamondCorrectSwapRouterSelectorsChanged struct {
	CorrectSwap common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCorrectSwapRouterSelectorsChanged is a free log retrieval operation binding the contract event 0x813bf5881670a05646b27acdfa0569a276a74e2c78a8489042e2583d2a303696.
//
// Solidity: event CorrectSwapRouterSelectorsChanged(address indexed correctSwap)
func (_SoDiamond *SoDiamondFilterer) FilterCorrectSwapRouterSelectorsChanged(opts *bind.FilterOpts, correctSwap []common.Address) (*SoDiamondCorrectSwapRouterSelectorsChangedIterator, error) {

	var correctSwapRule []interface{}
	for _, correctSwapItem := range correctSwap {
		correctSwapRule = append(correctSwapRule, correctSwapItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "CorrectSwapRouterSelectorsChanged", correctSwapRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondCorrectSwapRouterSelectorsChangedIterator{contract: _SoDiamond.contract, event: "CorrectSwapRouterSelectorsChanged", logs: logs, sub: sub}, nil
}

// WatchCorrectSwapRouterSelectorsChanged is a free log subscription operation binding the contract event 0x813bf5881670a05646b27acdfa0569a276a74e2c78a8489042e2583d2a303696.
//
// Solidity: event CorrectSwapRouterSelectorsChanged(address indexed correctSwap)
func (_SoDiamond *SoDiamondFilterer) WatchCorrectSwapRouterSelectorsChanged(opts *bind.WatchOpts, sink chan<- *SoDiamondCorrectSwapRouterSelectorsChanged, correctSwap []common.Address) (event.Subscription, error) {

	var correctSwapRule []interface{}
	for _, correctSwapItem := range correctSwap {
		correctSwapRule = append(correctSwapRule, correctSwapItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "CorrectSwapRouterSelectorsChanged", correctSwapRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondCorrectSwapRouterSelectorsChanged)
				if err := _SoDiamond.contract.UnpackLog(event, "CorrectSwapRouterSelectorsChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCorrectSwapRouterSelectorsChanged is a log parse operation binding the contract event 0x813bf5881670a05646b27acdfa0569a276a74e2c78a8489042e2583d2a303696.
//
// Solidity: event CorrectSwapRouterSelectorsChanged(address indexed correctSwap)
func (_SoDiamond *SoDiamondFilterer) ParseCorrectSwapRouterSelectorsChanged(log types.Log) (*SoDiamondCorrectSwapRouterSelectorsChanged, error) {
	event := new(SoDiamondCorrectSwapRouterSelectorsChanged)
	if err := _SoDiamond.contract.UnpackLog(event, "CorrectSwapRouterSelectorsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondDexAddedIterator is returned from FilterDexAdded and is used to iterate over the raw logs and unpacked data for DexAdded events raised by the SoDiamond contract.
type SoDiamondDexAddedIterator struct {
	Event *SoDiamondDexAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondDexAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondDexAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondDexAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondDexAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondDexAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondDexAdded represents a DexAdded event raised by the SoDiamond contract.
type SoDiamondDexAdded struct {
	DexAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDexAdded is a free log retrieval operation binding the contract event 0x7e0058dd0cbc0a8b7beaa013a4825655d8e9e81a5e2cc6582818deded0a41b99.
//
// Solidity: event DexAdded(address indexed dexAddress)
func (_SoDiamond *SoDiamondFilterer) FilterDexAdded(opts *bind.FilterOpts, dexAddress []common.Address) (*SoDiamondDexAddedIterator, error) {

	var dexAddressRule []interface{}
	for _, dexAddressItem := range dexAddress {
		dexAddressRule = append(dexAddressRule, dexAddressItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "DexAdded", dexAddressRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondDexAddedIterator{contract: _SoDiamond.contract, event: "DexAdded", logs: logs, sub: sub}, nil
}

// WatchDexAdded is a free log subscription operation binding the contract event 0x7e0058dd0cbc0a8b7beaa013a4825655d8e9e81a5e2cc6582818deded0a41b99.
//
// Solidity: event DexAdded(address indexed dexAddress)
func (_SoDiamond *SoDiamondFilterer) WatchDexAdded(opts *bind.WatchOpts, sink chan<- *SoDiamondDexAdded, dexAddress []common.Address) (event.Subscription, error) {

	var dexAddressRule []interface{}
	for _, dexAddressItem := range dexAddress {
		dexAddressRule = append(dexAddressRule, dexAddressItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "DexAdded", dexAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondDexAdded)
				if err := _SoDiamond.contract.UnpackLog(event, "DexAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDexAdded is a log parse operation binding the contract event 0x7e0058dd0cbc0a8b7beaa013a4825655d8e9e81a5e2cc6582818deded0a41b99.
//
// Solidity: event DexAdded(address indexed dexAddress)
func (_SoDiamond *SoDiamondFilterer) ParseDexAdded(log types.Log) (*SoDiamondDexAdded, error) {
	event := new(SoDiamondDexAdded)
	if err := _SoDiamond.contract.UnpackLog(event, "DexAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondDexRemovedIterator is returned from FilterDexRemoved and is used to iterate over the raw logs and unpacked data for DexRemoved events raised by the SoDiamond contract.
type SoDiamondDexRemovedIterator struct {
	Event *SoDiamondDexRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondDexRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondDexRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondDexRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondDexRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondDexRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondDexRemoved represents a DexRemoved event raised by the SoDiamond contract.
type SoDiamondDexRemoved struct {
	DexAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDexRemoved is a free log retrieval operation binding the contract event 0x78e0a2ffcdfbbb49ba5c8050d8630fab2176d825e8360809db049cd98f462a78.
//
// Solidity: event DexRemoved(address indexed dexAddress)
func (_SoDiamond *SoDiamondFilterer) FilterDexRemoved(opts *bind.FilterOpts, dexAddress []common.Address) (*SoDiamondDexRemovedIterator, error) {

	var dexAddressRule []interface{}
	for _, dexAddressItem := range dexAddress {
		dexAddressRule = append(dexAddressRule, dexAddressItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "DexRemoved", dexAddressRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondDexRemovedIterator{contract: _SoDiamond.contract, event: "DexRemoved", logs: logs, sub: sub}, nil
}

// WatchDexRemoved is a free log subscription operation binding the contract event 0x78e0a2ffcdfbbb49ba5c8050d8630fab2176d825e8360809db049cd98f462a78.
//
// Solidity: event DexRemoved(address indexed dexAddress)
func (_SoDiamond *SoDiamondFilterer) WatchDexRemoved(opts *bind.WatchOpts, sink chan<- *SoDiamondDexRemoved, dexAddress []common.Address) (event.Subscription, error) {

	var dexAddressRule []interface{}
	for _, dexAddressItem := range dexAddress {
		dexAddressRule = append(dexAddressRule, dexAddressItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "DexRemoved", dexAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondDexRemoved)
				if err := _SoDiamond.contract.UnpackLog(event, "DexRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDexRemoved is a log parse operation binding the contract event 0x78e0a2ffcdfbbb49ba5c8050d8630fab2176d825e8360809db049cd98f462a78.
//
// Solidity: event DexRemoved(address indexed dexAddress)
func (_SoDiamond *SoDiamondFilterer) ParseDexRemoved(log types.Log) (*SoDiamondDexRemoved, error) {
	event := new(SoDiamondDexRemoved)
	if err := _SoDiamond.contract.UnpackLog(event, "DexRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondDiamondCutIterator is returned from FilterDiamondCut and is used to iterate over the raw logs and unpacked data for DiamondCut events raised by the SoDiamond contract.
type SoDiamondDiamondCutIterator struct {
	Event *SoDiamondDiamondCut // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondDiamondCutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondDiamondCut)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondDiamondCut)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondDiamondCutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondDiamondCutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondDiamondCut represents a DiamondCut event raised by the SoDiamond contract.
type SoDiamondDiamondCut struct {
	DiamondCut []IDiamondCutFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDiamondCut is a free log retrieval operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_SoDiamond *SoDiamondFilterer) FilterDiamondCut(opts *bind.FilterOpts) (*SoDiamondDiamondCutIterator, error) {

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return &SoDiamondDiamondCutIterator{contract: _SoDiamond.contract, event: "DiamondCut", logs: logs, sub: sub}, nil
}

// WatchDiamondCut is a free log subscription operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_SoDiamond *SoDiamondFilterer) WatchDiamondCut(opts *bind.WatchOpts, sink chan<- *SoDiamondDiamondCut) (event.Subscription, error) {

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondDiamondCut)
				if err := _SoDiamond.contract.UnpackLog(event, "DiamondCut", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDiamondCut is a log parse operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_SoDiamond *SoDiamondFilterer) ParseDiamondCut(log types.Log) (*SoDiamondDiamondCut, error) {
	event := new(SoDiamondDiamondCut)
	if err := _SoDiamond.contract.UnpackLog(event, "DiamondCut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondFunctionSignatureApprovalChangedIterator is returned from FilterFunctionSignatureApprovalChanged and is used to iterate over the raw logs and unpacked data for FunctionSignatureApprovalChanged events raised by the SoDiamond contract.
type SoDiamondFunctionSignatureApprovalChangedIterator struct {
	Event *SoDiamondFunctionSignatureApprovalChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondFunctionSignatureApprovalChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondFunctionSignatureApprovalChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondFunctionSignatureApprovalChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondFunctionSignatureApprovalChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondFunctionSignatureApprovalChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondFunctionSignatureApprovalChanged represents a FunctionSignatureApprovalChanged event raised by the SoDiamond contract.
type SoDiamondFunctionSignatureApprovalChanged struct {
	FunctionSignature [32]byte
	Approved          bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFunctionSignatureApprovalChanged is a free log retrieval operation binding the contract event 0x0530b9c29ef48134f8041dfb0d833636aad6cfc43439c54ed4184067c3b1b278.
//
// Solidity: event FunctionSignatureApprovalChanged(bytes32 indexed functionSignature, bool indexed approved)
func (_SoDiamond *SoDiamondFilterer) FilterFunctionSignatureApprovalChanged(opts *bind.FilterOpts, functionSignature [][32]byte, approved []bool) (*SoDiamondFunctionSignatureApprovalChangedIterator, error) {

	var functionSignatureRule []interface{}
	for _, functionSignatureItem := range functionSignature {
		functionSignatureRule = append(functionSignatureRule, functionSignatureItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "FunctionSignatureApprovalChanged", functionSignatureRule, approvedRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondFunctionSignatureApprovalChangedIterator{contract: _SoDiamond.contract, event: "FunctionSignatureApprovalChanged", logs: logs, sub: sub}, nil
}

// WatchFunctionSignatureApprovalChanged is a free log subscription operation binding the contract event 0x0530b9c29ef48134f8041dfb0d833636aad6cfc43439c54ed4184067c3b1b278.
//
// Solidity: event FunctionSignatureApprovalChanged(bytes32 indexed functionSignature, bool indexed approved)
func (_SoDiamond *SoDiamondFilterer) WatchFunctionSignatureApprovalChanged(opts *bind.WatchOpts, sink chan<- *SoDiamondFunctionSignatureApprovalChanged, functionSignature [][32]byte, approved []bool) (event.Subscription, error) {

	var functionSignatureRule []interface{}
	for _, functionSignatureItem := range functionSignature {
		functionSignatureRule = append(functionSignatureRule, functionSignatureItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "FunctionSignatureApprovalChanged", functionSignatureRule, approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondFunctionSignatureApprovalChanged)
				if err := _SoDiamond.contract.UnpackLog(event, "FunctionSignatureApprovalChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFunctionSignatureApprovalChanged is a log parse operation binding the contract event 0x0530b9c29ef48134f8041dfb0d833636aad6cfc43439c54ed4184067c3b1b278.
//
// Solidity: event FunctionSignatureApprovalChanged(bytes32 indexed functionSignature, bool indexed approved)
func (_SoDiamond *SoDiamondFilterer) ParseFunctionSignatureApprovalChanged(log types.Log) (*SoDiamondFunctionSignatureApprovalChanged, error) {
	event := new(SoDiamondFunctionSignatureApprovalChanged)
	if err := _SoDiamond.contract.UnpackLog(event, "FunctionSignatureApprovalChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondGatewaySoFeeSelectorsChangedIterator is returned from FilterGatewaySoFeeSelectorsChanged and is used to iterate over the raw logs and unpacked data for GatewaySoFeeSelectorsChanged events raised by the SoDiamond contract.
type SoDiamondGatewaySoFeeSelectorsChangedIterator struct {
	Event *SoDiamondGatewaySoFeeSelectorsChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondGatewaySoFeeSelectorsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondGatewaySoFeeSelectorsChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondGatewaySoFeeSelectorsChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondGatewaySoFeeSelectorsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondGatewaySoFeeSelectorsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondGatewaySoFeeSelectorsChanged represents a GatewaySoFeeSelectorsChanged event raised by the SoDiamond contract.
type SoDiamondGatewaySoFeeSelectorsChanged struct {
	GatewayAddress common.Address
	SoFeeAddress   common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGatewaySoFeeSelectorsChanged is a free log retrieval operation binding the contract event 0x10cd97dcf8884aa032ee3617336995e80d74490b2f11d245945599361f03b85e.
//
// Solidity: event GatewaySoFeeSelectorsChanged(address indexed gatewayAddress, address indexed soFeeAddress)
func (_SoDiamond *SoDiamondFilterer) FilterGatewaySoFeeSelectorsChanged(opts *bind.FilterOpts, gatewayAddress []common.Address, soFeeAddress []common.Address) (*SoDiamondGatewaySoFeeSelectorsChangedIterator, error) {

	var gatewayAddressRule []interface{}
	for _, gatewayAddressItem := range gatewayAddress {
		gatewayAddressRule = append(gatewayAddressRule, gatewayAddressItem)
	}
	var soFeeAddressRule []interface{}
	for _, soFeeAddressItem := range soFeeAddress {
		soFeeAddressRule = append(soFeeAddressRule, soFeeAddressItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "GatewaySoFeeSelectorsChanged", gatewayAddressRule, soFeeAddressRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondGatewaySoFeeSelectorsChangedIterator{contract: _SoDiamond.contract, event: "GatewaySoFeeSelectorsChanged", logs: logs, sub: sub}, nil
}

// WatchGatewaySoFeeSelectorsChanged is a free log subscription operation binding the contract event 0x10cd97dcf8884aa032ee3617336995e80d74490b2f11d245945599361f03b85e.
//
// Solidity: event GatewaySoFeeSelectorsChanged(address indexed gatewayAddress, address indexed soFeeAddress)
func (_SoDiamond *SoDiamondFilterer) WatchGatewaySoFeeSelectorsChanged(opts *bind.WatchOpts, sink chan<- *SoDiamondGatewaySoFeeSelectorsChanged, gatewayAddress []common.Address, soFeeAddress []common.Address) (event.Subscription, error) {

	var gatewayAddressRule []interface{}
	for _, gatewayAddressItem := range gatewayAddress {
		gatewayAddressRule = append(gatewayAddressRule, gatewayAddressItem)
	}
	var soFeeAddressRule []interface{}
	for _, soFeeAddressItem := range soFeeAddress {
		soFeeAddressRule = append(soFeeAddressRule, soFeeAddressItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "GatewaySoFeeSelectorsChanged", gatewayAddressRule, soFeeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondGatewaySoFeeSelectorsChanged)
				if err := _SoDiamond.contract.UnpackLog(event, "GatewaySoFeeSelectorsChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGatewaySoFeeSelectorsChanged is a log parse operation binding the contract event 0x10cd97dcf8884aa032ee3617336995e80d74490b2f11d245945599361f03b85e.
//
// Solidity: event GatewaySoFeeSelectorsChanged(address indexed gatewayAddress, address indexed soFeeAddress)
func (_SoDiamond *SoDiamondFilterer) ParseGatewaySoFeeSelectorsChanged(log types.Log) (*SoDiamondGatewaySoFeeSelectorsChanged, error) {
	event := new(SoDiamondGatewaySoFeeSelectorsChanged)
	if err := _SoDiamond.contract.UnpackLog(event, "GatewaySoFeeSelectorsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondInitWormholeEventIterator is returned from FilterInitWormholeEvent and is used to iterate over the raw logs and unpacked data for InitWormholeEvent events raised by the SoDiamond contract.
type SoDiamondInitWormholeEventIterator struct {
	Event *SoDiamondInitWormholeEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondInitWormholeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondInitWormholeEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondInitWormholeEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondInitWormholeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondInitWormholeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondInitWormholeEvent represents a InitWormholeEvent event raised by the SoDiamond contract.
type SoDiamondInitWormholeEvent struct {
	TokenBridge        common.Address
	SrcWormholeChainId uint16
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterInitWormholeEvent is a free log retrieval operation binding the contract event 0xb42c2803db63f6a1bdd8d9d2f1114f85fb270af64df8e69b20abc2ac4b61b95d.
//
// Solidity: event InitWormholeEvent(address tokenBridge, uint16 srcWormholeChainId)
func (_SoDiamond *SoDiamondFilterer) FilterInitWormholeEvent(opts *bind.FilterOpts) (*SoDiamondInitWormholeEventIterator, error) {

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "InitWormholeEvent")
	if err != nil {
		return nil, err
	}
	return &SoDiamondInitWormholeEventIterator{contract: _SoDiamond.contract, event: "InitWormholeEvent", logs: logs, sub: sub}, nil
}

// WatchInitWormholeEvent is a free log subscription operation binding the contract event 0xb42c2803db63f6a1bdd8d9d2f1114f85fb270af64df8e69b20abc2ac4b61b95d.
//
// Solidity: event InitWormholeEvent(address tokenBridge, uint16 srcWormholeChainId)
func (_SoDiamond *SoDiamondFilterer) WatchInitWormholeEvent(opts *bind.WatchOpts, sink chan<- *SoDiamondInitWormholeEvent) (event.Subscription, error) {

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "InitWormholeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondInitWormholeEvent)
				if err := _SoDiamond.contract.UnpackLog(event, "InitWormholeEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitWormholeEvent is a log parse operation binding the contract event 0xb42c2803db63f6a1bdd8d9d2f1114f85fb270af64df8e69b20abc2ac4b61b95d.
//
// Solidity: event InitWormholeEvent(address tokenBridge, uint16 srcWormholeChainId)
func (_SoDiamond *SoDiamondFilterer) ParseInitWormholeEvent(log types.Log) (*SoDiamondInitWormholeEvent, error) {
	event := new(SoDiamondInitWormholeEvent)
	if err := _SoDiamond.contract.UnpackLog(event, "InitWormholeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondLogWithdrawIterator is returned from FilterLogWithdraw and is used to iterate over the raw logs and unpacked data for LogWithdraw events raised by the SoDiamond contract.
type SoDiamondLogWithdrawIterator struct {
	Event *SoDiamondLogWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondLogWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondLogWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondLogWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondLogWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondLogWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondLogWithdraw represents a LogWithdraw event raised by the SoDiamond contract.
type SoDiamondLogWithdraw struct {
	AssetAddress common.Address
	To           common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogWithdraw is a free log retrieval operation binding the contract event 0x9207361cc2a04b9c7a06691df1eb87c6a63957ae88bf01d0d18c81e3d1272099.
//
// Solidity: event LogWithdraw(address indexed assetAddress, address to, uint256 amount)
func (_SoDiamond *SoDiamondFilterer) FilterLogWithdraw(opts *bind.FilterOpts, assetAddress []common.Address) (*SoDiamondLogWithdrawIterator, error) {

	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "LogWithdraw", assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondLogWithdrawIterator{contract: _SoDiamond.contract, event: "LogWithdraw", logs: logs, sub: sub}, nil
}

// WatchLogWithdraw is a free log subscription operation binding the contract event 0x9207361cc2a04b9c7a06691df1eb87c6a63957ae88bf01d0d18c81e3d1272099.
//
// Solidity: event LogWithdraw(address indexed assetAddress, address to, uint256 amount)
func (_SoDiamond *SoDiamondFilterer) WatchLogWithdraw(opts *bind.WatchOpts, sink chan<- *SoDiamondLogWithdraw, assetAddress []common.Address) (event.Subscription, error) {

	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "LogWithdraw", assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondLogWithdraw)
				if err := _SoDiamond.contract.UnpackLog(event, "LogWithdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogWithdraw is a log parse operation binding the contract event 0x9207361cc2a04b9c7a06691df1eb87c6a63957ae88bf01d0d18c81e3d1272099.
//
// Solidity: event LogWithdraw(address indexed assetAddress, address to, uint256 amount)
func (_SoDiamond *SoDiamondFilterer) ParseLogWithdraw(log types.Log) (*SoDiamondLogWithdraw, error) {
	event := new(SoDiamondLogWithdraw)
	if err := _SoDiamond.contract.UnpackLog(event, "LogWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondMultiChainInitializedIterator is returned from FilterMultiChainInitialized and is used to iterate over the raw logs and unpacked data for MultiChainInitialized events raised by the SoDiamond contract.
type SoDiamondMultiChainInitializedIterator struct {
	Event *SoDiamondMultiChainInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondMultiChainInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondMultiChainInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondMultiChainInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondMultiChainInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondMultiChainInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondMultiChainInitialized represents a MultiChainInitialized event raised by the SoDiamond contract.
type SoDiamondMultiChainInitialized struct {
	FastRouter common.Address
	ChainId    uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMultiChainInitialized is a free log retrieval operation binding the contract event 0x86dc953ff32ec15f553729198d8edcce96b5ca6cba9a8f737c4d5a27c203838a.
//
// Solidity: event MultiChainInitialized(address fastRouter, uint64 chainId)
func (_SoDiamond *SoDiamondFilterer) FilterMultiChainInitialized(opts *bind.FilterOpts) (*SoDiamondMultiChainInitializedIterator, error) {

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "MultiChainInitialized")
	if err != nil {
		return nil, err
	}
	return &SoDiamondMultiChainInitializedIterator{contract: _SoDiamond.contract, event: "MultiChainInitialized", logs: logs, sub: sub}, nil
}

// WatchMultiChainInitialized is a free log subscription operation binding the contract event 0x86dc953ff32ec15f553729198d8edcce96b5ca6cba9a8f737c4d5a27c203838a.
//
// Solidity: event MultiChainInitialized(address fastRouter, uint64 chainId)
func (_SoDiamond *SoDiamondFilterer) WatchMultiChainInitialized(opts *bind.WatchOpts, sink chan<- *SoDiamondMultiChainInitialized) (event.Subscription, error) {

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "MultiChainInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondMultiChainInitialized)
				if err := _SoDiamond.contract.UnpackLog(event, "MultiChainInitialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMultiChainInitialized is a log parse operation binding the contract event 0x86dc953ff32ec15f553729198d8edcce96b5ca6cba9a8f737c4d5a27c203838a.
//
// Solidity: event MultiChainInitialized(address fastRouter, uint64 chainId)
func (_SoDiamond *SoDiamondFilterer) ParseMultiChainInitialized(log types.Log) (*SoDiamondMultiChainInitialized, error) {
	event := new(SoDiamondMultiChainInitialized)
	if err := _SoDiamond.contract.UnpackLog(event, "MultiChainInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the SoDiamond contract.
type SoDiamondOwnershipTransferRequestedIterator struct {
	Event *SoDiamondOwnershipTransferRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the SoDiamond contract.
type SoDiamondOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed _from, address indexed _to)
func (_SoDiamond *SoDiamondFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*SoDiamondOwnershipTransferRequestedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "OwnershipTransferRequested", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondOwnershipTransferRequestedIterator{contract: _SoDiamond.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed _from, address indexed _to)
func (_SoDiamond *SoDiamondFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *SoDiamondOwnershipTransferRequested, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "OwnershipTransferRequested", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondOwnershipTransferRequested)
				if err := _SoDiamond.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferRequested is a log parse operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed _from, address indexed _to)
func (_SoDiamond *SoDiamondFilterer) ParseOwnershipTransferRequested(log types.Log) (*SoDiamondOwnershipTransferRequested, error) {
	event := new(SoDiamondOwnershipTransferRequested)
	if err := _SoDiamond.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SoDiamond contract.
type SoDiamondOwnershipTransferredIterator struct {
	Event *SoDiamondOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondOwnershipTransferred represents a OwnershipTransferred event raised by the SoDiamond contract.
type SoDiamondOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SoDiamond *SoDiamondFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SoDiamondOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondOwnershipTransferredIterator{contract: _SoDiamond.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SoDiamond *SoDiamondFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SoDiamondOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondOwnershipTransferred)
				if err := _SoDiamond.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SoDiamond *SoDiamondFilterer) ParseOwnershipTransferred(log types.Log) (*SoDiamondOwnershipTransferred, error) {
	event := new(SoDiamondOwnershipTransferred)
	if err := _SoDiamond.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondPathPairUpdatedIterator is returned from FilterPathPairUpdated and is used to iterate over the raw logs and unpacked data for PathPairUpdated events raised by the SoDiamond contract.
type SoDiamondPathPairUpdatedIterator struct {
	Event *SoDiamondPathPairUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondPathPairUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondPathPairUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondPathPairUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondPathPairUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondPathPairUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondPathPairUpdated represents a PathPairUpdated event raised by the SoDiamond contract.
type SoDiamondPathPairUpdated struct {
	ConsumerChainId uint32
	BoolChainId     uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPathPairUpdated is a free log retrieval operation binding the contract event 0x79ebfff1780907354058c6a8bb57c06ca43a07b0eee4367325a11c515ed443bd.
//
// Solidity: event PathPairUpdated(uint32 indexed consumerChainId, uint32 indexed boolChainId)
func (_SoDiamond *SoDiamondFilterer) FilterPathPairUpdated(opts *bind.FilterOpts, consumerChainId []uint32, boolChainId []uint32) (*SoDiamondPathPairUpdatedIterator, error) {

	var consumerChainIdRule []interface{}
	for _, consumerChainIdItem := range consumerChainId {
		consumerChainIdRule = append(consumerChainIdRule, consumerChainIdItem)
	}
	var boolChainIdRule []interface{}
	for _, boolChainIdItem := range boolChainId {
		boolChainIdRule = append(boolChainIdRule, boolChainIdItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "PathPairUpdated", consumerChainIdRule, boolChainIdRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondPathPairUpdatedIterator{contract: _SoDiamond.contract, event: "PathPairUpdated", logs: logs, sub: sub}, nil
}

// WatchPathPairUpdated is a free log subscription operation binding the contract event 0x79ebfff1780907354058c6a8bb57c06ca43a07b0eee4367325a11c515ed443bd.
//
// Solidity: event PathPairUpdated(uint32 indexed consumerChainId, uint32 indexed boolChainId)
func (_SoDiamond *SoDiamondFilterer) WatchPathPairUpdated(opts *bind.WatchOpts, sink chan<- *SoDiamondPathPairUpdated, consumerChainId []uint32, boolChainId []uint32) (event.Subscription, error) {

	var consumerChainIdRule []interface{}
	for _, consumerChainIdItem := range consumerChainId {
		consumerChainIdRule = append(consumerChainIdRule, consumerChainIdItem)
	}
	var boolChainIdRule []interface{}
	for _, boolChainIdItem := range boolChainId {
		boolChainIdRule = append(boolChainIdRule, boolChainIdItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "PathPairUpdated", consumerChainIdRule, boolChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondPathPairUpdated)
				if err := _SoDiamond.contract.UnpackLog(event, "PathPairUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePathPairUpdated is a log parse operation binding the contract event 0x79ebfff1780907354058c6a8bb57c06ca43a07b0eee4367325a11c515ed443bd.
//
// Solidity: event PathPairUpdated(uint32 indexed consumerChainId, uint32 indexed boolChainId)
func (_SoDiamond *SoDiamondFilterer) ParsePathPairUpdated(log types.Log) (*SoDiamondPathPairUpdated, error) {
	event := new(SoDiamondPathPairUpdated)
	if err := _SoDiamond.contract.UnpackLog(event, "PathPairUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondRefundCelerTokenIterator is returned from FilterRefundCelerToken and is used to iterate over the raw logs and unpacked data for RefundCelerToken events raised by the SoDiamond contract.
type SoDiamondRefundCelerTokenIterator struct {
	Event *SoDiamondRefundCelerToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondRefundCelerTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondRefundCelerToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondRefundCelerToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondRefundCelerTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondRefundCelerTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondRefundCelerToken represents a RefundCelerToken event raised by the SoDiamond contract.
type SoDiamondRefundCelerToken struct {
	Token   common.Address
	Sender  common.Address
	Amount  *big.Int
	SrcTxId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRefundCelerToken is a free log retrieval operation binding the contract event 0xbc7bb07eb9d150e49466821273ec6365cd37f06ddab93621bfc87087d82409c9.
//
// Solidity: event RefundCelerToken(address indexed token, address sender, uint256 amount, bytes32 srcTxId)
func (_SoDiamond *SoDiamondFilterer) FilterRefundCelerToken(opts *bind.FilterOpts, token []common.Address) (*SoDiamondRefundCelerTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "RefundCelerToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondRefundCelerTokenIterator{contract: _SoDiamond.contract, event: "RefundCelerToken", logs: logs, sub: sub}, nil
}

// WatchRefundCelerToken is a free log subscription operation binding the contract event 0xbc7bb07eb9d150e49466821273ec6365cd37f06ddab93621bfc87087d82409c9.
//
// Solidity: event RefundCelerToken(address indexed token, address sender, uint256 amount, bytes32 srcTxId)
func (_SoDiamond *SoDiamondFilterer) WatchRefundCelerToken(opts *bind.WatchOpts, sink chan<- *SoDiamondRefundCelerToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "RefundCelerToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondRefundCelerToken)
				if err := _SoDiamond.contract.UnpackLog(event, "RefundCelerToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefundCelerToken is a log parse operation binding the contract event 0xbc7bb07eb9d150e49466821273ec6365cd37f06ddab93621bfc87087d82409c9.
//
// Solidity: event RefundCelerToken(address indexed token, address sender, uint256 amount, bytes32 srcTxId)
func (_SoDiamond *SoDiamondFilterer) ParseRefundCelerToken(log types.Log) (*SoDiamondRefundCelerToken, error) {
	event := new(SoDiamondRefundCelerToken)
	if err := _SoDiamond.contract.UnpackLog(event, "RefundCelerToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSetAllowedListIterator is returned from FilterSetAllowedList and is used to iterate over the raw logs and unpacked data for SetAllowedList events raised by the SoDiamond contract.
type SoDiamondSetAllowedListIterator struct {
	Event *SoDiamondSetAllowedList // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSetAllowedListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSetAllowedList)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSetAllowedList)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSetAllowedListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSetAllowedListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSetAllowedList represents a SetAllowedList event raised by the SoDiamond contract.
type SoDiamondSetAllowedList struct {
	Router    common.Address
	IsAllowed bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetAllowedList is a free log retrieval operation binding the contract event 0xa915695662b2872d8e64773d8301cf0e418b55f86ee42d4c8f5c07d8f0bb19b0.
//
// Solidity: event SetAllowedList(address router, bool isAllowed)
func (_SoDiamond *SoDiamondFilterer) FilterSetAllowedList(opts *bind.FilterOpts) (*SoDiamondSetAllowedListIterator, error) {

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SetAllowedList")
	if err != nil {
		return nil, err
	}
	return &SoDiamondSetAllowedListIterator{contract: _SoDiamond.contract, event: "SetAllowedList", logs: logs, sub: sub}, nil
}

// WatchSetAllowedList is a free log subscription operation binding the contract event 0xa915695662b2872d8e64773d8301cf0e418b55f86ee42d4c8f5c07d8f0bb19b0.
//
// Solidity: event SetAllowedList(address router, bool isAllowed)
func (_SoDiamond *SoDiamondFilterer) WatchSetAllowedList(opts *bind.WatchOpts, sink chan<- *SoDiamondSetAllowedList) (event.Subscription, error) {

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SetAllowedList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSetAllowedList)
				if err := _SoDiamond.contract.UnpackLog(event, "SetAllowedList", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetAllowedList is a log parse operation binding the contract event 0xa915695662b2872d8e64773d8301cf0e418b55f86ee42d4c8f5c07d8f0bb19b0.
//
// Solidity: event SetAllowedList(address router, bool isAllowed)
func (_SoDiamond *SoDiamondFilterer) ParseSetAllowedList(log types.Log) (*SoDiamondSetAllowedList, error) {
	event := new(SoDiamondSetAllowedList)
	if err := _SoDiamond.contract.UnpackLog(event, "SetAllowedList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSetAllowedList0Iterator is returned from FilterSetAllowedList0 and is used to iterate over the raw logs and unpacked data for SetAllowedList0 events raised by the SoDiamond contract.
type SoDiamondSetAllowedList0Iterator struct {
	Event *SoDiamondSetAllowedList0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSetAllowedList0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSetAllowedList0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSetAllowedList0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSetAllowedList0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSetAllowedList0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSetAllowedList0 represents a SetAllowedList0 event raised by the SoDiamond contract.
type SoDiamondSetAllowedList0 struct {
	MessageBus common.Address
	IsAllowed  bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetAllowedList0 is a free log retrieval operation binding the contract event 0xa915695662b2872d8e64773d8301cf0e418b55f86ee42d4c8f5c07d8f0bb19b0.
//
// Solidity: event SetAllowedList(address indexed messageBus, bool isAllowed)
func (_SoDiamond *SoDiamondFilterer) FilterSetAllowedList0(opts *bind.FilterOpts, messageBus []common.Address) (*SoDiamondSetAllowedList0Iterator, error) {

	var messageBusRule []interface{}
	for _, messageBusItem := range messageBus {
		messageBusRule = append(messageBusRule, messageBusItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SetAllowedList0", messageBusRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondSetAllowedList0Iterator{contract: _SoDiamond.contract, event: "SetAllowedList0", logs: logs, sub: sub}, nil
}

// WatchSetAllowedList0 is a free log subscription operation binding the contract event 0xa915695662b2872d8e64773d8301cf0e418b55f86ee42d4c8f5c07d8f0bb19b0.
//
// Solidity: event SetAllowedList(address indexed messageBus, bool isAllowed)
func (_SoDiamond *SoDiamondFilterer) WatchSetAllowedList0(opts *bind.WatchOpts, sink chan<- *SoDiamondSetAllowedList0, messageBus []common.Address) (event.Subscription, error) {

	var messageBusRule []interface{}
	for _, messageBusItem := range messageBus {
		messageBusRule = append(messageBusRule, messageBusItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SetAllowedList0", messageBusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSetAllowedList0)
				if err := _SoDiamond.contract.UnpackLog(event, "SetAllowedList0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetAllowedList0 is a log parse operation binding the contract event 0xa915695662b2872d8e64773d8301cf0e418b55f86ee42d4c8f5c07d8f0bb19b0.
//
// Solidity: event SetAllowedList(address indexed messageBus, bool isAllowed)
func (_SoDiamond *SoDiamondFilterer) ParseSetAllowedList0(log types.Log) (*SoDiamondSetAllowedList0, error) {
	event := new(SoDiamondSetAllowedList0)
	if err := _SoDiamond.contract.UnpackLog(event, "SetAllowedList0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSetAllowedList1Iterator is returned from FilterSetAllowedList1 and is used to iterate over the raw logs and unpacked data for SetAllowedList1 events raised by the SoDiamond contract.
type SoDiamondSetAllowedList1Iterator struct {
	Event *SoDiamondSetAllowedList1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSetAllowedList1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSetAllowedList1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSetAllowedList1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSetAllowedList1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSetAllowedList1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSetAllowedList1 represents a SetAllowedList1 event raised by the SoDiamond contract.
type SoDiamondSetAllowedList1 struct {
	Account   common.Address
	IsAllowed bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetAllowedList1 is a free log retrieval operation binding the contract event 0xa915695662b2872d8e64773d8301cf0e418b55f86ee42d4c8f5c07d8f0bb19b0.
//
// Solidity: event SetAllowedList(address account, bool isAllowed)
func (_SoDiamond *SoDiamondFilterer) FilterSetAllowedList1(opts *bind.FilterOpts) (*SoDiamondSetAllowedList1Iterator, error) {

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SetAllowedList1")
	if err != nil {
		return nil, err
	}
	return &SoDiamondSetAllowedList1Iterator{contract: _SoDiamond.contract, event: "SetAllowedList1", logs: logs, sub: sub}, nil
}

// WatchSetAllowedList1 is a free log subscription operation binding the contract event 0xa915695662b2872d8e64773d8301cf0e418b55f86ee42d4c8f5c07d8f0bb19b0.
//
// Solidity: event SetAllowedList(address account, bool isAllowed)
func (_SoDiamond *SoDiamondFilterer) WatchSetAllowedList1(opts *bind.WatchOpts, sink chan<- *SoDiamondSetAllowedList1) (event.Subscription, error) {

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SetAllowedList1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSetAllowedList1)
				if err := _SoDiamond.contract.UnpackLog(event, "SetAllowedList1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetAllowedList1 is a log parse operation binding the contract event 0xa915695662b2872d8e64773d8301cf0e418b55f86ee42d4c8f5c07d8f0bb19b0.
//
// Solidity: event SetAllowedList(address account, bool isAllowed)
func (_SoDiamond *SoDiamondFilterer) ParseSetAllowedList1(log types.Log) (*SoDiamondSetAllowedList1, error) {
	event := new(SoDiamondSetAllowedList1)
	if err := _SoDiamond.contract.UnpackLog(event, "SetAllowedList1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSetAllowedList2Iterator is returned from FilterSetAllowedList2 and is used to iterate over the raw logs and unpacked data for SetAllowedList2 events raised by the SoDiamond contract.
type SoDiamondSetAllowedList2Iterator struct {
	Event *SoDiamondSetAllowedList2 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSetAllowedList2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSetAllowedList2)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSetAllowedList2)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSetAllowedList2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSetAllowedList2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSetAllowedList2 represents a SetAllowedList2 event raised by the SoDiamond contract.
type SoDiamondSetAllowedList2 struct {
	BoolSwapPool common.Address
	IsAllowed    bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetAllowedList2 is a free log retrieval operation binding the contract event 0xa915695662b2872d8e64773d8301cf0e418b55f86ee42d4c8f5c07d8f0bb19b0.
//
// Solidity: event SetAllowedList(address boolSwapPool, bool isAllowed)
func (_SoDiamond *SoDiamondFilterer) FilterSetAllowedList2(opts *bind.FilterOpts) (*SoDiamondSetAllowedList2Iterator, error) {

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SetAllowedList2")
	if err != nil {
		return nil, err
	}
	return &SoDiamondSetAllowedList2Iterator{contract: _SoDiamond.contract, event: "SetAllowedList2", logs: logs, sub: sub}, nil
}

// WatchSetAllowedList2 is a free log subscription operation binding the contract event 0xa915695662b2872d8e64773d8301cf0e418b55f86ee42d4c8f5c07d8f0bb19b0.
//
// Solidity: event SetAllowedList(address boolSwapPool, bool isAllowed)
func (_SoDiamond *SoDiamondFilterer) WatchSetAllowedList2(opts *bind.WatchOpts, sink chan<- *SoDiamondSetAllowedList2) (event.Subscription, error) {

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SetAllowedList2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSetAllowedList2)
				if err := _SoDiamond.contract.UnpackLog(event, "SetAllowedList2", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetAllowedList2 is a log parse operation binding the contract event 0xa915695662b2872d8e64773d8301cf0e418b55f86ee42d4c8f5c07d8f0bb19b0.
//
// Solidity: event SetAllowedList(address boolSwapPool, bool isAllowed)
func (_SoDiamond *SoDiamondFilterer) ParseSetAllowedList2(log types.Log) (*SoDiamondSetAllowedList2, error) {
	event := new(SoDiamondSetAllowedList2)
	if err := _SoDiamond.contract.UnpackLog(event, "SetAllowedList2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSetMessageBusIterator is returned from FilterSetMessageBus and is used to iterate over the raw logs and unpacked data for SetMessageBus events raised by the SoDiamond contract.
type SoDiamondSetMessageBusIterator struct {
	Event *SoDiamondSetMessageBus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSetMessageBusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSetMessageBus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSetMessageBus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSetMessageBusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSetMessageBusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSetMessageBus represents a SetMessageBus event raised by the SoDiamond contract.
type SoDiamondSetMessageBus struct {
	MessageBus common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetMessageBus is a free log retrieval operation binding the contract event 0x1a9c8a2a5603565b380ea33f5d9b9bcc235223e646020e82da66c226ad48be14.
//
// Solidity: event SetMessageBus(address indexed messageBus)
func (_SoDiamond *SoDiamondFilterer) FilterSetMessageBus(opts *bind.FilterOpts, messageBus []common.Address) (*SoDiamondSetMessageBusIterator, error) {

	var messageBusRule []interface{}
	for _, messageBusItem := range messageBus {
		messageBusRule = append(messageBusRule, messageBusItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SetMessageBus", messageBusRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondSetMessageBusIterator{contract: _SoDiamond.contract, event: "SetMessageBus", logs: logs, sub: sub}, nil
}

// WatchSetMessageBus is a free log subscription operation binding the contract event 0x1a9c8a2a5603565b380ea33f5d9b9bcc235223e646020e82da66c226ad48be14.
//
// Solidity: event SetMessageBus(address indexed messageBus)
func (_SoDiamond *SoDiamondFilterer) WatchSetMessageBus(opts *bind.WatchOpts, sink chan<- *SoDiamondSetMessageBus, messageBus []common.Address) (event.Subscription, error) {

	var messageBusRule []interface{}
	for _, messageBusItem := range messageBus {
		messageBusRule = append(messageBusRule, messageBusItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SetMessageBus", messageBusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSetMessageBus)
				if err := _SoDiamond.contract.UnpackLog(event, "SetMessageBus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetMessageBus is a log parse operation binding the contract event 0x1a9c8a2a5603565b380ea33f5d9b9bcc235223e646020e82da66c226ad48be14.
//
// Solidity: event SetMessageBus(address indexed messageBus)
func (_SoDiamond *SoDiamondFilterer) ParseSetMessageBus(log types.Log) (*SoDiamondSetMessageBus, error) {
	event := new(SoDiamondSetMessageBus)
	if err := _SoDiamond.contract.UnpackLog(event, "SetMessageBus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSoSwappedGenericIterator is returned from FilterSoSwappedGeneric and is used to iterate over the raw logs and unpacked data for SoSwappedGeneric events raised by the SoDiamond contract.
type SoDiamondSoSwappedGenericIterator struct {
	Event *SoDiamondSoSwappedGeneric // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSoSwappedGenericIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSoSwappedGeneric)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSoSwappedGeneric)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSoSwappedGenericIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSoSwappedGenericIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSoSwappedGeneric represents a SoSwappedGeneric event raised by the SoDiamond contract.
type SoDiamondSoSwappedGeneric struct {
	TransactionId [32]byte
	FromAssetId   common.Address
	ToAssetId     common.Address
	FromAmount    *big.Int
	ToAmount      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSoSwappedGeneric is a free log retrieval operation binding the contract event 0x6447c738837bac15bb5e58f4b8751e2194914dd95b21fad428d86a95f846ed24.
//
// Solidity: event SoSwappedGeneric(bytes32 indexed transactionId, address fromAssetId, address toAssetId, uint256 fromAmount, uint256 toAmount)
func (_SoDiamond *SoDiamondFilterer) FilterSoSwappedGeneric(opts *bind.FilterOpts, transactionId [][32]byte) (*SoDiamondSoSwappedGenericIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SoSwappedGeneric", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondSoSwappedGenericIterator{contract: _SoDiamond.contract, event: "SoSwappedGeneric", logs: logs, sub: sub}, nil
}

// WatchSoSwappedGeneric is a free log subscription operation binding the contract event 0x6447c738837bac15bb5e58f4b8751e2194914dd95b21fad428d86a95f846ed24.
//
// Solidity: event SoSwappedGeneric(bytes32 indexed transactionId, address fromAssetId, address toAssetId, uint256 fromAmount, uint256 toAmount)
func (_SoDiamond *SoDiamondFilterer) WatchSoSwappedGeneric(opts *bind.WatchOpts, sink chan<- *SoDiamondSoSwappedGeneric, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SoSwappedGeneric", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSoSwappedGeneric)
				if err := _SoDiamond.contract.UnpackLog(event, "SoSwappedGeneric", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSoSwappedGeneric is a log parse operation binding the contract event 0x6447c738837bac15bb5e58f4b8751e2194914dd95b21fad428d86a95f846ed24.
//
// Solidity: event SoSwappedGeneric(bytes32 indexed transactionId, address fromAssetId, address toAssetId, uint256 fromAmount, uint256 toAmount)
func (_SoDiamond *SoDiamondFilterer) ParseSoSwappedGeneric(log types.Log) (*SoDiamondSoSwappedGeneric, error) {
	event := new(SoDiamondSoSwappedGeneric)
	if err := _SoDiamond.contract.UnpackLog(event, "SoSwappedGeneric", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSoTransferCompletedIterator is returned from FilterSoTransferCompleted and is used to iterate over the raw logs and unpacked data for SoTransferCompleted events raised by the SoDiamond contract.
type SoDiamondSoTransferCompletedIterator struct {
	Event *SoDiamondSoTransferCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSoTransferCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSoTransferCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSoTransferCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSoTransferCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSoTransferCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSoTransferCompleted represents a SoTransferCompleted event raised by the SoDiamond contract.
type SoDiamondSoTransferCompleted struct {
	TransactionId [32]byte
	ReceiveAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSoTransferCompleted is a free log retrieval operation binding the contract event 0x5272b980ac59723d5a8fe5be29daff5302abfaf057695289598e842dcf306e48.
//
// Solidity: event SoTransferCompleted(bytes32 indexed transactionId, uint256 receiveAmount)
func (_SoDiamond *SoDiamondFilterer) FilterSoTransferCompleted(opts *bind.FilterOpts, transactionId [][32]byte) (*SoDiamondSoTransferCompletedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SoTransferCompleted", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondSoTransferCompletedIterator{contract: _SoDiamond.contract, event: "SoTransferCompleted", logs: logs, sub: sub}, nil
}

// WatchSoTransferCompleted is a free log subscription operation binding the contract event 0x5272b980ac59723d5a8fe5be29daff5302abfaf057695289598e842dcf306e48.
//
// Solidity: event SoTransferCompleted(bytes32 indexed transactionId, uint256 receiveAmount)
func (_SoDiamond *SoDiamondFilterer) WatchSoTransferCompleted(opts *bind.WatchOpts, sink chan<- *SoDiamondSoTransferCompleted, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SoTransferCompleted", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSoTransferCompleted)
				if err := _SoDiamond.contract.UnpackLog(event, "SoTransferCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSoTransferCompleted is a log parse operation binding the contract event 0x5272b980ac59723d5a8fe5be29daff5302abfaf057695289598e842dcf306e48.
//
// Solidity: event SoTransferCompleted(bytes32 indexed transactionId, uint256 receiveAmount)
func (_SoDiamond *SoDiamondFilterer) ParseSoTransferCompleted(log types.Log) (*SoDiamondSoTransferCompleted, error) {
	event := new(SoDiamondSoTransferCompleted)
	if err := _SoDiamond.contract.UnpackLog(event, "SoTransferCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSoTransferCompleted0Iterator is returned from FilterSoTransferCompleted0 and is used to iterate over the raw logs and unpacked data for SoTransferCompleted0 events raised by the SoDiamond contract.
type SoDiamondSoTransferCompleted0Iterator struct {
	Event *SoDiamondSoTransferCompleted0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSoTransferCompleted0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSoTransferCompleted0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSoTransferCompleted0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSoTransferCompleted0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSoTransferCompleted0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSoTransferCompleted0 represents a SoTransferCompleted0 event raised by the SoDiamond contract.
type SoDiamondSoTransferCompleted0 struct {
	TransactionId [32]byte
	ReceiveAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSoTransferCompleted0 is a free log retrieval operation binding the contract event 0x5272b980ac59723d5a8fe5be29daff5302abfaf057695289598e842dcf306e48.
//
// Solidity: event SoTransferCompleted(bytes32 indexed transactionId, uint256 receiveAmount)
func (_SoDiamond *SoDiamondFilterer) FilterSoTransferCompleted0(opts *bind.FilterOpts, transactionId [][32]byte) (*SoDiamondSoTransferCompleted0Iterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SoTransferCompleted0", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondSoTransferCompleted0Iterator{contract: _SoDiamond.contract, event: "SoTransferCompleted0", logs: logs, sub: sub}, nil
}

// WatchSoTransferCompleted0 is a free log subscription operation binding the contract event 0x5272b980ac59723d5a8fe5be29daff5302abfaf057695289598e842dcf306e48.
//
// Solidity: event SoTransferCompleted(bytes32 indexed transactionId, uint256 receiveAmount)
func (_SoDiamond *SoDiamondFilterer) WatchSoTransferCompleted0(opts *bind.WatchOpts, sink chan<- *SoDiamondSoTransferCompleted0, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SoTransferCompleted0", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSoTransferCompleted0)
				if err := _SoDiamond.contract.UnpackLog(event, "SoTransferCompleted0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSoTransferCompleted0 is a log parse operation binding the contract event 0x5272b980ac59723d5a8fe5be29daff5302abfaf057695289598e842dcf306e48.
//
// Solidity: event SoTransferCompleted(bytes32 indexed transactionId, uint256 receiveAmount)
func (_SoDiamond *SoDiamondFilterer) ParseSoTransferCompleted0(log types.Log) (*SoDiamondSoTransferCompleted0, error) {
	event := new(SoDiamondSoTransferCompleted0)
	if err := _SoDiamond.contract.UnpackLog(event, "SoTransferCompleted0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSoTransferCompleted1Iterator is returned from FilterSoTransferCompleted1 and is used to iterate over the raw logs and unpacked data for SoTransferCompleted1 events raised by the SoDiamond contract.
type SoDiamondSoTransferCompleted1Iterator struct {
	Event *SoDiamondSoTransferCompleted1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSoTransferCompleted1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSoTransferCompleted1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSoTransferCompleted1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSoTransferCompleted1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSoTransferCompleted1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSoTransferCompleted1 represents a SoTransferCompleted1 event raised by the SoDiamond contract.
type SoDiamondSoTransferCompleted1 struct {
	TransactionId [32]byte
	ReceiveAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSoTransferCompleted1 is a free log retrieval operation binding the contract event 0x5272b980ac59723d5a8fe5be29daff5302abfaf057695289598e842dcf306e48.
//
// Solidity: event SoTransferCompleted(bytes32 indexed transactionId, uint256 receiveAmount)
func (_SoDiamond *SoDiamondFilterer) FilterSoTransferCompleted1(opts *bind.FilterOpts, transactionId [][32]byte) (*SoDiamondSoTransferCompleted1Iterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SoTransferCompleted1", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondSoTransferCompleted1Iterator{contract: _SoDiamond.contract, event: "SoTransferCompleted1", logs: logs, sub: sub}, nil
}

// WatchSoTransferCompleted1 is a free log subscription operation binding the contract event 0x5272b980ac59723d5a8fe5be29daff5302abfaf057695289598e842dcf306e48.
//
// Solidity: event SoTransferCompleted(bytes32 indexed transactionId, uint256 receiveAmount)
func (_SoDiamond *SoDiamondFilterer) WatchSoTransferCompleted1(opts *bind.WatchOpts, sink chan<- *SoDiamondSoTransferCompleted1, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SoTransferCompleted1", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSoTransferCompleted1)
				if err := _SoDiamond.contract.UnpackLog(event, "SoTransferCompleted1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSoTransferCompleted1 is a log parse operation binding the contract event 0x5272b980ac59723d5a8fe5be29daff5302abfaf057695289598e842dcf306e48.
//
// Solidity: event SoTransferCompleted(bytes32 indexed transactionId, uint256 receiveAmount)
func (_SoDiamond *SoDiamondFilterer) ParseSoTransferCompleted1(log types.Log) (*SoDiamondSoTransferCompleted1, error) {
	event := new(SoDiamondSoTransferCompleted1)
	if err := _SoDiamond.contract.UnpackLog(event, "SoTransferCompleted1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSoTransferCompleted2Iterator is returned from FilterSoTransferCompleted2 and is used to iterate over the raw logs and unpacked data for SoTransferCompleted2 events raised by the SoDiamond contract.
type SoDiamondSoTransferCompleted2Iterator struct {
	Event *SoDiamondSoTransferCompleted2 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSoTransferCompleted2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSoTransferCompleted2)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSoTransferCompleted2)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSoTransferCompleted2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSoTransferCompleted2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSoTransferCompleted2 represents a SoTransferCompleted2 event raised by the SoDiamond contract.
type SoDiamondSoTransferCompleted2 struct {
	TransactionId [32]byte
	ReceiveAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSoTransferCompleted2 is a free log retrieval operation binding the contract event 0x5272b980ac59723d5a8fe5be29daff5302abfaf057695289598e842dcf306e48.
//
// Solidity: event SoTransferCompleted(bytes32 indexed transactionId, uint256 receiveAmount)
func (_SoDiamond *SoDiamondFilterer) FilterSoTransferCompleted2(opts *bind.FilterOpts, transactionId [][32]byte) (*SoDiamondSoTransferCompleted2Iterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SoTransferCompleted2", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondSoTransferCompleted2Iterator{contract: _SoDiamond.contract, event: "SoTransferCompleted2", logs: logs, sub: sub}, nil
}

// WatchSoTransferCompleted2 is a free log subscription operation binding the contract event 0x5272b980ac59723d5a8fe5be29daff5302abfaf057695289598e842dcf306e48.
//
// Solidity: event SoTransferCompleted(bytes32 indexed transactionId, uint256 receiveAmount)
func (_SoDiamond *SoDiamondFilterer) WatchSoTransferCompleted2(opts *bind.WatchOpts, sink chan<- *SoDiamondSoTransferCompleted2, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SoTransferCompleted2", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSoTransferCompleted2)
				if err := _SoDiamond.contract.UnpackLog(event, "SoTransferCompleted2", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSoTransferCompleted2 is a log parse operation binding the contract event 0x5272b980ac59723d5a8fe5be29daff5302abfaf057695289598e842dcf306e48.
//
// Solidity: event SoTransferCompleted(bytes32 indexed transactionId, uint256 receiveAmount)
func (_SoDiamond *SoDiamondFilterer) ParseSoTransferCompleted2(log types.Log) (*SoDiamondSoTransferCompleted2, error) {
	event := new(SoDiamondSoTransferCompleted2)
	if err := _SoDiamond.contract.UnpackLog(event, "SoTransferCompleted2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSoTransferCompleted3Iterator is returned from FilterSoTransferCompleted3 and is used to iterate over the raw logs and unpacked data for SoTransferCompleted3 events raised by the SoDiamond contract.
type SoDiamondSoTransferCompleted3Iterator struct {
	Event *SoDiamondSoTransferCompleted3 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSoTransferCompleted3Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSoTransferCompleted3)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSoTransferCompleted3)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSoTransferCompleted3Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSoTransferCompleted3Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSoTransferCompleted3 represents a SoTransferCompleted3 event raised by the SoDiamond contract.
type SoDiamondSoTransferCompleted3 struct {
	TransactionId [32]byte
	ReceiveAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSoTransferCompleted3 is a free log retrieval operation binding the contract event 0x5272b980ac59723d5a8fe5be29daff5302abfaf057695289598e842dcf306e48.
//
// Solidity: event SoTransferCompleted(bytes32 indexed transactionId, uint256 receiveAmount)
func (_SoDiamond *SoDiamondFilterer) FilterSoTransferCompleted3(opts *bind.FilterOpts, transactionId [][32]byte) (*SoDiamondSoTransferCompleted3Iterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SoTransferCompleted3", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondSoTransferCompleted3Iterator{contract: _SoDiamond.contract, event: "SoTransferCompleted3", logs: logs, sub: sub}, nil
}

// WatchSoTransferCompleted3 is a free log subscription operation binding the contract event 0x5272b980ac59723d5a8fe5be29daff5302abfaf057695289598e842dcf306e48.
//
// Solidity: event SoTransferCompleted(bytes32 indexed transactionId, uint256 receiveAmount)
func (_SoDiamond *SoDiamondFilterer) WatchSoTransferCompleted3(opts *bind.WatchOpts, sink chan<- *SoDiamondSoTransferCompleted3, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SoTransferCompleted3", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSoTransferCompleted3)
				if err := _SoDiamond.contract.UnpackLog(event, "SoTransferCompleted3", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSoTransferCompleted3 is a log parse operation binding the contract event 0x5272b980ac59723d5a8fe5be29daff5302abfaf057695289598e842dcf306e48.
//
// Solidity: event SoTransferCompleted(bytes32 indexed transactionId, uint256 receiveAmount)
func (_SoDiamond *SoDiamondFilterer) ParseSoTransferCompleted3(log types.Log) (*SoDiamondSoTransferCompleted3, error) {
	event := new(SoDiamondSoTransferCompleted3)
	if err := _SoDiamond.contract.UnpackLog(event, "SoTransferCompleted3", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSoTransferCompleted4Iterator is returned from FilterSoTransferCompleted4 and is used to iterate over the raw logs and unpacked data for SoTransferCompleted4 events raised by the SoDiamond contract.
type SoDiamondSoTransferCompleted4Iterator struct {
	Event *SoDiamondSoTransferCompleted4 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSoTransferCompleted4Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSoTransferCompleted4)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSoTransferCompleted4)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSoTransferCompleted4Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSoTransferCompleted4Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSoTransferCompleted4 represents a SoTransferCompleted4 event raised by the SoDiamond contract.
type SoDiamondSoTransferCompleted4 struct {
	TransactionId [32]byte
	ReceiveAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSoTransferCompleted4 is a free log retrieval operation binding the contract event 0x5272b980ac59723d5a8fe5be29daff5302abfaf057695289598e842dcf306e48.
//
// Solidity: event SoTransferCompleted(bytes32 indexed transactionId, uint256 receiveAmount)
func (_SoDiamond *SoDiamondFilterer) FilterSoTransferCompleted4(opts *bind.FilterOpts, transactionId [][32]byte) (*SoDiamondSoTransferCompleted4Iterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SoTransferCompleted4", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondSoTransferCompleted4Iterator{contract: _SoDiamond.contract, event: "SoTransferCompleted4", logs: logs, sub: sub}, nil
}

// WatchSoTransferCompleted4 is a free log subscription operation binding the contract event 0x5272b980ac59723d5a8fe5be29daff5302abfaf057695289598e842dcf306e48.
//
// Solidity: event SoTransferCompleted(bytes32 indexed transactionId, uint256 receiveAmount)
func (_SoDiamond *SoDiamondFilterer) WatchSoTransferCompleted4(opts *bind.WatchOpts, sink chan<- *SoDiamondSoTransferCompleted4, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SoTransferCompleted4", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSoTransferCompleted4)
				if err := _SoDiamond.contract.UnpackLog(event, "SoTransferCompleted4", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSoTransferCompleted4 is a log parse operation binding the contract event 0x5272b980ac59723d5a8fe5be29daff5302abfaf057695289598e842dcf306e48.
//
// Solidity: event SoTransferCompleted(bytes32 indexed transactionId, uint256 receiveAmount)
func (_SoDiamond *SoDiamondFilterer) ParseSoTransferCompleted4(log types.Log) (*SoDiamondSoTransferCompleted4, error) {
	event := new(SoDiamondSoTransferCompleted4)
	if err := _SoDiamond.contract.UnpackLog(event, "SoTransferCompleted4", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSoTransferFailedIterator is returned from FilterSoTransferFailed and is used to iterate over the raw logs and unpacked data for SoTransferFailed events raised by the SoDiamond contract.
type SoDiamondSoTransferFailedIterator struct {
	Event *SoDiamondSoTransferFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSoTransferFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSoTransferFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSoTransferFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSoTransferFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSoTransferFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSoTransferFailed represents a SoTransferFailed event raised by the SoDiamond contract.
type SoDiamondSoTransferFailed struct {
	TransactionId [32]byte
	RevertReason  string
	OtherReason   []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSoTransferFailed is a free log retrieval operation binding the contract event 0x9f22c9d1796172ce7238087f7ac46b639876da8736bb55f0d957282f6cccd028.
//
// Solidity: event SoTransferFailed(bytes32 indexed transactionId, string revertReason, bytes otherReason)
func (_SoDiamond *SoDiamondFilterer) FilterSoTransferFailed(opts *bind.FilterOpts, transactionId [][32]byte) (*SoDiamondSoTransferFailedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SoTransferFailed", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondSoTransferFailedIterator{contract: _SoDiamond.contract, event: "SoTransferFailed", logs: logs, sub: sub}, nil
}

// WatchSoTransferFailed is a free log subscription operation binding the contract event 0x9f22c9d1796172ce7238087f7ac46b639876da8736bb55f0d957282f6cccd028.
//
// Solidity: event SoTransferFailed(bytes32 indexed transactionId, string revertReason, bytes otherReason)
func (_SoDiamond *SoDiamondFilterer) WatchSoTransferFailed(opts *bind.WatchOpts, sink chan<- *SoDiamondSoTransferFailed, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SoTransferFailed", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSoTransferFailed)
				if err := _SoDiamond.contract.UnpackLog(event, "SoTransferFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSoTransferFailed is a log parse operation binding the contract event 0x9f22c9d1796172ce7238087f7ac46b639876da8736bb55f0d957282f6cccd028.
//
// Solidity: event SoTransferFailed(bytes32 indexed transactionId, string revertReason, bytes otherReason)
func (_SoDiamond *SoDiamondFilterer) ParseSoTransferFailed(log types.Log) (*SoDiamondSoTransferFailed, error) {
	event := new(SoDiamondSoTransferFailed)
	if err := _SoDiamond.contract.UnpackLog(event, "SoTransferFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSoTransferFailed0Iterator is returned from FilterSoTransferFailed0 and is used to iterate over the raw logs and unpacked data for SoTransferFailed0 events raised by the SoDiamond contract.
type SoDiamondSoTransferFailed0Iterator struct {
	Event *SoDiamondSoTransferFailed0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSoTransferFailed0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSoTransferFailed0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSoTransferFailed0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSoTransferFailed0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSoTransferFailed0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSoTransferFailed0 represents a SoTransferFailed0 event raised by the SoDiamond contract.
type SoDiamondSoTransferFailed0 struct {
	TransactionId [32]byte
	RevertReason  string
	OtherReason   []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSoTransferFailed0 is a free log retrieval operation binding the contract event 0x9f22c9d1796172ce7238087f7ac46b639876da8736bb55f0d957282f6cccd028.
//
// Solidity: event SoTransferFailed(bytes32 indexed transactionId, string revertReason, bytes otherReason)
func (_SoDiamond *SoDiamondFilterer) FilterSoTransferFailed0(opts *bind.FilterOpts, transactionId [][32]byte) (*SoDiamondSoTransferFailed0Iterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SoTransferFailed0", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondSoTransferFailed0Iterator{contract: _SoDiamond.contract, event: "SoTransferFailed0", logs: logs, sub: sub}, nil
}

// WatchSoTransferFailed0 is a free log subscription operation binding the contract event 0x9f22c9d1796172ce7238087f7ac46b639876da8736bb55f0d957282f6cccd028.
//
// Solidity: event SoTransferFailed(bytes32 indexed transactionId, string revertReason, bytes otherReason)
func (_SoDiamond *SoDiamondFilterer) WatchSoTransferFailed0(opts *bind.WatchOpts, sink chan<- *SoDiamondSoTransferFailed0, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SoTransferFailed0", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSoTransferFailed0)
				if err := _SoDiamond.contract.UnpackLog(event, "SoTransferFailed0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSoTransferFailed0 is a log parse operation binding the contract event 0x9f22c9d1796172ce7238087f7ac46b639876da8736bb55f0d957282f6cccd028.
//
// Solidity: event SoTransferFailed(bytes32 indexed transactionId, string revertReason, bytes otherReason)
func (_SoDiamond *SoDiamondFilterer) ParseSoTransferFailed0(log types.Log) (*SoDiamondSoTransferFailed0, error) {
	event := new(SoDiamondSoTransferFailed0)
	if err := _SoDiamond.contract.UnpackLog(event, "SoTransferFailed0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSoTransferFailed1Iterator is returned from FilterSoTransferFailed1 and is used to iterate over the raw logs and unpacked data for SoTransferFailed1 events raised by the SoDiamond contract.
type SoDiamondSoTransferFailed1Iterator struct {
	Event *SoDiamondSoTransferFailed1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSoTransferFailed1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSoTransferFailed1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSoTransferFailed1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSoTransferFailed1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSoTransferFailed1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSoTransferFailed1 represents a SoTransferFailed1 event raised by the SoDiamond contract.
type SoDiamondSoTransferFailed1 struct {
	TransactionId [32]byte
	RevertReason  string
	OtherReason   []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSoTransferFailed1 is a free log retrieval operation binding the contract event 0x9f22c9d1796172ce7238087f7ac46b639876da8736bb55f0d957282f6cccd028.
//
// Solidity: event SoTransferFailed(bytes32 indexed transactionId, string revertReason, bytes otherReason)
func (_SoDiamond *SoDiamondFilterer) FilterSoTransferFailed1(opts *bind.FilterOpts, transactionId [][32]byte) (*SoDiamondSoTransferFailed1Iterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SoTransferFailed1", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondSoTransferFailed1Iterator{contract: _SoDiamond.contract, event: "SoTransferFailed1", logs: logs, sub: sub}, nil
}

// WatchSoTransferFailed1 is a free log subscription operation binding the contract event 0x9f22c9d1796172ce7238087f7ac46b639876da8736bb55f0d957282f6cccd028.
//
// Solidity: event SoTransferFailed(bytes32 indexed transactionId, string revertReason, bytes otherReason)
func (_SoDiamond *SoDiamondFilterer) WatchSoTransferFailed1(opts *bind.WatchOpts, sink chan<- *SoDiamondSoTransferFailed1, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SoTransferFailed1", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSoTransferFailed1)
				if err := _SoDiamond.contract.UnpackLog(event, "SoTransferFailed1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSoTransferFailed1 is a log parse operation binding the contract event 0x9f22c9d1796172ce7238087f7ac46b639876da8736bb55f0d957282f6cccd028.
//
// Solidity: event SoTransferFailed(bytes32 indexed transactionId, string revertReason, bytes otherReason)
func (_SoDiamond *SoDiamondFilterer) ParseSoTransferFailed1(log types.Log) (*SoDiamondSoTransferFailed1, error) {
	event := new(SoDiamondSoTransferFailed1)
	if err := _SoDiamond.contract.UnpackLog(event, "SoTransferFailed1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSoTransferFailed2Iterator is returned from FilterSoTransferFailed2 and is used to iterate over the raw logs and unpacked data for SoTransferFailed2 events raised by the SoDiamond contract.
type SoDiamondSoTransferFailed2Iterator struct {
	Event *SoDiamondSoTransferFailed2 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSoTransferFailed2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSoTransferFailed2)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSoTransferFailed2)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSoTransferFailed2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSoTransferFailed2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSoTransferFailed2 represents a SoTransferFailed2 event raised by the SoDiamond contract.
type SoDiamondSoTransferFailed2 struct {
	TransactionId [32]byte
	RevertReason  string
	OtherReason   []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSoTransferFailed2 is a free log retrieval operation binding the contract event 0x9f22c9d1796172ce7238087f7ac46b639876da8736bb55f0d957282f6cccd028.
//
// Solidity: event SoTransferFailed(bytes32 indexed transactionId, string revertReason, bytes otherReason)
func (_SoDiamond *SoDiamondFilterer) FilterSoTransferFailed2(opts *bind.FilterOpts, transactionId [][32]byte) (*SoDiamondSoTransferFailed2Iterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SoTransferFailed2", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondSoTransferFailed2Iterator{contract: _SoDiamond.contract, event: "SoTransferFailed2", logs: logs, sub: sub}, nil
}

// WatchSoTransferFailed2 is a free log subscription operation binding the contract event 0x9f22c9d1796172ce7238087f7ac46b639876da8736bb55f0d957282f6cccd028.
//
// Solidity: event SoTransferFailed(bytes32 indexed transactionId, string revertReason, bytes otherReason)
func (_SoDiamond *SoDiamondFilterer) WatchSoTransferFailed2(opts *bind.WatchOpts, sink chan<- *SoDiamondSoTransferFailed2, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SoTransferFailed2", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSoTransferFailed2)
				if err := _SoDiamond.contract.UnpackLog(event, "SoTransferFailed2", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSoTransferFailed2 is a log parse operation binding the contract event 0x9f22c9d1796172ce7238087f7ac46b639876da8736bb55f0d957282f6cccd028.
//
// Solidity: event SoTransferFailed(bytes32 indexed transactionId, string revertReason, bytes otherReason)
func (_SoDiamond *SoDiamondFilterer) ParseSoTransferFailed2(log types.Log) (*SoDiamondSoTransferFailed2, error) {
	event := new(SoDiamondSoTransferFailed2)
	if err := _SoDiamond.contract.UnpackLog(event, "SoTransferFailed2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSoTransferFailed3Iterator is returned from FilterSoTransferFailed3 and is used to iterate over the raw logs and unpacked data for SoTransferFailed3 events raised by the SoDiamond contract.
type SoDiamondSoTransferFailed3Iterator struct {
	Event *SoDiamondSoTransferFailed3 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSoTransferFailed3Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSoTransferFailed3)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSoTransferFailed3)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSoTransferFailed3Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSoTransferFailed3Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSoTransferFailed3 represents a SoTransferFailed3 event raised by the SoDiamond contract.
type SoDiamondSoTransferFailed3 struct {
	TransactionId [32]byte
	RevertReason  string
	OtherReason   []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSoTransferFailed3 is a free log retrieval operation binding the contract event 0x9f22c9d1796172ce7238087f7ac46b639876da8736bb55f0d957282f6cccd028.
//
// Solidity: event SoTransferFailed(bytes32 indexed transactionId, string revertReason, bytes otherReason)
func (_SoDiamond *SoDiamondFilterer) FilterSoTransferFailed3(opts *bind.FilterOpts, transactionId [][32]byte) (*SoDiamondSoTransferFailed3Iterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SoTransferFailed3", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondSoTransferFailed3Iterator{contract: _SoDiamond.contract, event: "SoTransferFailed3", logs: logs, sub: sub}, nil
}

// WatchSoTransferFailed3 is a free log subscription operation binding the contract event 0x9f22c9d1796172ce7238087f7ac46b639876da8736bb55f0d957282f6cccd028.
//
// Solidity: event SoTransferFailed(bytes32 indexed transactionId, string revertReason, bytes otherReason)
func (_SoDiamond *SoDiamondFilterer) WatchSoTransferFailed3(opts *bind.WatchOpts, sink chan<- *SoDiamondSoTransferFailed3, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SoTransferFailed3", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSoTransferFailed3)
				if err := _SoDiamond.contract.UnpackLog(event, "SoTransferFailed3", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSoTransferFailed3 is a log parse operation binding the contract event 0x9f22c9d1796172ce7238087f7ac46b639876da8736bb55f0d957282f6cccd028.
//
// Solidity: event SoTransferFailed(bytes32 indexed transactionId, string revertReason, bytes otherReason)
func (_SoDiamond *SoDiamondFilterer) ParseSoTransferFailed3(log types.Log) (*SoDiamondSoTransferFailed3, error) {
	event := new(SoDiamondSoTransferFailed3)
	if err := _SoDiamond.contract.UnpackLog(event, "SoTransferFailed3", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSoTransferFailed4Iterator is returned from FilterSoTransferFailed4 and is used to iterate over the raw logs and unpacked data for SoTransferFailed4 events raised by the SoDiamond contract.
type SoDiamondSoTransferFailed4Iterator struct {
	Event *SoDiamondSoTransferFailed4 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSoTransferFailed4Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSoTransferFailed4)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSoTransferFailed4)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSoTransferFailed4Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSoTransferFailed4Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSoTransferFailed4 represents a SoTransferFailed4 event raised by the SoDiamond contract.
type SoDiamondSoTransferFailed4 struct {
	TransactionId [32]byte
	RevertReason  string
	OtherReason   []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSoTransferFailed4 is a free log retrieval operation binding the contract event 0x9f22c9d1796172ce7238087f7ac46b639876da8736bb55f0d957282f6cccd028.
//
// Solidity: event SoTransferFailed(bytes32 indexed transactionId, string revertReason, bytes otherReason)
func (_SoDiamond *SoDiamondFilterer) FilterSoTransferFailed4(opts *bind.FilterOpts, transactionId [][32]byte) (*SoDiamondSoTransferFailed4Iterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SoTransferFailed4", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondSoTransferFailed4Iterator{contract: _SoDiamond.contract, event: "SoTransferFailed4", logs: logs, sub: sub}, nil
}

// WatchSoTransferFailed4 is a free log subscription operation binding the contract event 0x9f22c9d1796172ce7238087f7ac46b639876da8736bb55f0d957282f6cccd028.
//
// Solidity: event SoTransferFailed(bytes32 indexed transactionId, string revertReason, bytes otherReason)
func (_SoDiamond *SoDiamondFilterer) WatchSoTransferFailed4(opts *bind.WatchOpts, sink chan<- *SoDiamondSoTransferFailed4, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SoTransferFailed4", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSoTransferFailed4)
				if err := _SoDiamond.contract.UnpackLog(event, "SoTransferFailed4", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSoTransferFailed4 is a log parse operation binding the contract event 0x9f22c9d1796172ce7238087f7ac46b639876da8736bb55f0d957282f6cccd028.
//
// Solidity: event SoTransferFailed(bytes32 indexed transactionId, string revertReason, bytes otherReason)
func (_SoDiamond *SoDiamondFilterer) ParseSoTransferFailed4(log types.Log) (*SoDiamondSoTransferFailed4, error) {
	event := new(SoDiamondSoTransferFailed4)
	if err := _SoDiamond.contract.UnpackLog(event, "SoTransferFailed4", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSoTransferStartedIterator is returned from FilterSoTransferStarted and is used to iterate over the raw logs and unpacked data for SoTransferStarted events raised by the SoDiamond contract.
type SoDiamondSoTransferStartedIterator struct {
	Event *SoDiamondSoTransferStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSoTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSoTransferStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSoTransferStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSoTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSoTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSoTransferStarted represents a SoTransferStarted event raised by the SoDiamond contract.
type SoDiamondSoTransferStarted struct {
	TransactionId [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSoTransferStarted is a free log retrieval operation binding the contract event 0x0f0fd0ad174232a46f92a8d76b425830f45436483106ee965bbe91d3b7d1cd26.
//
// Solidity: event SoTransferStarted(bytes32 indexed transactionId)
func (_SoDiamond *SoDiamondFilterer) FilterSoTransferStarted(opts *bind.FilterOpts, transactionId [][32]byte) (*SoDiamondSoTransferStartedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SoTransferStarted", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondSoTransferStartedIterator{contract: _SoDiamond.contract, event: "SoTransferStarted", logs: logs, sub: sub}, nil
}

// WatchSoTransferStarted is a free log subscription operation binding the contract event 0x0f0fd0ad174232a46f92a8d76b425830f45436483106ee965bbe91d3b7d1cd26.
//
// Solidity: event SoTransferStarted(bytes32 indexed transactionId)
func (_SoDiamond *SoDiamondFilterer) WatchSoTransferStarted(opts *bind.WatchOpts, sink chan<- *SoDiamondSoTransferStarted, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SoTransferStarted", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSoTransferStarted)
				if err := _SoDiamond.contract.UnpackLog(event, "SoTransferStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSoTransferStarted is a log parse operation binding the contract event 0x0f0fd0ad174232a46f92a8d76b425830f45436483106ee965bbe91d3b7d1cd26.
//
// Solidity: event SoTransferStarted(bytes32 indexed transactionId)
func (_SoDiamond *SoDiamondFilterer) ParseSoTransferStarted(log types.Log) (*SoDiamondSoTransferStarted, error) {
	event := new(SoDiamondSoTransferStarted)
	if err := _SoDiamond.contract.UnpackLog(event, "SoTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSoTransferStarted0Iterator is returned from FilterSoTransferStarted0 and is used to iterate over the raw logs and unpacked data for SoTransferStarted0 events raised by the SoDiamond contract.
type SoDiamondSoTransferStarted0Iterator struct {
	Event *SoDiamondSoTransferStarted0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSoTransferStarted0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSoTransferStarted0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSoTransferStarted0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSoTransferStarted0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSoTransferStarted0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSoTransferStarted0 represents a SoTransferStarted0 event raised by the SoDiamond contract.
type SoDiamondSoTransferStarted0 struct {
	TransactionId [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSoTransferStarted0 is a free log retrieval operation binding the contract event 0x0f0fd0ad174232a46f92a8d76b425830f45436483106ee965bbe91d3b7d1cd26.
//
// Solidity: event SoTransferStarted(bytes32 indexed transactionId)
func (_SoDiamond *SoDiamondFilterer) FilterSoTransferStarted0(opts *bind.FilterOpts, transactionId [][32]byte) (*SoDiamondSoTransferStarted0Iterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SoTransferStarted0", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondSoTransferStarted0Iterator{contract: _SoDiamond.contract, event: "SoTransferStarted0", logs: logs, sub: sub}, nil
}

// WatchSoTransferStarted0 is a free log subscription operation binding the contract event 0x0f0fd0ad174232a46f92a8d76b425830f45436483106ee965bbe91d3b7d1cd26.
//
// Solidity: event SoTransferStarted(bytes32 indexed transactionId)
func (_SoDiamond *SoDiamondFilterer) WatchSoTransferStarted0(opts *bind.WatchOpts, sink chan<- *SoDiamondSoTransferStarted0, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SoTransferStarted0", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSoTransferStarted0)
				if err := _SoDiamond.contract.UnpackLog(event, "SoTransferStarted0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSoTransferStarted0 is a log parse operation binding the contract event 0x0f0fd0ad174232a46f92a8d76b425830f45436483106ee965bbe91d3b7d1cd26.
//
// Solidity: event SoTransferStarted(bytes32 indexed transactionId)
func (_SoDiamond *SoDiamondFilterer) ParseSoTransferStarted0(log types.Log) (*SoDiamondSoTransferStarted0, error) {
	event := new(SoDiamondSoTransferStarted0)
	if err := _SoDiamond.contract.UnpackLog(event, "SoTransferStarted0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSoTransferStarted1Iterator is returned from FilterSoTransferStarted1 and is used to iterate over the raw logs and unpacked data for SoTransferStarted1 events raised by the SoDiamond contract.
type SoDiamondSoTransferStarted1Iterator struct {
	Event *SoDiamondSoTransferStarted1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSoTransferStarted1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSoTransferStarted1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSoTransferStarted1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSoTransferStarted1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSoTransferStarted1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSoTransferStarted1 represents a SoTransferStarted1 event raised by the SoDiamond contract.
type SoDiamondSoTransferStarted1 struct {
	TransactionId [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSoTransferStarted1 is a free log retrieval operation binding the contract event 0x0f0fd0ad174232a46f92a8d76b425830f45436483106ee965bbe91d3b7d1cd26.
//
// Solidity: event SoTransferStarted(bytes32 indexed transactionId)
func (_SoDiamond *SoDiamondFilterer) FilterSoTransferStarted1(opts *bind.FilterOpts, transactionId [][32]byte) (*SoDiamondSoTransferStarted1Iterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SoTransferStarted1", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondSoTransferStarted1Iterator{contract: _SoDiamond.contract, event: "SoTransferStarted1", logs: logs, sub: sub}, nil
}

// WatchSoTransferStarted1 is a free log subscription operation binding the contract event 0x0f0fd0ad174232a46f92a8d76b425830f45436483106ee965bbe91d3b7d1cd26.
//
// Solidity: event SoTransferStarted(bytes32 indexed transactionId)
func (_SoDiamond *SoDiamondFilterer) WatchSoTransferStarted1(opts *bind.WatchOpts, sink chan<- *SoDiamondSoTransferStarted1, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SoTransferStarted1", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSoTransferStarted1)
				if err := _SoDiamond.contract.UnpackLog(event, "SoTransferStarted1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSoTransferStarted1 is a log parse operation binding the contract event 0x0f0fd0ad174232a46f92a8d76b425830f45436483106ee965bbe91d3b7d1cd26.
//
// Solidity: event SoTransferStarted(bytes32 indexed transactionId)
func (_SoDiamond *SoDiamondFilterer) ParseSoTransferStarted1(log types.Log) (*SoDiamondSoTransferStarted1, error) {
	event := new(SoDiamondSoTransferStarted1)
	if err := _SoDiamond.contract.UnpackLog(event, "SoTransferStarted1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSoTransferStarted2Iterator is returned from FilterSoTransferStarted2 and is used to iterate over the raw logs and unpacked data for SoTransferStarted2 events raised by the SoDiamond contract.
type SoDiamondSoTransferStarted2Iterator struct {
	Event *SoDiamondSoTransferStarted2 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSoTransferStarted2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSoTransferStarted2)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSoTransferStarted2)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSoTransferStarted2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSoTransferStarted2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSoTransferStarted2 represents a SoTransferStarted2 event raised by the SoDiamond contract.
type SoDiamondSoTransferStarted2 struct {
	TransactionId [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSoTransferStarted2 is a free log retrieval operation binding the contract event 0x0f0fd0ad174232a46f92a8d76b425830f45436483106ee965bbe91d3b7d1cd26.
//
// Solidity: event SoTransferStarted(bytes32 indexed transactionId)
func (_SoDiamond *SoDiamondFilterer) FilterSoTransferStarted2(opts *bind.FilterOpts, transactionId [][32]byte) (*SoDiamondSoTransferStarted2Iterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SoTransferStarted2", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondSoTransferStarted2Iterator{contract: _SoDiamond.contract, event: "SoTransferStarted2", logs: logs, sub: sub}, nil
}

// WatchSoTransferStarted2 is a free log subscription operation binding the contract event 0x0f0fd0ad174232a46f92a8d76b425830f45436483106ee965bbe91d3b7d1cd26.
//
// Solidity: event SoTransferStarted(bytes32 indexed transactionId)
func (_SoDiamond *SoDiamondFilterer) WatchSoTransferStarted2(opts *bind.WatchOpts, sink chan<- *SoDiamondSoTransferStarted2, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SoTransferStarted2", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSoTransferStarted2)
				if err := _SoDiamond.contract.UnpackLog(event, "SoTransferStarted2", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSoTransferStarted2 is a log parse operation binding the contract event 0x0f0fd0ad174232a46f92a8d76b425830f45436483106ee965bbe91d3b7d1cd26.
//
// Solidity: event SoTransferStarted(bytes32 indexed transactionId)
func (_SoDiamond *SoDiamondFilterer) ParseSoTransferStarted2(log types.Log) (*SoDiamondSoTransferStarted2, error) {
	event := new(SoDiamondSoTransferStarted2)
	if err := _SoDiamond.contract.UnpackLog(event, "SoTransferStarted2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSoTransferStarted3Iterator is returned from FilterSoTransferStarted3 and is used to iterate over the raw logs and unpacked data for SoTransferStarted3 events raised by the SoDiamond contract.
type SoDiamondSoTransferStarted3Iterator struct {
	Event *SoDiamondSoTransferStarted3 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSoTransferStarted3Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSoTransferStarted3)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSoTransferStarted3)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSoTransferStarted3Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSoTransferStarted3Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSoTransferStarted3 represents a SoTransferStarted3 event raised by the SoDiamond contract.
type SoDiamondSoTransferStarted3 struct {
	TransactionId [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSoTransferStarted3 is a free log retrieval operation binding the contract event 0x0f0fd0ad174232a46f92a8d76b425830f45436483106ee965bbe91d3b7d1cd26.
//
// Solidity: event SoTransferStarted(bytes32 indexed transactionId)
func (_SoDiamond *SoDiamondFilterer) FilterSoTransferStarted3(opts *bind.FilterOpts, transactionId [][32]byte) (*SoDiamondSoTransferStarted3Iterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SoTransferStarted3", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondSoTransferStarted3Iterator{contract: _SoDiamond.contract, event: "SoTransferStarted3", logs: logs, sub: sub}, nil
}

// WatchSoTransferStarted3 is a free log subscription operation binding the contract event 0x0f0fd0ad174232a46f92a8d76b425830f45436483106ee965bbe91d3b7d1cd26.
//
// Solidity: event SoTransferStarted(bytes32 indexed transactionId)
func (_SoDiamond *SoDiamondFilterer) WatchSoTransferStarted3(opts *bind.WatchOpts, sink chan<- *SoDiamondSoTransferStarted3, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SoTransferStarted3", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSoTransferStarted3)
				if err := _SoDiamond.contract.UnpackLog(event, "SoTransferStarted3", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSoTransferStarted3 is a log parse operation binding the contract event 0x0f0fd0ad174232a46f92a8d76b425830f45436483106ee965bbe91d3b7d1cd26.
//
// Solidity: event SoTransferStarted(bytes32 indexed transactionId)
func (_SoDiamond *SoDiamondFilterer) ParseSoTransferStarted3(log types.Log) (*SoDiamondSoTransferStarted3, error) {
	event := new(SoDiamondSoTransferStarted3)
	if err := _SoDiamond.contract.UnpackLog(event, "SoTransferStarted3", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondSoTransferStarted4Iterator is returned from FilterSoTransferStarted4 and is used to iterate over the raw logs and unpacked data for SoTransferStarted4 events raised by the SoDiamond contract.
type SoDiamondSoTransferStarted4Iterator struct {
	Event *SoDiamondSoTransferStarted4 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondSoTransferStarted4Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondSoTransferStarted4)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondSoTransferStarted4)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondSoTransferStarted4Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondSoTransferStarted4Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondSoTransferStarted4 represents a SoTransferStarted4 event raised by the SoDiamond contract.
type SoDiamondSoTransferStarted4 struct {
	TransactionId [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSoTransferStarted4 is a free log retrieval operation binding the contract event 0x0f0fd0ad174232a46f92a8d76b425830f45436483106ee965bbe91d3b7d1cd26.
//
// Solidity: event SoTransferStarted(bytes32 indexed transactionId)
func (_SoDiamond *SoDiamondFilterer) FilterSoTransferStarted4(opts *bind.FilterOpts, transactionId [][32]byte) (*SoDiamondSoTransferStarted4Iterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "SoTransferStarted4", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondSoTransferStarted4Iterator{contract: _SoDiamond.contract, event: "SoTransferStarted4", logs: logs, sub: sub}, nil
}

// WatchSoTransferStarted4 is a free log subscription operation binding the contract event 0x0f0fd0ad174232a46f92a8d76b425830f45436483106ee965bbe91d3b7d1cd26.
//
// Solidity: event SoTransferStarted(bytes32 indexed transactionId)
func (_SoDiamond *SoDiamondFilterer) WatchSoTransferStarted4(opts *bind.WatchOpts, sink chan<- *SoDiamondSoTransferStarted4, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "SoTransferStarted4", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondSoTransferStarted4)
				if err := _SoDiamond.contract.UnpackLog(event, "SoTransferStarted4", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSoTransferStarted4 is a log parse operation binding the contract event 0x0f0fd0ad174232a46f92a8d76b425830f45436483106ee965bbe91d3b7d1cd26.
//
// Solidity: event SoTransferStarted(bytes32 indexed transactionId)
func (_SoDiamond *SoDiamondFilterer) ParseSoTransferStarted4(log types.Log) (*SoDiamondSoTransferStarted4, error) {
	event := new(SoDiamondSoTransferStarted4)
	if err := _SoDiamond.contract.UnpackLog(event, "SoTransferStarted4", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondStargateInitializedIterator is returned from FilterStargateInitialized and is used to iterate over the raw logs and unpacked data for StargateInitialized events raised by the SoDiamond contract.
type SoDiamondStargateInitializedIterator struct {
	Event *SoDiamondStargateInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondStargateInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondStargateInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondStargateInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondStargateInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondStargateInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondStargateInitialized represents a StargateInitialized event raised by the SoDiamond contract.
type SoDiamondStargateInitialized struct {
	Stargate common.Address
	ChainId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStargateInitialized is a free log retrieval operation binding the contract event 0xdd1aad6fd60873172ecfdee635f5174c3a11648dd0769c10569f27cd596e02e8.
//
// Solidity: event StargateInitialized(address stargate, uint256 chainId)
func (_SoDiamond *SoDiamondFilterer) FilterStargateInitialized(opts *bind.FilterOpts) (*SoDiamondStargateInitializedIterator, error) {

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "StargateInitialized")
	if err != nil {
		return nil, err
	}
	return &SoDiamondStargateInitializedIterator{contract: _SoDiamond.contract, event: "StargateInitialized", logs: logs, sub: sub}, nil
}

// WatchStargateInitialized is a free log subscription operation binding the contract event 0xdd1aad6fd60873172ecfdee635f5174c3a11648dd0769c10569f27cd596e02e8.
//
// Solidity: event StargateInitialized(address stargate, uint256 chainId)
func (_SoDiamond *SoDiamondFilterer) WatchStargateInitialized(opts *bind.WatchOpts, sink chan<- *SoDiamondStargateInitialized) (event.Subscription, error) {

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "StargateInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondStargateInitialized)
				if err := _SoDiamond.contract.UnpackLog(event, "StargateInitialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStargateInitialized is a log parse operation binding the contract event 0xdd1aad6fd60873172ecfdee635f5174c3a11648dd0769c10569f27cd596e02e8.
//
// Solidity: event StargateInitialized(address stargate, uint256 chainId)
func (_SoDiamond *SoDiamondFilterer) ParseStargateInitialized(log types.Log) (*SoDiamondStargateInitialized, error) {
	event := new(SoDiamondStargateInitialized)
	if err := _SoDiamond.contract.UnpackLog(event, "StargateInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondTransferFromCelerIterator is returned from FilterTransferFromCeler and is used to iterate over the raw logs and unpacked data for TransferFromCeler events raised by the SoDiamond contract.
type SoDiamondTransferFromCelerIterator struct {
	Event *SoDiamondTransferFromCeler // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondTransferFromCelerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondTransferFromCeler)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondTransferFromCeler)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondTransferFromCelerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondTransferFromCelerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondTransferFromCeler represents a TransferFromCeler event raised by the SoDiamond contract.
type SoDiamondTransferFromCeler struct {
	CelerTransferId [32]byte
	SrcCelerChainId uint64
	DstCelerChainId uint64
	BridgeToken     common.Address
	BridgeAmount    *big.Int
	Nonce           uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferFromCeler is a free log retrieval operation binding the contract event 0x5d534905be115fbdc86285fe2a524b6ad9ef41edce8f7b1ab44d39aaf2828058.
//
// Solidity: event TransferFromCeler(bytes32 indexed celerTransferId, uint64 srcCelerChainId, uint64 dstCelerChainId, address bridgeToken, uint256 bridgeAmount, uint64 nonce)
func (_SoDiamond *SoDiamondFilterer) FilterTransferFromCeler(opts *bind.FilterOpts, celerTransferId [][32]byte) (*SoDiamondTransferFromCelerIterator, error) {

	var celerTransferIdRule []interface{}
	for _, celerTransferIdItem := range celerTransferId {
		celerTransferIdRule = append(celerTransferIdRule, celerTransferIdItem)
	}

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "TransferFromCeler", celerTransferIdRule)
	if err != nil {
		return nil, err
	}
	return &SoDiamondTransferFromCelerIterator{contract: _SoDiamond.contract, event: "TransferFromCeler", logs: logs, sub: sub}, nil
}

// WatchTransferFromCeler is a free log subscription operation binding the contract event 0x5d534905be115fbdc86285fe2a524b6ad9ef41edce8f7b1ab44d39aaf2828058.
//
// Solidity: event TransferFromCeler(bytes32 indexed celerTransferId, uint64 srcCelerChainId, uint64 dstCelerChainId, address bridgeToken, uint256 bridgeAmount, uint64 nonce)
func (_SoDiamond *SoDiamondFilterer) WatchTransferFromCeler(opts *bind.WatchOpts, sink chan<- *SoDiamondTransferFromCeler, celerTransferId [][32]byte) (event.Subscription, error) {

	var celerTransferIdRule []interface{}
	for _, celerTransferIdItem := range celerTransferId {
		celerTransferIdRule = append(celerTransferIdRule, celerTransferIdItem)
	}

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "TransferFromCeler", celerTransferIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondTransferFromCeler)
				if err := _SoDiamond.contract.UnpackLog(event, "TransferFromCeler", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferFromCeler is a log parse operation binding the contract event 0x5d534905be115fbdc86285fe2a524b6ad9ef41edce8f7b1ab44d39aaf2828058.
//
// Solidity: event TransferFromCeler(bytes32 indexed celerTransferId, uint64 srcCelerChainId, uint64 dstCelerChainId, address bridgeToken, uint256 bridgeAmount, uint64 nonce)
func (_SoDiamond *SoDiamondFilterer) ParseTransferFromCeler(log types.Log) (*SoDiamondTransferFromCeler, error) {
	event := new(SoDiamondTransferFromCeler)
	if err := _SoDiamond.contract.UnpackLog(event, "TransferFromCeler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondTransferFromWormholeIterator is returned from FilterTransferFromWormhole and is used to iterate over the raw logs and unpacked data for TransferFromWormhole events raised by the SoDiamond contract.
type SoDiamondTransferFromWormholeIterator struct {
	Event *SoDiamondTransferFromWormhole // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondTransferFromWormholeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondTransferFromWormhole)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondTransferFromWormhole)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondTransferFromWormholeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondTransferFromWormholeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondTransferFromWormhole represents a TransferFromWormhole event raised by the SoDiamond contract.
type SoDiamondTransferFromWormhole struct {
	SrcWormholeChainId uint16
	DstWormholeChainId uint16
	Sequence           uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTransferFromWormhole is a free log retrieval operation binding the contract event 0x54aac47748246a3ab5f96928b15860ffe57608cd8d74364ac38444a9093aa2ea.
//
// Solidity: event TransferFromWormhole(uint16 srcWormholeChainId, uint16 dstWormholeChainId, uint64 sequence)
func (_SoDiamond *SoDiamondFilterer) FilterTransferFromWormhole(opts *bind.FilterOpts) (*SoDiamondTransferFromWormholeIterator, error) {

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "TransferFromWormhole")
	if err != nil {
		return nil, err
	}
	return &SoDiamondTransferFromWormholeIterator{contract: _SoDiamond.contract, event: "TransferFromWormhole", logs: logs, sub: sub}, nil
}

// WatchTransferFromWormhole is a free log subscription operation binding the contract event 0x54aac47748246a3ab5f96928b15860ffe57608cd8d74364ac38444a9093aa2ea.
//
// Solidity: event TransferFromWormhole(uint16 srcWormholeChainId, uint16 dstWormholeChainId, uint64 sequence)
func (_SoDiamond *SoDiamondFilterer) WatchTransferFromWormhole(opts *bind.WatchOpts, sink chan<- *SoDiamondTransferFromWormhole) (event.Subscription, error) {

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "TransferFromWormhole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondTransferFromWormhole)
				if err := _SoDiamond.contract.UnpackLog(event, "TransferFromWormhole", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferFromWormhole is a log parse operation binding the contract event 0x54aac47748246a3ab5f96928b15860ffe57608cd8d74364ac38444a9093aa2ea.
//
// Solidity: event TransferFromWormhole(uint16 srcWormholeChainId, uint16 dstWormholeChainId, uint64 sequence)
func (_SoDiamond *SoDiamondFilterer) ParseTransferFromWormhole(log types.Log) (*SoDiamondTransferFromWormhole, error) {
	event := new(SoDiamondTransferFromWormhole)
	if err := _SoDiamond.contract.UnpackLog(event, "TransferFromWormhole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondUpdateCelerReserveIterator is returned from FilterUpdateCelerReserve and is used to iterate over the raw logs and unpacked data for UpdateCelerReserve events raised by the SoDiamond contract.
type SoDiamondUpdateCelerReserveIterator struct {
	Event *SoDiamondUpdateCelerReserve // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondUpdateCelerReserveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondUpdateCelerReserve)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondUpdateCelerReserve)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondUpdateCelerReserveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondUpdateCelerReserveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondUpdateCelerReserve represents a UpdateCelerReserve event raised by the SoDiamond contract.
type SoDiamondUpdateCelerReserve struct {
	ActualReserve   *big.Int
	EstimateReserve *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateCelerReserve is a free log retrieval operation binding the contract event 0xaed4361c2add4f4157489c992ce552e836b8846d4b7c7bcb107ddbda6e2f570a.
//
// Solidity: event UpdateCelerReserve(uint256 actualReserve, uint256 estimateReserve)
func (_SoDiamond *SoDiamondFilterer) FilterUpdateCelerReserve(opts *bind.FilterOpts) (*SoDiamondUpdateCelerReserveIterator, error) {

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "UpdateCelerReserve")
	if err != nil {
		return nil, err
	}
	return &SoDiamondUpdateCelerReserveIterator{contract: _SoDiamond.contract, event: "UpdateCelerReserve", logs: logs, sub: sub}, nil
}

// WatchUpdateCelerReserve is a free log subscription operation binding the contract event 0xaed4361c2add4f4157489c992ce552e836b8846d4b7c7bcb107ddbda6e2f570a.
//
// Solidity: event UpdateCelerReserve(uint256 actualReserve, uint256 estimateReserve)
func (_SoDiamond *SoDiamondFilterer) WatchUpdateCelerReserve(opts *bind.WatchOpts, sink chan<- *SoDiamondUpdateCelerReserve) (event.Subscription, error) {

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "UpdateCelerReserve")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondUpdateCelerReserve)
				if err := _SoDiamond.contract.UnpackLog(event, "UpdateCelerReserve", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateCelerReserve is a log parse operation binding the contract event 0xaed4361c2add4f4157489c992ce552e836b8846d4b7c7bcb107ddbda6e2f570a.
//
// Solidity: event UpdateCelerReserve(uint256 actualReserve, uint256 estimateReserve)
func (_SoDiamond *SoDiamondFilterer) ParseUpdateCelerReserve(log types.Log) (*SoDiamondUpdateCelerReserve, error) {
	event := new(SoDiamondUpdateCelerReserve)
	if err := _SoDiamond.contract.UnpackLog(event, "UpdateCelerReserve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondUpdateWormholeGasIterator is returned from FilterUpdateWormholeGas and is used to iterate over the raw logs and unpacked data for UpdateWormholeGas events raised by the SoDiamond contract.
type SoDiamondUpdateWormholeGasIterator struct {
	Event *SoDiamondUpdateWormholeGas // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondUpdateWormholeGasIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondUpdateWormholeGas)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondUpdateWormholeGas)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondUpdateWormholeGasIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondUpdateWormholeGasIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondUpdateWormholeGas represents a UpdateWormholeGas event raised by the SoDiamond contract.
type SoDiamondUpdateWormholeGas struct {
	DstWormholeChainId uint16
	BaseGas            *big.Int
	GasPerBytes        *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateWormholeGas is a free log retrieval operation binding the contract event 0x0aa61585b805fde53f9144e0b67e73fa5c022a2822052ae5803875f481b511d2.
//
// Solidity: event UpdateWormholeGas(uint16 dstWormholeChainId, uint256 baseGas, uint256 gasPerBytes)
func (_SoDiamond *SoDiamondFilterer) FilterUpdateWormholeGas(opts *bind.FilterOpts) (*SoDiamondUpdateWormholeGasIterator, error) {

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "UpdateWormholeGas")
	if err != nil {
		return nil, err
	}
	return &SoDiamondUpdateWormholeGasIterator{contract: _SoDiamond.contract, event: "UpdateWormholeGas", logs: logs, sub: sub}, nil
}

// WatchUpdateWormholeGas is a free log subscription operation binding the contract event 0x0aa61585b805fde53f9144e0b67e73fa5c022a2822052ae5803875f481b511d2.
//
// Solidity: event UpdateWormholeGas(uint16 dstWormholeChainId, uint256 baseGas, uint256 gasPerBytes)
func (_SoDiamond *SoDiamondFilterer) WatchUpdateWormholeGas(opts *bind.WatchOpts, sink chan<- *SoDiamondUpdateWormholeGas) (event.Subscription, error) {

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "UpdateWormholeGas")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondUpdateWormholeGas)
				if err := _SoDiamond.contract.UnpackLog(event, "UpdateWormholeGas", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateWormholeGas is a log parse operation binding the contract event 0x0aa61585b805fde53f9144e0b67e73fa5c022a2822052ae5803875f481b511d2.
//
// Solidity: event UpdateWormholeGas(uint16 dstWormholeChainId, uint256 baseGas, uint256 gasPerBytes)
func (_SoDiamond *SoDiamondFilterer) ParseUpdateWormholeGas(log types.Log) (*SoDiamondUpdateWormholeGas, error) {
	event := new(SoDiamondUpdateWormholeGas)
	if err := _SoDiamond.contract.UnpackLog(event, "UpdateWormholeGas", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoDiamondUpdateWormholeReserveIterator is returned from FilterUpdateWormholeReserve and is used to iterate over the raw logs and unpacked data for UpdateWormholeReserve events raised by the SoDiamond contract.
type SoDiamondUpdateWormholeReserveIterator struct {
	Event *SoDiamondUpdateWormholeReserve // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SoDiamondUpdateWormholeReserveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoDiamondUpdateWormholeReserve)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SoDiamondUpdateWormholeReserve)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SoDiamondUpdateWormholeReserveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoDiamondUpdateWormholeReserveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoDiamondUpdateWormholeReserve represents a UpdateWormholeReserve event raised by the SoDiamond contract.
type SoDiamondUpdateWormholeReserve struct {
	ActualReserve   *big.Int
	EstimateReserve *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateWormholeReserve is a free log retrieval operation binding the contract event 0x47eb9e20f6454ed3771e741f21d26f7f265ae9a1a06d722dd888e4585c0b7e95.
//
// Solidity: event UpdateWormholeReserve(uint256 actualReserve, uint256 estimateReserve)
func (_SoDiamond *SoDiamondFilterer) FilterUpdateWormholeReserve(opts *bind.FilterOpts) (*SoDiamondUpdateWormholeReserveIterator, error) {

	logs, sub, err := _SoDiamond.contract.FilterLogs(opts, "UpdateWormholeReserve")
	if err != nil {
		return nil, err
	}
	return &SoDiamondUpdateWormholeReserveIterator{contract: _SoDiamond.contract, event: "UpdateWormholeReserve", logs: logs, sub: sub}, nil
}

// WatchUpdateWormholeReserve is a free log subscription operation binding the contract event 0x47eb9e20f6454ed3771e741f21d26f7f265ae9a1a06d722dd888e4585c0b7e95.
//
// Solidity: event UpdateWormholeReserve(uint256 actualReserve, uint256 estimateReserve)
func (_SoDiamond *SoDiamondFilterer) WatchUpdateWormholeReserve(opts *bind.WatchOpts, sink chan<- *SoDiamondUpdateWormholeReserve) (event.Subscription, error) {

	logs, sub, err := _SoDiamond.contract.WatchLogs(opts, "UpdateWormholeReserve")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoDiamondUpdateWormholeReserve)
				if err := _SoDiamond.contract.UnpackLog(event, "UpdateWormholeReserve", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateWormholeReserve is a log parse operation binding the contract event 0x47eb9e20f6454ed3771e741f21d26f7f265ae9a1a06d722dd888e4585c0b7e95.
//
// Solidity: event UpdateWormholeReserve(uint256 actualReserve, uint256 estimateReserve)
func (_SoDiamond *SoDiamondFilterer) ParseUpdateWormholeReserve(log types.Log) (*SoDiamondUpdateWormholeReserve, error) {
	event := new(SoDiamondUpdateWormholeReserve)
	if err := _SoDiamond.contract.UnpackLog(event, "UpdateWormholeReserve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
