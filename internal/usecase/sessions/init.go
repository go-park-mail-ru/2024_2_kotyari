package sessions

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type sessionRepository interface {
	Create(ctx context.Context, session model.Session) (string, error)
	Get(ctx context.Context, sessionID string) (model.Session, error)
	Delete(ctx context.Context, session model.Session) error
}

type SessionService struct {
	SessionRepo sessionRepository
	log         *slog.Logger
}

func NewSessionService(sessionRepo sessionRepository, log *slog.Logger) *SessionService {
	return &SessionService{
		SessionRepo: sessionRepo,
		log:         log,
	}
}
