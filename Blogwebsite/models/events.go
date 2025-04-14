package models

import (
	"fmt"

	"github.com/goVendor/database"
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
