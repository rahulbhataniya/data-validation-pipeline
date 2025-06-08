package kafka

import (
	"context"
	"fmt"
	"log"

	kafkago "github.com/segmentio/kafka-go"
)

// MessageHandler defines a function type for processing consumed messages
type MessageHandler func(message []byte) error

func StartConsumer(brokers []string, topic string, groupID string, handler MessageHandler) error {
	r := kafkago.NewReader(kafkago.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
	})

	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		err = handler(m.Value)
		if err != nil {
			log.Printf("Handler error: %v", err)
			err := writeToDLQ(m.Value)
			if err != nil {
				return fmt.Errorf("failed to write to DLQ: %w", err)
			}
			// optionally send to DLQ or handle error here
		} else {
			log.Printf("Consumed message: %s", string(m.Value))
		}
	}
}
