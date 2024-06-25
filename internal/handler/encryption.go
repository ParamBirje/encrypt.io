package handler

import (
	"fmt"
	"os"
)

func EncryptionHandler() {

	// Checking if file path is passed
	if len(os.Args) < 3 {
		fmt.Println("File path not passed. For usage info, run `help`")
	}

}
