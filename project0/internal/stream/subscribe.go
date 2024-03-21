package stream

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"log"
	"project0/models"
	"project0/tools"
	"sync"
	"time"
)

type Saver interface {
	InsertJson(uid string, dataj []byte)
}

func SavetoDBandCache(saver Saver, cache map[string]models.Order) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	ch := make(chan []byte)
	var res []byte
	mu := sync.Mutex{}
loop:
	for {
		go SubscribeNATS(ch, ctx)
		select {
		case <-ctx.Done():
			break loop
		case res = <-ch:
			o := new(models.Order)

			err := json.Unmarshal(res, &o)
			if err != nil {
				log.Fatalf("unmarshal mistake,: %s", err)
			}
			if tools.Checker(cache, o.OrderUid) {
				mu.Lock()
				saver.InsertJson(o.OrderUid, res)

				cache[o.OrderUid] = *o
				fmt.Printf("Get order from stream. Uid: %s\n", o.OrderUid)
				mu.Unlock()
			}
		}
	}
	return nil
}

func SubscribeNATS(ch chan []byte, ctx context.Context) {
	// Connect to nats server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Errorf("could not connect to nats: %w", err)
	}

	// Create jetstream context from nats connection
	js, err := jetstream.New(nc)
	if err != nil {
		fmt.Errorf("jetstream mistake: %s", err)
	}

	// Get existing stream handle
	stream, err := js.Stream(ctx, "orders")
	if err != nil {
		fmt.Errorf("connect to stream mistake: %s", err)
	}

	// Retrieve consumer handle from a stream
	cons, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{Name: "cons"})
	if err != nil {
		fmt.Errorf("create consumer mistake: %s", err)
	}
	// Get messages from stream
	mCons, err := cons.Messages()
	if err != nil {
		panic(err)
	}

	for {
		msg, err := mCons.Next()
		if err != nil {
			panic(err)
		}
		select {
		case <-ctx.Done():
			fmt.Println("exit")
			close(ch)
			mCons.Stop()
			err = js.DeleteStream(ctx, "orders")
			if err != nil {
				panic(err)
			}
			return
		case ch <- msg.Data():
			err = msg.Ack()
			if err != nil {
				panic(err)
			}
		}
	}

}
