package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("welcome to the function in go..")
	var investmentamount, years, expectedreturnrate float64
	fmt.Print("Enter The Investment Amount:")
	fmt.Scan(&investmentamount)
	fmt.Print("Enter The Years:")
	fmt.Scan(&years)
	fmt.Print("Enter the expectedreturnrate:")
	fmt.Scan(&expectedreturnrate)
	futureAmount, futureRealAmount := Result(investmentamount, years, expectedreturnrate)
	fmt.Println("The futureAmount is:", futureAmount)
	fmt.Println("The futureRealAmount is:", futureRealAmount)
}

const inflationRate = 2.2

func Result(investmentamount, years, expectedreturnrate float64) (float64, float64) {
	futureAmount := investmentamount * math.Pow(1+expectedreturnrate/100, years)
	futureRealAmount := futureAmount / math.Pow(1+inflationRate/100, years)
	return futureAmount, futureRealAmount
}
