package sessions

import (
	"context"
	"errors"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/redis/go-redis/v9"
)

func (sr *SessionStore) Delete(ctx context.Context, session model.Session) error {
	err := sr.redis.Del(ctx, session.SessionID).Err()
	if errors.Is(err, redis.Nil) {
		return errs.SessionNotFound
	}

	return nil
}
