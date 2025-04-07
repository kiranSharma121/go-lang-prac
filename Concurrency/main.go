package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go Greeter("How are you??", done)
	go slowGreeter("welcome to the golang", done)
	go Greeter("how....are...you???", done)
	for range done {

	}
}
func Greeter(info string, done chan bool) {
	fmt.Println("Hello", info)
	done <- true
}
func slowGreeter(info string, done chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello", info)
	done <- true
	close(done)
}
