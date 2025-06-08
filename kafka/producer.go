package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/rahulbhataniya/data-validation-pipeline/model"
	"github.com/segmentio/kafka-go"
	kafkago "github.com/segmentio/kafka-go"
)

func SendToDLQ(brokers []string, data []byte) {
	writer := kafkago.NewWriter(kafkago.WriterConfig{
		Brokers:  brokers,
		Topic:    "product.dlq",
		Balancer: &kafkago.LeastBytes{},
	})
	defer writer.Close()

	err := writer.WriteMessages(context.Background(),
		kafkago.Message{
			Value: data,
			Time:  time.Now(),
		},
	)
	if err != nil {
		log.Println("Failed to write to DLQ:", err)
	}
}

// SendProductToKafka sends a product message to Kafka topic
func SendProductToKafka(brokers []string, topic string, product model.Product) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	defer writer.Close()

	data, err := json.Marshal(product)
	if err != nil {
		return err
	}

	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(product.ID),
			Value: data,
			Time:  time.Now(),
		},
	)

	if err != nil {
		log.Printf("❌ Failed to send message: %v", err)
		return err
	}

	log.Printf("✅ Message sent to topic %s: %s", topic, string(data))
	return nil
}
