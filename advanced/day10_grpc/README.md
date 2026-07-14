# Day 10 — gRPC + protobuf

`server.go` and `server_test.go` reference a generated package `echopb` that does not
exist yet — generate it before implementing, or the package won't compile.

## 1. Install the codegen plugins (once)

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# ensure $(go env GOPATH)/bin is on your PATH; you also need `protoc` (brew install protobuf)
```

## 2. Generate echopb (run from the advanced/ module root)

```bash
protoc \
  --go_out=. --go_opt=module=github.com/kb/go-in-a-week/advanced \
  --go-grpc_out=. --go-grpc_opt=module=github.com/kb/go-in-a-week/advanced \
  day10_grpc/proto/echo.proto
```

The `module=` option routes output by the file's `go_package`, so the files land in
`day10_grpc/echopb/`. (Alternatively use [buf](https://buf.build) with a `buf.gen.yaml`.)

## 3. Implement & test

```bash
go mod tidy            # pulls google.golang.org/grpc and protobuf
go test ./day10_grpc/  # bufconn, in-process, no network
```

## What you're learning

The `.proto` is the contract; `protoc` generates the messages, the server interface
(`echopb.EchoServer`), and a typed client (`echopb.EchoClient`). You implement the
interface; gRPC handles serialization, HTTP/2 framing, and streaming. `context` flows
through every RPC (deadlines/cancellation propagate across the wire) — the same context
discipline from days 4 and 8.
