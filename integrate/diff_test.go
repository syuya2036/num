package integrate

import (
	"testing"
)

// BEGIN: 5d8c9f7d6b4c
func TestFuncArray_Euler(t *testing.T) {
	f := func(x, y float64) float64 {
		return x*x*x
	}
    got := Euler(f, 1, 3, 100, 0)
    want := 20.0
    if got != want {
        t.Errorf("Euler() = %v, want %v", got, want)
    }
}
// END: 5d8c9f7d6b4c