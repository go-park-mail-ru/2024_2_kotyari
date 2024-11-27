package user

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/api/protos/user/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

type sessionCreator interface {
	Create(ctx context.Context, userID uint32) (string, error)
}

type UsersDelivery struct {
	userClientGrpc grpc_gen.UserServiceClient
	inputValidator *utils.InputValidator
	sessionService sessionCreator
	errResolver    errs.GetErrorCode
	log            *slog.Logger
}

func NewUsersDelivery(userManager grpc_gen.UserServiceClient, inputValidator *utils.InputValidator, sessionService sessionCreator, errResolver errs.GetErrorCode, log *slog.Logger) *UsersDelivery {
	return &UsersDelivery{
		userClientGrpc: userManager,
		inputValidator: inputValidator,
		sessionService: sessionService,
		errResolver:    errResolver,
		log:            log,
	}
}
