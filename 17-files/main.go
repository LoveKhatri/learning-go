package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files")

	content := "This is supposed to be moved inside a file."

	file, err := os.Create("./myfile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Printf("Length: %v\n", length)
	defer file.Close()
	//^^^ This will close the file after the main function ends

	readFile("./myfile.txt")
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)

	checkNilError(err)

	fmt.Println("Data read from file: ", string(data))
	//                                   ^^^^^^ This is a byte array, so we need to convert it to string
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
