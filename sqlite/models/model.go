package models

import "time"

type Events struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Userid   int    `json:"userid" binding:"required"`
	Time     time.Time
}
type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
