package main

import (
    "fmt"
    "math"
)

//feeling happy!
func main(){
    hypot := func(x, y float64) float64{
        return math.Sqrt(x*x + y*y)
    }

    fmt.Println(hypot(3, 4))
}
