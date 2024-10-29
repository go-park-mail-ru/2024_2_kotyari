package sessions

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (sd *SessionHandler) Get(ctx context.Context, sessionID string) (model.Session, error) {
	return sd.sessionManager.Get(ctx, sessionID)
}
