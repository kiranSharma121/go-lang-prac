package model

type User struct {
	Id       int    `json:"id" gorm:"primarykey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
