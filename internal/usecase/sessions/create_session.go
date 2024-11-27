package sessions

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/google/uuid"
)

func (ss *SessionService) Create(ctx context.Context, userID uint32) (string, error) {
	id := uuid.New()

	session := model.Session{
		SessionID: id.String(),
		UserID:    userID,
	}

	return ss.SessionRepo.Create(ctx, session)
}
