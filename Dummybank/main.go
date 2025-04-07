package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const fileName = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, err
	}
	balancestring := string(data)
	balance, err := strconv.ParseFloat(balancestring, 64)
	if err != nil {
		return 1000, err
	}
	return balance, nil

}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}
func main() {
	var CurrentBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("-----")
	}
	var passcode string
	var pincode = "1234"
	var userName string
	fmt.Print("Enter Your FullName:")
	reader := bufio.NewReader(os.Stdin)
	userName, _ = reader.ReadString('\n')
	userName = strings.TrimSpace(userName)
	fmt.Println("Welcome", userName, "to the Go bank!!!")

	for {
		fmt.Print("Enter your Pincode:")
		fmt.Scan(&passcode)
		if passcode != pincode {
			fmt.Println("Wrong pincode")
			continue
		}
		PresentCode()
		var choice int
		fmt.Print("Enter your choice:")
		fmt.Scan(&choice)
		// var isyourchoice = choice == 1
		if choice == 1 {
			fmt.Println("Your current balance is:", CurrentBalance)
		} else if choice == 2 {
			var Deposite float64
			fmt.Print("Enter the Amount:")
			fmt.Scan(&Deposite)
			if Deposite > 500000 || Deposite < 1 {
				fmt.Println("Can't deposite this amount of money")
				continue
			}
			CurrentBalance += Deposite
			fmt.Println("Balance updated!Your new balance is:", CurrentBalance)
			writeBalanceToFile(CurrentBalance)
		} else if choice == 3 {
			var Withdraw float64
			fmt.Print("Enter the Amount:")
			fmt.Scan(&Withdraw)
			if Withdraw > CurrentBalance {
				fmt.Println("Insufficient Balance")
				continue
			} else if Withdraw < 1 {
				fmt.Println("can't withdraw this amount of money")
				continue

			}
			CurrentBalance -= Withdraw
			fmt.Println("Balance updated!Your new balance is:", CurrentBalance)
			writeBalanceToFile(CurrentBalance)

		} else {
			fmt.Println("Welcome Back to Home page")
			break
		}
		fmt.Println("Thanks for choosing our bank")

	}
}
