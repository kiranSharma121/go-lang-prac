package main

import (
	"encoding/json"
	"fmt"
)

type Courses struct {
	Name     string
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to the decode JSON in golang")
	// encodingJSON()
	decodingJSON()
}
func encodingJSON() {
	onlineCourse := []Courses{

		{"Reactjs", 0, "Youtube.com", "abcd1234", []string{"Fronted web development", "javascript framework"}},
		{"Golang", 1000, "Google.com", "qwert`234", []string{"Backend web development", "Best for backend"}},
		{"Python", 2000, "Instagram.com", "asdf1234", []string{"Ai and machine learning", "Beginner friendly course"}},
		{"Ruby", 220, "Freecodecamp.com", "cvzx234", nil},
	}
	detailCourses, err := json.MarshalIndent(onlineCourse, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(detailCourses))

}
func decodingJSON() {
	jsonFromWeb := []byte(`
	{
                "Name": "Golang",
                "Price": 1000,
                "Platform": "Google.com",
                "tags": [
                        "Backend web development",
                        "Best for backend"
                ]
        }
	`)
	var hiteshCourses Courses
	checkValid := json.Valid(jsonFromWeb)
	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonFromWeb, &hiteshCourses)
		fmt.Printf("%#v\n", hiteshCourses)
	} else {
		fmt.Println("The JSON is invalid")
	}
	var jsonData map[string]interface{}
	json.Unmarshal(jsonFromWeb, &jsonData)
	fmt.Printf("%#v\n", jsonData)
	// for key, value := range jsonData {
	// 	fmt.Printf("The %v for the data is %v\n", key, value)

	// }

}
