package sessions

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type sessionRemover interface {
	Delete(ctx context.Context, session model.Session) error
}

type SessionDelivery struct {
	sessionRemover sessionRemover
	errResolver    errs.GetErrorCode
}

func NewSessionDelivery(sessionRemover sessionRemover, errResolver errs.GetErrorCode) *SessionDelivery {
	return &SessionDelivery{
		sessionRemover: sessionRemover,
		errResolver:    errResolver,
	}
}
