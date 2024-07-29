package kafka

import (
	"github.com/segmentio/kafka-go"
	"log"
	"messaggioService/models"
	"os"
)

var producer *kafka.Writer

func InitProducer() {
	producer = &kafka.Writer{
		Addr:     kafka.TCP(os.Getenv("KAFKA_BROKER")),
		Topic:    "messages",
		Balancer: &kafka.LeastBytes{},
	}
}

func CloseProducer() {
	producer.Close()
}

func SendMessage(message models.Message) error {
	msg := kafka.Message{
		Key:   []byte(string(message.ID)),
		Value: []byte(message.Content),
	}

	if err := producer.WriteMessages(nil, msg); err != nil {
		log.Printf("Could not write message to kafka: %v", err)
		return err
	}
	return nil
}
