package main

// this program takes a customers name
// calculates the price of their order and // prints out the total price
// in a receipt Formatting

import (
	"fmt"
	"time"
)

func calculateOrder(name string, order string, price float64) {
	currentTime := time.Now()
	fmt.Printf("Receipt for %s\n", name)
	fmt.Printf("Order: %s\n", order)
	fmt.Printf("Price: $%.2f\n", price)
	fmt.Printf("Date: %s\n", currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println("Thank you for your order!")
}

func main() {
	// variable declaration
	var customerName string
	var order string
	var price float64

	// user prompts

	fmt.Println("Enter your name: ")
	fmt.Scanln(&customerName)
	fmt.Println("\n")

	fmt.Println("Enter your order: ")
	fmt.Scanln(&order)

	/*
	* I just learnt that in Go after a var is
	* declared with ':=', the call after doesn't
	* need the ':=' operator again, you can simply
	* use '=' to assign a new value to the variable.
	 */

	// assuming a fixed price for simplicity
	price = 12.99

	// printing empty lines
	fmt.Println("\n")
	fmt.Println("Processing your order...")
	fmt.Println("\n")
	fmt.Println("\n")

	fmt.Println("========================================")
	fmt.Println("Thanks for shopping Subway!!!")
	fmt.Println("========================================")

	calculateOrder(customerName, order, price)
	fmt.Println("----------------------------------------")
	fmt.Println("\n")
	fmt.Println("Order processed successfully!")
}
