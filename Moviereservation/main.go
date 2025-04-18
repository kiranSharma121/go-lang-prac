package main

import (
	"fmt"

	"github.com/movie/database"
	"github.com/movie/routers"
)

func main() {
	database.InitDB()
	server := routers.Router()
	server.Run(":8080")
	fmt.Println("welcome to movie reservation in golang")
}
