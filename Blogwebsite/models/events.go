package models

import (
	"errors"
	"fmt"

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
func (p *Post) Save() error {
	query := `INSERT INTO posts(author,title,content)VALUES(?,?,?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(p.Author, p.Title, p.Content)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	p.Postid = id
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

var secretKey = []byte("kiran is my name")

func GenerateJwtToken(id int64, username, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"email":    email,
	})
	return token.SignedString(secretKey)
}
func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unauthorized signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}
	if parsedToken == nil || !parsedToken.Valid {
		return 0, err
	}
	return 0, nil
}
