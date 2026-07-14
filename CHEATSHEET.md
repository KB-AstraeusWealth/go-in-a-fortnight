# Scala → Go Cheat Sheet

Fast translation table for a principal Scala engineer. Left column is your instinct;
right column is what Go actually wants. When they disagree, Go wins — the ecosystem is
built around the Go way.

## Mindset / philosophy

| Scala instinct | Go reality |
|---|---|
| Rich types express intent; abstraction is virtuous | Simplicity is virtuous; abstraction is a cost you justify |
| Expression-oriented (everything returns a value) | Statement-oriented; `if`/`for`/`switch` don't return values |
| DRY at almost any cost | "A little copying is better than a little dependency" — some duplication is fine |
| Powerful type system catches bugs | Simple type system + explicit code + tests + `-race` catches bugs |
| Prefer the elegant one-liner | Prefer the obvious ten lines a junior can modify at 2am |

## Types & data

| Scala | Go |
|---|---|
| `case class Point(x: Int, y: Int)` | `type Point struct { X, Y int }` |
| `val` (immutable) | No language-level immutability. Convention + unexported fields + copies. |
| `Option[A]` | `(A, bool)` (comma-ok), or `*A` where nil means absent, or a zero value |
| `Either[E, A]` / `Try[A]` | `(A, error)` return tuple |
| `sealed trait` + case classes | interface + concrete types + `switch v := x.(type)`; no exhaustiveness check |
| Tuples `(A, B)` | multiple return values (not a first-class tuple type) |
| Enums / `sealed` ADT | `const` + `iota` for simple enums; interface+marker method for closed unions |
| Collections: `List`, `Vector`, `Map`, `Set` | `[]T` (slice), `map[K]V`. **No built-in Set** — use `map[T]struct{}`. |
| Generics everywhere, HKTs, variance | Generics 1.18+: type params + constraints. **No HKT, no variance, no method type params.** |
| Implicits / given / typeclasses | Pass explicit args or small interfaces. There is no implicit resolution. |

## Functions & control flow

| Scala | Go |
|---|---|
| `def f(x: Int): Int = x + 1` | `func f(x int) int { return x + 1 }` |
| Multiple return via tuple | `func f() (int, error)` — genuinely multiple returns |
| `x.map(f).filter(g)` fluent chains | explicit `for` loops; no lazy collection ops in stdlib (helpers in `slices`/`maps` 1.21+) |
| Pattern matching | `switch` (incl. type switch); no destructuring bind on ADTs |
| `for`-comprehension / monadic flow | plain `for`; errors handled with `if err != nil { return ..., err }` |
| Named/default/implicit params | none — use a config struct or functional options |
| Currying / partial application | closures, manually |
| `lazy val` | `sync.Once` or compute-on-first-use by hand |

## Error handling

| Scala | Go |
|---|---|
| `throw` / `try`/`catch` | Errors are values: `return nil, err`. `panic` exists but is for *programmer* errors / unrecoverable state, not control flow. |
| `Try` / `Either` / `MonadError` | `(T, error)`; wrap with `fmt.Errorf("doing X: %w", err)` |
| Match on error subtype | `errors.Is(err, ErrSentinel)` and `errors.As(err, &target)` |
| Custom exception classes | custom types implementing `error` (an `Error() string` method) |
| `finally` | `defer` (runs on function exit, LIFO) |
| Stack traces by default | no stack trace by default; you build context via wrapping (or use a lib) |

## Concurrency (the big one)

| Scala | Go |
|---|---|
| `Future[A]` / `IO[A]` | a goroutine + a channel to carry the result; `go f()` |
| `.map`/`.flatMap`/`for` over `Future`/`IO` | wire stages together with channels; no monadic composition |
| `ExecutionContext` / thread pools | the runtime scheduler (M:N); you rarely manage threads |
| Akka actor + mailbox | goroutine owning state + a channel as mailbox (often overkill — prefer a mutex) |
| `Promise` | an unbuffered channel, or a channel you send one value on |
| Cancellation (`CancellationToken`, fiber cancel) | `context.Context` + `ctx.Done()` |
| Backpressure (Akka Streams, fs2) | a **buffered channel** with bounded capacity |
| `Ref`/`AtomicReference` | `sync/atomic`, or a `sync.Mutex`-guarded field |
| `par`/parallel collections | fan-out goroutines + `sync.WaitGroup` or an errgroup |
| Structured concurrency (fibers scoped) | `errgroup.Group` (`golang.org/x/sync/errgroup`) or a `WaitGroup`+`context` by hand |
| fs2 / Akka Streams pipeline | goroutines connected by channels (generator → stages → sink) |

## Modules, packages, build

| Scala/sbt | Go |
|---|---|
| `build.sbt`, groupId/artifactId | `go.mod`, module path is a URL-like string |
| Package = namespace, files independent | Package = directory; all files in a dir share the package + its scope |
| `private`/`protected`/`public` | Capitalized identifier = exported; lowercase = package-private. That's it. |
| Companion objects | package-level funcs/vars; `NewX()` constructor functions by convention |
| Implicit imports / wildcard | explicit imports; unused import = **compile error** |
| `import x.{A => B}` | `import b "path"` (import alias) |
| Publish to Maven | push a tagged git repo; consumers `go get module@version` |
| Cross-building, many compilers | one toolchain, `GOOS`/`GOARCH` for cross-compile |

## Testing

| Scala | Go |
|---|---|
| ScalaTest / MUnit / specs2 | stdlib `testing`, `func TestXxx(t *testing.T)` in `_test.go` |
| `FunSuite`, matchers, DSLs | table-driven tests + `t.Run` subtests; `testify` optional |
| property testing (ScalaCheck) | `go test -fuzz` (fuzzing), or `testing/quick` |
| mocking frameworks | define a small interface at the consumer, pass a fake struct |
| `sbt test` | `go test ./...` (add `-race`, `-bench`, `-cover`) |

## Quick idiom reminders

- **Accept interfaces, return structs.**
- **Define interfaces where they're used** (consumer side), keep them tiny.
- **`ctx context.Context` is the first parameter** of anything that blocks or does IO.
- **The zero value should be useful** — design types so `var x T` is often ready to go.
- **Don't panic** across package boundaries; return errors. `recover` only at process/goroutine boundaries.
- **`gofmt` is law** — never argue about formatting.
- **Handle every error** or explicitly ignore with `_ =` and a reason. `errcheck`/`go vet` will catch you.
