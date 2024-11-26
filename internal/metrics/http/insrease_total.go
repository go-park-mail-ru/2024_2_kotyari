package http

import "log"

func (m *HTTPMetrics) IncreaseTotal(path, code string) {
	m.totalHits.WithLabelValues(path, m.serviceName, code).Inc()
	log.Println(path, m.serviceName, code)
}
