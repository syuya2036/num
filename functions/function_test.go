// BEGIN: 7d5c8f7d7c8c
package function_test

import (
	"testing"
	"math"

	c "github.com/syuya2036/num/constant"
	"github.com/syuya2036/num/functions"
)

func TestFactorial(t *testing.T) {
	testCases := []struct {
		input int
		want  int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
	}

	for _, tc := range testCases {
		got := function.Factorial(tc.input)
		if got != tc.want {
			t.Errorf("Factorial(%d) = %d; want %d", tc.input, got, tc.want)
		}
	}
}

func TestSin(t *testing.T) {
	testCases := []struct {
		input float64
		want  float64
	}{
		{0, 0},
		{c.Pi / 2, 1},
		{c.Pi, 0},
		{3 * c.Pi / 2, -1},
		{2 * c.Pi, 0},
		{c.Pi / 6, 0.5},
		{c.Pi / 4, math.Sqrt(2) / 2},
		{ 7, math.Sin(7)},
		{ 2423589, math.Sin(2423589)},
		{ -576949, math.Sin(-576949)},
	}

	eps := 1e-6
	for _, tc := range testCases {
		got := function.Sin(tc.input)
		if math.Abs(got - tc.want) > eps {
			t.Errorf("Sin(%f) = %f; want %f", tc.input, got, tc.want)
		}
	}
}
// END: 7d5c8f7d7c8c// BEGIN: 8f3c5d6b7a8e

func TestCos(t *testing.T) {
	testCases := []struct {
		input float64
		want  float64
	}{
		{0, 1},
		{c.Pi / 2, 0},
		{c.Pi, -1},
		{3 * c.Pi / 2, 0},
		{2 * c.Pi, 1},
		{c.Pi / 6, math.Sqrt(3) / 2},
		{c.Pi / 4, math.Sqrt(2) / 2},
		{ 7, math.Cos(7)},
		{ 2423589, math.Cos(2423589)},
		{ -576949, math.Cos(-576949)},
	}

	eps := 1e-6
	for _, tc := range testCases {
		got := function.Cos(tc.input)
		if math.Abs(got - tc.want) > eps {
			t.Errorf("Cos(%f) = %f; want %f", tc.input, got, tc.want)
		}
	}
}