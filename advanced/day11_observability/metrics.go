// Package day11 covers observability: Prometheus metrics + slog middleware (and OTel
// tracing + pprof, described in README). YOUR JOB: implement so
// `go test ./day11_observability/` passes. Run `go mod tidy` first (pulls the
// prometheus client).
package day11

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

// Metrics holds the service's Prometheus collectors. (Struct provided.)
type Metrics struct {
	Requests *prometheus.CounterVec // label: "method"
}

// NewMetrics builds the collectors and registers them with reg.
// HINT:
//   c := prometheus.NewCounterVec(
//       prometheus.CounterOpts{Name: "http_requests_total", Help: "total HTTP requests"},
//       []string{"method"})
//   reg.MustRegister(c)
//   return &Metrics{Requests: c}
// Taking a prometheus.Registerer (rather than the global registry) makes it testable.
func NewMetrics(reg prometheus.Registerer) *Metrics { panic("TODO: implement NewMetrics") }

// Middleware records one request, then calls next. Standard func(http.Handler) http.Handler.
// HINT:
//   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//       m.Requests.WithLabelValues(r.Method).Inc()
//       next.ServeHTTP(w, r)
//   })
func (m *Metrics) Middleware(next http.Handler) http.Handler { panic("TODO: implement Middleware") }
