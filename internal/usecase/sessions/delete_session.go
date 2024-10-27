package sessions

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (ss *SessionService) Delete(ctx context.Context, session model.Session) error {
	return ss.SessionRepo.Delete(ctx, session)
}
