# Exercises (days 1-6, core)

One Go module, **stdlib-only** (nothing to download). Each day is its own package of
**stub functions you implement**. The `_test.go` files are the spec: they **fail**
until your code is correct. Every stub carries a `// HINT:` comment.

The advanced track (days 7-13: external libs, Postgres, Pub/Sub, gRPC, observability,
testing, config) is a separate module under `../advanced/` because it needs real
dependencies and services. Finish these first.

## The loop

```bash
cd exercises
go test ./day01_basics/          # RED: panics with "TODO: implement ..."
# ... open the .go files, replace each panic("TODO") with real code ...
go test ./day01_basics/          # GREEN when you're done
```

Run the whole core suite (race detector matters from Day 3 on):

```bash
go vet ./...
go test -race -count=1 ./...
```

Day 5 also builds a runnable server once implemented:

```bash
go run ./cmd/api          # :8080, Ctrl-C for graceful shutdown
curl -s -XPOST localhost:8080/items -d '{"id":"1","name":"widget"}'
curl -s localhost:8080/items/1
```

## What you implement

| Dir | Package | You implement | Tests |
|---|---|---|---|
| `day01_basics` | `day1` | `Set` methods, `Map`/`Filter`/`Reduce`, `Intersect`/`Difference` | `basics_test.go`, `set_challenge_test.go` |
| `day02_interfaces_errors` | `day2` | `ValidationError.Error`, `Store.Lookup`, shape methods, `Describe` | `day2_test.go` |
| `day03_concurrency` | `day3` | `WorkerPool`, `gen`/`sq`/`SumOfSquares` | `day3_test.go` |
| `day04_context` | `day4` | `FetchAll` (cancel-on-first-error) | `day4_test.go` |
| `day05_http` | `day5` | `Store.Get`/`Put`, `NewRouter` | `server_test.go` |
| `cmd/api` | `main` | (uses the day05 code you write) | run it manually |
| `day06_streaming` | `day6` | `ProcessNDJSON` | `stream_test.go` |

## Notes

- Struct definitions, constructors, type declarations, and tests are provided. Your job
  is the function bodies marked `panic("TODO")`.
- Some stub files intentionally omit imports you'll need (`fmt`, `math`, `sync`,
  `encoding/json`, `bufio`, `log/slog`). Add them as you go — `goimports`/your editor
  does it automatically, or add by hand.
- `go test` caches passes; use `-count=1` to force a fresh run (important for the
  timing/concurrency tests on days 3, 4, 6).
