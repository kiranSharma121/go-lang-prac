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
	CreateTables()
}
func CreateTables() {
	createEventTable := `CREATE TABLE IF NOT EXISTS events(
Id INTEGER PRIMARY KEY AUTOINCREMENT,
username TEXT NOT NULL,
email TEXT NOT NULL,
password TEXT NOT NULL,
userid INTEGER,
time DATETIME NOT NULL
)`
	_, err := DB.Exec(createEventTable)
	if err != nil {
		panic("error in creating the table")
	}
}
