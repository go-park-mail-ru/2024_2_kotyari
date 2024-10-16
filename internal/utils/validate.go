package utils

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func ValidateRegistration(userCredentials model.UserSignupRequestDTO) error {
	if err := ValidateEmailAndPassword(userCredentials.Email, userCredentials.Password); err != nil {
		return err
	}

	if userCredentials.Password != userCredentials.RepeatPassword {
		return errs.PasswordsDoNotMatch
	}

	if !isValidUsername(userCredentials.Username) {
		return errs.InvalidUsernameFormat
	}

	return nil
}

func ValidateEmailAndPassword(email string, password string) error {
	switch {
	case !isValidEmail(email):
		return errs.InvalidEmailFormat
	case !isValidPassword(password):
		return errs.InvalidPasswordFormat
	}

	return nil
}

// isValidEmail проверяет, является ли email действительным
func isValidEmail(email string) bool {
	const emailRegex = `(?i)^[a-z0-9а-яё._%+-]+@[a-z0-9а-яё.-]+\.[a-zа-я]{2,}$`
	re := regexp.MustCompile(emailRegex)

	return re.MatchString(email)
}

func isValidUsername(username string) bool {
	re := regexp.MustCompile(`[a-zA-Zа-яА-ЯёЁ0-9 _-]{2,40}$`)

	return re.MatchString(username)
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
		case isInGroup(char, "!@#$%^:&?*."):
			hasSpecial = true
		}
	}

	return hasMinLen && hasNumber && hasUpper && hasLower && hasSpecial
}
