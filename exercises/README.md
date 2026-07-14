# Exercises

One Go module, **stdlib-only** (nothing to download). Each day is its own package of
**stub functions you implement**. The `_test.go` files are the spec: they **fail**
until your code is correct. There are no solutions in this repo — the tests define
"correct", and `go vet` + the race detector keep you honest.

## The loop

```bash
cd exercises
go test ./day01_basics/          # RED: panics with "TODO: implement ..."
# ... open the .go files, replace each panic("TODO") with real code ...
go test ./day01_basics/          # GREEN when you're done
```

Then move to the next day. Run the whole suite (with the race detector, which matters
from Day 3 on):

```bash
go vet ./...
go test -race ./...
```

Day 5 also builds a runnable server once you've implemented it:

```bash
go run ./cmd/api          # :8080, Ctrl-C for graceful shutdown
curl -s -XPOST localhost:8080/items -d '{"id":"1","name":"widget"}'
curl -s localhost:8080/items/1
```

## What you implement

| Dir | Package | You implement | Tests |
|---|---|---|---|
| `day01_basics` | `day1` | `Set` methods, `Map`/`Filter`/`Reduce`, and the `Intersect`/`Difference` challenge | `basics_test.go`, `set_challenge_test.go` |
| `day02_interfaces_errors` | `day2` | `ValidationError.Error`, `Store.Lookup`, the shape methods, `Describe` | `day2_test.go` |
| `day03_concurrency` | `day3` | `WorkerPool`, the `gen`/`sq`/`SumOfSquares` pipeline | `day3_test.go` |
| `day04_context` | `day4` | `FetchAll` (cancel-on-first-error) | `day4_test.go` |
| `day05_http` | `day5` | `Store.Get`/`Put`, `NewRouter` | `server_test.go` |
| `cmd/api` | `main` | (uses the day05 code you write) | run it manually |
| `day06_streaming` | `day6` | `ProcessNDJSON` | `stream_test.go` |
| `day07_capstone` | — | build the whole thing yourself (spec in its README) | write your own |

## Notes

- The struct definitions, constructors (`NewSet`, `NewStore`), type declarations, and
  test files are provided. Your job is the function bodies marked `panic("TODO")`.
- Some stub files intentionally **omit imports** you'll need (e.g. `fmt`, `math`,
  `sync`, `encoding/json`). Add them as you go — `goimports`/your editor will do it
  automatically, or add them by hand. A stub file compiles as given; it stops
  compiling the moment you reference an unimported package, which is your cue.
- Each function's doc comment tells you the intended behavior and drops a hint.
- Stuck? The test is the exact contract — read it. Or ask me for a hint on a specific one.
