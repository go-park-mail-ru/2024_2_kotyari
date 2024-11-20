package user

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/sessions"
)

type usersRepository interface {
	CreateUser(ctx context.Context, userModel model.User) (model.User, error)
	GetUserByEmail(ctx context.Context, userModel model.User) (model.User, error)
	GetUserByUserID(ctx context.Context, id uint32) (model.User, error)
}

type UsersService struct {
	userRepo       usersRepository
	inputValidator *utils.InputValidator
	sessionService *sessions.SessionService
}

func NewUserService(usersRepository usersRepository, inputValidator *utils.InputValidator, sessionService *sessions.SessionService) *UsersService {
	return &UsersService{
		userRepo:       usersRepository,
		inputValidator: inputValidator,
		sessionService: sessionService,
	}
}
