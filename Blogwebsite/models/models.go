package models

type Post struct {
	Postid  int    `json:"postid" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}
type User struct {
	Id       int64  `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
