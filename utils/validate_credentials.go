package utils

import (
	"errors"

	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/types"
	"golang.org/x/crypto/bcrypt"
)

func ValidateCredentials(input types.Credentials, credential *store.Credential) error {
	if input.Username != credential.Username {
		return errors.New("invalid username")
	}

	hashedPassword := []byte(credential.Password)
	passwordBytes := []byte(input.Password)

	err := bcrypt.CompareHashAndPassword(hashedPassword, passwordBytes)
	if err != nil {
		return err
	}

	return nil
}
