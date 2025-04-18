package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HasedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
func CompareHasedPassword(password, hasedpasword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hasedpasword), []byte(password))
	return err == nil

}
