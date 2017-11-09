package main

import "fmt"

func MultReturn() (name string, age int) {
	name = "Henrique Lopes"
	age = 35

	return
}

func main() {
	name, age := MultReturn()

	fmt.Println(name, age)
}
