package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	welcome := "Welcome to file handing in goLang"
	fmt.Println(welcome)
	context := "Dipen muji randi hoo"
	file, err := os.Create("./mytext.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, context)
	fmt.Println("The length is", length)
	defer file.Close()
	databyte, err := os.ReadFile("./mytext.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("The contain in the file is:", string(databyte))
}
