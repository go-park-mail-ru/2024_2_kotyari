package utils

import (
	"context"
	"net/http"
	"strconv"
	"time"
)

const (
	SessionName            = "session-id"
	UserSessionID          = "user-id"
	DefaultSessionLifetime = 10 * time.Hour // 10 часов в секундах
	deleteSessionLifetime  = -1             // Удалить сессию
)

var DefaultSessionLifetimeString = strconv.Itoa(int(DefaultSessionLifetime.Seconds()))

func SetSessionCookie(cookieValue string) *http.Cookie {
	return &http.Cookie{
		Name:     SessionName,
		MaxAge:   int(DefaultSessionLifetime.Seconds()),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Value:    cookieValue,
	}
}

func RemoveSessionCookie() *http.Cookie {
	return &http.Cookie{
		Name:     SessionName,
		MaxAge:   deleteSessionLifetime,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
}

func SetContextSessionUserID(parentCtx context.Context, userID uint32) context.Context {
	return context.WithValue(parentCtx, UserSessionID, userID)
}

func GetContextSessionUserID(ctx context.Context) (uint32, bool) {
	userId, ok := ctx.Value(UserSessionID).(uint32)

	return userId, ok
}
