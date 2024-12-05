package main

import (
	"fmt"
	"section3/fileOps"
)

const accountBalanceFile = "balance.txt"

func main() {
	//var accountBalance, err =fileops.
	accountBalance, err:=fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("-----------")
		//panic("can't execute ,sorry")
	}

	fmt.Println("Welcome to Go Bank")
	fmt.Println("Account Balance is: ", accountBalance)
	for {
		presentOptions()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		// fmt.Println("Your choice is: ", choice)

		if choice == 1 {
			balance, err :=fileops.GetFloatFromFile(accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance,accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance,accountBalanceFile)
			fmt.Println("Balance updated! New Balance", accountBalance)
		} else {
			fmt.Println("Goodbye!")
			//return
			break
		}
	}
	fmt.Println("Thanks for choosing Go Bank")
}
