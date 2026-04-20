package utils

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
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

func ValidateMPIN(pin string) error {
	if len(pin) != 6 {
		return errors.New("MPIN must be exactly 6 digits")
	}

	for _, char := range pin {
		if !unicode.IsDigit(char) {
			return errors.New("MPIN must contain only numbers")
		}
	}

	//repitive check
	isRepetitive := true
	for i := 1; i < len(pin); i++ {
		if pin[i] != pin[0] {
			isRepetitive = false
			break
		}
	}
	if isRepetitive {
		return errors.New("MPIN cannot consist of the same repeating digit")
	}

	// // 4. Sequential Digits Check (e.g., 1234 or 4321)
	// isForward := true
	// isBackward := true
	// for i := 1; i < len(pin); i++ {
	// 	if pin[i] != pin[i-1]+1 {
	// 		isForward = false
	// 	}
	// 	if pin[i] != pin[i-1]-1 {
	// 		isBackward = false
	// 	}
	// }
	// if isForward || isBackward {
	// 	return errors.New("MPIN cannot be a simple numerical sequence")
	// }

	return nil
}

var indianMobileRegex = regexp.MustCompile(`^[6-9]\d{9}$`)

func ValidatePhoneNumber(phone string) (string, error) {
	cleanPhone := strings.TrimSpace(phone)
	cleanPhone = strings.TrimPrefix(cleanPhone, "+91")
	cleanPhone = strings.TrimPrefix(cleanPhone, "91")
	cleanPhone = strings.TrimPrefix(cleanPhone, "0")

	if len(cleanPhone) != 10 {
		return "", errors.New("phone number must be exactly 10 digits")
	}

	if !indianMobileRegex.MatchString(cleanPhone) {
		return "", errors.New("invalid mobile number: must start with 6, 7, 8, or 9")
	}

	return cleanPhone, nil
}