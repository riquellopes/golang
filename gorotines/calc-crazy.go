package calccrazy

import "fmt"

func Calc(c chan float64) {
	c <- 2.3423 / 534.23423

}

func UserGoRoutine() float64 {
	canal := make(chan float64)

	go Calc(canal)

	return <-canal
}

func main() {
	valor := UserGoRoutine()
	fmt.Printf("The result was: %v \n", valor)
}
