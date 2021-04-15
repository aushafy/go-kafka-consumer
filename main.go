package main

import (
	"context"
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
)

var (
	topic           = os.Getenv("KAFKA_TOPIC")
	kafkaBroker     = os.Getenv("KAFKA_HOST")
	consumerGroupId = os.Getenv("CONSUMER_GROUP_ID")
)

func main() {
	consume(context.Background())
}

func consume(ctx context.Context) {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaBroker},
		Topic:   topic,
		GroupID: consumerGroupId,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
	}
}
