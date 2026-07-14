package day10

import (
	"context"
	"net"
	"testing"

	"github.com/kb/go-in-a-week/advanced/day10_grpc/echopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

// bufconn gives an in-memory net.Conn, so client and server talk over a real gRPC
// stack without opening a socket — fast, deterministic, no port conflicts.
func TestSay(t *testing.T) {
	lis := bufconn.Listen(1024 * 1024)
	srv := grpc.NewServer()
	echopb.RegisterEchoServer(srv, &EchoServer{})
	go func() { _ = srv.Serve(lis) }()
	defer srv.Stop()

	dialer := func(context.Context, string) (net.Conn, error) { return lis.Dial() }
	conn, err := grpc.NewClient(
		"passthrough:///bufnet",
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer conn.Close()

	client := echopb.NewEchoClient(conn)
	resp, err := client.Say(context.Background(), &echopb.SayRequest{Message: "hi"})
	if err != nil {
		t.Fatalf("Say: %v", err)
	}
	if resp.GetMessage() != "echo: hi" {
		t.Errorf("Message = %q, want %q", resp.GetMessage(), "echo: hi")
	}
}
