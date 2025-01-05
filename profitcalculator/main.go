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
	fmt.Printf("The Earning Before Tax is:%.1f\n", EBT)
	fmt.Printf("The Profit is:%.1f\n", Profit)
	fmt.Printf("The Ratio is:%.1f\n", Ratio)

}
