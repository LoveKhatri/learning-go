package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")

	// No Inheritance in Go
	// No Super or parent

	amour := User{"Amour", "amour@mail.com", true, 20}

	fmt.Printf("Details: %+v\n", amour)
	//                   ^^^ -> Printf will print the struct fields

	amour.GetStatus()
	amour.NewMail()
	fmt.Println(amour)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User Active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "newmail@mail.com"
	fmt.Println("New Email is: ", u.Email)
}
