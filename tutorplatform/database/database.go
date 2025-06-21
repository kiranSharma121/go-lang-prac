package database

import (
	"log"

	"github.com/tutorplatform/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := "host=localhost user=postgres password=kiran123 dbname=tutoring port=5432 sslmode=disable TimeZone=Asia/Kathmandu"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Couldn't connect to the database %v", err)
	}
	DB = database
	log.Println("connected to the database sucessfully")
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

}
