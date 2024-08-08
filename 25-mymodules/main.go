package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to modules")

	greet()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func greet() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to modules</h1>"))
}

// go mod tidy
// go list all
// go list -m all
// go list -m -versions github.com/gorilla/mux
// go mod why github.com/gorilla/mux
// go mod graph
// go mod edit -go 1.22.0
// go mod vendor
// go run -mod=vendor main.go
