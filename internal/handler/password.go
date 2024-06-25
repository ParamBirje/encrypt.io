package handler

import (
	"fmt"

	"github.com/parambirje/encrypt.io/internal/validation"
	term "golang.org/x/term"
)

func CreatePassword() []byte {

	fmt.Print("\nCreate a password: ")
	password, _ := term.ReadPassword(0)

	fmt.Print("\nConfirm the password: ")
	confirmPassword, _ := term.ReadPassword(0)

	// Matching the passwords
	if !validation.ValidatePassword(password, confirmPassword) {
		fmt.Println("\nPasswords don't match. Please try again.")
		return GetPassword()
	}

	return password
}

func GetPassword() []byte {

	fmt.Print("\nEnter password: ")
	password, _ := term.ReadPassword(0)

	return password
}
