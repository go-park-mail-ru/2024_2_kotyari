package user

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/sessions"
)

type userRepository interface {
	CreateUser(ctx context.Context, userModel model.User) (uint32, string, error)
	GetUserByEmail(ctx context.Context, userModel model.User) (model.User, error)
}

type UsersService struct {
	userRepo       userRepository
	sessionService sessions.SessionService
}

func NewUserService(userRepository userRepository, sessionService sessions.SessionService) *UsersService {
	return &UsersService{
		userRepo:       userRepository,
		sessionService: sessionService,
	}
}
