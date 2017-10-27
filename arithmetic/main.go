package main

import (
    "log"
    "os"
    "strconv"
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
    x, err := strconv.Atoi(os.Args[1])
    y, err := strconv.Atoi(os.Args[2])

    if err != nil {
      log.Println(err)
      os.Exit(2)
    }

    ari := Arithmetic{x, y}

    log.Println(ari.ToCalc())
}
