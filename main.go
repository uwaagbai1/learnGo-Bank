package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("failed to read file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}
	return value, nil
}

func writeFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprintf("%.2f", value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func main() {
	accountBalance, err := getFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR: ")
		fmt.Println(err)
		fmt.Println("-------------------------")
		panic("Can't Continue, sorry")
	}
	fmt.Println("Welcome to Go Bank")
	for {
		presentOptions()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		fmt.Println("Your Choice is", choice)

		switch choice {

		case 1:
			fmt.Println("Your Balance is: ", accountBalance)

		case 2:
			fmt.Print("How much do you want to deposit? ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid Amount, balance must be greater then or equal to 0, try again")
				continue
			}
			accountBalance += depositAmount
			writeFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Your Balance after the deposit is: ", accountBalance)

		case 3:
			fmt.Print("How much do you want to withdraw? ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid Amount, withdrawal must be greater than 0, try again")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("You have insufficient funds, try again")
				continue
			}
			accountBalance -= withdrawAmount
			writeFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Your Balance after withdraw is: ", accountBalance)

		default:
			fmt.Println("Successfully Exited")
			fmt.Println("Thanks for choosing our bank")
			//break
			return
		}

	}
}
