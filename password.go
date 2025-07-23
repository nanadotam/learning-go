package main

import "fmt"
func password() {
	var password string
	fmt.Println("enter your password: \n->")
	fmt.Scanln(&password)


	// get password length
	var passwordLength int
	passwordLength = getLength(password)
	
	if passwordLength < 8 {
		fmt.Println("password must be 8 chars long")
	} else {
		fmt.Println("password is valid")
	}
}

// okay so functons can be declared anywhere in the file
// and they can be called even after a function is defined

func getLength(s string) int{
	return len(s)
}
func main(){
	fmt.Println("----------")
	password()
	fmt.Println("----------")
}