package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	content, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		// log.Println("Error while reading the file", err)
		return 1000, errors.New("Failed to find balance file")
	}

	//fileContent := fmt.Sprintf("%s", content)
	fileContent := string(content)

	//fmt.Println(fileContent)

	balance, _ := strconv.ParseFloat(fileContent, 64)
	if err!= nil {
		return 1000, errors.New("Failed to parse the stored balance value.")
	}
	return balance, nil
}

func WriteBalanceToFile(balance float64) {
	balanceStr := fmt.Sprintf("%.2f", balance)
	os.WriteFile(accountBalanceFile, []byte(balanceStr), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("-----------")
		//panic("can't execute ,sorry")
	}

	fmt.Println("Welcome to Go Bank")
	fmt.Println("Account Balance is: ", accountBalance)
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		// fmt.Println("Your choice is: ", choice)

		if choice == 1 {
			balance, _ := getBalanceFromFile()
			fmt.Println("Your account balancce is: ", balance)
		} else if choice == 2 {
			fmt.Print("Your Deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount.Must be grater than 0.")
				//return
				continue
			}
			accountBalance += depositAmount
			WriteBalanceToFile(accountBalance)
			fmt.Println("Balance updated! New Balance", accountBalance)
		} else if choice == 3 {
			fmt.Print("Your Withdrawal amount: ")
			var Withdraw float64
			fmt.Scan(&Withdraw)
			if Withdraw <= 0 {
				fmt.Println("Invalid amount.Must be grater than 0.")
				return
			}
			if Withdraw > accountBalance {
				fmt.Println("Transcation Declined! Not enough funds")
				return
			}
			accountBalance -= Withdraw
			WriteBalanceToFile(accountBalance)
			fmt.Println("Balance updated! New Balance", accountBalance)
		} else {
			fmt.Println("Goodbye!")
			//return
			break
		}

	}
	fmt.Println("Thanks for choosing Go Bank")
}
