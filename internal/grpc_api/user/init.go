package user

import (
	"context"
	proto "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/user/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type usersManager interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	LoginUser(ctx context.Context, user model.User) (model.User, error)
	GetUserBySessionID(ctx context.Context, sessionID string) (model.User, error)
}

type userGetter interface {
	GetUserByUserID(ctx context.Context, id uint32) (model.User, error)
}

type GrpcUserManager struct {
	proto.UnimplementedUserServiceServer
	usersManager usersManager
	userGetter   userGetter
}

func NewUsersHandler(usersManager usersManager) *GrpcUserManager {
	return &GrpcUserManager{
		usersManager: usersManager,
	}
}
