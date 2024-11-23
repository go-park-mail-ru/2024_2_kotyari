package grpc

func (m *Metrics) IncreaseTotal(code, method string) {
	m.totalHits.WithLabelValues(m.serviceName, method, code).Inc()
}
