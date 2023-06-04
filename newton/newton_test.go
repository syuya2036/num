package newton

import (
	"testing"
	"os"

	"github.com/syuya2036/num/functions"
	c "github.com/syuya2036/num/constant"
)

func TestNewton(t *testing.T) {
	// 課題のテストケース: f(x) = x^2 - 3, x0 = 1.5
	f1 := func(x float64) float64 {
		return x*x - 3
	}
	x1, precs := Newton(f1, 1.5)
	if functions.Abs(x1-1.73205080756) > 1e-6 {
		t.Errorf("Test case 1 failed: expected %v, but got %v", 1.73205080756, x1)
	}

	file, err := os.Create("../docs/newton.md")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	defer file.Close()

	file.WriteString(precs)


	// Test case 2: f(x) = x^3 - 2x - 5, x0 = 2
	f2 := func(x float64) float64 {
		return x*x*x - 2*x - 5
	}
	x2, _ := Newton(f2, 2)
	if functions.Abs(x2-2.0945514815423265) > 1e-6 {
		t.Errorf("Test case 2 failed: expected %v, but got %v", 2.0945514815423265, x2)
	}
}

func TestPincer(t *testing.T) {
	// Test case 1: f(x) = x^2 - 2, a = 0, b = 2
	f1 := func(x float64) float64 {
		return x*x - 2
	}
	x1 := Pincer(f1, 1, 0, 2)
	if functions.Abs(x1-1.4142135623731) > 1e-6 {
		t.Errorf("Test case 1 failed: expected %v, but got %v", 1.4142135623731, x1)
	}

	// 課題のテストケース: f(x) = 1.25sin(0.4x)sin(0.6x) - xsin(x), a = 2.5, b = 3
	f2 := func(x float64) float64 {
		c :=  1.25*functions.Sin(0.4*x) * functions.Sin(0.6*x) - x*functions.Sin(x)
		return c
	}
	x2 := Pincer(f2, 2.7, 2.5, 3)
	if functions.Abs(x2-2.72353213966994) > c.Eps {
		t.Errorf("Test case 2 failed: expected %v, but got %v", 2.72353213966994, x2)
	}

	// 適当でない初期値を与えると収束しないことを確認
	x3 := Pincer(f2, -2, 0, 3)
	if functions.Abs(x3-2.72353213966994) < c.Eps {
		t.Errorf("Test case 3 failed: expected %v, but got %v", 2.72353213966994, x3)
	}
}