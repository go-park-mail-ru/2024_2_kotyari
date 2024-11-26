package grpc

import (
	"log"
	"time"
)

func (m *Metrics) AddDuration(code, method string, duration time.Duration) {
	m.duration.WithLabelValues(m.serviceName, method, code).Observe(duration.Seconds())
	log.Println(m.serviceName, code)
}
