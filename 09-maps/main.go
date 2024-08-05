package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"
	languages["TS"] = "TypeScript"

	fmt.Println("Languages:", languages)
	fmt.Println("JS:", languages["JS"])

	delete(languages, "PY")
	fmt.Println("Languages:", languages)

	// Looping
	for key, value := range languages {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}
}
