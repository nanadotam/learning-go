package main

import (
	// "fmt"
	"time"
)

func calculateAge(birthYear int) int {
	currentYear := time.Now().Year()
	return currentYear - birthYear
}

// func main() {
// 	name := "Nana"
// 	birthYear := 2004
// 	age := calculateAge(birthYear)

// 	fmt.Printf("%s is %d years old. He was born in %d.\n", name, age, birthYear)
// }
