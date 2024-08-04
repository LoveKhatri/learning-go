package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Pizza App")
	fmt.Print("Please rate our pizza between 1 and 5: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating: ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	//                                   ^^^^^^^ -> package to manipulate strings
	//                ^^^^^^^ -> package to parse strings into numbers

	if err != nil {
		fmt.Println(err)
	} else {
		numRating = numRating + 1
	}

	fmt.Println("Added 1 to your rating: ", numRating)
}
