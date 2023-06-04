// BEGIN: 8d7f6e5a7b3c
package regression

import (
	"math"
	"testing"

	"github.com/syuya2036/num/matrix"
	c "github.com/syuya2036/num/constant"
)

func TestLinearRegression(t *testing.T) {
	tests := []struct {
		x      []float64
		y      []float64
		a, b   float64
	}{
		{
			x:      []float64{1, 2, 3, 4, 5},
			y:      []float64{2, 4, 6, 8, 10},
			a:      0,
			b:      2,
		},
		{
			x:      []float64{1, 2, 3, 4, 5},
			y:      []float64{10, 8, 6, 4, 2},
			a:      12,
			b:      -2,
		},
		{
			x:      []float64{1, 2, 3, 4, 5},
			y:      []float64{1, 2, 3, 4, 5},
			a:      0,
			b:      1,
		},
		{
			// 課題のテストケース
			x:      []float64{10, 15, 31, 41, 50, 62},
			y:      []float64{31.5, 31.0, 27.3, 25.0, 21.8, 19.3},
			a:	    34.43281083,
			b: 		-0.242568732,
		},
	}

	for _, tt := range tests {
		x := matrix.NewVec(tt.x)
		y := matrix.NewVec(tt.y)

		a, b, err := LinearRegression(x, y)
		if err != nil {
			t.Errorf("LinearRegression(%v, %v) is failed. Erorr: %v", tt.x, tt.y, err)
		}

		if math.Abs(a-tt.a) > c.Eps || math.Abs(b-tt.b) > 1e-10 {
			t.Errorf("LinearRegression(%v, %v) = (%v, %v), want (%v, %v)", tt.x, tt.y, a, b, tt.a, tt.b)
		}
	}
}
