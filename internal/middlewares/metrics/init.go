package metrics

import (
	metrics "github.com/go-park-mail-ru/2024_2_kotyari/internal/metrics/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ClientError = status.Error(codes.InvalidArgument, "invalid ID, fail to cast uuid")
	ServerError = status.Error(codes.Internal, "internal server error")
)

type GrpcMiddleware struct {
	metrics metrics.Metrics
}

func NewGrpcMiddleware(metrics metrics.Metrics) *GrpcMiddleware {
	return &GrpcMiddleware{
		metrics: metrics,
	}
}
