package main

import (
	"flag"
	"fmt"

	"github.com/hadihammurabi/belajar-go-kafka/consumer"
	"github.com/hadihammurabi/belajar-go-kafka/producer"
	"github.com/segmentio/kafka-go"
)

func main() {
	var run string
	flag.StringVar(&run, "run", "consumer", "what you want to run")

	var topic string
	flag.StringVar(&topic, "topic", "main", "what topic you want to use")

	var message string
	flag.StringVar(&message, "message", "", "what message you want to send")

	var interactive bool
	flag.BoolVar(&interactive, "interactive", false, "enable interactive mode")

	flag.Parse()

	if run == "consumer" {
		c := consumer.New(
			consumer.WithTopic(topic),
		)
		c.Start(func(message kafka.Message) {
			fmt.Println(string(message.Value))
		})
	}
	if run == "producer" {
		p := producer.New(
			producer.WithTopic(topic),
			producer.WithMessage(message),
			producer.WithInteractive(interactive),
		)
		if err := p.Start(); err != nil {
			panic(err)
		}
	}

}
