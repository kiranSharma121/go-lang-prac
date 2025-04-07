// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Println("welcome to profit calculator in go")
// 	revenue := userInfo("Enter the revenue amount:")
// 	expenses := userInfo("Enter the expenses amount:")
// 	taxRate := userInfo("Enter the tax rate:")
// 	caluculation(revenue, expenses, taxRate)
// }
// func writeDataToFile(EBT, profit, ratio float64) {
// 	dataTxt := fmt.Sprintf("EBT:%2f\nprofit:%2f\nratio:%2f\n", EBT, profit, ratio)
// 	os.WriteFile("Balance.txt", []byte(dataTxt), 0644)
// }
// func caluculation(revenue, expenses, taxRate float64) {
// 	EBT := revenue - expenses
// 	fmt.Println("The earning before the tax is:", EBT)
// 	profit := EBT * (1 - taxRate/100)
// 	fmt.Println("The profit is:", profit)
// 	ratio := EBT / profit
// 	fmt.Println("The ratio is:", ratio)
// 	writeDataToFile(EBT, profit, ratio)
// }
// func userInfo(info string) float64 {
// 	fmt.Print(info)
// 	var data float64
// 	fmt.Scan(&data)
// 	if data <= 0 {
// 		fmt.Println("Invalid amount")
// 	}
// 	return data
// }

package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	fileops "github.com/hello/fileOps"
)

var fileName = "Balance.txt"

func main() {
	fmt.Println("Welcome to Go Bank")
	fmt.Println("Reach out us 24/7", randomdata.Address())
	var accountBalance float64 = fileops.GetDataFromFile(fileName)
	for {
		greet()
		var choice int
		fmt.Scan(&choice)
		if choice == 1 {
			fmt.Println("Your Balance is:", accountBalance)
		} else if choice == 2 {
			fmt.Print("Enter the amount that you want to deposite:")
			var amount float64
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}
			accountBalance += amount
			fileops.WriteDataToFile(accountBalance, fileName)
			fmt.Println("Your new amount is:", accountBalance)
		} else if choice == 3 {
			fmt.Print("Enter the amount that you want to withdraw:")
			var amount float64
			fmt.Scan(&amount)
			if amount <= 0 || amount > accountBalance {
				fmt.Println("Invalid amount")
				continue
			}
			accountBalance -= amount
			fileops.WriteDataToFile(accountBalance, fileName)
			fmt.Println("Your new balance is:", accountBalance)

		} else if choice == 4 {
			fmt.Println("welcome back to the home page")
			break
		} else {
			fmt.Println("Invalid entry")
			break
		}
	}
	fmt.Println("Thank you for visting our bank...")
}
