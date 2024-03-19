package main

import (
	"os"
	"fmt"
	"errors"
	"strconv"
)

const balanceFile string = "balance.txt"

func storeBalanceToFile(balance float64) {
	os.WriteFile(balanceFile, []byte(fmt.Sprint(balance)), 0644)
}

func getBalanceFromFile() (float64, error) {
	balance, err := os.ReadFile(balanceFile)
	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	result, err := strconv.ParseFloat(string(balance), 64)
	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	return result, nil
}

func main() {
	balance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Whitdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is", balance)
		} else if choice == 2 {
			fmt.Print("Deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			balance += depositAmount
			storeBalanceToFile(balance)
			fmt.Println("Balance updated! New amount:", balance)
		} else if choice == 3 {
			fmt.Print("Withdrawal amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			} else if withdrawAmount > balance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}

			balance -= withdrawAmount
			storeBalanceToFile(balance)
			fmt.Println("Balance updated! New amount:", balance)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}
}
