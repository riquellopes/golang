package main

import (
	"fmt"

	"github.com/riquellopes/golang/clean-arch/usecase/user"
)

func main() {
	fmt.Println("Start cmd")

	repository := user.NewRepository()
	service := user.NewService(repository)

	// Create one
	service.CreateUser("Henrique", 40)

	users := service.ListUsers()

	for _, u := range users {
		fmt.Printf("%s %d \n", u.Name, u.Age)
	}
}
