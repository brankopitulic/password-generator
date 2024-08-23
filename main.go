package main

import (
	"bufio"
	"fmt"
	"github.com/atotto/clipboard"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func main() {

	fmt.Println("Menu!")
	fmt.Println("1.) Generate new password")
	fmt.Println("2.) Password validation")
	fmt.Println("Enter unit: ")

	var input string
	fmt.Scanln(&input)

	switch input {
	case "1":
		generateNewPassword()
	case "2":
		validateNewPassword()
	default:
		fmt.Println("Number is not on the list!")
	}

}

func generateNewPassword() {

	newPassword := generatePass()

NewPassword:
	fmt.Println(newPassword)
RepeatDecision:
	fmt.Println("Do you like Y/N: ")

	var input string
	strings.ToLower(input)
	fmt.Scanln(&input)

	switch input {
	case "y":
		err := clipboard.WriteAll(newPassword)
		if err != nil {
			fmt.Println("Failed to copy to clipboard:", err)
		} else {
			fmt.Println("Copied to clipboard successfully!")
		}
	case "n":
		goto NewPassword
	default:
		fmt.Println("You can use only Y(yes) and N(no)!")
		goto RepeatDecision
	}
}

func generatePass() string {

	var loweCase = []rune("abcdefghijklmnopqrstuvwxyz")
	var upperCase = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var digits = []rune("0123456789")
	var specialCharacter = []rune(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]+`)
	//var character = []rune("!@#$%^&*()_+-=")

	b := make([]rune, 12)
	for i := range b {
		if i < 3 {
			b[i] = loweCase[rand.Intn(len(loweCase))]
		} else if i < 7 {
			b[i] = upperCase[rand.Intn(len(upperCase))]
		} else if i < 11 {
			b[i] = digits[rand.Intn(len(digits))]
		} else {
			b[i] = specialCharacter[rand.Intn(len(specialCharacter))]
		}
	}
	return string(b)
}

func validateNewPassword() {
	//print rules
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Your password must contain:\nLowe Case\nUpper Case\nDigits\nSpecial Character\nSpace not allowed ")
	fmt.Println("Enter new password: ")
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]
	//var input string
	//fmt.Scanln(&input)2

	var ruleLowerCase = false
	var ruleUpperCase = false
	var ruleDigit = false
	var ruleSpecialChar = false
	var ruleSpace = false

	for _, r := range input {

		if unicode.IsLower(r) {
			ruleLowerCase = true
			continue
		} else if unicode.IsUpper(r) {
			ruleUpperCase = true
			continue
		} else if unicode.IsDigit(r) {
			ruleDigit = true
			continue
		} else if containsSpecialChars(input) {
			ruleSpecialChar = true
			continue
		} else if containsWhitespace(input) {
			ruleSpace = true
			fmt.Println("You contain space is not allowed")
			break
		}
	}

	if !ruleSpace {
		if !ruleLowerCase {
			fmt.Println("You miss lower case letter")
		}
		if !ruleUpperCase {
			fmt.Println("You miss upper case letter")
		}
		if !ruleDigit {
			fmt.Println("You miss digit")
		}
		if !ruleSpecialChar {
			fmt.Println("You miss special character")
		}
	}

	//input your password
	//check validation rules
	//if right print pass is valid, if not print pass is not valid
	//if valid ask do you want new pass validation if yes, start again, if not copy the current one
}

func containsSpecialChars(str string) bool {
	// Define a regex pattern that matches any of the special characters
	specialCharPattern := `[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]+`

	// Compile the regex pattern
	re := regexp.MustCompile(specialCharPattern)

	// Check if the string contains any special characters
	return re.MatchString(str)
}

func containsWhitespace(str string) bool {
	// Regular expression to check for single or multiple whitespaces
	re := regexp.MustCompile(`\s+`)
	return re.MatchString(str)
}
