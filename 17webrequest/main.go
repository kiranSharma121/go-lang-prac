package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "https://example.com"

func main() {
	fmt.Println("Welcome the web request handing in go lang")
	res, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	dataByte, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("The web request we get", string(dataByte))
}
