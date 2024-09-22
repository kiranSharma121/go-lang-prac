package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcomet to the goifelse")
	userinput := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your marks:")
	input, err := userinput.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.TrimSpace(input)
	passMarks, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	var result string
	if passMarks < 40 {
		result = "You got less then 40 marks!Fail"

	} else if passMarks > 40 {
		result = "You got more then 40 marks!! pass"

	} else {
		result = "you got exact 40 marks"
	}
	fmt.Println(result)

}
