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
	"time"
)

type Saver interface {
	InsertJson(uid string, dataj []byte)
}

func SavetoDBandCache(saver Saver, cache map[string]models.Order) error {
	order, err := SubscribeNATS()
	if err != nil {
		log.Fatalf("read from stream mistake,: %s", err)
	}

	o := new(models.Order)

	err = json.Unmarshal(order, &o)
	if err != nil {
		log.Fatalf("unmarshal mistake,: %s", err)
	}
	if tools.Checker(cache, o.OrderUid) {
		saver.InsertJson(o.OrderUid, order)

		cache[o.OrderUid] = *o
	}
	return err
}

func SubscribeNATS() ([]byte, error) {
	// connect to nats server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return []byte{}, fmt.Errorf("could not connect to nats: %w", err)
	}

	// create jetstream context from nats connection
	js, err := jetstream.New(nc)
	if err != nil {
		return []byte{}, fmt.Errorf("jetstream mistake: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// get existing stream handle
	stream, err := js.Stream(ctx, "orders")
	if err != nil {
		return []byte{}, fmt.Errorf("connect to stream mistake: %s", err)
	}

	// retrieve consumer handle from a stream
	cons, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{Name: "cons"})
	if err != nil {
		return []byte{}, fmt.Errorf("create consumer mistake: %s", err)
	}

	// consume messages from the consumer in callback
	cc, err := cons.Consume(handleMyMessage)
	if err != nil {
		return []byte{}, fmt.Errorf("consumer mistake: %s", err)
	}

	a, err := stream.GetLastMsgForSubject(ctx, "subject")
	if err != nil {
		return []byte{}, fmt.Errorf("get message mistake: %s", err)
	}
	id := a.Data

	defer js.DeleteStream(ctx, "orders")
	defer cc.Stop()
	//fmt.Scanln()

	return id, err

}

func handleMyMessage(msg jetstream.Msg) {
	fmt.Println("Get Uid")
	msg.Ack()
}
