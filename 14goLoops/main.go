package main

import "fmt"

func main() {
	fmt.Println("Welcome to the loops in Go")
	lang := []string{"python", "Java", "Javascript", "Golang", "Ruby", "Rust"}
	//
	for l := 0; l < len(lang); l++ {
		fmt.Println(lang[l])
	}
	goto kiran
kiran:
	fmt.Println("Hello kiran welcome to the loop in go lang")

}
