package handlers

import (
	"grafana-rbac/metrics"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct {
	PrometheusRegistry *prometheus.Registry
	PrometheusMetrics  *metrics.Metrics
}

func NewMetrics() *Metrics {
	reg := prometheus.NewRegistry()
	m := metrics.NewMetrics(reg)
	return &Metrics{
		PrometheusRegistry: reg,
		PrometheusMetrics:  m,
	}
}

func (m *Metrics) PromHandler(w http.ResponseWriter, req *http.Request) {
	promhttp.HandlerFor(m.PrometheusRegistry, promhttp.HandlerOpts{}).ServeHTTP(w, req)
}
