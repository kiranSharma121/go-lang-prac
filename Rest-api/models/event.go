package models

import (
	"time"

	"github.com/restapi/database"
)

type Events struct {
	Id       int64
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserId   int
	Time     time.Time
}

func (e *Events) Save() error {
	query := `INSERT INTO events(
	email,username,password,userid,time
	) VALUES(?,?,?,?,?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Email, e.Username, e.Password, e.UserId, e.Time)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.Id = id
	return err
}

func GetAllEvents() ([]Events, error) {
	query := "SELECT id, email, username, password, userid, time FROM events"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Events
	for rows.Next() {
		var event Events
		err := rows.Scan(
			&event.Id,
			&event.Email,
			&event.Username,
			&event.Password,
			&event.UserId,
			&event.Time,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
func GetEventById(id int) ([]Events, error) {
	query := "SELECT * FROM events WHERE id=?"
	row := database.DB.QueryRow(query, id)
	var event Events
	err := row.Scan(&event.Id, &event.Email, &event.Username, &event.Password, &event.UserId, &event.Time)
	if err != nil {
		return nil, err
	}
	return []Events{event}, nil
}
func (event Events) Update() error {
	query := `UPDATE events
	SET email=?, username=?, password=?, userid=?, time=?
	WHERE ID=?`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Email, event.Username, event.Password, event.UserId, event.Time, event.Id)
	return err
}
func (event Events) Delete() error {
	query := `DELETE FROM events WHERE id = ?`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Id)
	return err
}
