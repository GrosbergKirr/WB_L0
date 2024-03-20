package stream

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"time"
)

func PublishToNATS(msg string) error {
	// connect to nats server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return fmt.Errorf("could not connect to nats: %w", err)
	}

	// create jetstream context from nats connection
	js, err := jetstream.New(nc)
	if err != nil {
		return fmt.Errorf("stream inition: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	// get existing stream handle
	_, err = js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{Name: "orders", Subjects: []string{"subject"}})
	if err != nil {
		return fmt.Errorf("stream creation: %w", err)
	}

	_, err = js.Publish(ctx, "subject", []byte(msg))
	if err != nil {
		return fmt.Errorf("publishing: %w", err)
	}

	return nil
}
