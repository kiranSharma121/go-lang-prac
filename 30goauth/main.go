package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kiransharma121/goauth/router"
)

func main() {
	fmt.Println("Dada ho dada, Sargam Dada!!!")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
}
