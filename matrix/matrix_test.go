package matrix

import (
	"testing"
	"math"

	c "github.com/syuya2036/num/constant"
)

func TestGaussianElimination(t *testing.T) {
	testCases := []struct {
		input [][]float64
		want []float64
	} {
		{[][]float64{
			{-2, 1, 0, 0, 0, 1},
			{1, -2, 1, 0, 0, 2},
			{0, 1, -2, 1, 0, 3},
			{0, 0, 1, -2, 1, 4},
			{0, 0, 0, 1, -2, 5},
		},
		[]float64{-5.833333, -10.666666, -13.499999, -13.333333, -9.166666},},
	}

	for _, tc := range testCases {
		x := NewMatrixFromArray(tc.input)
		got := x.GaussianElimination()
		for i := range tc.want {
			if math.Abs(got.Get(i, 6-1) - tc.want[i]) > c.Eps {
				t.Errorf("GaussianElimination(%v) = %v, want %v", tc.input, got, tc.want)
			}
		}
	}
}

func TestGaussSeidel(t *testing.T) {
	testCases := []struct {
		inputX [][]float64
		inputY []float64
		want []float64
	} {
		{[][]float64{
			{-2, 1, 0, 0, 0},
			{1, -2, 1, 0, 0},
			{0, 1, -2, 1, 0},
			{0, 0, 1, -2, 1},
			{0, 0, 0, 1, -2},
		},
		[]float64{1,2,3,4,5},
		[]float64{-5.833333, -10.666666, -13.499999, -13.333333, -9.166666},},
	}

	for _, tc := range testCases {
		x := NewMatrixFromArray(tc.inputX)
		b := NewMatrix(tc.inputY, 1, 5)

		got := GaussSeidel(x, b)
		for i := range got.Mat {
			if math.Abs(got.Mat[i] - tc.want[i]) > c.Eps {
				t.Errorf("GaussSeidel(%v, %v) = %v, want %v", tc.inputX, tc.inputY, got, tc.want)
			}
		}
	}
}