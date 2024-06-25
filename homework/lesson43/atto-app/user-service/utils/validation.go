package utils

import (
	"fmt"
	"regexp"

	"github.com/Go11Group/Javokhir-A/homework/lesson43/atto-app/user-service/repository"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s - %s", e.Field, e.Message)
}

func ValidateUser(user repository.CreateUser) *[]ValidationError {
	var validationErrors []ValidationError

	if len(user.Email) == 0 || !isValidEmail(user.Email) {
		validationErrors = append(validationErrors, ValidationError{Field: "Email", Message: "Invalid email address"})
	}

	if !isValidPassword(user.Password) {
		validationErrors = append(validationErrors, ValidationError{Field: "Password", Message: "Password does not meet complexity requirements"})
	}

	if len(user.Phone) == 0 || !isValidPhone(user.Phone) {
		validationErrors = append(validationErrors, ValidationError{Field: "Phone", Message: "Phone does not meet complexity requirements"})
	}

	if len(validationErrors) > 0 {
		return &validationErrors
	}

	return nil
}

func isValidPassword(pass string) bool {
	if len(pass) < 8 {
		return false
	}

	lower := regexp.MustCompile(`[a-z]`)
	upper := regexp.MustCompile(`[A-Z]`)
	digit := regexp.MustCompile(`\d`)
	special := regexp.MustCompile(`[@$!%*?&]`)

	return lower.MatchString(pass) && upper.MatchString(pass) && digit.MatchString(pass) && special.MatchString(pass)
}

func isValidEmail(email string) bool {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func isValidPhone(phone string) bool {
	var phoneRegex = regexp.MustCompile(`^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\s\./0-9]*$`)
	return phoneRegex.MatchString(phone)
}
