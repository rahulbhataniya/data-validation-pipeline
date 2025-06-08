package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/rahulbhataniya/data-validation-pipeline/kafka"
	"github.com/rahulbhataniya/data-validation-pipeline/model"
)

func main() {
	brokers := []string{"localhost:9092"}
	topic := "product.ingest"

	product := model.Product{
		ID:                uuid.NewString(),
		Name:              "Rahul's Product",
		QuantityAvailable: 100,
		QuantityOrdered:   20,
		LeadTime:          5,
		SupplierID:        "supplier-123",
		LastUpdated:       time.Now().Format("2006-01-02"),
	}

	err := kafka.SendProductToKafka(brokers, topic, product)
	if err != nil {
		panic(err)
	}
}
