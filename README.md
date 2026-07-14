# Go in a Week — for a Principal Scala Engineer

A high-intensity, 7-day ramp to write Go at staff level. Written assuming you already
have deep systems intuition (Scala/JVM, C++, Z80), so it spends **zero** time on
"what is a variable" and **all** its time on: how Go thinks, where your Scala reflexes
will actively mislead you, and the idioms that separate "Java-in-Go" from real Go.

Target output: you build backend/API services, concurrent/distributed systems, and
data/streaming pipelines. The plan and exercises are tuned to that.

---

## How to use this

- Each day is ~6–8 focused hours: **read → run the exercise → do the "Extend it" challenges → read the idiom notes.**
- The `exercises/` directory is a single Go module, **stdlib-only** (no external deps to fight with). Each day is its own package of stub functions for you to implement — the tests are the executable spec and stay red (failing) until your code makes them green.
- Run everything with the race detector on from day one: `go test -race ./...`
- Keep `CHEATSHEET.md` open in a split. It's the Scala→Go translation table you'll reach for constantly in week one.

### The one mental reset before you start

Go is not a "better Java" and definitely not a small Scala. Its entire design is a
reaction *against* the things you love in Scala: rich type systems, implicits, deep
abstraction, expression-orientation. Go optimizes for **reading code you didn't write,
at 2am, under an incident.** The language is deliberately boring. Your job as a staff
engineer is not to smuggle Scala patterns in — it's to internalize *why* the constraints
exist and lead a team in producing code that a mid-level engineer can safely modify.

The five reflexes to unlearn immediately:

1. **There are no exceptions.** Errors are ordinary values you return and handle explicitly. `if err != nil` is not boilerplate to abstract away — it *is* the control flow.
2. **There is no inheritance.** Composition (struct embedding) + small interfaces. Interfaces are satisfied *structurally and implicitly* — no `implements`, no `extends`.
3. **There are no implicits, no typeclass machinery, no HKTs.** Generics exist (1.18+) but are intentionally weak. Don't try to rebuild cats.
4. **`nil` is real and has sharp edges** (especially nil interfaces and nil maps). You'll respect it like you respected raw pointers in C++.
5. **Concurrency is CSP, not `Future`/effect systems.** Goroutines + channels + `context`. No monadic composition; you wire it up by hand, and that's the point.

---

## Day 1 — The Go model, tooling, and types

**Goal:** read and write basic Go fluently; internalize the module/package/tooling loop; get structs, methods, and value-vs-pointer semantics correct.

**Read/learn:**
- The tour of Go (go.dev/tour) — skim, you'll clear it in ~90 min. Don't dwell.
- *Effective Go* (go.dev/doc/effective_go) — this is the closest thing to a style constitution. Read it properly; you'll re-read it all week.
- Tooling: `go mod init`, `go build`, `go test`, `go vet`, `gofmt`/`goimports`, `go doc`. There is **one** formatter and **no** debate about style — this is a feature. Wire `gofmt`/`goimports` into your editor now.
- Types: `struct`, methods on value vs pointer receivers, `slice` vs `array` (and why slices are the sharp one), `map`, `string` vs `[]byte` vs `[]rune`, zero values.

**The things that will bite you:**
- **Value semantics by default.** Assigning or passing a struct *copies* it. Coming from JVM reference-everywhere, this is the biggest silent behavior change. Methods with pointer receivers vs value receivers change whether mutation is visible. Pick a receiver kind per type and be consistent.
- **Slices share backing arrays.** `append` may or may not mutate the original depending on capacity. This is a classic aliasing bug. Understand `len` vs `cap` and when `append` reallocates. (Your C++ background helps here — think of it as a `{ptr, len, cap}` view.)
- **Zero values are the idiom.** A `struct` with all fields at their zero value should ideally be usable (`sync.Mutex`, `bytes.Buffer` work this way). Design types so the zero value is meaningful; avoid mandatory constructors where you can.
- **Exported = capitalized.** Visibility is by identifier case, per-package. There's no `private`/`public` keyword.

**Exercise:** `exercises/day01_basics` — a generic `Set[T comparable]` and slice utilities, with tests. Confirms you've got generics-lite, methods, and slice semantics.

---

## Day 2 — Interfaces, composition, errors, generics

**Goal:** master Go's abstraction model (the part most unlike Scala) and idiomatic error handling.

**Read/learn:**
- Interfaces: implicit satisfaction, "accept interfaces, return structs", the empty interface `any`, type switches and type assertions. **Keep interfaces tiny** (`io.Reader`/`io.Writer` are one method). Define interfaces at the *consumer*, not the producer.
- Composition via **struct embedding** (promotion of fields/methods) — this is your inheritance replacement. It's delegation, not subtyping.
- Errors as values: sentinel errors (`errors.New`, `var ErrNotFound = ...`), wrapping with `fmt.Errorf("...: %w", err)`, `errors.Is` / `errors.As`, and when to define a custom error type. Study when to wrap vs annotate vs return bare.
- Generics: type parameters, constraints (`comparable`, `constraints`-style interfaces, union types). Understand the deliberate limits: **no method-level type parameters, no higher-kinded types, no variance.** You can't and shouldn't port your typeclass hierarchy.

**Scala→Go framing:**
- Scala `sealed trait` + pattern match ≈ Go interface + type switch, *but* it's not exhaustive-checked. For closed sets, a common idiom is an interface with an unexported marker method so only your package can implement it.
- `Either[E, A]` / `Try` → just `(A, error)` return tuples. Don't build a `Result` monad; the whole ecosystem expects the two-value return.
- Typeclasses/implicits → pass explicit values/functions, or use small interfaces. Explicit is the culture.

**Exercise:** `exercises/day02_interfaces_errors` — error wrapping with sentinels + custom typed errors (`errors.Is`/`errors.As`), and a `Shape` interface with a type switch. Tests included.

---

## Day 3 — Concurrency I: goroutines, channels, select

**Goal:** think in CSP. Build the canonical concurrency patterns from scratch.

**Read/learn:**
- Goroutines are cheap (~KB stacks, grow on demand) and scheduled by the runtime onto OS threads (M:N). This is *not* one-goroutine-per-thread; it's closer to how you'd think about green threads / fibers.
- Channels: unbuffered (synchronization/handoff) vs buffered (bounded queue). Send/receive semantics, closing channels, ranging over a channel, the comma-ok receive.
- `select` for multiplexing, including `default` for non-blocking and a `done` channel for cancellation.
- The canonical patterns: **generator, fan-out/fan-in, pipeline, worker pool.** Learn them cold — you will assemble these constantly.
- "Share memory by communicating" — the slogan — plus its honest caveat: sometimes a `sync.Mutex` is simpler and correct. Don't be dogmatic.

**Scala→Go framing:**
- No `Future`/`IO`. A goroutine is a running computation; a channel is how you get results out. There's no `.map`/`.flatMap` composition — you wire stages with channels.
- Akka actors ≈ a goroutine owning state + a channel as its mailbox. You can build an actor in ~15 lines; often you shouldn't (a mutex-guarded struct is simpler).
- Backpressure isn't a library feature — it's a *bounded channel*. This matters enormously for your data/streaming work.

**The things that will bite you:**
- **Goroutine leaks.** A goroutine blocked forever on a channel never gets GC'd. Every goroutine needs a guaranteed exit path (usually via `context` or a closed channel).
- **Closing a channel from the receiver, or double-closing** → panic. Rule: the *sender* closes, and only once.
- **Loop variable capture** — historically a top-3 Go bug (`for ... { go func(){ use v }() }`). Go 1.22 changed loop-var scoping to per-iteration, which fixes the classic case; know both the old footgun and the new behavior so you can read old code.

**Exercise:** `exercises/day03_concurrency` — a bounded worker pool (fan-out/fan-in) and a pipeline, with tests that run under `-race`.

---

## Day 4 — Concurrency II: context, cancellation, sync, and correctness

**Goal:** write concurrent code that shuts down cleanly and passes the race detector. This is the day that most separates staff-level Go from tutorial Go.

**Read/learn:**
- `context.Context`: the standard mechanism for cancellation, deadlines/timeouts, and request-scoped values. It threads through *every* blocking/IO call in real backends. Learn `WithCancel`, `WithTimeout`, `WithDeadline`, and the convention that `ctx` is always the first parameter.
- `sync`: `Mutex`/`RWMutex`, `WaitGroup`, `Once`, and `atomic`. When each is the right tool.
- `golang.org/x/sync/errgroup` conceptually (we reimplement a tiny version stdlib-only so the exercise has no deps) — structured concurrency for "run N things, cancel all on first error."
- The **race detector** (`-race`) and the Go **memory model** (what "happens-before" guarantees you actually have). Read the memory model doc; your C++ mental model of atomics/ordering transfers well but Go's guarantees are narrower and more specific.
- `pprof` goroutine profile to find leaks.

**The things that will bite you:**
- **Storing values in `context` beyond request-scoped metadata** — an anti-pattern; don't use it as a DI container.
- **Forgetting `defer cancel()`** leaks the context's timer/goroutine.
- **`WaitGroup.Add` inside the goroutine** (race) — always `Add` before you launch.

**Exercise:** `exercises/day04_context` — a context-aware concurrent fetcher with timeout + cancel-on-first-error (mini errgroup), tested for both success and cancellation paths.

---

## Day 5 — Backend & APIs: net/http, routing, JSON, services

**Goal:** stand up a production-shaped HTTP service: routing, middleware, JSON, config, graceful shutdown.

**Read/learn:**
- `net/http`: `Handler`/`HandlerFunc`, `ServeMux`. As of **Go 1.22** the stdlib mux supports method + path patterns (`GET /items/{id}`), which covers a lot of what you'd previously reach for `chi`/`gorilla` for. Know the stdlib first; reach for `chi` when you need more.
- **Middleware** as `func(http.Handler) http.Handler` — composition by wrapping. This is the idiom for logging, auth, recovery, request IDs.
- JSON with `encoding/json` (struct tags, `omitempty`, streaming with `json.Decoder`/`Encoder`). Know its gotchas: unexported fields ignored, `int` vs `float64` for `any`, unknown-field handling.
- **Graceful shutdown**: `http.Server.Shutdown(ctx)` driven by a signal (`os/signal.NotifyContext`). Non-negotiable for real services.
- `database/sql` mental model + `pgx` for Postgres, and where `sqlc` fits (generate type-safe code from SQL). Connection pooling, `QueryContext`, scanning rows.
- Structured logging with `log/slog` (stdlib, 1.21+).

**Scala→Go framing:**
- No Play/http4s/tapir magic. Routing is explicit, handlers are plain functions, dependency wiring is constructor functions returning structs. It feels low-level; that's intentional and it reads beautifully in review.
- Config: plain structs + env/flags. Resist a heavyweight framework.

**Exercise:** `exercises/day05_http` — a small REST service (in-memory store) using the 1.22 mux, middleware, JSON, and graceful shutdown, tested with `net/http/httptest` (no network, no deps).

---

## Day 6 — Distributed & data/streaming patterns

**Goal:** apply concurrency to real data movement: streaming decode, batching, backpressure, and the observability you need in distributed systems.

**Read/learn:**
- **Streaming, not slurping.** Process `io.Reader` streams incrementally: `bufio.Scanner`, `json.Decoder` in a loop over NDJSON, `io.Pipe`. Bounded channels give you backpressure between producer and consumer.
- **Batching & flushing:** accumulate N items or flush every T with a `time.Ticker` + `select`. The canonical batcher pattern for writing to a DB/queue efficiently.
- Queues/brokers conceptually (Kafka/NATS/SQS): at-least-once delivery, idempotency, consumer groups, offset/ack management. The Go clients (`segmentio/kafka-go`, `franz-go`, NATS) all expose `context`-aware, channel-friendly APIs — the patterns from days 3–4 map directly.
- **Observability**: `slog` structured logs, Prometheus metrics (`promhttp`), distributed tracing (OpenTelemetry) at a conceptual level, and `net/http/pprof` for live profiling. In distributed systems, observability *is* part of the design, not an afterthought.
- Retries, timeouts, and idempotency as first-class design concerns; simple backoff.

**Exercise:** `exercises/day06_streaming` — a streaming NDJSON processor with a bounded-channel pipeline, batching+flush, and graceful drain on cancel. Tested end-to-end with an in-memory reader.

---

## Day 7 — Staff-level concerns & capstone

**Goal:** the things you'll be *evaluated on* as a staff engineer: performance, project structure, testing discipline, and leading on idiom.

**Read/learn:**
- **Performance & memory:** benchmarks (`go test -bench . -benchmem`), `pprof` (CPU/heap/goroutine/block/mutex profiles), **escape analysis** (`go build -gcflags=-m`) — when a value escapes to the heap vs stays on the stack. Your C++ instincts about allocation cost are directly relevant; Go just hides allocation behind the GC, so you profile to see it.
- **GC model:** concurrent, low-latency, non-generational, non-compacting mark-sweep. Tune via `GOGC` and the soft memory limit `GOMEMLIMIT` (1.19+). Usually you reduce allocations rather than tune the GC.
- **Project layout:** package = unit of design (organize by capability/domain, not by "models/controllers"). Avoid circular imports (the compiler forbids them — a useful forcing function). Know the `internal/` convention. Read the "Style Guide" / "Go Code Review Comments" wiki — this is the vocabulary you'll use in reviews.
- **Testing discipline:** table-driven tests (the dominant idiom), `t.Run` subtests, `testing.T` helpers, `httptest`, golden files, fuzzing (`go test -fuzz`), and when integration tests earn their keep. `testify` is common but stdlib-only is very achievable and often preferred.
- **Anti-patterns Scala devs bring:** over-abstraction, premature interfaces (define them when you have 2+ implementers or need to mock at a boundary), generic soup, `panic` as control flow, giant "utils" packages, and getter/setter ceremony.

**Capstone (`exercises/day07_capstone`):** a spec, not a solution — combine days 4–6 into a small service that ingests a stream over HTTP, processes it through a bounded concurrent pipeline with per-request cancellation, batches results to an (in-memory) store, exposes health + metrics endpoints, and shuts down gracefully. Success = it passes `go vet`, `go test -race ./...`, and reads like code a mid-level engineer could safely extend.

---

## After the week (staying honest about the gap)

A week gets you *fluent and dangerous-in-a-good-way*. The remaining staff-level depth comes from reps:
- Read high-quality Go source. The **standard library itself** is the best style teacher — read `net/http`, `io`, `context`, `sync`. Then a real codebase: Kubernetes is huge and idiosyncratic; prefer something like `cockroachdb`, `nats-server`, or `grpc-go` for cleaner exemplars.
- Watch a few talks: Pike's "Concurrency is not Parallelism", "Go Proverbs"; anything by Dave Cheney on practical Go.
- Do 2–3 real code reviews with a strong Go reviewer and *ask why* on every idiom nit. The nits are the culture.

Proverbs worth tattooing on the inside of your eyelids:
> "Clear is better than clever." · "Errors are values." · "Don't communicate by sharing memory; share memory by communicating." · "A little copying is better than a little dependency." · "The bigger the interface, the weaker the abstraction."
