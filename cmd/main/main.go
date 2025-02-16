package main

import (
	"fmt"
	"os"

	"github.com/parambirje/encrypt.io/internal/handler"
)

func main() {

	// If no args are passed, help text is displayed
	if len(os.Args) < 2 {
		help()
		os.Exit(0)
	}

	handlerFunction := os.Args[1]
	switch handlerFunction {
	case "encrypt":
		handler.EncryptionHandler()
	case "decrypt":
		handler.DecryptionHandler()
	case "help":
		help()
	default:
		fmt.Printf("\nCommand: %v not found.\n", handlerFunction)
		fmt.Println("Run `help` command for usage intructions.")
		os.Exit(1)
	}
}

func help() {
	fmt.Println("\n\nencrypt.io")
	fmt.Println("A simple file encryption service.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("\tencryptio encrypt /path/to/your/file")
	fmt.Println("\tencryptio decrypt /path/to/your/file")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println()
	fmt.Println("\t help\t\tDisplays help text")
	fmt.Println("\t encrypt\tEncrypts a file")
	fmt.Println("\t decrypt\tDecrypts a file from a password")
	fmt.Println()
}
