package profile

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	profilegrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/profile/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/golang/protobuf/ptypes/empty"
)

func (p *ProfilesGrpc) ChangePassword(ctx context.Context, in *profilegrpc.ChangePasswordRequest) (*empty.Empty, error) {
	err := p.manager.ChangePassword(ctx,
		in.GetUserId(),
		in.GetOldPassword(),
		in.GetNewPassword(),
		in.GetRepeatPassword(),
	)
	if err != nil {
		switch {
		case errors.Is(err, errs.WrongPassword):
			p.log.Error("[ ProfilesGrpc.ChangePassword ] Неправильный пароль", "err", err.Error())

			return nil, status.Error(codes.InvalidArgument, "Неправильный пароль")
		case errors.Is(err, errs.PasswordsDoNotMatch):
			p.log.Error("[ ProfilesGrpc.ChangePassword ] Пароли не совпадали", "err", err.Error())

			return nil, status.Error(codes.Unauthenticated, "Пароли не совпадали")
		default:
			p.log.Error("[ ProfilesGrpc.ChangePassword ] Неизвестная ошибка", "err", err.Error())

			return nil, status.Error(codes.Internal, "Неизвестная ошибка")
		}

	}

	return &empty.Empty{}, nil
}
