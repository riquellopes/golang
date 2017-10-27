package main

import (
    "log"
    // "os"
    "flag"
)

// Arithmetic -
type Arithmetic struct {
  x int
  y int
}

// ToCalc -
func (a *Arithmetic) ToCalc () (int, int, int){
  return a.Sum(), a.Difference(), a.Product()
}

// Sum -
func (a *Arithmetic) Sum() int {
  return (a.x + a.y)
}

// Difference -
func (a *Arithmetic) Difference() int {
  return (a.x - a.y)
}

// Product -
func (a *Arithmetic) Product() int {
  return (a.x * a.y)
}

func main() {
    x := flag.Int("x", 0, "Should be a integer.")
    y := flag.Int("y", 0, "Should be a integer.")

    flag.Parse()

    ari := Arithmetic{*x, *y}
    log.Println(ari.ToCalc())
}
