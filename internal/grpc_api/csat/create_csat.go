package csat

import (
	"context"
	proto "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/csat/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"log/slog"
)

func (cg *CsatsGrpc) CreateCsat(ctx context.Context, in *proto.CreateCsatRequest) (*proto.CreateCsatResponse, error) {
	userID, ok := utils.GetContextSessionUserID(ctx)
	if !ok {
		cg.log.Error("[ UsersDelivery.GetUserById ] Пользователь не авторизован")
	}

	csatModel, err := cg.csatManager.CreateCsat(ctx, model.CSAT{
		UserID: userID,
		Rating: in.Rating,
		Type:   model.CSATType(in.Type),
		Text:   in.Text,
	})

	if err != nil {
		cg.log.Error("[ UsersGrpc.GetUserById ] Ошибка при отдаче на уровень usecase", slog.String("error", err.Error()))
		return &proto.CreateCsatResponse{}, err
	}
	return &proto.CreateCsatResponse{
		Type:   string(csatModel.Type),
		Rating: csatModel.Rating,
		Text:   csatModel.Text,
	}, nil
}
