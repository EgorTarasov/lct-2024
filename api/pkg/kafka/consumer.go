package kafka

import (
	"time"

	"github.com/IBM/sarama"
)

// Reader обработчик сообщений логера
type Reader struct {
	brokers  []string
	Consumer sarama.Consumer
}

// NewConsumer конструктор обработчика сообщений логера
func NewConsumer(cfg Config) (*Reader, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = false
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 5 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumer, err := sarama.NewConsumer(cfg.Brokers, config)

	if err != nil {
		return nil, err
	}

	return &Reader{
		brokers:  cfg.Brokers,
		Consumer: consumer,
	}, err
}
