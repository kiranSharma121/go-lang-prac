package main

import (
	"fmt"

	"github.com/kiransharma121/Authorization/router"
)

func main() {
	fmt.Println("Hello world")
	r := router.Router()
	fmt.Println("Server is running")
	r.Run(":4000")
}
