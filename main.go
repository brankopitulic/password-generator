package main

import (
	"PasswordGenerator/menu"
	"fmt"
)

func main() {

RunAgain:
	fmt.Println("Menu!")
	fmt.Println("1.) Generate new password")
	fmt.Println("2.) Create new password")
	fmt.Println("q ) Quit")
	fmt.Println("Enter unit: ")

	var input string
	fmt.Scanln(&input)

	switch input {
	case "1":
		menu.GenerateNewPassword()
	case "2":
		menu.CreateNewPassword()
	case "q":
		break
	default:
		fmt.Println("Number is not on the list!")
		goto RunAgain
	}

}
