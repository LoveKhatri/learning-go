package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://lovekhatri.github.io:443/path?name=love&age=21"

func main() {
	fmt.Println("URLs")
	fmt.Println(myUrl)

	result, err := url.Parse(myUrl)
	checkNil(err)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	queryParams := result.Query()
	fmt.Println(queryParams)
	fmt.Println(queryParams["name"])

	for _, value := range queryParams {
		fmt.Println(value)
	}

	partsOfUrl := &url.URL{
		//        ^ Always need to pass reference
		Scheme:   "https",
		Host:     "lovekhatri.github.io",
		Path:     "/account",
		RawQuery: "name=love&age=21",
	}

	fmt.Println(partsOfUrl.String())
}

func checkNil(err error) {
	if err != nil {
		panic(err)
	}
}
