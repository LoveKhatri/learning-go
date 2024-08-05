package main

import "fmt"

func main() {
	fmt.Println("Loops")

	// days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Printf("%v\t", days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, value := range days {
	// 	fmt.Printf("%v %v\n ", index, value)
	// }

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 2 {
			goto amour
		}

		if rogueValue == 5 {
			rogueValue++
			// ^^^^^^^^ -> if not done then it will go into infinite loop
			// cause the value is never incremented and the condition is always true
			continue
		}
		// if rogueValue == 5 {
		// 	break
		// }

		fmt.Println(rogueValue)
		rogueValue++
	}

amour:
	fmt.Println("Jumping at amour label")

}
