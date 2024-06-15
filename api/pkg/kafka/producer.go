package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
)

// Producer для кафки в данном сервисе кафка будет использоваться только для логов
// из-за этого используется только sarama.AsyncProducer
// нам не очень критично потерять несколько логов
type Producer struct {
	brokers       []string
	asyncProducer sarama.AsyncProducer
}

func newAsyncProducer(brokers []string) (sarama.AsyncProducer, error) {
	asyncProducerConfig := sarama.NewConfig()

	asyncProducerConfig.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	asyncProducerConfig.Producer.RequiredAcks = sarama.NoResponse
	asyncProducerConfig.Producer.Idempotent = false

	asyncProducerConfig.Producer.Return.Successes = false
	asyncProducerConfig.Producer.Return.Errors = true
	asyncProducerConfig.Net.MaxOpenRequests = 5

	asyncProducer, err := sarama.NewAsyncProducer(brokers, asyncProducerConfig)

	if err != nil {
		return nil, fmt.Errorf("error with async kafka-producer")
	}

	go func() {
		for e := range asyncProducer.Errors() {
			fmt.Printf("error during writing to logs %s", e.Error())
		}
	}()

	return asyncProducer, nil
}

// NewProducer конструктор продюсера
func NewProducer(brokers []string) (*Producer, error) {
	asyncProducer, err := newAsyncProducer(brokers)
	if err != nil {
		return nil, fmt.Errorf("error with async kafka-producer")
	}

	producer := &Producer{
		brokers:       brokers,
		asyncProducer: asyncProducer,
	}

	return producer, nil
}

// SendAsyncMessage отправка сообщения
func (k *Producer) SendAsyncMessage(message *sarama.ProducerMessage) {
	k.asyncProducer.Input() <- message
}

// Close завершение работы с кафкой
func (k *Producer) Close() error {
	err := k.asyncProducer.Close()
	if err != nil {
		return fmt.Errorf("kafka.Connector.Close")
	}
	return nil
}
