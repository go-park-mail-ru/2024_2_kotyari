package sessions

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type sessionManager interface {
	Create(ctx context.Context, userID uint32) (string, error)
	Delete(ctx context.Context, session model.Session) error
}

type SessionHandler struct {
	sessionManager sessionManager
	errResolver    errs.GetErrorCode
}

func NewSessionDelivery(sessionCreator sessionManager, errResolver errs.GetErrorCode) *SessionHandler {
	return &SessionHandler{
		sessionManager: sessionCreator,
		errResolver:    errResolver,
	}
}
