package main

import "log"

// Arithmetic -
type Arithmetic struct {
  x int
  y int
}

// ToCalc -
func (a *Arithmetic) ToCalc () (int, int, int){
  return (a.x + a.y), (a.x - a.y), (a.x * a.y)
}


func main() {
    log.Println("Shows")
}
