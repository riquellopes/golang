package main

import "fmt"

var pow = []int{1,2,4,6,16,32,64,128}

func main(){

    for i, j := range pow{
        fmt.Printf("2**%d = %d\n", i,j)
    }
}
