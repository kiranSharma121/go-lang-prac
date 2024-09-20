package main

import "fmt"

func main() {
	var personAge int = 25
	fmt.Println("The age of the person is:", personAge)
	fmt.Printf("The variable type of the person age is:%T\n", personAge)

	var isSignedup bool = true
	fmt.Println("The signeup status is:", isSignedup)
	fmt.Printf("The variable type is:%T\n", isSignedup)

	var nameOfPerson = "kiran sharma"
	fmt.Println("Name of the person is:", nameOfPerson)
	fmt.Printf("The variable type is:%T\n", nameOfPerson)

	hostelfee := 100000
	fmt.Println("The hostel amount is:", hostelfee)
	fmt.Printf("The variable type is:%T\n", hostelfee)
}
