package grpc

import "log"

func (m *Metrics) IncreaseTotal(code, method string) {
	log.Println("increase total", m.serviceName, method, code)

	m.totalHits.WithLabelValues(m.serviceName, method, code).Inc()
}
