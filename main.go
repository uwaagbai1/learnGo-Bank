package main

import (
	"bank/fileOps"
	"fmt"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileOps.GetFloatFromFile(accountBalanceFile)
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
			fileOps.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileOps.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Your Balance after withdraw is: ", accountBalance)

		default:
			fmt.Println("Successfully Exited")
			fmt.Println("Thanks for choosing our bank")
			//break
			return
		}

	}
}
