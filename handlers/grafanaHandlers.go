package handlers

import (
	"fmt"
	"grafana-rbac/metrics"
	"net"
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

	// get request ip address
	ip, _, _ := net.SplitHostPort(req.RemoteAddr)

	p.metrics.RequestInfo.With(prometheus.Labels{
		"host":     req.Host,
		"received": req.URL.Path,
		"method":   req.Method,
		"ip":       ip,
	}).Inc()

	// check for cookie
	cookie, err := req.Cookie("grafana_rbac_session")
	if err != nil {
		// create new grafana_rbac_session cookie
		cookie = &http.Cookie{
			Name:  "grafana_rbac_session",
			Value: "annoymous",
		}

		// add cookie to response
		http.SetCookie(w, cookie)
	}

	// get cookie value
	user := cookie.Value

	// get context from request
	// ctx := req.Context()
	// ctx = context.WithValue(ctx, "user", user)

	fmt.Printf("User: %s\n", user)

	req.Header.Set("X-WEBAUTH-USER", user)

	// update req context with new context
	// req = req.WithContext(ctx)

	p.proxy.ServeHTTP(w, req)
}
