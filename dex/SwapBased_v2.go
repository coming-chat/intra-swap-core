package dex

type SwapBasedV2 struct {
	UniswapV2
}

func NewSwapBasedV2() SwapBasedV2 {
	return SwapBasedV2{
		UniswapV2: NewUniswapV2(nil, nil, nil),
	}
}
