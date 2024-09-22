package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	welcome := "welcome to the get method in golang"
	fmt.Println(welcome)
	// getMethod("http://localhost:8080/get")
	getMethod("http://localhost:8080")

}
func getMethod(url string) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	dataByte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("The containt we get is:", string(dataByte))
	fmt.Println("The status of the contain is:", response.StatusCode)
	fmt.Println("The length of the contain is:", response.ContentLength)

}
