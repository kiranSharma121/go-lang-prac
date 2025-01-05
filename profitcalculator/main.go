package main

import "fmt"

func main() {
	Revenue, Expenses, Tax := Userinput()
	EBT, Profit, Ratio := Result(Revenue, Expenses, Tax)
	fmt.Printf("The Earning Before Tax is:%.1f\n", EBT)
	fmt.Printf("The Profit is:%.1f\n", Profit)
	fmt.Printf("The Ratio is:%.1f\n", Ratio)

}
func Userinput() (float64, float64, float64) {
	var Revenue, Expenses, Tax float64
	fmt.Print("Enter The Revenue:")
	fmt.Scan(&Revenue)
	fmt.Print("Enter The Expenses:")
	fmt.Scan(&Expenses)
	fmt.Print("Enter the Tax Rate:")
	fmt.Scan(&Tax)
	return Revenue, Expenses, Tax

}
func Result(Revenue, Expenses, Tax float64) (float64, float64, float64) {
	EBT := Revenue - Expenses
	Profit := EBT * (1 - Tax/100)
	Ratio := EBT / Profit
	return EBT, Profit, Ratio

}
