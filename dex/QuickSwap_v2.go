package dex

type QuickSwapV2 struct {
	UniswapV2
}

func NewQuickSwapV2() QuickSwapV2 {
	return QuickSwapV2{
		UniswapV2: NewUniswapV2(nil, nil, nil),
	}
}
