package controller

import (
	"errors"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
func CompareHashPassword(password, hashedpassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
	return err == nil
}
func GenerateJwtToken(id uint, name, email, role string) (string, error) {
	secretKey := os.Getenv("JWT_SECRECT")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"name":  name,
		"email": email,
		"role":  role,
	})
	return token.SignedString([]byte(secretKey))
}
func VerifyJwtToken(tokenString string) (string, int, error) {
	secretKey := os.Getenv("JWT_SECRECT")
	tokenString = strings.TrimPrefix(tokenString, "Bearer")
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
		return "", 0, errors.New("invalid jwt tokens claims")
	}
	role, ok := claims["role"].(string)
	if !ok {
		return "", 0, errors.New("role is not found in the token")
	}
	idFloat, ok := claims["id"].(float64)
	if !ok {
		return "", 0, errors.New("id not found in the token")
	}
	id := int(idFloat)
	return role, id, err
}
