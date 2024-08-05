package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")

	greet()

	// func greetTwo() {
	// 	fmt.Println("Hello from GoLang")
	// }
	// ! ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ -> functions can't be defined inside another function

	result := adder(3, 4)
	fmt.Println("Result is: ", result)

	proResult := proAdder(1, 2, 3, 4, 5)
	fmt.Println("Pro Result is: ", proResult)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	//               ^^^^^ -> all values are of int type
	//        ^^^^^^ -> is a slice
	result := 0
	for _, val := range values {
		result += val
	}
	return result
}

func greet() {
	fmt.Println("Hello from GoLang")
}
