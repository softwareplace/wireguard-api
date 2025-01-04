package user

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"
)

func checkPassword(inputPassword, storedPassword, salt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(inputPassword+salt))
	return err == nil
}

func validateUserFields(user models.User) error {
	if user.Username == "" || user.Password == "" || user.Email == "" || user.Role == "" {
		return fmt.Errorf("All fields (username, password, email, role) are required")
	}

	if !isValidEmail(user.Email) {
		return fmt.Errorf("Invalid email format")
	}

	if !isValidPassword(user.Password) {
		return fmt.Errorf("Password must be at least 8 characters, include uppercase, lowercase, number, and special character")
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

func hashPassword(password string) (string, string, error) {
	salt := generateSalt()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}
	return string(hashedPassword), salt, nil
}

func generateSalt() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
