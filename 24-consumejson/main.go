package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string   `json:"courseName"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON")

	DecodeJson()
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		  {
		    "courseName": "Docker",
		    "price": 100,
		    "website": "Youtube",
		    "tags": [
		      "container",
		      "docker"
		    ]
		  }`)

	var myCourse Course
	isValid := json.Valid(jsonDataFromWeb)

	if isValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("JSON is not valid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

}
