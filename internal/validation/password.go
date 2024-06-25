package validation

import "bytes"

func ValidatePassword(password, password2 []byte) bool {
	return bytes.Equal(password, password2)
}
