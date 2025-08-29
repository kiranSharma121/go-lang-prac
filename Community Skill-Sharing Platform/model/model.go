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
	Role      string         `gorm:"type:varchar(20);check:role;IN('learner','mentor','admin';default;'learner')" json:"role"`
	Bio       string         `gorm:"size:200" json:"bio"`
	Skills    []Skill        `gorm:"foreginKey:User_id" json:"skills"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
type Skill struct {
}
