package grpc

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"
)

func CreateGrpcMetrics(service string) *Metrics {
	metrics := &Metrics{
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
				Name:    service + "_duration_seconds",
				Help:    "Request duration in seconds",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"service", "method", "code"},
		),

		cpuUsage: prometheus.NewGaugeFunc(
			prometheus.GaugeOpts{
				Name: service + "_cpu_usage_percent",
				Help: "Current CPU usage as a percentage",
			},
			getCPUUsage,
		),
		memoryUsage: prometheus.NewGaugeFunc(
			prometheus.GaugeOpts{
				Name: service + "_memory_usage_bytes",
				Help: "Current memory usage in bytes",
			},
			getMemoryUsage,
		),
		diskUsageTotal: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: service + "_disk_usage_total_bytes",
				Help: "Total disk space in bytes",
			},
			[]string{"path"},
		),
		diskUsageFree: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: service + "_disk_usage_free_bytes",
				Help: "Free disk space in bytes",
			},
			[]string{"path"},
		),
	}

	return metrics
}

func InitGrpcMetrics(metrics *Metrics) error {
	if err := prometheus.Register(metrics.totalHits); err != nil {
		log.Printf("Failed to register totalHits: %v", err)
		return err
	}
	if err := prometheus.Register(metrics.duration); err != nil {
		return err
	}
	if err := prometheus.Register(metrics.cpuUsage); err != nil {
		return err
	}
	if err := prometheus.Register(metrics.memoryUsage); err != nil {
		return err
	}
	if err := prometheus.Register(metrics.diskUsageTotal); err != nil {
		return err
	}
	if err := prometheus.Register(metrics.diskUsageFree); err != nil {
		return err
	}

	log.Printf("grpc metrics initialized")
	return nil
}
