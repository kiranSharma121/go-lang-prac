package main

import "fmt"

type transformfn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println("The slice of number after double is:", tranformNumber(&numbers, double))
	fmt.Println("The slice of number after double is:", tranformNumber(&numbers, triple))

}
func tranformNumber(number *[]int, transform transformfn) []int {
	dbNumber := []int{}
	for _, value := range *number {
		dbNumber = append(dbNumber, transform(value))
	}
	return dbNumber
}
func double(number int) int {
	return number * 2
}
func triple(number int) int {
	return number * 3
}
