package main

import "fmt"

// apparently you cant have more than one main function per package in Go

/* testing out this comment style
okayyyy pretty cool
*/

func numtypes() {
	// types in Go
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermision bool
	var userName string

	fmt.Printf(
		"SMS Sending Limit: %d\nCost per SMS: %.2f\nHas Permission: %t\nUser Name: %s\n",
		smsSendingLimit,
		costPerSMS,
		hasPermision,
		userName,
	)
}