package utils

import (
	"errors"
	"fmt"
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

func VerifyToken(tokenString string) (string, int64, error) {
	secretKey := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unauthorized signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", 0, errors.New("invalid token claims")
	}

	fmt.Println("this is first one")
	fmt.Println(claims)

	role, ok := claims["role"].(string)
	if !ok {
		return "", 0, errors.New("role not found in token")
	}

	fmt.Println("role is ", role)
	userid, ok := claims["userid"].(float64)
	if !ok {
		return "", 0, errors.New("userid not found in token")
	}
	fmt.Println("from verifytoken function")
	fmt.Println("role is ", role, " and userid is ", userid)

	return role, int64(userid), err
}
