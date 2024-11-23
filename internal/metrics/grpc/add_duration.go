package grpc

import "time"

func (m *Metrics) AddDuration(code, method string, duration time.Duration) {
	m.duration.WithLabelValues(m.serviceName, method, code).Observe(duration.Seconds())
}
