package functions

import (
	"testing"
)

func TestNewton(t *testing.T) {
	// Test case 1: f(x) = x^2 - 2, x0 = 1
	f1 := func(x float64) float64 {
		return x*x - 3
	}
	x1 := Newton(f1, 1)
	if Abs(x1-1.73205080756) > 1e-6 {
		t.Errorf("Test case 1 failed: expected %v, but got %v", 1.73205080756, x1)
	}

	// Test case 2: f(x) = x^3 - 2x - 5, x0 = 2
	f2 := func(x float64) float64 {
		return x*x*x - 2*x - 5
	}
	x2 := Newton(f2, 2)
	if Abs(x2-2.0945514815423265) > 1e-6 {
		t.Errorf("Test case 2 failed: expected %v, but got %v", 2.0945514815423265, x2)
	}
}

func TestPincer(t *testing.T) {
	// Test case 1: f(x) = x^2 - 2, a = 0, b = 2
	f1 := func(x float64) float64 {
		return x*x - 2
	}
	x1 := Pincer(f1, 1, 0, 2)
	if Abs(x1-1.4142135623731) > 1e-6 {
		t.Errorf("Test case 1 failed: expected %v, but got %v", 1.4142135623731, x1)
	}

	// Test case 2: f(x) = 1.25sin(0.4x)sin(0.6x) - xsin(x), a = 2.5, b = 3
	f2 := func(x float64) float64 {
		c :=  1.25*Sin(0.4*x) * Sin(0.6*x) - x*Sin(x)
		return c
	}
	x2 := Pincer(f2, 2.7, 2.5, 3)
	if Abs(x2-2.72353213966994) > 1e-6 {
		t.Errorf("Test case 2 failed: expected %v, but got %v", 2.72353213966994, x2)
	}
}