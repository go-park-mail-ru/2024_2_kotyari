package metrics

import (
	"context"
	"google.golang.org/grpc"
	"strconv"
	"time"
)

func (m *Interceptor) ServerMetricsInterceptor(ctx context.Context,
	req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	h, err := handler(ctx, req)
	tm := time.Since(start)
	if err != nil {
		_, code := m.errResolver.Get(err)
		m.metrics.IncreaseTotal(strconv.Itoa(code), info.FullMethod)

		m.metrics.AddDuration(strconv.Itoa(code), info.FullMethod, tm)
		return h, err
	}

	m.metrics.AddDuration("200", info.FullMethod, tm)
	m.metrics.IncreaseTotal(strconv.Itoa(200), info.FullMethod)

	return h, err
}
