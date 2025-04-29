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
	Movieid     int64  `json:"movieid"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Genre       string `json:"genre" binding:"required"`
	Userid      int64  `json:"userid"`
}
type Showtime struct {
	Showtimeid int64     `json:"showtimeid"`
	Movieid    int64     `json:"movieid"`
	Userid     int64     `json:"userid"`
	StartAT    time.Time `json:"startat"`
	Capacity   int64     `json:"capacity" binding:"required"`
}
type Seat struct {
	Seatid     int64   `json:"seatid"`
	SeatNumber string `json:"seatnumber" binding:"required"`
	IsBooked   bool   `json:"isbooked"`
}
type Booking struct {
	BookingID  int64     `json:"bookingid"`
	UserID     int64     `json:"userid"`
	SeatID     int       `json:"seatid" binding:"required"`
	ShowtimeID int       `json:"showtimeid" binding:"required"`
	BookedAt   time.Time `json:"bookedat"`
}
