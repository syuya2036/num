package main

import (
	"log"

	"github.com/syuya2036/num/matrix"
) 

func main() {
	err := matrix.Calculate()
	if err != nil {
		log.Fatal(err)
	}
}