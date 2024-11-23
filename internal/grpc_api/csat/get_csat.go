package csat

import (
	"context"
	proto "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/csat/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"log/slog"
)

func (cg *CsatsGrpc) GetCsat(ctx context.Context, in *proto.GetCsatRequest) (*proto.GetCsatResponse, error) {

	question, err := cg.csatRepo.GetSurveyQuestion(ctx, model.CSATType(in.GetType()))
	if err != nil {
		cg.log.Error("[ csatsGrpc.GetCsat ] GetCsat error", slog.String("err", err.Error()))
		return nil, err
	}

	return &proto.GetCsatResponse{
		Text: question.QuestionText,
	}, nil
}
