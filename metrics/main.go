package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	CpuTemp     prometheus.Gauge
	HdFailures  *prometheus.CounterVec
	RequestInfo *prometheus.CounterVec
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		CpuTemp: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "cpu_temperature_celsius",
			Help: "Current temperature of the CPU.",
		}),
		HdFailures: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "hd_errors_total",
				Help: "Number of hard-disk errors.",
			},
			[]string{"device"},
		),
		RequestInfo: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "grafana_requests_info",
				Help: "Number of HTTP requests.",
			},
			[]string{"host", "url", "method"},
		),
	}
	reg.MustRegister(m.CpuTemp)
	reg.MustRegister(m.HdFailures)
	reg.MustRegister(m.RequestInfo)
	return m
}
