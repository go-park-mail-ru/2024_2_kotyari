package handlers

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"net/http"
	"regexp"
	"strings"
	"unicode"
)

func validateLogin(w *http.ResponseWriter, creds credsApiRequest) bool {
	return validateEmailAndPassword(w, creds)
}

func validateRegistration(w *http.ResponseWriter, creds credsApiRequest) bool {
	if !validateEmailAndPassword(w, creds) {
		return false
	}

	if creds.Password != creds.RepeatPassword {
		writeJSON(*w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.PasswordsDoNotMatch.Error(),
		})
		return false
	}

	if !isValidUsername(creds.Username) {
		writeJSON(*w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.InvalidUsernameFormat.Error(),
		})
		return false
	}

	return true
}

func validateEmailAndPassword(w *http.ResponseWriter, creds credsApiRequest) bool {
	switch {
	case !isValidEmail(creds.Email):
		writeJSON(*w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.InvalidEmailFormat.Error(),
		})
		return false
	case !isValidPassword(creds.Password):
		writeJSON(*w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.InvalidPasswordFormat.Error(),
		})
		return false
	}

	return true
}

// isValidEmail проверяет, является ли email действительным
func isValidEmail(email string) bool {
	const emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func isValidUsername(username string) bool {
	return len(username) > 2 && len(username) < 20
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
