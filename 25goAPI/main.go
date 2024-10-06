package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake database
var Courses []Course

// middleware
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("The server is loading...")
	r := mux.NewRouter()
	Courses = append(Courses, Course{CourseId: "2", CourseName: "ReactJs", CoursePrice: 6000, Author: &Author{Fullname: "Kiran sharma", Website: "Kirantube.com"}})
	Courses = append(Courses, Course{CourseId: "3", CourseName: "Javascript", CoursePrice: 7000, Author: &Author{Fullname: "Karuna sharma", Website: "Youtube.com"}})
	Courses = append(Courses, Course{CourseId: "1", CourseName: "Golang", CoursePrice: 15000, Author: &Author{Fullname: "Rudra prasad sharma", Website: "geeksforgeeks.com"}})
	Courses = append(Courses, Course{CourseId: "4", CourseName: "Ruby", CoursePrice: 55000, Author: &Author{Fullname: "Kalpana Aryal", Website: "porche.com"}})
	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getallData).Methods("GET")
	r.HandleFunc("/course/{Id}", getonecourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{Id}", updateonecourse).Methods("PUT")
	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to the serveHome</h1>"))
}

func getallData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all the dataBase")
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(Courses)
}

func getonecourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	Params := mux.Vars(r)
	for _, course := range Courses {
		if course.CourseId == Params["Id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")

}
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send the data")
		return
	}

	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		json.NewEncoder(w).Encode("Error in decoding JSON")
		return
	}

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	Courses = append(Courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateonecourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	Params := mux.Vars(r)
	for index, course := range Courses {
		if course.CourseId == Params["Id"] {
			Courses = append(Courses[:index], Courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = Params["Id"]
			Courses = append(Courses, course)
			json.NewEncoder(w).Encode(course)
		}
	}
}
