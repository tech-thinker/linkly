package utils

import (
	"fmt"
	"os"
	"strings"
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
