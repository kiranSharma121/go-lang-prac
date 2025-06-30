package controller

import (
	"errors"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Hashpassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
func Comparepassword(password, hashedpassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
	return err == nil
}
func Generatejwttoken(id int, name, email, role string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"name":  name,
		"email": email,
		"role":  role,
	})
	return token.SignedString([]byte(secretKey))
}
func Verifyjwttoken(tokenString string) (string, int, error) {
	secretKey := os.Getenv("JWT_SECRET")

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

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
		return "", 0, errors.New("invalid jwt token claims")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return "", 0, errors.New("role not found in the token")
	}

	idFloat, ok := claims["id"].(float64)
	if !ok {
		return "", 0, errors.New("id not found in the token")
	}
	id := int(idFloat)

	return role, id, nil
}
