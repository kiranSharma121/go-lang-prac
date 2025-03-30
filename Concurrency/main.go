package main

import (
	"fmt"
	"time"
)

func main() {
	// dones := make([]chan bool, 3)
	done := make(chan bool)
	// dones[0] = make(chan bool)
	go slowGreet("How are you!", done)
	// dones[1] = make(chan bool)
	go Greet("Nice to meet you!", done)
	// dones[2] = make(chan bool)
	go Greet("How...are...you...???", done)
	for range done {

	}

}
func slowGreet(info string, done chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", info)
	done <- true
	close(done)
}
func Greet(info string, done chan bool) {
	fmt.Println("Hello!", info)
	done <- true

}
