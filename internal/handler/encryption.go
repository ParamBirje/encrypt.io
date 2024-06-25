package handler

import (
	"fmt"
	"os"

	"github.com/parambirje/encrypt.io/internal/validation"
)

func EncryptionHandler() {

	// Checking if file path is passed
	if len(os.Args) < 3 {
		fmt.Println("File path not passed. For usage info, run `help` command")
		os.Exit(1)
	}

	// Validating file
	filePath := os.Args[2]
	if !validation.ValidateFile(filePath) {
		panic("Fatal: File not found.")
	}
}
