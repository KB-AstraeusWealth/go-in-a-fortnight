package day09

import (
	"context"
	"os"
	"sync"
	"testing"

	"cloud.google.com/go/pubsub"
)

func testClient(t *testing.T) *pubsub.Client {
	t.Helper()
	if os.Getenv("PUBSUB_EMULATOR_HOST") == "" {
		t.Fatal("PUBSUB_EMULATOR_HOST not set — start the emulator (see advanced/README.md)")
	}
	proj := os.Getenv("PUBSUB_PROJECT_ID")
	if proj == "" {
		proj = "go-week"
	}
	c, err := pubsub.NewClient(context.Background(), proj)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	return c
}

func TestPublishReceive(t *testing.T) {
	ctx := context.Background()
	c := testClient(t)
	defer c.Close()

	topic, err := c.CreateTopic(ctx, "t-"+t.Name())
	if err != nil {
		t.Fatalf("CreateTopic: %v", err)
	}
	defer topic.Delete(ctx)

	sub, err := c.CreateSubscription(ctx, "s-"+t.Name(), pubsub.SubscriptionConfig{Topic: topic})
	if err != nil {
		t.Fatalf("CreateSubscription: %v", err)
	}
	defer sub.Delete(ctx)

	want := []string{"a", "b", "c"}
	for _, m := range want {
		if _, err := Publish(ctx, c, topic.ID(), []byte(m)); err != nil {
			t.Fatalf("Publish: %v", err)
		}
	}

	var mu sync.Mutex
	got := map[string]bool{}
	err = ReceiveN(ctx, c, sub.ID(), len(want), func(b []byte) {
		mu.Lock()
		got[string(b)] = true
		mu.Unlock()
	})
	if err != nil {
		t.Fatalf("ReceiveN: %v", err)
	}
	for _, w := range want {
		if !got[w] {
			t.Errorf("missing message %q (got %v)", w, got)
		}
	}
}
