package user

import (
	"errors"
	"fmt"
)

type User struct {
	firstName   string
	secondName  string
	dateOfBirth string
}

func (u User) PrintUserDetails() {
	fmt.Printf("Hi %v %v welcome to the struct in go and your date of birth is %v\n", u.firstName, u.secondName, u.dateOfBirth)

}

// func (u *user) clearUserDetails() {
// 	u.firstName = ""
// 	u.secondName = ""

// }
func NewUser(userfirstName, usersecondName, userdateOfBirth string) (*User, error) {
	if userfirstName == "" || usersecondName == "" || userdateOfBirth == "" {
		return nil, errors.New("this field can't be empty")
	}
	return &User{
		userfirstName,
		usersecondName,
		userdateOfBirth,
	}, nil
}
