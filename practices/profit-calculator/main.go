package main

import (
	"os"
	"fmt"
	"errors"
)

func main() {
	revenue, err := readUserInput("Revenue: ")
	if err != nil {
		fmt.Printf("Error while reading input for revenue!\n")
		panic(err)
	}

	expenses, err := readUserInput("Expenses: ")
	if err != nil {
		fmt.Printf("Error while reading input for expenses!\n")
		panic(err)
	}

	taxRate, err := readUserInput("taxRate: ")
	if err != nil {
		fmt.Printf("Error while reading input for tax rate!\n")
		panic(err)
	}


	ebt, profit, ratio := runCalculations(revenue, expenses, taxRate)

	fmt.Printf("| %-10s | %-10s | %-10s |\n", "EBT", "Profit", "Ratio")
	fmt.Printf("|------------|------------|------------|\n")
 	fmt.Printf("| %10.2f | %10.2f | %10.2f |\n", ebt, profit, ratio)

	storeValuesToFile(ebt, profit, ratio)
}

func storeValuesToFile(ebt, profit, ratio float64) {
	data := fmt.Sprintf("%-10s %10.2f", "EBT:", ebt)
	data = fmt.Sprintf("%s\n%-10s %10.2f", data, "Profit:", profit)
	data = fmt.Sprintf("%s\n%-10s %10.2f\n", data, "Ratio:", ratio)
	os.WriteFile("calculations.txt", []byte(data), 0644)
}

func readUserInput(message string) (float64, error) {
	var input float64
	fmt.Print(message)
	fmt.Scan(&input)
	if input <= 0 {
		return 0, errors.New("Input cannot be 0 or negative.")
	}
	return input, nil
}

func runCalculations(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate / 100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
