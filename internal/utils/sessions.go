package utils

import (
	"net/http"
	"time"
)

const (
	SessionName            = "session-id"
	DefaultSessionLifetime = 10 * time.Hour // 10 часов в секундах
	deleteSessionLifetime  = -1             // Удалить сессию
)

func SetSessionCookie(cookieValue string) *http.Cookie {
	return &http.Cookie{
		Name:     SessionName,
		MaxAge:   int(DefaultSessionLifetime.Seconds()),
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Value:    cookieValue,
	}
}

func RemoveSessionCookie() *http.Cookie {
	return &http.Cookie{
		Name:     SessionName,
		MaxAge:   deleteSessionLifetime,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
}
