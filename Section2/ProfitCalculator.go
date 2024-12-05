package main

import (
	"errors"
	"fmt"
	"os"
)

func WriteDataToFile(ebt, profit, ratio float64) {
	// Format the values into a single string
	data := fmt.Sprintf("EBT: %f\nProfit: %f\nRatio: %f\n", ebt, profit, ratio)

	// Write the string to the file
	err := os.WriteFile("data.txt", []byte(data), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func ProfitCalculator() {
	// var revenue float64
	// var expenses float64
	// var tax float64

	revenue, err := getUserInput("revenue: ")

	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("expenses: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	tax, err := getUserInput("tax: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, tax)
	WriteDataToFile(ebt, profit, ratio)
	fmt.Println("ebt: ", ebt)
	fmt.Println("profit: ", profit)
	fmt.Printf("%.1f", ratio)
}

func getUserInput(inputText string) (float64, error) {
	var userInput float64
	fmt.Print(inputText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		//    panic("User Input cannot be less than 0")
		return 0, errors.New("value msut be a positive number")
	}
	return userInput, nil
}

func calculateFinancials(revenue, expenses, tax float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - tax/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
