package consumer

type OptionFunc func(consumer *Consumer)

func WithTopic(topic string) OptionFunc {
	return func(consumer *Consumer) {
		consumer.Topic = topic
	}
}
