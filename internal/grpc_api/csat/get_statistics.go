package csat

import (
	"context"
	proto "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/csat/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"log/slog"
)

func (cg *CsatsGrpc) GetStatistics(ctx context.Context, in *proto.GetStatisticsRequest) (*proto.GetStatisticsResponse, error) {

	statistics, err := cg.csatRepo.GetStatistics(ctx, model.CSATType(in.GetType()))
	if err != nil {
		cg.log.Error("[ csatsGrpc.GetStatistics ] GetStatistics error", slog.String("err", err.Error()))
		return nil, err
	}

	var grpcStats []*proto.GetStatisticsResponse_StarVotes
	for _, stat := range statistics.Ratings {
		grpcStats = append(grpcStats, &proto.GetStatisticsResponse_StarVotes{
			StarNumber: stat.RatingName,
			VoteCount:  stat.RatingValue,
		})
	}

	return &proto.GetStatisticsResponse{
		Stats:   grpcStats,
		Average: statistics.Avg,
	}, nil
}
