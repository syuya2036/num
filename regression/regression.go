package regression

import "fmt"

type Array1D interface {
	Get(i ...int) float64
	Len() int
	Mean() float64
}

func LinearRegression(x, y Array1D) (float64, float64, error) {
	if x.Len() != y.Len() {
		err := fmt.Errorf("array length is not same")
		return 0.0, 0.0, err
	}
	meanX := x.Mean()
	meanY := y.Mean()
	covXY, _ := Cov(x, y)
	varianceX, _ := Cov(x, x)

	// ゼロ除算を防ぐ
	if varianceX == 0 {
		err := fmt.Errorf("variance of x is zero")
		return 0.0, 0.0, err
	}

	// 最小二乗法
	a := covXY / varianceX
	b := meanY - a * meanX

	return a, b, nil
}

func Cov(a, b Array1D) (float64, error) {
	if a.Len() != b.Len() {
		err := fmt.Errorf("array length is not same")
		return 0.0, err
	}
	meanA := a.Mean()
	meanB := b.Mean()
	sum := 0.0
	for i := 0; i < a.Len(); i++ {
		sum += (a.Get(i) - meanA) * (b.Get(i) - meanB)
	}
	return sum / float64(a.Len()), nil
}