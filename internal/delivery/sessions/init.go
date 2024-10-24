package sessions

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type sessionRemover interface {
	Delete(ctx context.Context, session model.Session) error
}

type SessionDelivery struct {
	sessionRemover sessionRemover
}

func NewSessionDelivery(sessionRemover sessionRemover) *SessionDelivery {
	return &SessionDelivery{
		sessionRemover: sessionRemover,
	}
}
