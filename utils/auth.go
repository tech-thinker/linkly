package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/tech-thinker/linkly/models"
	"golang.org/x/crypto/bcrypt"
)

// GetEnv gets the environment variable
func GetEnv(key string) string {
	return os.Getenv(key)
}

// ParseToken parses the token from authorization header
func ParseToken(authorization string) (string, error) {
	if strings.HasPrefix(authorization, "Bearer ") {
		return strings.TrimPrefix(authorization, "Bearer "), nil
	}
	return "", fmt.Errorf("invalid authorization header")
}

// ToMaxAge converts the duration to max age
func ToMaxAge(expire time.Time) int {
	maxAge := time.Until(expire).Seconds()
	return int(maxAge)
}

// HashAndSalt generates a hashed password
func HashAndSalt(password string) (string, error) {
	// Generate a hashed password with bcypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyHash verifies the hashed password
func VerifyHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err == nil
	}
	return true
}

// ValidateUser the user informations
func ValidateUser(u *models.User) error {
	// if u.FirstName == "" || u.LastName == "" {
	// 	return errors.New("first name and last name are required")
	// }
	if u.Username == "" {
		return errors.New("username is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}

	// check if password contains at least 7 characters,
	// at least one letter, one number and one special character
	if !IsValidPassword(u.Password) {
		return errors.New("password should contain at least seven characters, one number and one special character")
	}
	return nil
}

// IsRestrictedUser checks if the user is restricted
func IsRestrictedUser(username string) bool {
	restrictedList := []string{"admin", "root", "me", "system", "search"}
	for _, restrictedUser := range restrictedList {
		if username == restrictedUser {
			return true
		}
	}
	return false
}

// IsValidUserName checks if the username is valid
func IsValidUserName(username string) bool {
	if username == "" {
		return false
	}
	// username can be only alphanumeric
	for _, char := range username {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			return false
		}
	}

	if IsRestrictedUser(username) {
		return false
	}

	// return true
	return len(username) >= 3
}

// IsValidPassword checks if the password (Strength) is valid
func IsValidPassword(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 8 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}
