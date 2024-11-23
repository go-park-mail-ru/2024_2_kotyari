package csat

import grpc_gen "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/csat/gen"

func convertGrpcStatsToHTTP(grpcStats []*grpc_gen.GetStatisticsResponse_StarVotes) []*StarVotes {
	httpStats := make([]*StarVotes, len(grpcStats))
	for i, grpcStat := range grpcStats {
		httpStats[i] = &StarVotes{
			StarNumber: grpcStat.GetStarNumber(),
			VoteCount:  grpcStat.GetVoteCount(),
		}
	}
	return httpStats
}
