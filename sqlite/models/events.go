package models

import (
	"time"

	"github.com/sqlite/database"
)

type Events struct {
	Id       int64
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Userid   int    `json:"userid" binding:"required"`
	Time     time.Time
}

func (e *Events) Save() error {
	query := "INSERT INTO events(username,email,password,userid,time)VALUES(?,?,?,?,?)"
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Username, e.Email, e.Password, e.Userid, e.Time) //when we insert data in the database
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.Id = id
	return err

}

func GetAllEvents() ([]Events, error) {
	query := `SELECT * FROM events`
	row, err := database.DB.Query(query) //when we fetch data from the database
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var events []Events
	for row.Next() {
		var event Events
		err := row.Scan(&event.Id, &event.Username, &event.Email, &event.Password, &event.Userid, &event.Time)
		if err != nil {
			return nil, err
		}
		events = append(events, event)

	}

	return events, nil
}
func GetEventByID(id int) (*Events, error) {
	query := "SELECT * FROM events WHERE id=?"
	row := database.DB.QueryRow(query, id)
	var event Events
	err := row.Scan(&event.Id, &event.Username, &event.Email, &event.Password, &event.Userid, &event.Time)
	if err != nil {
		return nil, err
	}
	return &event, err
}
func (e Events) UpDateEvents() error {
	query := `UPDATE events SET username=?,email=?,password=?,userid=?,time=?
	WHERE id=?`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Id, e.Username, e.Email, e.Password, e.Userid, e.Time) //when we insert data in the database
	return err
}
