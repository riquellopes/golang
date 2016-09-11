package main

import "fmt"

func main(){
    defer fmt.Println("World")

    fmt.Println("Hello")
}
