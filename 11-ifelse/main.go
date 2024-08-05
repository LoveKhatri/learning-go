package main

import "fmt"

func main() {
	fmt.Println("If Else in Go")

	loginCount := 23

	var result string

	if loginCount < 30 {
		result = "Regular User"
	} else if loginCount < 100 {
		result = "Special User"
	} else {
		result = "Super User"
	}

	fmt.Println("Result:", result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 0 {
		// ^^^ -> initialize variable
		fmt.Println("Number is negative")
	} else {
		fmt.Println("Number is positive")
	}

}
