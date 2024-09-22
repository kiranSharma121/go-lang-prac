package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the golang slice")

	myFruits := []string{"Apple", "Banana", "pineapple", "Orange"}
	fmt.Println("The slice is:", myFruits)
	fmt.Printf("The variable type of the myFruits is:%T\n", myFruits)

	examMarks := make([]int, 4)
	examMarks[0] = 94
	examMarks[1] = 45
	examMarks[2] = 34
	examMarks[3] = 76
	fmt.Println("The slice of examMarks is:", examMarks)
	examMarks = append(examMarks, 89, 70, 85)
	fmt.Println("The slice of examMarks after append:", examMarks)
	sort.Ints(examMarks)
	fmt.Println("The slice of examMarks after sort:", examMarks)
	fmt.Println(sort.IntsAreSorted(examMarks))
	examMarks = append(examMarks[1:])
	fmt.Println("The sliced examMarks is:", examMarks)

	var myVegis = []string{"potato", "Tomato", "ginger", "radice"}
	myVegis = append(myVegis[:1], myVegis[2:]...)
	fmt.Println("The removed vegi", myVegis)

}
