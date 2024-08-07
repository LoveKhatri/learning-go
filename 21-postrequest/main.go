package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to GoLang")

	PerformPostJsonReq()
}

func PerformPostJsonReq() {
	const myUrl = "http://localhost:8000/post"

	// Fake JSON payload
	requestBody := strings.NewReader(`
	{
		"key1": "value1",
		"key2": "value2"
	}
	`)

	res, err := http.Post(myUrl, "application/json", requestBody)
	checkNil(err)

	content, err := io.ReadAll(res.Body)
	checkNil(err)
	fmt.Println(string(content))

	defer res.Body.Close()
}

func checkNil(err error) {
	if err != nil {
		panic(err)
	}
}
