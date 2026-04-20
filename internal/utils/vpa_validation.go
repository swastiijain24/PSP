package utils

import (
	"errors"
	"regexp"
	"strings"
)

var (
	// Regex: alphanumeric, dots, underscores, hyphens before @; alphanumeric handle after @
	vpaRegex = regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9]+$`)
	
	// Reserved usernames that shouldn't be allowed as VPAs
	reservedUsernames = map[string]bool{
		"admin":    true,
		"support":  true,
		"npci":     true,
		"root":     true,
		"nexus":    true,
		"official": true,
		"system":   true,
	}
)

// ValidateVPA performs structural and logic checks on a VPA string
func ValidateVPA(vpa string) error {
	if len(vpa) < 3 || len(vpa) > 100 {
		return errors.New("vpa length must be between 3 and 100 characters")
	}

	if !vpaRegex.MatchString(vpa) {
		return errors.New("invalid vpa format: must follow 'username@handle' pattern")
	}

	parts := strings.Split(vpa, "@")
	username := strings.ToLower(parts[0])
	if reservedUsernames[username] {
		return errors.New("the username part of this vpa is reserved and cannot be used")
	}

	return nil
}