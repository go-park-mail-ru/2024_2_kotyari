package user

import (
	"context"
	proto "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/user/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"log/slog"
)

type usersManager interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	LoginUser(ctx context.Context, user model.User) (model.User, error)
}

type userGetter interface {
	GetUserByUserID(ctx context.Context, id uint32) (model.User, error)
}

type UsersGrpc struct {
	proto.UnimplementedUserServiceServer
	usersManager usersManager
	userGetter   userGetter
	log          *slog.Logger
}

func NewUsersGrpc(usersManager usersManager, userGetter userGetter, log *slog.Logger) *UsersGrpc {
	return &UsersGrpc{
		usersManager: usersManager,
		userGetter:   userGetter,
		log:          log,
	}
}
