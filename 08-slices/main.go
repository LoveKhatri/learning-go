package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Welcome to Slices")

	var fruits = []string{"Apple", "Banana", "Mango"}

	fmt.Printf("Type of Fruits: %T\n", fruits)
	fmt.Printf("Fruits: %v\n", fruits)
	fmt.Println("Fruits Length:", len(fruits))

	// Append
	fruits = append(fruits, "Pineapple")
	fmt.Println("Fruits after append:", fruits)
	fmt.Println("Fruits Length after append:", len(fruits))

	// Slice
	fruits = fruits[1:3]
	fmt.Println("Fruits after slice:", fruits)
	fmt.Println("Fruits Length after slice:", len(fruits))

	prices := make([]int, 3)
	prices[0] = 900
	prices[1] = 700
	prices[2] = 1000
	// prices[3] = 400
	// ^^^^^^^^^^^^ -> index out of bounds error

	fmt.Println("Prices:", prices)

	prices = append(prices, 490, 60, 7000)
	//       ^^^^^^ -> re-allocates the memory
	// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ -> append multiple values (NO ERROR)

	fmt.Println("Prices after append:", prices)

	// sort
	slices.Sort(prices)
	fmt.Println("Prices after sort:", prices)
}
