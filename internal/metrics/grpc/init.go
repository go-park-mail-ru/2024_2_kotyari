package grpc

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	serviceName string
	totalHits   *prometheus.CounterVec
	duration    *prometheus.HistogramVec
}

func NewGrpcMetrics(service string) (*Metrics, error) {
	metrics := CreateGrpcMetrics(service)

	if err := InitGrpcMetrics(metrics); err != nil {
		return nil, fmt.Errorf("failed to register gRPC metrics: %w", err)
	}

	return metrics, nil
}
