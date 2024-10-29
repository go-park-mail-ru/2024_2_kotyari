package utils

import (
	"regexp"
	"unicode"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
)

func ValidateRegistration(email string, username string, password string, repeatPassword string) error {
	if err := ValidateEmailAndPassword(email, password); err != nil {
		return err
	}

	if password != repeatPassword {
		return errs.PasswordsDoNotMatch
	}

	if !isValidUsername(username) {
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

// isValidPassword проверяет, соответствует ли пароль критериям
func isValidPassword(password string) bool {
	var hasMinLen = len(password) >= 8
	var hasNumber, hasUpper, hasLower bool
	for _, char := range password {
		switch {
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		}
	}

	return hasMinLen && hasNumber && hasUpper && hasLower
}
