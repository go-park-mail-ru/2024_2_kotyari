package user

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type userManager interface {
	CreateUser(ctx context.Context, user model.User) (string, error)
	GetUserByEmail(ctx context.Context, user model.User) (string, error)
}

type UsersDelivery struct {
	userManager userManager
}

func NewUsersHandler(userManager userManager) *UsersDelivery {
	return &UsersDelivery{
		userManager: userManager,
	}
}
