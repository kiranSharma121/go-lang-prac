package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/movie/database"
	"github.com/movie/routers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.InitDB()
	server := routers.Router()
	server.Run(":8080")
	fmt.Println("welcome to movie reservation in golang")
}
