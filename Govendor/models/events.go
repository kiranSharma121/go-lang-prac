package models

import "github.com/goVendor/database"

func (u *User) Save() error {
	query := `INSERT INTO users(email,password)VALUES(?,?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.Id = id
	return err
}
