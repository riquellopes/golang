package main

import (
	"fmt"
	"math"
)

func main() {

	for n := 0; n < 5; n++ {
		if n >= 0 {
			fmt.Println(math.Pow(float64(n), 2.0))
		}
	}
}
