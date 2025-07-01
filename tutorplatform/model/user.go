package model

type User struct {
	ID      uint    `json:"id" gorm:"primarykey"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" gorm:"unique" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}
type Course struct {
	ID     uint    `json:"id" gorm:"primarykey"`
	Title   string `json:"title" gorm:"unique"`
	Content string `json:"content"`
	TutorID uint    `json:"tutorid"`
}
type Enrollment struct {
	Id        uint    `gorm:"primarykey"`
	CourseID  uint    `json:"courseid"`
	StudentID uint    `json:"studentid"`
	Course    Course `gorm:"foreignKey:CourseID"`
	Student   User   `gorm:"foreignKey:StudentID"`
}
