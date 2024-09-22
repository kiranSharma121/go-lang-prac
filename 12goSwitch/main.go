package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to the switch case")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter The number from 1 to 6:")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.TrimSpace(input)
	Userinput, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	if Userinput < 0 || Userinput > 6 {
		fmt.Println("please enter the correct value")
	}
	switch Userinput {
	case 1:
		fmt.Println("you got one ")
	case 2:
		fmt.Println("you got two ")
	case 3:
		fmt.Println("you got three ")
	case 4:
		fmt.Println("you got four ")
	case 5:
		fmt.Println("you got five  ")
	case 6:
		fmt.Println("you got six ")

	}

}
