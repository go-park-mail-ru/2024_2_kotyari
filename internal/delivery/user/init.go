package user

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type userManager interface {
	CreateUser(ctx context.Context, user model.User) (string, model.User, error)
	LoginUser(ctx context.Context, user model.User) (string, model.User, error)
	GetUserBySessionID(ctx context.Context, sessionID string) (string, string, error)
}

type UsersDelivery struct {
	userManager userManager
}

func NewUsersHandler(userManager userManager) *UsersDelivery {
	return &UsersDelivery{
		userManager: userManager,
	}
}
