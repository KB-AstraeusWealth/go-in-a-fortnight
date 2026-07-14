// Package day10 covers gRPC + protobuf.
//
// PREREQUISITE: generate the echopb package first (see README.md) — this file won't
// compile until echopb exists. Then implement Say. The test runs the server and client
// in-process over bufconn, so no network or ports are involved.
package day10

import (
	"context"

	"github.com/kb/go-in-a-week/advanced/day10_grpc/echopb"
)

// EchoServer implements the generated echopb.EchoServer interface.
// Embedding UnimplementedEchoServer is the forward-compatible convention: if the .proto
// gains a new method, your type still satisfies the interface (returns Unimplemented).
type EchoServer struct {
	echopb.UnimplementedEchoServer
}

// Say echoes the request message back, prefixed with "echo: ".
// HINT: return &echopb.SayReply{Message: "echo: " + req.GetMessage()}, nil
func (s *EchoServer) Say(ctx context.Context, req *echopb.SayRequest) (*echopb.SayReply, error) {
	panic("TODO: implement Say")
}
