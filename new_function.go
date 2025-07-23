package main

// import "fmt"

// type comes after name of variable
func concat(s1 string, s2 string) string {
	return s1 + s2
}

// An idiomatic way to write the above function in Go is to use a single type for both parameters:

// func concat(s1, s2 string) string {
// 	return s1 + s2
// }

// func main() {
// 	fmt.Println(concat("Hello, ", "World!"))
// 	fmt.Println(concat("Nana ", "Amoako"))
// }

