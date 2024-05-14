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
	// Create new metrics and register them using the custom registry.
	// Set values for the new created metrics.
	m.PrometheusMetrics.CpuTemp.Set(65.3)
	m.PrometheusMetrics.HdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
	// Serve the metrics from the custom registry.
	promhttp.HandlerFor(m.PrometheusRegistry, promhttp.HandlerOpts{}).ServeHTTP(w, req)
}
