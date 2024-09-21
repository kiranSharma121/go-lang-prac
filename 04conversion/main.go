package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the conversion int go lang")
	userinput := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the hostel:")
	input, err := userinput.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.TrimSpace(input)
	ratings, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	if ratings < 1 || ratings > 5 {
		fmt.Println("Invalid ratings")
	} else {
		fmt.Println("Thank you for rating:", ratings+1)
	}
}
