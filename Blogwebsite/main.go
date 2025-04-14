package main

import (
	"fmt"

	"github.com/goVendor/database"
	"github.com/goVendor/router"
)

func main() {
	database.InitDB()
	fmt.Println("welcome to the goVendor")
	server := router.Router()
	fmt.Println("server is running....")
	server.Run(":8080")

}
