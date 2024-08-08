package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Post Form Request")

	PostFormReq()
}

func PostFormReq() {
	const myUrl = "http://localhost:8000/postform"

	// Fake form data
	data := url.Values{}
	data.Add("firstName", "Amour")
	data.Add("age", "20")
	data.Add("email", "amour@mail.com")

	res, err := http.PostForm(myUrl, data)
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
