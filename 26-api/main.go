package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB
var courses []Course

// middlewares
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Hello from API")

	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "React", CoursePrice: 100, Author: &Author{Fullname: "John Doe", Website: "https://johndoe.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Angular", CoursePrice: 200, Author: &Author{Fullname: "Jane Doe", Website: "https://janedoe.com"}})

	// Routes
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourse).Methods("GET")
	r.HandleFunc("/course/{id}", getCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	// Listen to a port
	log.Fatal(http.ListenAndServe(":3000", r))
}

// Controllers
func serveHome(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Welcome to Home page"))
}

func getAllCourse(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("Get all courses")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get course")

	w.Header().Set("Content-Type", "application/json")

	// Get the id from request (r)
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course found with given ID")
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create course")

	w.Header().Set("Content-Type", "application/json")

	// Empty body
	if r.Body == nil {
		json.NewEncoder(w).Encode("Empty Body")
		return
	}

	var course Course

	// {} body
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No Course found in body")
		return
	}

	// Generate unique ID
	course.CourseId = strconv.Itoa(rand.Intn(1000000))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update course")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			course.CourseId = params["id"]
			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course found with given ID")
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete course")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			break
		}
	}

	json.NewEncoder(w).Encode("No Course found with given ID")
}
