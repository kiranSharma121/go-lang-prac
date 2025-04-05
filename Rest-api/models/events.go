package models

import "time"

type Events struct {
	Id       int
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserId   int
	Time     time.Time
}

var events = []Events{}

func (e Events) Save() {
	events = append(events, e)
}

func GetAllEvents() []Events {
	return events
}
