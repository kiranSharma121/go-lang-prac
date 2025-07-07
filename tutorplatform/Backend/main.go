package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/tutorplatform/database"
	"github.com/tutorplatform/router"
)

func main() {
	fmt.Println("tutor platform using golang")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.Connection()
	r := router.Router()
	r.Run(":8080")

}
