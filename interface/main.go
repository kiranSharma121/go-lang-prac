package main

import "fmt"

func main() {
	printSomeThing("my name is kiran sharma and i'm from nepal")
	printSomeThing(2)
	printSomeThing(3.0)
}
func printSomeThing(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("It is an integer:", value)
	case float64:
		fmt.Println("It is an float type:", value)
	case string:
		fmt.Println("The value is string type:", value)
	}
}
