package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:100;not null" json:"name"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"not null" json:"password"`
	Role      string         `gorm:"type:varchar(20);check:role IN ('learner','mentor','admin');default:'learner'" json:"role"`
	Bio       string         `gorm:"size:200" json:"bio"`
	Skills    []Skill        `gorm:"foreignKey:UserID" json:"skills"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Skill struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Session struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	LearnerID uint      `json:"learner_id"`
	MentorID  uint      `json:"mentor_id"`
	Time      time.Time `json:"time"`
	Status    string    `gorm:"type:varchar(20);check:status IN ('pending','accepted','rejected','completed');default:'pending'" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Message struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	SenderID   uint      `json:"sender_id"`
	ReceiverID uint      `json:"receiver_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}

type Notification struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	Type      string    `json:"type"`
	Content   string    `json:"content"`
	IsRead    bool      `gorm:"default:false" json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}
type Enrollment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	SkillID   uint      `json:"skill_id"`
	Skill     Skill     `gorm:"foreignKey:SkillID" json:"skill"`
	LearnerID uint      `json:"learner_id"`
	Learner   User      `gorm:"foreignKey:LearnerID" json:"learner"`
	CreatedAt time.Time `json:"created_at"`
}
