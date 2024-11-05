package csrf

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"time"
)

type sessionGetter interface {
	Get(ctx context.Context, sessionID string) (model.Session, error)
}

type csrfCreator interface {
	CreateCsrfToken(session model.Session, now time.Time) (string, error)
}

type CsrfDelivery struct {
	csrfCreator   csrfCreator
	sessionGetter sessionGetter
}

func NewCsrfDelivery(creator csrfCreator, sessionGetter sessionGetter) *CsrfDelivery {
	return &CsrfDelivery{
		csrfCreator:   creator,
		sessionGetter: sessionGetter,
	}
}
