package integrate

import (
	"math/rand"
)

func CircleMonteCarlo3D(r float64, N int) float64 {
	var sum float64
	r2 := r*r
	for i := 0; i < N; i++ {
		x := rand.Float64() * 2 * r - r
		y := rand.Float64() * 2 * r - r
		z := rand.Float64() * 2 * r - r
		if x*x+y*y+z*z <= r2 {
			sum += 1
		}
	}
	return sum / float64(N)
}