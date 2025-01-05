package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var CurrentBalance = 1000.0
	var passcode int
	var pincode = 1234
	var userName string
	fmt.Print("Enter Your FullName:")
	reader := bufio.NewReader(os.Stdin)
	userName, _ = reader.ReadString('\n')
	userName = strings.TrimSpace(userName)
	fmt.Println("Welcome", userName, "to the Go bank!!!")
	for i := 0; i < 200; i++ {
		fmt.Print("Enter your Pincode:")
		fmt.Scan(&passcode)
		if passcode != pincode {
			fmt.Println("Wrong pincode")
			continue
		}
		fmt.Println("What do you want to do?")
		fmt.Println("1.Check Bank Balance")
		fmt.Println("2.Deposite Money")
		fmt.Println("3.Withdraw Money")
		fmt.Println("4.Exit")
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

		} else {
			fmt.Println("Welcome Back to Home page")
			break
		}
		fmt.Println("Thanks for choosing our bank")

	}
}
