package sessions

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (sr *SessionRepo) Create(ctx context.Context, session model.Session) (string, error) {
	/// TODO: Remove magic constant
	err := sr.redis.Set(ctx, session.SessionID, session.UserID, 3600*time.Second).Err()
	if err != nil {
		return "", errs.SessionCreationError
	}

	return session.SessionID, nil
}
