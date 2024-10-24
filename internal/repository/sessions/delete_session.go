package sessions

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (sr *SessionRepo) Delete(ctx context.Context, session model.Session) error {
	return sr.redis.Del(ctx, session.SessionID).Err()
}
