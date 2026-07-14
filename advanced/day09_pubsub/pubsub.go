// Package day09 covers GCP Pub/Sub. INTEGRATION-ONLY: run the emulator and set
// PUBSUB_EMULATOR_HOST (+ PUBSUB_PROJECT_ID). See ../README.md. Run `go mod tidy` first.
//
// YOUR JOB: implement the stubs so `go test ./day09_pubsub/` passes against the emulator.
package day09

import (
	"context"

	"cloud.google.com/go/pubsub"
)

// Publish publishes data to topicID and returns the server-assigned message ID.
// HINT:
//   res := client.Topic(topicID).Publish(ctx, &pubsub.Message{Data: data})
//   id, err := res.Get(ctx)   // Get blocks until the publish is acknowledged
//   return id, err
func Publish(ctx context.Context, client *pubsub.Client, topicID string, data []byte) (string, error) {
	panic("TODO: implement Publish")
}

// ReceiveN receives exactly n messages from subID, passing each payload to handle and
// Acking it, then returns nil. Receive runs its callback concurrently, so guard shared
// state and cancel once n have arrived.
// HINT:
//   cctx, cancel := context.WithCancel(ctx); defer cancel()
//   var mu sync.Mutex; count := 0
//   err := client.Subscription(subID).Receive(cctx, func(_ context.Context, m *pubsub.Message) {
//       handle(m.Data)
//       m.Ack()
//       mu.Lock(); count++; if count >= n { cancel() }; mu.Unlock()
//   })
//   if err != nil { return err }   // Receive returns nil when the ctx is canceled
//   return nil
func ReceiveN(ctx context.Context, client *pubsub.Client, subID string, n int, handle func([]byte)) error {
	panic("TODO: implement ReceiveN")
}
