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
		panic("Unable to connect to the database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	CreateTable()
}

func CreateTable() {
	createUserTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		username TEXT NOT NULL UNIQUE ,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	);`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic("Couldn't create the users table")
	}

	createPostTable := `CREATE TABLE IF NOT EXISTS posts (
		postid INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL
	);`

	_, err = DB.Exec(createPostTable)
	if err != nil {
		panic("Couldn't create the posts table")
	}
}
