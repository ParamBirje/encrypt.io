package cypher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"

	"github.com/parambirje/encrypt.io/internal/validation"
	"golang.org/x/crypto/pbkdf2"
)

func Decrypt(filePath string, password []byte) {

	// Validating file.
	if !validation.ValidateFile(filePath) {
		panic("Fatal: File not found.")
	}

	// Opening file for reading
	source, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
	}

	defer source.Close()

	cipherText, err := io.ReadAll(source)
	if err != nil {
		panic(err.Error())
	}

	// Getting the attached nonce from the ciphertext
	salt := cipherText[len(cipherText)-12:]
	saltStr := hex.EncodeToString(salt)
	nonce, err := hex.DecodeString(saltStr)
	if err != nil {
		panic(err.Error())
	}

	// Deriving a secure key for initiating decryption
	// params: key, salt, iterations, key length, hashing function
	derivedKey := pbkdf2.Key(password, nonce, 4096, 32, sha1.New)

	// Starting decryption
	cipherBlock, err := aes.NewCipher(derivedKey)
	if err != nil {
		panic(err.Error())
	}

	// Using the Galois Counter Mode for checking data integrity
	// by checking attached tag to the ciphertext
	aesgcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		panic(err.Error())
	}

	// Converting ciphertext to plaintext
	// and ignoring the nonce attached at the end
	plainText, err := aesgcm.Open(nil, nonce, cipherText[:len(cipherText)-12], nil)
	if err != nil {
		panic(err.Error())
	}

	// Writing the ciphertext back to the file
	destinationFile, err := os.Create(filePath)
	if err != nil {
		panic(err.Error())
	}

	defer destinationFile.Close()

	_, err = destinationFile.Write(plainText)
	if err != nil {
		panic(err.Error())
	}

}
