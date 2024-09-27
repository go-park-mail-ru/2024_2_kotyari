package handlers

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"net/http"
	"regexp"
	"strings"
	"unicode"
)

// validateCredentials проверяет учетные данные пользователя
func validateCredentials(w *http.ResponseWriter, creds credsApiRequest, requireUsername bool) bool {
	switch {
	case !isValidEmail(creds.Email):
		writeJSON(*w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: errs.InvalidEmailFormat.Error(),
		})
		return false
	case !isValidPassword(creds.Password):
		writeJSON(*w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: errs.InvalidPasswordFormat.Error(),
		})
		return false
	}

	if requireUsername {
		if !isValidUsername(creds.Username) {
			writeJSON(*w, http.StatusBadRequest, errs.HTTPErrorResponse{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: errs.InvalidUsernameFormat.Error(),
			})
			return false
		}
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
