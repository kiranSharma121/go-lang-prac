package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sqlite/database"
	"golang.org/x/crypto/bcrypt"
)

func (u *User) Save() error {
	query := `INSERT INTO users(email,password) VALUES(?,?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hasedpassword, err := HasedPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hasedpassword)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.Id = id
	return err
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
func (e Events) Delete() error {
	query := `DELETE FROM events WHERE id=?`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Id)
	return err
}
func HasedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CompareHasedPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
func (u User) ValidateCredentials() error {
	query := `SELECT id,email, password FROM users WHERE email=?`
	row := database.DB.QueryRow(query, u.Email)

	var email, retrievedPassword string
	err := row.Scan(&u.Id, &email, &retrievedPassword)
	if err != nil {
		return err
	}

	passwordIsValid := CompareHasedPassword(u.Password, retrievedPassword)
	if !passwordIsValid {
		return fmt.Errorf("invalid password")
	}

	return nil
}

var secretKey = []byte("meronamekiranho")

func GenerateJwtToken(email string, id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"id":     id,
		"expire": time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString(secretKey)
}
func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// Check if the signing method is HMAC
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unauthorized signing method")
		}
		return []byte(secretKey), nil
	})

	// Check for parsing error
	if err != nil {
		return 0, err
	}

	// Check if token is nil or invalid
	if parsedToken == nil || !parsedToken.Valid {
		return 0, err
	}
	return 0, err

}
