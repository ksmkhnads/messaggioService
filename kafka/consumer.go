package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"messaggioService/db"
	"messaggioService/models"
	"os"
)

func ConsumeMessages() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{os.Getenv("KAFKA_BROKER")},
		Topic:   "messages",
		GroupID: "message-consumer-group",
	})
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Could not read message: %v", err)
			continue
		}

		var message models.Message
		db.DB.First(&message, "id = ?", msg.Key)
		message.Processed = true
		db.DB.Save(&message)
		log.Printf("Processed message: %s", msg.Value)
	}
}
