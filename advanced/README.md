# Advanced track (days 7-13)

Separate Go module from `../exercises` because these days use real third-party
dependencies and, for some, external services. Same format: stub functions marked
`panic("TODO")` with `// HINT:` comments, and tests as the spec.

## One-time setup

```bash
cd advanced
go mod tidy          # resolves imports, writes go.mod requires + go.sum (needs network)
```

`go.mod` ships with no `require` lines on purpose — `go mod tidy` scans the source and
adds the right versions, so you always pull current releases rather than whatever I
pinned. Run it before anything else. (Verify major versions in the import paths still
match the current releases, e.g. `pgx/v5`.)

## Per-day prerequisites

| Day | Topic | Needs | How tests run |
|---|---|---|---|
| 07 | External libs (`errgroup`, `google/uuid`) | network (`go mod tidy`) | offline unit tests, red→green |
| 08 | PostgreSQL (`pgx`) | a running Postgres | integration: reads `DATABASE_URL` |
| 09 | GCP Pub/Sub | the Pub/Sub emulator | integration: reads `PUBSUB_EMULATOR_HOST` |
| 10 | gRPC + protobuf | `protoc`/`buf` codegen | in-process (`bufconn`), no network |
| 11 | Observability (Prometheus/OTel/slog) | network only | offline unit tests, red→green |
| 12 | Testing (`testcontainers`, fuzz) | Docker | integration: spins Postgres in-process |
| 13 | Config & options | none | offline unit tests, red→green |

Integration days (08, 09, 12) are **integration-only by design**: their tests require
the service to be up and will fail with a clear message if it isn't.

### Postgres (day 08, and day 12 uses testcontainers instead)

```bash
docker run --rm -e POSTGRES_PASSWORD=pw -e POSTGRES_USER=go -e POSTGRES_DB=goweek \
  -p 5432:5432 postgres:16
export DATABASE_URL='postgres://go:pw@localhost:5432/goweek?sslmode=disable'
go test ./day08_postgres/
```

### Pub/Sub emulator (day 09)

```bash
# via gcloud:  gcloud beta emulators pubsub start --host-port=localhost:8085
# or docker:
docker run --rm -p 8085:8085 gcr.io/google.com/cloudsdktool/google-cloud-cli:stable \
  gcloud beta emulators pubsub start --host-port=0.0.0.0:8085
export PUBSUB_EMULATOR_HOST=localhost:8085
export PUBSUB_PROJECT_ID=go-week
go test ./day09_pubsub/
```

### gRPC codegen (day 10)

Install the plugins and generate before implementing — see `day10_grpc/README.md`.

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# then, from day10_grpc/:
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/echo.proto
```

### Docker (day 12) — testcontainers talks to your local Docker daemon; just have it running.

## Heads-up

I could not compile any of this in the authoring sandbox (no Go toolchain, and the
proxy blocks module downloads), and third-party APIs drift between versions. Treat the
code as careful-but-unverified: run `go mod tidy` then `go vet ./...`, and if a symbol
or signature is off against the version you pull, it'll surface immediately — tell me
and I'll correct it.
