package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in GoLang")

	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	// fruits[2] = "Mango"
	fruits[3] = "Pineapple"

	fmt.Println("Fruits: ", fruits)
	fmt.Println("Fruits: ", len(fruits))
	//                      ^^^ -> total size of the array (not the number of items inside the array)

	var vegies = [3]string{"Potato", "Tomato", "Onion"}

	fmt.Println("Vegies: ", vegies)
}
