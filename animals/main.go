package main

import (
    "fmt"
)

// Animal --
type Animal struct {
    Name string
    mean bool
}

// AnimalSounder --
type AnimalSounder interface {
    MakeNoise()
}

// Dog --
type Dog struct {
    Animal
    BarkStrength int
}

// Cat --
type Cat struct {
    Basics Animal
    MeowStrength int
}

func main(){
    myDog := &Dog{
        Animal{
            "Rover",
            false,
        },
        2,
    }

    myCat := &Cat{
        Basics: Animal{
            Name: "Julius",
            mean: true,
        },
        MeowStrength: 3,
    }

    MakeSomeNoise(myDog)
    MakeSomeNoise(myCat)
}

// PerformNoise --
func (animal *Animal) PerformNoise(strength int, sound string){
    if animal.mean == true {
        strength = strength * 5
    }

    for voice :=0; voice < strength; voice++ {
        fmt.Printf("%s ", sound)
    }

    fmt.Println()
}

// MakeNoise -
func (dog *Dog) MakeNoise(){
    dog.PerformNoise(dog.BarkStrength, "BARK")
}

// MakeNoise --
func (cat *Cat) MakeNoise(){
    cat.Basics.PerformNoise(cat.MeowStrength, "MEOW")
}
// MakeSomeNoise --
func MakeSomeNoise(animalSounder AnimalSounder){
    animalSounder.MakeNoise()
}
