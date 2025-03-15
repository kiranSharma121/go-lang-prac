package main

import "fmt"

type user struct {
	firstName   string
	secondName  string
	dateOfBirth string
}

func main() {
	fmt.Println("welcome to the struct in go")
	firstName := getUserDetails("Enter your firstName:")
	secondName := getUserDetails("Enter your secondName:")
	dateOfBirth := getUserDetails("Enter your date of birth in MM//DD//YY format:")
	var appUser = user{
		firstName,
		secondName,
		dateOfBirth,
	}
	fmt.Println(appUser)
}
func getUserDetails(text string) string {
	fmt.Print(text)
	var value string
	fmt.Scanln(&value)
	return value
}
