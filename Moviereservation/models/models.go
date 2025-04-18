package models

import "time"

type User struct {
	Userid   int64  `json:"userid"`
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}
type Movie struct {
	Movieid     int    `json:"movieid"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Genre       string `json:"genre" binding:"required"`
}
type Showtime struct {
	Showtimeid int `json:"showtimeid"`
	Movieid    int `json:"movieid"`
	StartAT    time.Time
	Capacity   int `json:"capacity" binding:"required"`
}
type Seat struct {
	Seatid     int    `json:"seatid"`
	SeatNumber string `json:"seatnumber" binding:"required"`
	IsBooked   bool   `json:"isbooked"`
}
