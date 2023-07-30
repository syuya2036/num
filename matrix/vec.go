package matrix

// rows=1のmat型にする

type Vec struct {
	Matrix
}

func NewVec(values []float64) *Vec {
	return &Vec{
		Matrix: NewMatrix(values, 1, len(values)),
	}
}

func (v *Vec) Mean() float64 {
	sum := v.Sum()
	return sum / float64(len(v.Mat))
}

func (v *Vec) Sum() float64 {
	sum := 0.0
	for _, value := range v.Mat {
		sum += value
	}
	return sum
}

func (v *Vec) Len() int {
	return v.Cols
}