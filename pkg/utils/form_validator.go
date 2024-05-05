package utils

import "regexp"

func IsValidPhoneNumber(phoneNumber string) bool {
	// Remove any non-digit characters
	reg := regexp.MustCompile("[^0-9]+")
	digitsOnly := reg.ReplaceAllString(phoneNumber, "")

	// Check that the remaining string has between 7 and 15 digits
	reg = regexp.MustCompile(`^\d{7,15}$`)
	return reg.MatchString(digitsOnly)
}

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
