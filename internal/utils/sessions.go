package utils

import (
	"context"
	"net/http"
	"time"
)

const (
	SessionName            = "session-id"
	UserSessionID          = "user-id"
	DefaultSessionLifetime = 10 * time.Hour // 10 часов в секундах
	deleteSessionLifetime  = -1             // Удалить сессию
)

func SetSessionCookie(cookieValue string) *http.Cookie {
	return &http.Cookie{
		Name:     SessionName,
		MaxAge:   int(DefaultSessionLifetime.Seconds()),
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Value:    cookieValue,
	}
}

func RemoveSessionCookie() *http.Cookie {
	return &http.Cookie{
		Name:     SessionName,
		MaxAge:   deleteSessionLifetime,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
}

func SetContextSessionUserID(parentCtx context.Context, userID uint32) context.Context {
	return context.WithValue(parentCtx, UserSessionID, userID)
}

func GetContextSessionUserID(ctx context.Context) (uint32, bool) {
	userId, ok := ctx.Value(UserSessionID).(uint32)

	return userId, ok
}
