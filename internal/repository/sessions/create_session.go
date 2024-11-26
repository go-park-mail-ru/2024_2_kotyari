package sessions

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (sr *SessionStore) Create(ctx context.Context, session model.Session) (string, error) {
	err := sr.redis.Set(ctx, session.SessionID, session.UserID, utils.DefaultSessionLifetime).Err()
	if err != nil {
		return "", errs.SessionCreationError
	}

	return session.SessionID, nil
}
