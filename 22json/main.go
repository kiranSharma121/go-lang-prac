package main

import (
	"encoding/json"
	"fmt"
)

type Courses struct {
	Name     string `json:"Courses"`
	Price    int
	Platform string   `json:"Website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	welcome := "Welcome to the creating json in golang"
	fmt.Println(welcome)
	onlineCourses := []Courses{
		{"Reactjs", 0, "Youtube.com", "123ads", []string{"frontend web development", "js framework"}},
		{"Golang", 1000, "Hiteshchoudhary.com", "1245ads", []string{"Backend web development", "Created by google for the backend programming and develops and Ai and Ml too"}},
		{"Ruby", 0, "gigsforgigs.com", "1345ads", nil},
		{"JavaScript", 0, "freecodecamp.com", "12r4ads", []string{"full-stack web development", "lord js"}},
	}
	detailJson, err := json.MarshalIndent(onlineCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(detailJson))
}
