package dex

type BaseSwapV2 struct {
	UniswapV2
}

func NewBaseSwapV2() BaseSwapV2 {
	return BaseSwapV2{
		UniswapV2: NewUniswapV2(nil, nil, nil),
	}
}
