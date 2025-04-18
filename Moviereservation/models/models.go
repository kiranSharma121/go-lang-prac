package models

import "time"

type User struct {
	Userid   int    `json:"userid"`
	UserName string `json:"username" binding:"require"`
	Email    string `json:"email" binding:"require"`
	Password string `json:"password" binding:"require"`
	Role     string `json:"role" binding:"require"`
}
type Movie struct {
	Movieid     int    `json:"movieid"`
	Title       string `json:"title" binding:"require"`
	Description string `json:"description" binding:"require"`
	Genre       string `json:"genre" binding:"require"`
}
type Showtime struct {
	Showtimeid int `json:"showtimeid"`
	Movieid    int `json:"movieid"`
	StartAT    time.Time
	Capacity   int `json:"capacity" binding:"require"`
}
type Seat struct {
	Seatid     int    `json:"seatid"`
	SeatNumber string `json:"seatnumber" binding:"require"`
	IsBooked   bool   `json:"isbooked"`
}
