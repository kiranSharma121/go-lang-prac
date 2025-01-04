package main

import "fmt"

func main() {
	var Revenue, Expenses, Tax float64
	fmt.Print("Enter The Revenue:")
	fmt.Scan(&Revenue)
	fmt.Print("Enter The Expenses:")
	fmt.Scan(&Expenses)
	fmt.Print("Enter the Tax Rate:")
	fmt.Scan(&Tax)
	EBT := Revenue - Expenses
	Profit := EBT * (1 - Tax/100)
	Ratio := EBT / Profit
	fmt.Println("The Earning Before Tax is:", EBT)
	fmt.Println("The Profit is:", Profit)
	fmt.Println("The Ratio is:", Ratio)

}
