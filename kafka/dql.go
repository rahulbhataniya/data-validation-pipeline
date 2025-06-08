package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func writeToDLQ(value []byte) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "product.dlq",
	})
	defer writer.Close()

	msg := kafka.Message{
		Key:   []byte("invalid-product"),
		Value: value,
	}

	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Printf("❌ Failed to write to DLQ: %v", err)
		return err
	}

	log.Println("📤 Sent invalid message to DLQ")
	return nil
}
