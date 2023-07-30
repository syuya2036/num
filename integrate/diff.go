package integrate

func Euler(f func(x, y float64) float64, a, b float64, N int, eta float64) float64 {
	h := (b - a) / float64(N)
	x, y := a, eta
	for x < b {
		y += h * f(x, y)
		x += h
	}

	return y
}

func RungeKutta(f func(x, y float64) float64, a, b float64, N int, eta float64) float64 {
	h := (b - a) / float64(N)
	x, y := a, eta
	for x < b {
		k1 := h * f(x, y)
		k2 := h * f(x+h/2, y+k1/2)
		k3 := h * f(x+h/2, y+k2/2)
		k4 := h * f(x+h, y+k3)
		y += (k1 + 2*k2 + 2*k3 + k4) / 6
		x += h
	}

	return y
}