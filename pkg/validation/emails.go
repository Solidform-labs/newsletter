package validation

import (
	"net"
	"regexp"
	"strings"
)

func IsValidEmail(email string) bool {
	// Simple regular expression for basic email syntax
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !emailRegex.MatchString(email) {
		return false
	}

	// Split email into domain
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]

	// Check for MX records
	_, err := net.LookupMX(domain)
	return err == nil
}
