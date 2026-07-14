# Day 11 — Observability

Implement the Prometheus counter + middleware in `metrics.go` (offline, red→green).

Then wire the rest of the observability stack into a real service (reading, not tested here):

- **Expose metrics:** mount `promhttp.Handler()` at `/metrics` and scrape it.
- **Tracing (OpenTelemetry):** create a `TracerProvider` (stdout exporter for local, OTLP
  in prod), and in a middleware start a span per request:
  `ctx, span := tracer.Start(r.Context(), "http.request"); defer span.End()`. Propagate
  `ctx` down through your handlers and DB/RPC calls (this is why day 8 puts `ctx` first).
- **Structured logs:** `slog` with the request's trace ID as an attribute ties logs to traces.
- **Profiling (`pprof`):** `import _ "net/http/pprof"` exposes `/debug/pprof/*`. Capture a
  CPU/heap/goroutine profile under load; combine with `go test -bench . -benchmem` and
  `go build -gcflags=-m` (escape analysis) to find allocations. Tune with `GOGC` /
  `GOMEMLIMIT` only after profiling says so.
