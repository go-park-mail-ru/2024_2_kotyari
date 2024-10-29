package sessions

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type sessionManager interface {
	Delete(ctx context.Context, session model.Session) error
	Get(ctx context.Context, sessionID string) (model.Session, error)
}

type SessionHandler struct {
	sessionManager sessionManager
	errResolver    errs.GetErrorCode
}

func NewSessionDelivery(sessionManager sessionManager, errResolver errs.GetErrorCode) *SessionHandler {
	return &SessionHandler{
		sessionManager: sessionManager,
		errResolver:    errResolver,
	}
}
