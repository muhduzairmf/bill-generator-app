package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Helper function for making user input
func getInput(prompt string, inputReader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	userInput, err := inputReader.ReadString('\n')
	
	return strings.TrimSpace(userInput), err
}

// Configure a new bill
func createBill() bill {

	inputReader := bufio.NewReader(os.Stdin)
	// Initialize the reader from standard input from terminal

	// fmt.Print("Create a new bill name : ")
	// name, _ := inputReader.ReadString('\n')
	// When the user press enter, it will reader what user was enter

	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name : ", inputReader)

	theBill := generateBill(name)
	fmt.Println("Created the bill - ", theBill.name)

	return theBill
}

// Prompt several options to user
func promptOptions(theBill *bill)  {
	inputReader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose an option (a - add item, t - add tip s - save bill) : ", inputReader)

	switch option {
		case "a":
			itemName, _ := getInput("Item name : ", inputReader)
			itemPrice_input, _ := getInput("Item price : ", inputReader)
			// itemPrice_input is string data type

			itemPrice, err := strconv.ParseFloat(itemPrice_input, 64)
			// Convert itemPrice_input to float64

			if err != nil {
				// Check if it was an error
				fmt.Println("The price must be a number")
				promptOptions(theBill)
			}

			theBill.addItem(itemName, itemPrice)

			fmt.Println("Item added -", itemName, itemPrice)
			promptOptions(theBill)

		case "t":
			tip_input, _ := getInput("Tip amount : ", inputReader)

			tip, err := strconv.ParseFloat(tip_input, 64)

			if err != nil {
				fmt.Println("The tip must be a number")
				promptOptions(theBill)
			}

			theBill.updateTip(tip)

			fmt.Println("Tip updated -", tip)
			promptOptions(theBill)

		case "s":
			fmt.Print("\n")
			fmt.Println(theBill.format())
			theBill.saveBill()

		default:
			fmt.Println("Invalid input...")
			promptOptions(theBill)
			// Keep repeat this function if user enter wrong input
	}
}

func main() {

	myBill := createBill()

	promptOptions(&myBill)

}