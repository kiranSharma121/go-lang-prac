package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	welcome := "welcome to the post request in golang"
	fmt.Println(welcome)
	postRequest("http://localhost:8080/post")
}
func postRequest(url string) {
	requestBody := strings.NewReader(`
	{
	"Name":"Kiran Sharma",
	"Age":22

	}`)
	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	databyte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("The data added to the server is:", string(databyte))

}
