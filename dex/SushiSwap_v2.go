package dex

type SushiSwapV2 struct {
	UniswapV2
}

func NewSushiSwapV2() SushiSwapV2 {
	return SushiSwapV2{
		UniswapV2: NewUniswapV2(nil, nil, nil),
	}
}
