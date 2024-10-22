package user

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type userManager interface {
	CreateUser(ctx context.Context, userModel model.User) (string, error)
	GetUserByEmail(ctx context.Context, userModel model.User) (model.User, error)
}

type UsersService struct {
	userManager userManager
}

func NewUserService(userManager userManager) *UsersService {
	return &UsersService{
		userManager: userManager,
	}
}
