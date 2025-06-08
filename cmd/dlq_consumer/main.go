package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/segmentio/kafka-go"
)

const (
	dlqTopic   = "product.dlq"
	groupID    = "dlq-consumer-group"
	brokerAddr = "localhost:9092" // change if you're using Docker host IP
)

func main() {
	log.Println("🚨 Starting DLQ consumer...")

	// Kafka reader config
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{brokerAddr},
		GroupID:     groupID,
		Topic:       dlqTopic,
		StartOffset: kafka.LastOffset,
	})

	defer func() {
		if err := r.Close(); err != nil {
			log.Fatalf("❌ Failed to close reader: %v", err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())

	// Graceful shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		log.Println("🛑 DLQ consumer shutting down...")
		cancel()
	}()

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				break // graceful shutdown
			}
			log.Printf("❌ Error reading DLQ message: %v\n", err)
			continue
		}

		log.Printf("📨 [DLQ] Received message at offset %d: %s\n", m.Offset, string(m.Value))

		// 🚨 Simulate alert (log or HTTP webhook)
		err = sendAlert(m.Key, m.Value)
		if err != nil {
			log.Printf("⚠️ Failed to send alert: %v\n", err)
		}
	}
}

// sendAlert sends alert for invalid message
func sendAlert(key, value []byte) error {
	// Simulated alerting (can integrate with Slack, Email, etc.)
	log.Printf("🚨 Alert: Invalid product data [key=%s]: %s\n", string(key), string(value))

	// You can integrate webhook calls like this:
	/*
		payload := map[string]string{"message": string(value)}
		data, _ := json.Marshal(payload)
		http.Post("http://alert-service/webhook", "application/json", bytes.NewBuffer(data))
	*/
	return nil
}
