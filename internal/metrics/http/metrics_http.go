package http

//func CreateHTTPMetrics(service string) (*HTTPMetrics, error) {
//	var metric HTTPMetrics
//	metric.serviceName = service
//
//	metric.totalHits = prometheus.NewCounterVec(
//		prometheus.CounterOpts{
//			Name: service + "_total_hits_count",
//			Help: "Number of total http requests",
//		},
//		[]string{"path", "service", "code"})
//	if err := prometheus.Register(metric.totalHits); err != nil {
//		return nil, err
//	}
//
//	metric.duration = prometheus.NewHistogramVec(
//		prometheus.HistogramOpts{
//			Name:    service + "_request_duration_seconds",
//			Help:    "Request time",
//			Buckets: prometheus.DefBuckets,
//		},
//		[]string{"path", "service", "code"})
//	if err := prometheus.Register(metric.duration); err != nil {
//		return nil, err
//	}
//
//	metric.cpuUsage = prometheus.NewGaugeFunc(
//		prometheus.GaugeOpts{
//			Name: service + "_cpu_usage_percent",
//			Help: "Current CPU usage as a percentage",
//		},
//		getCPUUsage,
//	)
//	if err := prometheus.Register(metric.cpuUsage); err != nil {
//		return nil, err
//	}
//
//	metric.memoryUsage = prometheus.NewGaugeFunc(
//		prometheus.GaugeOpts{
//			Name: service + "_memory_usage_bytes",
//			Help: "Current memory usage in bytes",
//		},
//		getMemoryUsage,
//	)
//	if err := prometheus.Register(metric.memoryUsage); err != nil {
//		return nil, err
//	}
//
//	metric.diskUsageTotal = prometheus.NewGaugeVec(
//		prometheus.GaugeOpts{
//			Name: service + "_disk_usage_total_bytes",
//			Help: "Total disk space in bytes",
//		},
//		[]string{"path"},
//	)
//	if err := prometheus.Register(metric.diskUsageTotal); err != nil {
//		return nil, err
//	}
//
//	metric.diskUsageFree = prometheus.NewGaugeVec(
//		prometheus.GaugeOpts{
//			Name: service + "_disk_usage_free_bytes",
//			Help: "Free disk space in bytes",
//		},
//		[]string{"path"},
//	)
//	if err := prometheus.Register(metric.diskUsageFree); err != nil {
//		return nil, err
//	}
//
//	go updateDiskMetrics(metric.diskUsageTotal, metric.diskUsageFree)
//
//	return &metric, nil
//}
