package validation

import "os"

func ValidateFile(path string) bool {
	_, err := os.Stat(path)

	// Returns a bool depending on if the err reports a path that does not exist.
	return !os.IsNotExist(err)
}
