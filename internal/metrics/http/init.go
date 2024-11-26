package http

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

type HTTPMetrics struct {
	totalHits      *prometheus.CounterVec
	serviceName    string
	duration       *prometheus.HistogramVec
	cpuUsage       prometheus.GaugeFunc
	memoryUsage    prometheus.GaugeFunc
	diskUsageTotal *prometheus.GaugeVec
	diskUsageFree  *prometheus.GaugeVec
}

func NewHTTPMetrics(service string) (*HTTPMetrics, error) {
	metrics := CreateHTTPMetrics(service)

	if err := InitHTTPMetrics(metrics); err != nil {
		return nil, fmt.Errorf("failed to register HTTP metrics: %w", err)
	}

	go updateDiskMetrics(metrics.diskUsageTotal, metrics.diskUsageFree)

	return metrics, nil
}
