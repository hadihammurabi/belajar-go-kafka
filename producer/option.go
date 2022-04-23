package producer

type OptionFunc func(producer *Producer)

func WithTopic(topic string) OptionFunc {
	return func(producer *Producer) {
		producer.Topic = topic
	}
}

func WithMessage(message string) OptionFunc {
	return func(producer *Producer) {
		producer.Message = message
	}
}

func WithInteractive(interactive bool) OptionFunc {
	return func(producer *Producer) {
		producer.Interactive = interactive
	}
}
