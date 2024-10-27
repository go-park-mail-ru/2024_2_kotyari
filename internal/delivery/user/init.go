package user

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type usersManager interface {
	CreateUser(ctx context.Context, user model.User) (string, model.User, error)
	LoginUser(ctx context.Context, user model.User) (string, model.User, error)
	GetUserBySessionID(ctx context.Context, sessionID string) (model.User, error)
}

type UsersHandler struct {
	userManager usersManager
	errResolver errs.GetErrorCode
}

func NewUsersHandler(userManager usersManager, errResolver errs.GetErrorCode) *UsersHandler {
	return &UsersHandler{
		userManager: userManager,
		errResolver: errResolver,
	}
}
