package main

import "fmt"

// Constants in Go

func constants() {
	const name = "Saul Goodman"
	const age = 45
	const isLawyer = true
	const salary = 100000.50

	msg := fmt.Sprintf(
		"Hi, my name is: %s\nI am %d years old.\nIs Lawyer: %t\nSalary: %.2f\n",
		name,
		age,
		isLawyer,
		salary,
	)

	fmt.Println(msg)
}

func main() {
	constants()
}

/*
Formatting verbs:
- %s for string
- %d for integer
- %t for boolean
- %.2f for float with 2 decimal places
- %v for any value
- %T for type of value
// - %p for pointer
- %x for hexadecimal
- %b for binary
- %o for octal
// - %e for scientific notation
// - %c for character
- %q for quoted string
// - %f for float
- %g for float with no trailing zeros
*/
