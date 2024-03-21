package stream

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"math/rand"
	"project0/models"
	"time"
)

func Publish() {
	uid := NewRandomString(20)
	d := models.Delivery{uid, "Test Testov", "+9720000000", 2639809, "Kiryat Mozkin",
		"Ploshad Mira 15", "Kraiot", "est@gmail.com"}

	p := models.Payment{uid, uid,
		"", "USD", "wbpay", 1817, 1637907727, "alpha",
		1500, 317, 0}
	i := models.Items{uid, 9934930, "WBILMTESTTRACK", 453, "ab4219087a764ae0btest",
		"Mascaras", 30, 0, 317, 2389212, "Vivienne Sabo", 202}
	var li []models.Items
	li = append(li, i)

	o := models.Order{uid, "WBILMTESTTRACK", "WBIL", d, p,
		li, "en", "", "test",
		"meest", 9, 99, "2024-03-14 18:59:34.013135", 1}

	a, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}

	err = PublishToNATS(a)
	if err != nil {
		fmt.Errorf("publish mistake", err)
	}

	fmt.Printf("Publish to stream: %s", uid)

}
func NewRandomString(size int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []rune("abcdefghijklmnopqrstuvwxyz" +
		"0123456789")

	b := make([]rune, size)
	for i := range b {
		b[i] = chars[rnd.Intn(len(chars))]
	}

	return string(b)
}

func PublishToNATS(msg []byte) error {
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

	_, err = js.Publish(ctx, "subject", msg)
	if err != nil {
		return fmt.Errorf("publishing: %w", err)
	}

	return nil
}
