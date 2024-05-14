package handlers

import (
	"fmt"
	"grafana-rbac/metrics"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/prometheus/client_golang/prometheus"
)

type Proxy struct {
	grafanaURL *url.URL
	proxy      *httputil.ReverseProxy
	metrics    *metrics.Metrics
}

func NewProxy(grafanaURL string, m *metrics.Metrics) (Proxy, error) {
	url, err := url.Parse(grafanaURL)
	if err != nil {
		return Proxy{}, err
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	return Proxy{
		grafanaURL: url,
		proxy:      proxy,
		metrics:    m,
	}, nil
}

func (p Proxy) GrafanaProxyHandler(w http.ResponseWriter, req *http.Request) {

	p.metrics.RequestInfo.With(prometheus.Labels{
		"host":     req.Host,
		"received": req.URL.Path,
		"method":   req.Method,
	}).Inc()

	fmt.Printf("Request received: %v\n", req.URL.Path)
	fmt.Printf("Request host: %v\n", req.Host)
	fmt.Printf("Request URL: %v\n", req.URL)
	fmt.Printf("Request method: %v\n", req.Method)
	fmt.Printf("Request headers: %v\n", req.Header)
	fmt.Printf("Request cookies: %v\n", req.Cookies())
	fmt.Printf("Request body: %v\n", req.Body)
	fmt.Printf("Request context: %v\n", req.Context())
	fmt.Printf("Request form: %v\n", req.Form)
	fmt.Println("")

	p.proxy.ServeHTTP(w, req)
}
