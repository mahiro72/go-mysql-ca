package helper

import "fmt"

func CreateErrorMessage(message string, err error) error {
	return fmt.Errorf("%s : %w", message, err)
}
