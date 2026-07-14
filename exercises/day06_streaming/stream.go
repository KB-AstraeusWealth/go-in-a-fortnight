// Package day6 covers Day 6: streaming, bounded channels (backpressure), and batching.
//
// YOUR JOB: implement ProcessNDJSON so `go test -race ./day06_streaming/` passes.
// You'll add the bufio and encoding/json imports yourself.
package day6

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
)

// Event is one record in the NDJSON stream. (Provided.)
type Event struct {
	ID   int    `json:"id"`
	Kind string `json:"kind"`
}

// ProcessNDJSON reads newline-delimited JSON from r, streams events through a BOUNDED
// channel (capacity == batchSize gives backpressure), and calls flush on batches of up
// to batchSize. Stop promptly if ctx is canceled (best-effort flush first). Return the
// number of events processed.
//
// HINT — one PRODUCER goroutine + an inline CONSUMER loop:
//
//	Producer:
//	  - events := make(chan Event, batchSize); errc := make(chan error, 1).
//	  - defer close(events) so EVERY exit path closes it exactly once (close nowhere else).
//	  - bufio.Scanner over r; json.Unmarshal each line into an Event.
//	  - send with: select { case events <- e: case <-ctx.Done(): return }.
//	  - on decode error OR scanner.Err(): errc <- err; return.
//	Consumer (this goroutine):
//	  - batch := make([]Event, 0, batchSize) declared ONCE, outside the loop.
//	  - for { select { case <-ctx.Done(): flush remainder; return count, ctx.Err()
//	                    case e, ok := <-events: ... } }.
//	  - a SINGLE-case select blocks forever; always pair the channel op with <-ctx.Done().
//	  - on !ok (channel closed): flush remainder, non-blocking read errc, return.
//	  - append e; when len(batch)==batchSize -> flush(batch); count += len; batch = batch[:0].
//	  - reusing batch's backing array is why callers must copy (the test relies on it).
func ProcessNDJSON(ctx context.Context, r io.Reader, batchSize int, flush func([]Event) error) (int, error) {

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	events := make(chan Event, batchSize)
	errc := make(chan error, 1)

	go func() {
		defer close(events)
		for scanner.Scan() {
			var e Event
			if err := json.Unmarshal(scanner.Bytes(), &e); err != nil {
				errc <- err
				return
			}
			select {
			case events <- e:
			case <-ctx.Done():
				return
			}
		}
		if err := scanner.Err(); err != nil {
			errc <- err
		}
	}()

	batch := make([]Event, 0, batchSize)
	count := 0
	doFlush := func() error {
		if len(batch) == 0 {
			return nil
		}
		if err := flush(batch); err != nil {
			return err
		}
		count += len(batch)
		batch = batch[:0] // reuse backing array — this is the "callers must copy" contract
		return nil
	}

	for {
		select {
		case <-ctx.Done():
			_ = doFlush()
			return count, ctx.Err()
		case e, ok := <-events:
			if !ok {
				if err := doFlush(); err != nil {
					return count, err
				}
				select { // did the producer report an error?
				case err := <-errc:
					return count, err
				default:
					return count, nil
				}
			}
			batch = append(batch, e)
			if len(batch) == batchSize {
				if err := doFlush(); err != nil {
					return count, err
				}
			}
		}
	}
}
