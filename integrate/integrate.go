package integrate

import (
	"fmt"
	"math/rand"

	mat "github.com/syuya2036/num/matrix"
)

type MathFunction func(float64) float64

type FuncArray struct {
	Mf MathFunction
	mat.Vec
	h float64
}

func NewFunc(mf MathFunction) *FuncArray {
	return &FuncArray{
		Mf: mf,
	}
}

// 関数の結果を事前に計算する
func (f *FuncArray) Init(a, b float64, N int) error {
	if N <= 0 {
		return fmt.Errorf("Nは正の整数にしてください。N: %d", N)
	}
	// 幅を定義
	h := (b - a) / float64(N)
	f.h = h

	vec := make([]float64, N)
	for i := range vec {
		vec[i] = f.Mf(a + float64(i) * h)
	}

	f.Mat = vec
	f.Rows, f.Cols = 1, N

	return nil
}

func (f *FuncArray) Trapezoid() float64 {
	sum := f.Sum()
	I := (f.h / 2) * (f.Get(0) + f.Get(f.Cols-1) + sum * 2)
	return I
}

func (f *FuncArray) Simpson() (float64, error) {
	n := f.Len()
	if n % 2 != 0 {
		return 0, fmt.Errorf("Nを偶数にしてください。N: %d", n)
	}

	sum := 0.0
	for i:=1; i<=n-1; i+=2 {
		sum += 4 * f.Get(i)
	}
	for j:=2; j<=n-2; j+=2 {
		sum += 2 * f.Get(j)
	}

	I := (sum + f.Get(0) + f.Get(n-1))*f.h/3
	return I, nil
}

func MonteCarlo(a, b float64, N int, f func(float64) float64) float64 {
	sum := 0.0
	for i:=0; i<N; i++ {
		x := rand.Float64() * (b - a) + a
		sum += f(x)
	}

	I := (b - a) * sum / float64(N)
	return I
}