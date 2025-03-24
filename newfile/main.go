package main

import "fmt"

func main() {
	fmt.Println("welcome to the struct embedding in golang")
	admin := Admin{
		User: User{
			firstName:  "kiran",
			secondName: "sharma",
		},
		userName: "poudelkiran.np@gmail.com",
		password: "Kiran@123",
	}
	fmt.Printf("Hi %v %v your user name is %v and password is %v\n", admin.firstName, admin.secondName, admin.userName, admin.password)
}

type User struct {
	firstName  string
	secondName string
}

type Admin struct {
	User
	userName string
	password string
}

// package main

// import "fmt"

// type str string

// func (text str) log() {
// 	fmt.Println(text)
// }

// func main() {
// 	fmt.Println("welcome to the custome type in go")
// 	var name str = "kiran sharma"
// 	name.log()
// }
