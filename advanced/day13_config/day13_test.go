package day13

import (
	"testing"
	"time"
)

func TestDefaults(t *testing.T) {
	s := New()
	if s.Addr != ":8080" || s.ReadTimeout != 5*time.Second || s.MaxConns != 100 {
		t.Fatalf("bad defaults: %+v", s)
	}
}

func TestOptionsApplied(t *testing.T) {
	s := New(WithAddr(":9090"), WithMaxConns(10), WithReadTimeout(time.Second))
	if s.Addr != ":9090" {
		t.Errorf("Addr = %q", s.Addr)
	}
	if s.MaxConns != 10 {
		t.Errorf("MaxConns = %d", s.MaxConns)
	}
	if s.ReadTimeout != time.Second {
		t.Errorf("ReadTimeout = %v", s.ReadTimeout)
	}
}

func TestLaterOptionWins(t *testing.T) {
	s := New(WithAddr(":1"), WithAddr(":2"))
	if s.Addr != ":2" {
		t.Errorf("Addr = %q, want :2 (options apply in order)", s.Addr)
	}
}
