package main

import "fmt"

type User struct {
	FirstName   string
	SecondName  string
	DateOfBirth string
}

func main() {
	fmt.Println("welcome to the golang series")
	firstName := getUserInfo("Enter the firstname:")
	secondName := getUserInfo("Enter the secondname:")
	dateOfBirth := getUserInfo("Enter the dateOfBirth:")
	var appUser User
	appUser = User{
		FirstName:   firstName,
		SecondName:  secondName,
		DateOfBirth: dateOfBirth,
	}
	userDetails(appUser)
}
func userDetails(u User) {
	fmt.Println(u.FirstName, u.SecondName, u.DateOfBirth)
}
func getUserInfo(info string) string {
	fmt.Print(info)
	var data string
	fmt.Scan(&data)
	return data
}
