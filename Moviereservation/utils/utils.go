package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
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

func GeneratejwtToken(userid int64, username, email, role string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":   userid,
		"username": username,
		"email":    email,
		"role":     role,
	})
	return token.SignedString([]byte(secretKey))
}
