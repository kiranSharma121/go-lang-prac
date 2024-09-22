package main

import "fmt"

func main() {
	fmt.Println("welcome to the function in Go lang")
	result := masterAdder(2, 2, 2, 3, 3, 23, 2, 32, 3, 232, 11)
	fmt.Println("The result is:", result)
}
func masterAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total

}
