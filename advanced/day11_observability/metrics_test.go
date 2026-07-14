package day11

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestMiddlewareCountsRequests(t *testing.T) {
	reg := prometheus.NewRegistry()
	m := NewMetrics(reg)

	h := m.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	for i := 0; i < 3; i++ {
		h.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/x", nil))
	}

	if got := testutil.ToFloat64(m.Requests.WithLabelValues("GET")); got != 3 {
		t.Fatalf("GET counter = %v, want 3", got)
	}
}
