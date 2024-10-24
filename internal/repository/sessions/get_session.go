package sessions

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (sr *SessionRepo) Get(ctx context.Context, sessionID string) (model.Session, error) {
	userID, err := sr.redis.Get(ctx, sessionID).Uint64()
	if err != nil {
		return model.Session{}, err
	}

	return model.Session{
		UserID:    uint32(userID),
		SessionID: sessionID,
	}, nil
}
