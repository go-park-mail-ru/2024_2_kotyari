package user

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (us *UsersService) GetUserBySessionID(ctx context.Context, sessionID string) (model.User, error) {
	session, err := us.sessionGetter.Get(ctx, sessionID)
	if err != nil {
		return model.User{}, err
	}

	user, err := us.userRepo.GetUserByUserID(ctx, session.UserID)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
