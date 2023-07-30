// BEGIN: 5f8c9d8fj3d9
package integrate

import (
	"math"
	"testing"
)

func TestIntegral(t *testing.T) {
	f := func(x float64) float64 {
		return math.Sin(x)
	}
	fn := NewFunc(f)

	tests := []struct {
		a, b float64
		N    int
		want float64
	}{
		{0, math.Pi, 10e6, 2.0},
		{0, math.Pi / 2, 10e6, 1.0},
		{0, math.Pi / 4, 10e6, 0.293},
	}

	for _, tt := range tests {
		fn.Init(tt.a, tt.b, tt.N)
		if fn.Len() != tt.N {
			t.Errorf("Invalid f size, got %d", fn.Len())
		}

		got1 := fn.Trapezoid()
		if math.Abs(got1-tt.want) > 0.001 {
			t.Errorf("Trapezoid(%g, %g, %d) = %g; want %g", tt.a, tt.b, tt.N, got1, tt.want)
		}
		got2, err := fn.Simpson()
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if math.Abs(got2-tt.want) > 0.001 {
			t.Errorf("Simpson(%g, %g, %d) = %g; want %g", tt.a, tt.b, tt.N, got2, tt.want)
		}

		got := MonteCarlo(tt.a, tt.b, tt.N, f)
		if math.Abs(got-tt.want) > 0.001 {
			t.Errorf("MonteCarlo(%g, %g, %d) = %g; want %g", tt.a, tt.b, tt.N, got, tt.want)
		}
	}
}
// END: 5f8c9d8fj3d9