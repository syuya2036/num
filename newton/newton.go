package newton

import (
	"github.com/syuya2036/num/functions"
	c "github.com/syuya2036/num/constant"
)

func Newton(f func(float64) float64, x float64) float64 {
	maxIter := 1000 // 最大反復回数

	for i := 0; i < maxIter; i++ {
		fx := f(x)

		if functions.Abs(fx) <= c.Eps {
			break // 収束判定を満たした場合、反復を終了
		}

		x = x - fx / Differential(f, x)
	}

	return x
}

func Differential(f func(float64) float64, x float64) float64 {
	return (f(x + c.Eps) - f(x - c.Eps)) / (2 * c.Eps)
}

func Pincer(f func(float64) float64, xc float64, a, b float64) float64 {
	maxIter := 1000 // 最大反復回数

	for i := 0; i < maxIter; i++ {
		fa := f(a)
		fb := f(b)

		xc = (a*fb - b*fa) / (fb - fa)
		fxc := f(xc)

		if functions.Abs(fxc) <= c.Eps {
			break // 収束判定を満たした場合、反復を終了
		}

		if fa*fxc < 0 {
			b = xc // 根が範囲 [a, xc] にある場合、xcをbに更新
		} else {
			a = xc // 根が範囲 [xc, b] にある場合、xcをaに更新
		}
	}

	return xc
}