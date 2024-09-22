package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	welcome := "Welcome to the postform in golang"
	fmt.Println(welcome)
	postform("http://localhost:8080/postform")

}
func postform(myurl string) {
	data := url.Values{}
	data.Add("Name", "Kiran Sharma")
	data.Add("Age", "22")
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	databyte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("The postform request is given to the server is:", string(databyte))

}
