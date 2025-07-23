package main

import "fmt"

func conditionals() {
	// Program to check if a student is a minor
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	if age < 18 {
		fmt.Println("you are underage. please get out of the bar")
	} else {
		fmt.Println("Ah chale! wey drink you dey want?")
	}
}

// func main() {
// 	fmt.Println("----------")
// 	conditionals()
// 	fmt.Println("----------")
// }
