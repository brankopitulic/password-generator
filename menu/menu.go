package menu

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

func GenerateNewPassword() {

GenerateNewPassword:
	newPassword := generatePass()

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
		goto GenerateNewPassword
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

func CreateNewPassword() {

	fmt.Println("Your password must contain:\n - Lower Case\n - Upper Case\n - Digits\n - Special Character\n - White space are not allowed\n - Not less than 12 characters")
	fmt.Println("Enter new password: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]

	validatePassword(input)
}

func validatePassword(password string) {

	var ruleLowerCase = false
	var ruleUpperCase = false
	var ruleDigit = false
	var ruleSpecialChar = false
	var ruleSpace = false
	var rulePasswordLength = false

	for _, l := range password {

		if unicode.IsLower(l) {
			ruleLowerCase = true
			continue
		} else if unicode.IsUpper(l) {
			ruleUpperCase = true
			continue
		} else if unicode.IsDigit(l) {
			ruleDigit = true
			continue
		}
	}

	if containsSpecialChars(password) {
		ruleSpecialChar = true
	}

	if containsWhitespace(password) {
		ruleSpace = true
		fmt.Println("White space are not allowed")
	} else if len(password) < 12 {
		rulePasswordLength = true
		fmt.Println("Password length must include at least 12 characters")
	}

	if !ruleSpace && !rulePasswordLength {
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
