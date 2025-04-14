package main

import (
	"fmt"

	"github.com/goVendor/database"
)

func main() {
	database.InitDB()
	fmt.Println("welcome to the goVendor")
}
