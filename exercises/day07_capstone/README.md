# Day 7 — Capstone (build this yourself)

No solution provided. Combine days 4–6 into one small service. Aim for code that a
mid-level engineer could safely extend, and that passes `go vet` and `go test -race ./...`.

## Spec

Build an **event ingestion service**:

1. `POST /ingest` accepts a streaming NDJSON body of events (reuse the Day 6 `Event`).
2. Decode the stream incrementally and push events through a **bounded** channel
   (backpressure) into a concurrent worker stage (Day 3 fan-out) that "processes"
   each event (e.g. validates + transforms).
3. **Batch** processed events (Day 6) and write them to an in-memory store,
   flushing on size OR a time interval (`time.Ticker`).
4. Everything is **context-aware** (Day 4): if the client disconnects or the request
   times out, all stages stop promptly and drain (`r.Context()`).
5. `GET /healthz` returns 200; `GET /stats` returns JSON counts (events ingested,
   batches flushed, errors).
6. `main` runs it with **graceful shutdown** on SIGINT/SIGTERM (Day 5).

## Acceptance checklist

- [ ] `go vet ./...` is clean.
- [ ] `go test -race ./...` passes, including a test that cancels mid-stream and
      asserts no goroutine leak (compare `runtime.NumGoroutine()` before/after, or
      use a goroutine-leak check).
- [ ] No goroutine can block forever — every one has a guaranteed exit via context
      or a closed channel.
- [ ] Interfaces are defined at the consumer and are small.
- [ ] The store is race-free (mutex or single-owner goroutine).

## Stretch

- Add Prometheus metrics via `promhttp` (this pulls a dependency — a good moment to
  practice `go get` and `go.sum`).
- Add a `/debug/pprof` endpoint (`import _ "net/http/pprof"`) and capture a goroutine
  profile while a stream is in flight.
- Swap the in-memory store for Postgres via `pgx`, using `context` on every query.
