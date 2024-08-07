package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to GoLang")

	PerformGetReq()
}

func PerformGetReq() {
	const myUrl = "http://localhost:8000/get"

	res, err := http.Get(myUrl)
	checkNil(err)

	fmt.Println(res.Status)
	fmt.Println(res.ContentLength)

	var resString strings.Builder
	content, _ := io.ReadAll(res.Body)
	byteCount, _ := resString.Write(content)

	fmt.Println(byteCount)
	fmt.Println(resString.String())
	fmt.Println(string(content))

	defer res.Body.Close()
}

func checkNil(err error) {
	if err != nil {
		panic(err)
	}
}
