package user

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

type usersManager interface {
	CreateUser(ctx context.Context, user model.User) (string, model.User, error)
	LoginUser(ctx context.Context, user model.User) (string, model.User, error)
	GetUserBySessionID(ctx context.Context, sessionID string) (model.User, error)
}

type UsersHandler struct {
	grpcClient     grpc_gen.UserServiceClient
	inputValidator utils.InputValidator
	errResolver    errs.Resolver
}

func NewUsersHandler(userManager usersManager, inputValidator *utils.InputValidator, errResolver errs.GetErrorCode) *UsersHandler {
	return &UsersHandler{
		userManager:    userManager,
		inputValidator: inputValidator,
		errResolver:    errResolver,
	}
}
