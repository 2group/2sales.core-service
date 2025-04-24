package kafka

import (
	"github.com/IBM/sarama"
)

type KafkaPublisher struct {
	producer sarama.SyncProducer
}

func NewKafkaPublisher(brokers string) (*KafkaPublisher, error) {
	brokerList := []string{brokers}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		return nil, err
	}

	return &KafkaPublisher{
		producer: producer,
	}, nil
}

func (kp *KafkaPublisher) Publish(topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	_, _, err := kp.producer.SendMessage(msg)
	return err
}

func (kp *KafkaPublisher) Close() error {
	return kp.producer.Close()
}
