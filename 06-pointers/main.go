package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")

	var one = 2

	var ptr *int = &one
	//             ^ -> used to get reference of variable
	//      ^ -> used to declare pointer variable

	fmt.Println("Value of ptr:", ptr)
	fmt.Println("Value of *ptr:", *ptr)

	*ptr = 3
	fmt.Println("Value of one:", one)
}
