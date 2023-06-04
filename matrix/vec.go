package matrix

type Vec struct {
	Values []float64
}

func NewVec(values []float64) *Vec {
	return &Vec{values}
}

func (v *Vec) Get(i int) float64 {
	return v.Values[i]
}

func (v *Vec) Len() int {
	return len(v.Values)
}

func (v *Vec) Mean() float64 {
	sum := 0.0
	for _, value := range v.Values {
		sum += value
	}
	return sum / float64(len(v.Values))
}