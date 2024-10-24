package main

import (
	"fmt"

	"github.com/kiransharma121/gin/router"
)

func main() {
	fmt.Println("welcome to gin framework")
	r := router.Router()
	r.Run()
}
