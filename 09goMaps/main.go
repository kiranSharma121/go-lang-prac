package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to maps in golang")
	myFruits := [3]string{"Apple", "Mango", "Orange"}
	fmt.Println("The list of myFruits", myFruits)
	fmt.Printf("The variable type of myFruits is:%T\n", myFruits)
	myVegis := []string{"Potato", "Tomato", "Ginger"}
	fmt.Println("The list of myFruits", myVegis)
	fmt.Printf("The variable type of myFruits is:%T\n", myVegis)
	myLang := make([]int, 3)
	myLang[0] = 12
	myLang[1] = 13
	myLang[2] = 23
	fmt.Println("", myLang)
	fmt.Printf("The variable type of mylang is:%T\n", myLang)
	///maps in golang
	mySubs := make(map[string]string)
	mySubs["JS"] = "Javascripts"
	mySubs["PY"] = "python"
	mySubs["RB"] = "Ruby"
	fmt.Println("The js stands for", mySubs["JS"])
	fmt.Printf("The variable type of the mysubs is:%T\n", mySubs)
	// delete(mySubs, "JS")
	// fmt.Println("The new maps is:", mySubs)
	for key, index := range mySubs {
		fmt.Printf("For the key %v\n,The value is %v\n", key, index)
	}

}
