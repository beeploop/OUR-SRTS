package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func ValidateCredentials(input, credential string) error {
	hashedPassword := []byte(credential)
	passwordBytes := []byte(input)

	err := bcrypt.CompareHashAndPassword(hashedPassword, passwordBytes)
	if err != nil {
		return err
	}

	return nil
}
