package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lovekhatri.github.io"

func main() {
	fmt.Println("Hello Web")

	response, err := http.Get(url)

	checkNilError(err)

	defer response.Body.Close()
	//                  ^^^^^ caller's responsibility to close the body connection

	data, err := io.ReadAll(response.Body)
	checkNilError(err)

	fmt.Println(string(data))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
