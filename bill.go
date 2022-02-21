package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Generate new bills
func generateBill(name string) bill {

	newbill := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return newbill
}

// Format the bill
func (theBill *bill) format() string {
	// This format() function is only associated to bill struct

	itemList := "Bill breakdown : \n"
	var total float64 = 0

	// list the items
	for itemName, itemPrice := range theBill.items {
		itemList += fmt.Sprintf("%-25v .....$%v \n", itemName + " :", itemPrice)
		total += itemPrice
	}

	itemList += fmt.Sprintf("%-25v .....$%.2f \n", "TIP :", theBill.tip)

	itemList += fmt.Sprintf("\n%-25v .....$%.2f \n", "TOTAL :", total+theBill.tip)

	return itemList
}

// Update tip
func (theBill *bill) updateTip(tip float64) {
	
	theBill.tip = tip

}

// Add item to the bill
func (theBill *bill) addItem(name string, price float64) {
	
	theBill.items[name] = price

}

// Save bill to the file
func (theBill *bill) saveBill() {
	
	data := []byte(theBill.format())

	err := os.WriteFile("bills/" + theBill.name + ".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("The bill successfully saved.")

}
