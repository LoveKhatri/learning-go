package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")

	// No Inheritance in Go
	// No Super or parent

	amour := User{"Amour", "amour@mail.com", true, 20}

	fmt.Printf("Details: %+v\n", amour)
	//                   ^^^ -> Printf will print the struct fields
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
