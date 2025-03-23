package main

import (
	"fmt"

	"github.com/structs/user"
)

func main() {
	fmt.Println("welcome to the struct in go")
	userfirstName := getUserDetails("Enter your firstName:")
	usersecondName := getUserDetails("Enter your secondName:")
	userdateOfBirth := getUserDetails("Enter your date of birth in MM//DD//YY format:")
	var appUser, err = user.NewUser(userfirstName, usersecondName, userdateOfBirth)
	if err != nil {
		fmt.Println("error in getting the user values...")
		return
	}
	appUser.PrintUserDetails()
	admin := user.Newadmin("Kiransharma@gmail.com", "Kiran@123")
	fmt.Println(admin)
	// appUser.clearUserDetails()
	// appUser.printUserDetails()
}

func getUserDetails(text string) string {
	fmt.Print(text)
	var value string
	fmt.Scanln(&value)
	return value
}
