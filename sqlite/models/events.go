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
	UserId   int
	Time     time.Time
}

func (e *Events) Save() error {
	query := `INSERT INTO events(username, email, password, userid, time) VALUES (?, ?, ?, ?, ?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Username, e.Email, e.Password, e.UserId, e.Time)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.Id = id
	return nil
}

func GetAllEvents() ([]Events, error) {
	query := `SELECT * FROM events`
	row, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var events []Events
	for row.Next() {
		var event Events
		err = row.Scan(&event.Id, &event.Username, &event.Email, &event.Password, &event.UserId, &event.Time)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
