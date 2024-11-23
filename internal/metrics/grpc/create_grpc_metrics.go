package grpc

import (
	"github.com/prometheus/client_golang/prometheus"
)

func CreateGrpcMetrics(service string) *Metrics {
	return &Metrics{
		serviceName: service,
		totalHits: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: service + "_total_hits_count",
				Help: "Number of total gRPC requests",
			},
			[]string{"service", "method", "code"},
		),
		duration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: service + "_duration_seconds",
				Help: "Request duration in seconds",
			},
			[]string{"service", "method", "code"},
		),
	}
}

func InitGrpcMetrics(metrics *Metrics) error {
	if err := prometheus.Register(metrics.totalHits); err != nil {
		return err
	}
	if err := prometheus.Register(metrics.duration); err != nil {
		return err
	}
	return nil
}
