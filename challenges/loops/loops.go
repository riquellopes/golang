package main

import (
	"fmt"
	"math"
)

// Calculate -
func Calculate(number int) float64 {
	return math.Pow(float64(number), 2.0)
}

func main() {

	for n := 0; n < 5; n++ {
		if n >= 0 {
			fmt.Println(Calculate(n))
		}
	}
}
