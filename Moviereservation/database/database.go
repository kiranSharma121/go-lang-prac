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
	CreateDatabase()
}
func CreateDatabase() {
	CreateUserTable := `CREATE TABLE IF NOT EXISTS users(
userid INTEGER PRIMARY KEY AUTOINCREMENT,
username TEXT NOT NULL,
email TEXT NOT NULL UNIQUE,
password TEXT NOT NULL,
role TEXT NOT NULL
)`
	_, err := DB.Exec(CreateUserTable)
	if err != nil {
		panic("couldn't able to create user table")
	}
	CreateMovieTable := `CREATE TABLE IF NOT EXISTS movies(
	movieid INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL UNIQUE,
	description TEXT NOT NULL,
	genre TEXT NOT NULL,
	userid INTEGER NOT NULL, 
	FOREIGN KEY (userid) REFERENCES users(userid)
	
	)`
	_, err = DB.Exec(CreateMovieTable)
	if err != nil {

		panic("couldn't able to create movie table")
	}
	CreateShowTimeTable := `CREATE TABLE IF NOT EXISTS shows(
	showtimeid INTEGER PRIMARY KEY AUTOINCREMENT,
	movieid INTEGER NOT NULL,
	userid INTEGER NOT NULL,
	startat DATETIME NOT NULL,
	capacity INTEGER NOT NULL,
	FOREIGN KEY (movieid) REFERENCES movies(movieid),
	FOREIGN KEY(userid)REFERENCES users(usersid)
	)`
	_, err = DB.Exec(CreateShowTimeTable)
	if err != nil {
		panic("couldn;t able to create the showtime table")
	}
	CreateSeatTable := `CREATE TABLE IF NOT EXISTS seats(
	seatid INTEGER PRIMARY KEY AUTOINCREMENT,
	seatnumber TEXT NOT NULL,
	isbooked BOOL NOT NULL
	)`
	_, err = DB.Exec(CreateSeatTable)
	if err != nil {
		panic("couldn't able to create the seat table")
	}
	CreateTableBooking := `CREATE TABLE IF NOT EXISTS booking(
	bookingid INTEGER PRIMARY KEY AUTOINCREMENT,
	userid INTEGER NOT NULL,
	seatid INTEGER NOT  NULL,
	showtimeid INTEGER NOT NULL,
	bookat DATETIME NOT NULL,
	FOREIGN KEY (userid) REFERENCES users(userid) ON DELETE CASCADE,
    FOREIGN KEY (seatid) REFERENCES seats(seatid) ON DELETE CASCADE,
    FOREIGN KEY (showtimeid) REFERENCES showtimes(showtimeid) ON DELETE CASCADE
	)`
	_, err = DB.Exec(CreateTableBooking)
	if err != nil {
		panic("couldn't able to creat the Table Booking")
	}

}
