package user

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/api/protos/user/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	sessionsServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/sessions"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

type UsersHandler struct {
	userClientGrpc grpc_gen.UserServiceClient
	inputValidator *utils.InputValidator
	sessionService sessionsServiceLib.SessionService
	errResolver    errs.GetErrorCode
}

func NewUsersHandler(userManager grpc_gen.UserServiceClient, inputValidator *utils.InputValidator, errResolver errs.GetErrorCode) *UsersHandler {
	return &UsersHandler{
		userClientGrpc: userManager,
		inputValidator: inputValidator,
		errResolver:    errResolver,
	}
}
