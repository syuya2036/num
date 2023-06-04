// BEGIN: 9f5c3f8d6b7d
package matrix

import (
	"math"
	"testing"
)

func TestNewVec(t *testing.T) {
	values := []float64{1, 2, 3}
	v := NewVec(values)
	if len(v.Values) != len(values) {
		t.Errorf("Expected length of %d, but got %d", len(values), len(v.Values))
	}
	for i, value := range values {
		if v.Values[i] != value {
			t.Errorf("Expected value of %f at index %d, but got %f", value, i, v.Values[i])
		}
	}
}

func TestVecGet(t *testing.T) {
	values := []float64{1, 2, 3}
	v := NewVec(values)
	for i, value := range values {
		if v.Get(i) != value {
			t.Errorf("Expected value of %f at index %d, but got %f", value, i, v.Get(i))
		}
	}
}

func TestVecLen(t *testing.T) {
	values := []float64{1, 2, 3}
	v := NewVec(values)
	if v.Len() != len(values) {
		t.Errorf("Expected length of %d, but got %d", len(values), v.Len())
	}
}

func TestVecMean(t *testing.T) {
	values := []float64{1, 2, 3}
	v := NewVec(values)
	expectedMean := (values[0] + values[1] + values[2]) / float64(len(values))
	if math.Abs(v.Mean()-expectedMean) > 1e-9 {
		t.Errorf("Expected mean of %f, but got %f", expectedMean, v.Mean())
	}
}
// END: 9f5c3f8d6b7d