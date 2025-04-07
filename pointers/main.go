package main

import "fmt"

func main() {
	var age int = 32
	var agePointer *int = &age
	fmt.Println("Age:", age)
	realage := ageCalculator(agePointer)
	fmt.Println("The realAge is:", realage)

}
func ageCalculator(age *int) int {
	return *age - 18
}
