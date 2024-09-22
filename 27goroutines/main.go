package main

import (
	"fmt"
)

func main() {
	go say("Hello")
	say("World")

}
func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}
