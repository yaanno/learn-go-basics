package main

import (
	f "bank/fileops"
	"fmt"

	"github.com/pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var balance, err = f.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		printInvalidChoice(err)
	}
	fmt.Printf("Hello %s, welcome back to Go Bank !!!\n", randomdata.FullName(1))
	fmt.Println("Call us at", randomdata.PhoneNumber())

	var choice int

	for choice != 4 {
		PresentOptions()
		fmt.Print("~ Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("$$$ Your balance: ", balance)
		case 2:
			fmt.Print("How much do you want to deposit? ")
			var deposit float64
			fmt.Scan(&deposit)
			if deposit <= 0 {
				fmt.Println("Deposit must be greater than 0")
				continue
			}
			balance += deposit
			f.WriteFloatToFile(balance, accountBalanceFile)
			fmt.Println("$$$ Your new balance is: ", balance)
		case 3:
			fmt.Print("How much do you want to withdraw? ")
			var withdraw float64
			fmt.Scan(&withdraw)
			if withdraw <= 0 {
				fmt.Println("Withrawal must be greater than 0")
				continue
			}
			if withdraw > balance {
				fmt.Println("Not enough money in balance")
				continue
			}
			balance -= withdraw
			f.WriteFloatToFile(balance, accountBalanceFile)
			fmt.Println("Your new balance is: ", balance)
		case 4:
			fmt.Println("Thank you and goodbye")
			return
		default:
			fmt.Println("Please select an option below")
			continue
		}
	}
}

func printInvalidChoice(err error) {
	fmt.Println("--------------------------------->")
	fmt.Println("Error: ", err)
	fmt.Println("<---------------------------------")
	panic("Please check the balance store")
}
