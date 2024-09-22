package main

import "fmt"

type userInfo struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func main() {
	fmt.Println("Welcome to the goStructs")
	userFirst := userInfo{"Kiran Sharma", "poudelkiran.np@gmail.com", 21, true}
	fmt.Println("The userinfo of the userFirst is:", userFirst)
	fmt.Printf("Hello %v your Email address is %v\n", userFirst.Name, userFirst.Email)
}
