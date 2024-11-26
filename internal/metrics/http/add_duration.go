package http

import (
	"log"
	"time"
)

func (m *HTTPMetrics) AddDuration(path, code string, duration time.Duration) {
	m.duration.WithLabelValues(path, m.serviceName, code).Observe(duration.Seconds())
	log.Println(path, m.serviceName, code)
}
