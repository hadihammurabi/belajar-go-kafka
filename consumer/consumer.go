package consumer

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	Topic string
}

func New(opts ...OptionFunc) *Consumer {
	consumer := &Consumer{}

	for _, opt := range opts {
		opt(consumer)
	}

	return consumer
}

func (c *Consumer) Start(onMessage func(message kafka.Message)) error {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", c.Topic, 0)
	if err != nil {
		return err
	}
	defer conn.Close()

	for {
		msg, err := conn.ReadMessage(10e3)
		if err != nil {
			break
		}
		onMessage(msg)
	}

	return nil
}
