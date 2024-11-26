package metrics

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	metrics "github.com/go-park-mail-ru/2024_2_kotyari/internal/metrics/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ClientError = status.Error(codes.InvalidArgument, "")
	ServerError = status.Error(codes.Internal, "internal server error")
)

type Interceptor struct {
	metrics     metrics.Metrics
	errResolver errs.GetErrorCode
}

func NewGrpcMiddleware(metrics metrics.Metrics, errResolver errs.GetErrorCode) *Interceptor {
	return &Interceptor{
		metrics:     metrics,
		errResolver: errResolver,
	}
}
