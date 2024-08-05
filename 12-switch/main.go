package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch Case")

	// rand.Seed(time.Now().Unix())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
		fallthrough
		// ^^^ -> fallthrough keyword is used to execute the next case block
	case 2:
		fmt.Println("Work hard")
	case 3:
		fmt.Println("Take a break")
	case 4:
		fmt.Println("Go to sleep")
	case 5:
		fmt.Println("Have a coffee")
	case 6:
		fmt.Println("Go for a walk")
	default:
		fmt.Println("What was that!")
	}

}
