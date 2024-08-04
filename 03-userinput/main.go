package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to GoLang"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	//        ^^^^^ -> package to handle input from the user
	fmt.Println("Enter the rating for our Pizza: ")

	// Comma Ok syntax || error ok syntax
	input, _ := reader.ReadString('\n')
	//     ^ -> placeholder for error
	// _ -> ignore the error
	fmt.Println("Thanks for rating, ", input)
}
