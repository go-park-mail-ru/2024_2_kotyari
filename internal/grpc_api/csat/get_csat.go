package csat

import (
	"context"
	proto "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/csat/gen"
	"log/slog"
)

func (cg *CsatsGrpc) GetCsat(ctx context.Context, in *proto.GetCsatRequest) (*proto.GetCsatResponse, error) {

	question, err := cg.csatRepo.GetCsat(ctx, in.GetType())
	if err != nil {
		cg.log.Error("[ csatsGrpc.GetCsat ] GetCsat error", slog.String("err", err.Error()))
		return nil, err
	}

	return &proto.GetCsatResponse{
		Text: question,
	}, nil
}
