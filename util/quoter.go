package util

import (
	"errors"
	"fmt"
	"github.com/coming-chat/intra-swap-core/base_entities"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
	"reflect"
)

/**
 * Converts a route to a hex encoded path
 * @param route the v3 path to convert to an encoded path
 * @param exactOutput whether the route should be encoded in reverse, for making exact output swaps
 */
func EncodeRouteToPath(pools []base_entities.Pool, inputToken *entities.Token, exactOutput bool) ([]byte, error) {
	var (
		types []string
		path  []interface{}

		addressTy = "address"
		uint24Ty  = "uint24"
	)

	for i, pool := range pools {
		p, ok := pool.(*base_entities.V3Pool)
		if !ok {
			return nil, errors.New("EncodeRouteToPath only support v3 pools")
		}
		var outputToken *entities.Token
		if p.Token0().Equal(inputToken) {
			outputToken = p.Token1()
		} else {
			outputToken = p.Token0()
		}
		if i == 0 {
			types = []string{addressTy, uint24Ty, addressTy}
			path = []interface{}{inputToken.Address, uint64(p.Fee), outputToken.Address}
		} else {
			types = append(types, uint24Ty, addressTy)
			path = append(path, uint64(p.Fee), outputToken.Address)
		}
		inputToken = outputToken
	}

	if exactOutput {
		reverse(types)
		reverse(path)
	}

	// tight pack the path
	var packedPath [][]byte
	for i, t := range types {
		switch t {
		case addressTy:
			packedPath = append(packedPath, path[i].(common.Address).Bytes())
		case uint24Ty:
			packedPath = append(packedPath, common.LeftPadBytes(PutUint24(path[i].(uint64)), 24/8))
		default:
			return nil, fmt.Errorf("unknown type %s", t)
		}

	}

	var pack []byte
	for _, p := range packedPath {
		pack = append(pack, p...)
	}
	return pack, nil
}

// PutUint24 put bigendian uint24
func PutUint24(i uint64) []byte {
	b := make([]byte, 3)
	b[0] = byte(i >> 16)
	b[1] = byte(i >> 8)
	b[2] = byte(i)
	return b
}

func reverse(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
