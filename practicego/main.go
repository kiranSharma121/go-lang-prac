// package main

// import (
// 	"fmt"
// 	"math"

// 	"github.com/kiransharma121/greetings"
// )

// func main() {
// 	var name string
// 	fmt.Print("Enter your name:")
// 	fmt.Scan(&name)
// 	fmt.Println(greetings.Sayhello(name))
// 	const inflationRate = 2.2
// 	var investmentamount, years, expectedreturnrate float64
// 	fmt.Print("Enter the Investment Amount:")
// 	fmt.Scan(&investmentamount)
// 	fmt.Print("Enter the Years:")
// 	fmt.Scan(&years)
// 	fmt.Print("Enter the Expected Return Rate:")
// 	fmt.Scan(&expectedreturnrate)
// 	futureAmount := investmentamount * math.Pow(1+expectedreturnrate/100, years)
// 	futureRealAmount := futureAmount / math.Pow(1+inflationRate/100, float64(years))
// 	fmt.Println("The Future Amount is:", futureAmount)
// 	fmt.Println("The Future Real Amount is:", futureRealAmount)

// }
// package main

// import "fmt"

//	func main() {
//		fmt.Println("welcome to the slice in go")
//		var Hobbies = [3]string{"playing cricket", "Table Tennis", "Badminton"}
//		fmt.Println("The hobbies are:", Hobbies)
//		fmt.Println(Hobbies[0:1])
//		fmt.Println(Hobbies[1:3])
//	}
package main

import "fmt"

func main() {
	fmt.Println("welcome to the map in go")
	// websites := make(map[string]string)
	// websites["Google"] = "https://google.com"
	// websites["FaceBook"] = "https://facebook.com"
	// fmt.Println(websites)
	// websites[0] = 1
	// websites[1] = 2
	// websites[2] = 3
	// websites[3] = 4
	// fmt.Println(websites)
	// fmt.Printf("%T\n", websites)
	websites := map[string]string{
		"Google":   "https://google.com",
		"Facebook": "https://facebook.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Facebook"])
	Fruits := make([]int, 4)
	Fruits[0] = 1
	Fruits[1] = 2
	Fruits[2] = 3
	Fruits[3] = 4
	fmt.Println(Fruits)

}
