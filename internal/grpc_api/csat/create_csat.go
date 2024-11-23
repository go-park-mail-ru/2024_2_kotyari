package csat

import (
	"context"
	proto "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/csat/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

func (cg *CsatsGrpc) CreateCsat(ctx context.Context, in *proto.CreateCsatRequest) (*emptypb.Empty, error) {
	userID, ok := utils.GetContextSessionUserID(ctx)
	if !ok {
		cg.log.Error("[ UsersDelivery.GetUserById ] Пользователь не авторизован")
	}

	err := cg.csatManager.CreateCSAT(ctx, model.CSAT{
		UserID: userID,
		Rating: in.Rating,
		Type:   model.CSATType(in.Type),
		Text:   in.Text,
	})

	if err != nil {
		cg.log.Error("[ UsersGrpc.GetUserById ] Ошибка при отдаче на уровень usecase", slog.String("error", err.Error()))
		return &empty.Empty{}, err
	}
	return &empty.Empty{}, nil
}
