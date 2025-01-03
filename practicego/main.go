package main

import (
	"fmt"
	"math"

	"github.com/kiransharma121/greetings"
)

func main() {
	var name string
	fmt.Print("Enter your name:")
	fmt.Scan(&name)
	fmt.Println(greetings.Sayhello(name))
	const inflationRate = 2.2
	var investmentamount, years, expectedreturnrate float64
	fmt.Print("Enter the Investment Amount:")
	fmt.Scan(&investmentamount)
	fmt.Print("Enter the Years:")
	fmt.Scan(&years)
	fmt.Print("Enter the Expected Return Rate:")
	fmt.Scan(&expectedreturnrate)
	futureAmount := investmentamount * math.Pow(1+expectedreturnrate/100, years)
	futureRealAmount := futureAmount / math.Pow(1+inflationRate/100, float64(years))
	fmt.Println("The Future Amount is:", futureAmount)
	fmt.Println("The Future Real Amount is:", futureRealAmount)

}
