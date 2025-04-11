package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("couldn't connect to the database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	CreateEventTable()
}
func CreateEventTable() {

	CreateUserTable := `CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`
	_, err := DB.Exec(CreateUserTable)
	if err != nil {
		panic("couldn't create the  user table")
	}
	CreateEvent := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		userid INTEGER NOT NULL,
		time DATETIME NOT NULL,
		FOREIGN KEY (userid) REFERENCES users(id)
	);
	`
	_, err = DB.Exec(CreateEvent) //exec when we create the database
	if err != nil {
		panic("couldn't create the event table")
	}

}
