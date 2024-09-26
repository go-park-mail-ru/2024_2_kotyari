package handlers

import (
	"regexp"
	"strings"
	"unicode"
)

// isValidEmail проверяет, является ли email действительным
func isValidEmail(email string) bool {
	const emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func isValidUsername(username string) bool {
	return len(username) > 5 && len(username) < 20
}

func isInGroup(char rune, group string) bool {
	return strings.ContainsRune(group, char)
}

// isValidPassword проверяет, соответствует ли пароль критериям
func isValidPassword(password string) bool {
	var hasMinLen = len(password) >= 8
	var hasNumber, hasUpper, hasLower, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case isInGroup(char, "@$%*?&#."):
			hasSpecial = true
		}
	}
	return hasMinLen && hasNumber && hasUpper && hasLower && hasSpecial
}
