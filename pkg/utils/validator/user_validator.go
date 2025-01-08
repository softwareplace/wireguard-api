package validator

import (
	"fmt"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

func CheckPassword(inputPassword, storedPassword, salt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(inputPassword+salt))
	return err == nil
}

func ValidateUserFields(user models.User) error {
	if user.Username == "" || user.Password == "" || user.Email == "" || len(user.Roles) == 0 {
		return fmt.Errorf("all fields (username, password, email, roles) are required")
	}

	if !isValidEmail(user.Email) {
		return fmt.Errorf("invalid email format")
	}

	if !isValidPassword(user.Password) {
		return fmt.Errorf("password must be at least 8 characters, include uppercase, lowercase, number, and special character")
	}

	return nil
}

func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(emailRegex).MatchString(email)
}

func isValidPassword(password string) bool {
	passwordRegex := `^[A-Za-z\d@$!%*?&]{8,}$`
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[@$!%*?&]`).MatchString(password)

	return regexp.MustCompile(passwordRegex).MatchString(password) &&
		hasUppercase && hasLowercase && hasDigit && hasSpecial
}
