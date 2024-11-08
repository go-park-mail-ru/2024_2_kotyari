package middlewares

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type sessionGetter interface {
	Get(ctx context.Context, sessionID string) (model.Session, error)
}

type csrfValidator interface {
	IsValidCSRFToken(session model.Session, token string) (bool, error)
}
