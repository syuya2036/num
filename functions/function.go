package functions

import (
	c "github.com/syuya2036/num/constant"
)

func Sin(x float64) float64 {
	if x == 0 {
		return 0
	} else if x == c.Pi/2 {
		return 1
	} else if x < 2 * c.Pi && x >= c.Pi {
		x = x - c.Pi
		return -Sin(x)
	} else if x >= 2 * c.Pi {
		q := int(x / (2 * c.Pi))
		x = x - 2*c.Pi*float64(q)
	} else if x < 0 {
		return -Sin(-x)
	} else if x >= c.Pi/2 && x < c.Pi {
		return Sin(c.Pi - x)
	}

	return x - x*x*x/6 + x*x*x*x*x/120 - x*x*x*x*x*x*x/5040 + x*x*x*x*x*x*x*x*x/362880 - x*x*x*x*x*x*x*x*x*x*x/39916800
}

func Cos(x float64) float64 {
	if x == 0 {
		return 1
	} else if x == c.Pi * 0.5 {
		return 0
	} else if x < 2 * c.Pi && x >= c.Pi {
		x = x - c.Pi
		return -1 * Cos(x)
	} else if x >= 2 * c.Pi {
		q := int(x / (2 * c.Pi))
		x = x - 2*c.Pi*float64(q)
	} else if x < 0 {
		return Cos(-x)
	} else if x >= c.Pi *0.5 && x < c.Pi {
		return -1 * Cos(c.Pi - x)
	}

	return 1 - x*x/2 + x*x*x*x/24 - x*x*x*x*x*x/720 + x*x*x*x*x*x*x*x/40320 - x*x*x*x*x*x*x*x*x*x/3628800 + x*x*x*x*x*x*x*x*x*x*x*x/479001600
}

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}

	return x
}