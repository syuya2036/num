package main

import (
	"fmt"
	"math"
	"bytes"

	"github.com/syuya2036/num/integrate"
)

func main() {
	Ns := []int{10, 10e2, 10e3, 10e4, 10e5, 10e6}
	expected := 0.5236
	var precs bytes.Buffer

	precs.WriteString("| N | 結果 | 誤差 |\n")
	precs.WriteString("| --- | --- | --- |\n")

	for _, N := range Ns {
		got := integrate.CircleMonteCarlo3D(0.5, N)

		e := math.Abs(got - expected)
		precs.WriteString(fmt.Sprintf("| %d | %f | %f | \n", N, got, e))
	}

	fmt.Println(precs.String())
}