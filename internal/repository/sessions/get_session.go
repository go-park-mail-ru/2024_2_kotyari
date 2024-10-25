package sessions

import (
	"context"
	"errors"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/redis/go-redis/v9"
)

func (sr *SessionRepo) Get(ctx context.Context, sessionID string) (model.Session, error) {
	value := sr.redis.Get(ctx, sessionID)
	if errors.Is(value.Err(), redis.Nil) {
		return model.Session{}, errs.SessionNotFound
	}

	userID, err := value.Uint64()
	if err != nil {
		return model.Session{}, errs.InternalServerError
	}

	return model.Session{
		UserID:    uint32(userID),
		SessionID: sessionID,
	}, nil
}
