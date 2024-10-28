package sessions

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type sessionCreator interface {
	Create(ctx context.Context, userID uint32) (string, error)
}

type sessionRemover interface {
	Delete(ctx context.Context, session model.Session) error
}

type SessionHandler struct {
	sessionCreator sessionCreator
	sessionRemover sessionRemover
	errResolver    errs.GetErrorCode
}

func NewSessionDelivery(sessionCreator sessionCreator, sessionRemover sessionRemover, errResolver errs.GetErrorCode) *SessionHandler {
	return &SessionHandler{
		sessionCreator: sessionCreator,
		sessionRemover: sessionRemover,
		errResolver:    errResolver,
	}
}
