// BEGIN: 2j3k4l5j6h7g
package integrate

import (
	"math"
	"testing"
)

func TestCircleMonteCarlo3D(t *testing.T) {
	r := 0.5
	N := 1000000
	expected := math.Pi * r * r * r * 4 / 3
	got := CircleMonteCarlo3D(r, N)
	if math.Abs(got-expected) > 0.01 {
		t.Errorf("CircleMonteCarlo3D(%v, %v) = %v, want %v", r, N, got, expected)
	}
}
// END: 2j3k4l5j6h7g