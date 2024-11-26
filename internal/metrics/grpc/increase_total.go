package grpc

import "log"

func (m *Metrics) IncreaseTotal(code, method string) {
	m.totalHits.WithLabelValues(m.serviceName, method, code).Inc()
	log.Println(m.serviceName, code)
}
