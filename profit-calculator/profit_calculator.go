package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const profitFile = "profit.txt"

func getProfitValueFromFile() (float64, error) {
	data, err := os.ReadFile(profitFile)
	if err != nil {
		formattedError := fmt.Sprintf("FILE ERROR: %s", err)
		return 0, errors.New(formattedError)
	}
	textValue := string(data)
	floatValue, err := strconv.ParseFloat(textValue, 64)
	if err != nil {
		formattedError := fmt.Sprintf("FILE ERROR: %s", err)
		return 0, errors.New(formattedError)
	}
	return floatValue, nil
}

func writeProfitValueToFile(values float64) error {
	valuesText := fmt.Sprint(values)
	err := os.WriteFile(profitFile, []byte(valuesText), 0644)
	if err != nil {
		formattedError := fmt.Sprintf("FILE ERROR: %s", err)
		return errors.New(formattedError)
	}
	return nil
}

func main() {
	profit, err := getProfitValueFromFile()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Current balance: %.2f\n", profit)
	revenue, err := getUserInput("Revenue")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Expense")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Tax rate")
	if err != nil {
		fmt.Println(err)
		return
	}
	earningsBeforeTax, earningsAfterTax, earningsRatio := calculateFinancials(revenue, expenses, taxRate)
	profit += earningsAfterTax

	fmt.Printf("Your earnings before tax: %.2f \n", earningsBeforeTax)
	fmt.Printf("Your earnings after tax: %.2f\n", earningsAfterTax)
	fmt.Printf("Ratio: %.2f\n", earningsRatio)
	fmt.Printf("Current balance: %.2f\n", profit)

	err = writeProfitValueToFile(profit)
	if err != nil {
		fmt.Println(err)
	}
}

func calculateFinancials(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (1 - taxRate/100)
	earningsRatio := earningsBeforeTax / earningsAfterTax
	return earningsBeforeTax, earningsAfterTax, earningsRatio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText, ": ")
	_, err := fmt.Scan(&userInput)
	if err != nil {
		formattedError := fmt.Sprintf("INPUT ERROR: %s", err)
		return 0, errors.New(formattedError)
	}
	if userInput <= 0 {
		formattedError := fmt.Sprintf("%s must be greater than 0", infoText)
		return 0, errors.New(formattedError)
	}
	return userInput, nil
}
