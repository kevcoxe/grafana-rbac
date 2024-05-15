package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	RequestInfo *prometheus.CounterVec
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		RequestInfo: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "grafana_requests_info",
				Help: "Number of HTTP requests.",
			},
			[]string{"host", "received", "method", "ip"},
		),
	}
	reg.MustRegister(m.RequestInfo)
	return m
}
