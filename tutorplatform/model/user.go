package model

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" gorm:"unique" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}
type Course struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Title   string `json:"title" gorm:"unique"`
	Content string `json:"content"`
	TutorID uint   `json:"tutorid"`
}
type Enrollment struct {
	Id        uint   `gorm:"primaryKey"`
	CourseID  uint   `json:"courseid"`
	StudentID uint   `json:"studentid"`
	Course    Course `gorm:"foreignKey:CourseID;references:ID"`
	Student   User   `gorm:"foreignKey:StudentID;references:ID"`
}
