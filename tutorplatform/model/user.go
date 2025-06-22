package model

type User struct {
	Id       int    `json:"id" gorm:"primarykey"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" gorm:"unique" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}
