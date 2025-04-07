package main

import (
	"fmt"
	"os"
)

const fileName = "result.txt"

func main() {
	Revenue, Expenses, Tax := Userinput()
	EBT, Profit, Ratio := Result(Revenue, Expenses, Tax)
	fmt.Printf("The Earning Before Tax is:%.1f\n", EBT)
	fmt.Printf("The Profit is:%.1f\n", Profit)
	fmt.Printf("The Ratio is:%.1f\n", Ratio)
	storeResult(EBT, Profit, Ratio)

}
func storeResult(EBT, Profit, Ratio float64) {
	ResultText := fmt.Sprintf("EBT:%1.f\nProfit:%1.f\nRatio:%1.f", EBT, Profit, Ratio)
	os.WriteFile(fileName, []byte(ResultText), 0644)

}

func Userinput() (float64, float64, float64) {
	var Revenue, Expenses, Tax float64
	fmt.Print("Enter The Revenue:")
	fmt.Scan(&Revenue)
	if Revenue <= 0 {
		panic("Invalid input")
	}
	fmt.Print("Enter The Expenses:")
	fmt.Scan(&Expenses)
	if Expenses <= 0 {
		panic("Invalid input")
	}
	fmt.Print("Enter the Tax Rate:")
	fmt.Scan(&Tax)
	if Tax <= 0 {
		panic("Invalid input")
	}
	return Revenue, Expenses, Tax

}
func Result(Revenue, Expenses, Tax float64) (float64, float64, float64) {
	EBT := Revenue - Expenses
	Profit := EBT * (1 - Tax/100)
	Ratio := EBT / Profit
	return EBT, Profit, Ratio

}

// func main() {
// 	Revenue := Userinput("Enter the Revenue:")
// 	Expenses := Userinput("Enter the Expenses:")
// 	Tax := Userinput("Enter the Tax:")
// 	EBT, Profit, Ratio := Result(Revenue, Expenses, Tax)

// 	fmt.Printf("The Earning Before Tax is:%.1f\n", EBT)
// 	fmt.Printf("The Profit is:%.1f\n", Profit)
// 	fmt.Printf("The Ratio is:%.1f\n", Ratio)

// }
// func Userinput(text string) float64 {
// 	var userinput float64
// 	fmt.Print(text)
// 	fmt.Scan(&userinput)
// 	return userinput

// }
// func Result(Revenue, Expenses, Tax float64) (float64, float64, float64) {
// 	EBT := Revenue - Expenses
// 	Profit := EBT * (1 - Tax/100)
// 	Ratio := EBT / Profit
// 	return EBT, Profit, Ratio

// }
