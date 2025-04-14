package models

import (
	"fmt"
	"time"

	"github.com/goVendor/database"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (u *User) Save() error {
	query := `INSERT INTO users(username,email,password)VALUES(?,?,?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hasedPassword, err := HasedPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.UserName, u.Email, hasedPassword) //insert data in the database exec
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.Id = id
	return err
}
func HasedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func ComparePassword(password, hasedpassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hasedpassword), []byte(password))
	return err == nil

}
func (u User) ValidateCredentials() error {
	query := `SELECT email,password FROM users WHERE email=?`
	row := database.DB.QueryRow(query, u.Email)
	var email, retrivePassword string
	err := row.Scan(&email, &retrivePassword)
	if err != nil {
		return err
	}
	isPasswordValid := ComparePassword(u.Password, retrivePassword)
	if !isPasswordValid {
		return fmt.Errorf("invalid password")
	}
	return nil

}

var secretKey = []byte("kiran is my name")

func GenerateJwtToken(username, email string, id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"email":    email,
		"id":       id,
		"expire":   time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString(secretKey)
}
