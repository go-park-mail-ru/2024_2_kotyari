package metrics

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"time"
)

func (m *GrpcMiddleware) ServerMetricsInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	start := time.Now()
	h, err := handler(ctx, req)
	tm := time.Since(start)

	if err != nil {
		if errors.Is(err, ClientError) {
			m.metrics.IncreaseTotal("400", info.FullMethod)
			m.metrics.AddDuration("400", info.FullMethod, tm)
		}
		if errors.Is(err, ServerError) {
			m.metrics.IncreaseTotal("429", info.FullMethod)
			m.metrics.AddDuration("429", info.FullMethod, tm)
		}
	} else {
		m.metrics.AddDuration("200", info.FullMethod, tm)
	}

	return h, err

}
