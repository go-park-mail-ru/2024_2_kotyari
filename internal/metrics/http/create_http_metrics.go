package http

import "github.com/prometheus/client_golang/prometheus"

func CreateHTTPMetrics(service string) *HTTPMetrics {
	return &HTTPMetrics{
		serviceName: service,
		totalHits: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: service + "_total_hits_count",
				Help: "Number of total HTTP requests",
			},
			[]string{"path", "service", "code"},
		),
		duration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    service + "_request_duration_seconds",
				Help:    "Request time",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"path", "service", "code"},
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
}

func InitHTTPMetrics(metrics *HTTPMetrics) error {
	if err := prometheus.Register(metrics.totalHits); err != nil {
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
	return nil
}
