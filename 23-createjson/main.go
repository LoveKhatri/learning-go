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

	EncodeJson()
}

func EncodeJson() {
	newCourse := []Course{
		{"Docker", 100, "Youtube", "123456", []string{"container", "docker"}},
		{"Next", 100, "Freecodecamp", "abcdefg", []string{"web", "backend"}},
		{"React", 100, "Coursera", "zyxwvut", nil},
	}

	// finalJson, err := json.Marshal(newCourse)
	finalJson, err := json.MarshalIndent(newCourse, "", "  ")
	checkNil(err)

	fmt.Printf("%s\n", finalJson)
}

func checkNil(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
