package models

import (
	"github.com/movie/database"
	"github.com/movie/utils"
)

func (u *User) Save() error {
	query := `INSERT INTO users(username,email,password,role)VALUES(?,?,?,?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	hasedpassword, err := utils.HasedPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.UserName, u.Email, hasedpassword, u.Role)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.Userid = id
	return err
}
