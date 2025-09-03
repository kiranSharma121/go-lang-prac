package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/skillplatform/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("fail to load .env file")
	}
	dsn := os.Getenv("DATABASE_URL")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("couldn't connect to the database %v", err)
	}
	DB = database
	log.Println("Database connect successfully")
	err = database.AutoMigrate(&model.User{}, &model.Skill{}, &model.Session{}, &model.Message{}, &model.Notification{}, &model.Enrollment{})
	if err != nil {
		fmt.Println("Migration fail", err)
	}
}
