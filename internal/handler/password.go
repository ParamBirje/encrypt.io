package handler

import (
	"fmt"

	"github.com/parambirje/encrypt.io/internal/validation"
	term "golang.org/x/term"
)

func CreatePassword() []byte {

	fmt.Println("\nCreate a password:")
	password, _ := term.ReadPassword(0)

	fmt.Println("\tConfirm the password:")
	confirmPassword, _ := term.ReadPassword(0)

	// Matching the passwords
	if !validation.ValidatePassword(password, confirmPassword) {
		fmt.Println("\nPasswords don't match. Please try again.")
		return GetPassword()
	}

	return password
}

func GetPassword() []byte {

	fmt.Println("\nEnter password:")
	password, _ := term.ReadPassword(0)

	return password
}
