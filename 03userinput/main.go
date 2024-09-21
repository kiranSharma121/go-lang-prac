package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the user input")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the ratings:")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.TrimSpace(input)
	ratings, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	if ratings < 1 || ratings > 5 {
		fmt.Println("Please enter the valid ratings")
	} else {
		fmt.Println("Thank you for the ratings", ratings, "stars")

	}
	fmt.Printf("The type of the rating is:%T\n", ratings)

}
