package sessions

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (sr *SessionStore) Create(ctx context.Context, session model.Session) (string, error) {
	userIDStr := fmt.Sprintf("%d", session.UserID)

	err := sr.redis.Set(ctx, session.SessionID, userIDStr, utils.DefaultSessionLifetime).Err()
	if err != nil {
		return "", errs.SessionCreationError
	}

	return session.SessionID, nil
}
