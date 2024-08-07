package main

import "fmt"

func main() {
	defer fmt.Println("one")
	// |  one  | <- TOP
	// -----
	defer fmt.Println("two")
	// |  two  | <- TOP
	// |  one  |
	// -----
	defer fmt.Println("three")
	//               LIFO order
	// | three | <- TOP
	// |  two  |
	// |  one  |
	// ---------
	fmt.Println("four")
	// ^^^^ -> Not affected by defer
	fmt.Println("__________________________")
	myDefer()

	// Final Stack
	// |   4   | <- TOP
	// |   3   |
	// |   2   |
	// |   1   |
	// |   0   |
	// | three |
	// |  two  |
	// |  one  |
	// ---------

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	// | 4 | <- TOP
	// | 3 |
	// | 2 |
	// | 1 |
	// | 0 |
	// ------
}
