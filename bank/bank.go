package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 0, errors.New("falied to open balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0, errors.New("failed to parse stored balance value")
	}
	return balance, nil
}

func main() {
	var balance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("--------------------------------->")
		fmt.Println("Error: ", err)
		fmt.Println("<---------------------------------")
		panic("Please check the balance store")
	}
	fmt.Println("!!! Welcome to Go Bank !!!")

	var choice int

	for choice != 4 {
		fmt.Println(">>> What do you want to do?")
		fmt.Println("------------------------")
		fmt.Println("> 1: Check balance")
		fmt.Println("> 2: Deposit money")
		fmt.Println("> 3: Withdraw money")
		fmt.Println("> 4. Exit")

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
			writeBalanceToFile(balance)
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
			writeBalanceToFile(balance)
			fmt.Println("Your new balance is: ", balance)
		default:
			fmt.Println("Please select from the specified options")
			continue
		}
	}
	fmt.Println("Thank you and goodbye")
}
