package utils

import (
	"strings"
)

func ValidPassword(password string) bool {
	// pasword must be at least 8 characters
	if len(password) < 8 {
		return false
	}

	// password must have at least 1 uppercase letter
	if password == strings.ToLower(password) {
		return false
	}

	// password must have at least 1 lowercase letter
	if password == strings.ToUpper(password) {
		return false
	}

	// password must have at least 1 number
	if !strings.ContainsAny(password, "0123456789") {
		return false
	}

	// password must have at least 1 special character
	if !strings.ContainsAny(password, "!@#$%^&*()_+-={}[]|\\:;\"'<>,.?/") {
		return false
	}

	// password must not have any whitespace
	if strings.ContainsAny(password, " ") {
		return false
	}

	return true
}
