package main

import "fmt"

func main() {
	fmt.Println("Welcome to the goArrays")
	var myFruits [4]string
	myFruits[0] = "Apple"
	myFruits[1] = "pineapple"
	myFruits[2] = "Orange"
	myFruits[3] = "Banana"
	fmt.Println("The arrays of the myfruits is:", myFruits)

	var myvegis = [3]string{"Potato", "cauliflower", "ladyfingers"}
	fmt.Println("The array list of the myVegis is:", myvegis)
}
