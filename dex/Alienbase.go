package dex

type Alienbase struct {
	UniswapV2
}

func NewAlienbase() Alienbase {
	return Alienbase{
		UniswapV2: NewUniswapV2(nil, nil, nil),
	}
}
