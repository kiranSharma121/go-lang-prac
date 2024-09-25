package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Subject struct {
	SubjectName    string   `json:"subjectName"`
	SubjectCode    string   `json:"subjectCode"`
	SubjectTeacher string   `json:"subjectTeacher"`
	SubjectFee     int      `json:"subjectFee"`
	Periods        *Periods `json:"Periods"`
}
type Periods struct {
	Sunday bool `json:"sunday"`
}

func (s *Subject) IsEmpty() bool {
	return s.SubjectName == ""
}

var subject []Subject

func main() {
	fmt.Println("The server is starting...")
	subject = append(subject, Subject{SubjectName: "English", SubjectCode: "22JUCH003", SubjectTeacher: "Harilal sharma", SubjectFee: 4545, Periods: &Periods{Sunday: true}})
	subject = append(subject, Subject{SubjectName: "Nepali", SubjectCode: "21HUSSS354", SubjectTeacher: "Ramlal Gyawali", SubjectFee: 5454, Periods: &Periods{Sunday: false}})
	subject = append(subject, Subject{SubjectName: "Science", SubjectCode: "20UFF454", SubjectTeacher: "Ramu Rai", SubjectFee: 7676, Periods: &Periods{Sunday: true}})
	subject = append(subject, Subject{SubjectName: "Chemistry", SubjectCode: "23CHS393", SubjectTeacher: "Gulab tharu", SubjectFee: 8843, Periods: &Periods{Sunday: false}})
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/subjects", getAllSubjects).Methods("GET")
	r.HandleFunc("/subject/{code}", getOneSubject).Methods("GET")
	r.HandleFunc("/subject", createOneSubject).Methods("POST")
	r.HandleFunc("/subject/{code}", upDateOneSubject).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", r))

}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to the page...</h1>"))
}
func getAllSubjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(subject)
}
func getOneSubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	for _, value := range subject {
		value.SubjectCode = params["code"]
		json.NewEncoder(w).Encode(value)
		return
	}
	json.NewEncoder(w).Encode("Not found the course with the id")
}
func createOneSubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("please enter the data")
		return
	}

	var subjects Subject
	err := json.NewDecoder(r.Body).Decode(&subjects)
	if err != nil {
		json.NewEncoder(w).Encode("Error in decoding data")
		return
	}
	if subjects.IsEmpty() {
		json.NewEncoder(w).Encode("No data in the json")
	}
	subject = append(subject, subjects)
	json.NewEncoder(w).Encode(subjects)
}

func upDateOneSubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	for index, value := range subject {
		if value.SubjectCode == params["code"] {
			subject = append(subject[:index], subject[index+1:]...)
			var Subjects Subject
			json.NewDecoder(r.Body).Decode(&Subjects)
			value.SubjectCode = params["code"]
			subject = append(subject, Subjects)
			json.NewEncoder(w).Encode(Subjects)
		}

	}
}
