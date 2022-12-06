package kafka_producers

import (
	"context"
	"github.com/Shopify/sarama"
	"log"
)

type KafkaConfig struct {
	Brokers []string
}

type Message struct {
	Topic string
	Key   []byte
	Value []byte
}

type SyncProducer interface {
	SendMessage(ctx context.Context, msg Message) (err error)
	Close() (err error)
}

type syncProducer struct {
	Producer sarama.SyncProducer
}

func NewSyncProducer(kafkaConfig KafkaConfig) (SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Version = sarama.V2_0_0_0
	producer, err := sarama.NewSyncProducer(kafkaConfig.Brokers, config)
	if err != nil {
		log.Print(nil, "FAILED to create producer: ", err)
		return nil, err
	}
	obj := syncProducer{Producer: producer}
	return obj, nil
}

func (sp syncProducer) SendMessage(ctx context.Context, msg Message) error {

	partition, offset, err := sp.Producer.SendMessage(convertToProducerMessage(msg))
	if err != nil {
		log.Fatal(ctx, "FAILED to send message: ", err)
		return err
	}
	log.Print(ctx, "> message sent to partition ", partition, "at offset ", offset)

	return nil
}

func convertToProducerMessage(msg Message) *sarama.ProducerMessage {
	message := &sarama.ProducerMessage{Topic: msg.Topic}

	if msg.Key != nil {
		message.Key = sarama.ByteEncoder(msg.Key)
	}
	if msg.Value != nil {
		message.Value = sarama.ByteEncoder(msg.Value)
	}
	return message
}

func (sp syncProducer) Close() error {
	if err := sp.Producer.Close(); err != nil {
		log.Fatal(nil, "FAILED to close the producer: ", err)
		return err
	}
	return nil
}
