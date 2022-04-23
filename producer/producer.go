package producer

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	kafka "github.com/segmentio/kafka-go"
)

type Producer struct {
	Topic       string
	Message     string
	Interactive bool
}

func New(opts ...OptionFunc) *Producer {
	producer := &Producer{}

	for _, opt := range opts {
		opt(producer)
	}

	return producer
}

func (s *Producer) Start() error {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", s.Topic, 0)
	if err != nil {
		return err
	}
	defer conn.Close()

	if s.Interactive {
		for {
			inreader := bufio.NewReader(os.Stdin)
			fmt.Print("> ")
			messageString, err := inreader.ReadString('\n')
			if err != nil {
				return err
			}
			message := kafka.Message{
				Value: []byte(strings.ReplaceAll(messageString, "\n", "")),
			}
			_, err = conn.WriteMessages(message)
			if err != nil {
				return err
			}
		}
	} else {
		message := kafka.Message{
			Value: []byte(s.Message),
		}
		_, err = conn.WriteMessages(message)
		if err != nil {
			return err
		}
	}

	return nil
}
