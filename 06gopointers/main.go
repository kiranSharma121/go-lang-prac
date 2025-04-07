package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointers chapter")
	myNumber := 12
	myPointer := &myNumber
	// fmt.Println("The memory address of the myNumber is:", myPointer)
	fmt.Println("The actual value stored in the myPointer is:", *myPointer)
	fmt.Println("The memory address of the myNumber is:", &myNumber)

}
