package cypher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"io"
	"os"

	"github.com/parambirje/encrypt.io/internal/validation"
	"golang.org/x/crypto/pbkdf2"
)

func Encrypt(filePath string, password []byte) {

	// Validating file again.
	if !validation.ValidateFile(filePath) {
		panic("Fatal: File not found.")
	}

	// Opening file for reading
	source, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
	}

	defer source.Close()

	plainText, err := io.ReadAll(source)
	if err != nil {
		panic(err.Error())
	}

	// creating the 12 bytes long randomized nonce
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	// Deriving a secure key for encryption
	// params: key, salt, iterations, key length, hashing function
	derivedKey := pbkdf2.Key(password, nonce, 4096, 32, sha1.New)

	// Starting encryption
	cipherBlock, err := aes.NewCipher(derivedKey)
	if err != nil {
		panic(err.Error())
	}

	// Using the Galois Counter Mode for checking data integrity
	// by attaching a tag to ciphertext
	aesgcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		panic(err.Error())
	}

	// Converting plaintext to ciphertext
	cipherText := aesgcm.Seal(nil, nonce, plainText, nil)
	cipherText = append(cipherText, nonce...)

	// Writing the ciphertext back to the file
	destinationFile, err := os.Create(filePath)
	if err != nil {
		panic(err.Error())
	}

	defer destinationFile.Close()

	_, err = destinationFile.Write(cipherText)
	if err != nil {
		panic(err.Error())
	}

}
