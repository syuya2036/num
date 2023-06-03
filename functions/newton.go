package functions

import (
	c "github.com/syuya2036/num/constant"
)

type Function interface {
	F(float64) float64
	Diff(float64) float64
}

func Newton(f func(float64) float64, x float64) float64 {
	return x - f(x)/Differential(f, x)
}

func Differential(f func(float64) float64, x float64) float64 {
	return (f(x + c.Eps) - f(x - c.Eps)) / (2 * c.Eps)
}