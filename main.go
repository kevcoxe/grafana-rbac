package main

import (
	"grafana-rbac/handlers"
	"net/http"
	"os"
)

func main() {

	m := handlers.NewMetrics()

	grafanaURL := os.Getenv("GRAFANA_URL")
	grafanaProxy, err := handlers.NewProxy(grafanaURL, m.PrometheusMetrics)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", grafanaProxy.GrafanaProxyHandler)
	http.HandleFunc("/hello", handlers.TestHandler)
	// Expose metrics and custom registry via an HTTP server
	// using the HandleFor function. "/metrics" is the usual endpoint for that.
	http.HandleFunc("/metrics", m.PromHandler)
	http.ListenAndServe(":8090", nil)
}
