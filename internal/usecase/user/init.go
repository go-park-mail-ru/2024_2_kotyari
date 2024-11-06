package user

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type usersRepository interface {
	CreateUser(ctx context.Context, userModel model.User) (model.User, error)
	GetUserByEmail(ctx context.Context, userModel model.User) (model.User, error)
	GetUserByUserID(ctx context.Context, id uint32) (model.User, error)
}

type sessionGetter interface {
	Get(ctx context.Context, sessionID string) (model.Session, error)
}

type sessionCreator interface {
	Create(ctx context.Context, userID uint32) (string, error)
}

type UsersService struct {
	userRepo       usersRepository
	sessionGetter  sessionGetter
	sessionCreator sessionCreator
}

func NewUserService(usersRepository usersRepository, getter sessionGetter, creator sessionCreator) *UsersService {
	return &UsersService{
		userRepo:       usersRepository,
		sessionGetter:  getter,
		sessionCreator: creator,
	}
}
