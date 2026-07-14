# Makefile for go-in-a-week
# Core exercises (days 1-6) are stdlib-only and offline.
# advanced/ (days 7-13) needs third-party deps (`make tidy`) and, for some days,
# a running service (Postgres / Pub/Sub emulator / Docker).
#
# Override defaults on the command line, e.g.:  make test-pg DATABASE_URL=...

DATABASE_URL         ?= postgres://go:pw@localhost:5432/goweek?sslmode=disable
PUBSUB_EMULATOR_HOST ?= localhost:8085
PUBSUB_PROJECT_ID    ?= go-week

PG_CONTAINER     := goweek-pg
PUBSUB_CONTAINER := goweek-pubsub

.DEFAULT_GOAL := help
.PHONY: help test-core vet-core test-advanced tidy fmt proto \
        pg-up pg-down test-pg pubsub-up pubsub-down test-pubsub test-docker

help: ## Show this help
	@grep -hE '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | awk 'BEGIN{FS=":.*?## "}{printf "  \033[36m%-13s\033[0m %s\n", $$1, $$2}'

# ---- core (days 1-6, offline, no deps) ----

test-core: ## Run the core exercises with the race detector
	cd exercises && go test -race -count=1 ./...

vet-core: ## go vet the core module
	cd exercises && go vet ./...

# ---- advanced (days 7-13) ----

tidy: ## Resolve advanced deps (writes go.mod requires + go.sum; needs network)
	cd advanced && go mod tidy

test-advanced: ## Run the OFFLINE advanced days (7 extlibs, 11 metrics, 12 unit+fuzz, 13 config)
	cd advanced && go test -short -count=1 ./day07_extlibs/ ./day11_observability/ ./day12_testing/ ./day13_config/

fmt: ## gofmt both modules in place
	gofmt -l -w exercises advanced

proto: ## Generate the gRPC echopb package for day 10 (needs protoc + plugins)
	cd advanced && protoc --go_out=. --go_opt=module=github.com/kb/go-in-a-week/advanced --go-grpc_out=. --go-grpc_opt=module=github.com/kb/go-in-a-week/advanced day10_grpc/proto/echo.proto

# ---- Postgres (day 8) ----

pg-up: ## Start a Postgres container and wait until it's ready
	docker run -d --name $(PG_CONTAINER) -e POSTGRES_USER=go -e POSTGRES_PASSWORD=pw -e POSTGRES_DB=goweek -p 5432:5432 postgres:16
	@echo "waiting for postgres..."
	@until docker exec $(PG_CONTAINER) pg_isready -U go -d goweek >/dev/null 2>&1; do sleep 1; done
	@echo "ready: DATABASE_URL=$(DATABASE_URL)"

pg-down: ## Stop and remove the Postgres container
	-docker rm -f $(PG_CONTAINER)

test-pg: ## Run the day-8 Postgres integration test (start it first with pg-up)
	cd advanced && DATABASE_URL="$(DATABASE_URL)" go test -count=1 ./day08_postgres/

# ---- Pub/Sub emulator (day 9) ----

pubsub-up: ## Start the Pub/Sub emulator container
	docker run -d --name $(PUBSUB_CONTAINER) -p 8085:8085 gcr.io/google.com/cloudsdktool/google-cloud-cli:stable gcloud beta emulators pubsub start --host-port=0.0.0.0:8085
	@echo "ready: PUBSUB_EMULATOR_HOST=$(PUBSUB_EMULATOR_HOST)"

pubsub-down: ## Stop and remove the Pub/Sub emulator container
	-docker rm -f $(PUBSUB_CONTAINER)

test-pubsub: ## Run the day-9 Pub/Sub integration test (start it first with pubsub-up)
	cd advanced && PUBSUB_EMULATOR_HOST="$(PUBSUB_EMULATOR_HOST)" PUBSUB_PROJECT_ID="$(PUBSUB_PROJECT_ID)" go test -count=1 ./day09_pubsub/

# ---- day 12 testcontainers (spins its own Postgres via Docker) ----

test-docker: ## Run the day-12 testcontainers test (just needs Docker running)
	cd advanced && go test -count=1 ./day12_testing/
