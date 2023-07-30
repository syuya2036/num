package matrix

import (
	"math"
	"testing"
)

func TestVecMean(t *testing.T) {
	values := []float64{1, 2, 3}
	v := NewVec(values)

	expectedMean := (1 + 2 + 3) / 3.0
	if math.Abs(v.Mean()-expectedMean) > 1e-9 {
		t.Errorf("Expected mean to be %f, but got %f", expectedMean, v.Mean())
	}
}

func TestVecSum(t *testing.T) {
	values := []float64{1, 2, 3}
	v := NewVec(values)

	expectedSum := 1 + 2 + 3
	if math.Abs(v.Sum()-float64(expectedSum)) > 1e-9 {
		t.Errorf("Expected sum to be %d, but got %f", expectedSum, v.Sum())
	}
}

func TestLen(t *testing.T) {
	values := []float64{1, 2, 3}
	v := NewVec(values)

	expectedLen := 3
	if v.Len() != expectedLen {
		t.Errorf("Expected len to be %d, but got %d", expectedLen, v.Len())
	}
}