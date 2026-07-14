package day6

import (
	"context"
	"strings"
	"testing"
)

func TestProcessNDJSONBatches(t *testing.T) {
	input := `{"id":1,"kind":"a"}
{"id":2,"kind":"b"}
{"id":3,"kind":"c"}
{"id":4,"kind":"d"}
{"id":5,"kind":"e"}`

	var batches [][]Event
	flush := func(b []Event) error {
		// The processor reuses b's backing array, so copy before retaining.
		cp := make([]Event, len(b))
		copy(cp, b)
		batches = append(batches, cp)
		return nil
	}

	n, err := ProcessNDJSON(context.Background(), strings.NewReader(input), 2, flush)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 5 {
		t.Fatalf("processed = %d, want 5", n)
	}
	// batchSize 2 over 5 events -> batches of 2, 2, 1
	if len(batches) != 3 {
		t.Fatalf("num batches = %d, want 3", len(batches))
	}
	if len(batches[0]) != 2 || len(batches[2]) != 1 {
		t.Errorf("batch sizes = %d,%d,%d; want 2,2,1",
			len(batches[0]), len(batches[1]), len(batches[2]))
	}
	if batches[0][0].Kind != "a" {
		t.Errorf("first event kind = %q, want a", batches[0][0].Kind)
	}
}

func TestProcessNDJSONBadLine(t *testing.T) {
	input := `{"id":1,"kind":"a"}
{oops not json}`
	_, err := ProcessNDJSON(context.Background(), strings.NewReader(input), 10, func([]Event) error { return nil })
	if err == nil {
		t.Fatal("expected a decode error")
	}
}
