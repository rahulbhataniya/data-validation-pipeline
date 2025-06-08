package main

import (
	"encoding/json"
	"log"

	"github.com/rahulbhataniya/data-validation-pipeline/kafka"
	"github.com/rahulbhataniya/data-validation-pipeline/model"
	"github.com/rahulbhataniya/data-validation-pipeline/validation"
)

func handleMessage(msg []byte) error {
	var p model.Product
	if err := json.Unmarshal(msg, &p); err != nil {
		log.Printf("❌ JSON unmarshal error: %v", err)
		// send to DLQ here
		return err
	}

	v := validation.Validator{}
	if err := v.Validate(p); err != nil {
		log.Printf("❌ Validation failed: %v", err)
		// send to DLQ here
		return err
	}

	log.Printf("✅ Product is valid: %v", p)
	return nil
	// store or further process product here
}

func main() {
	kafka.StartConsumer([]string{"localhost:9092"}, "product.ingest", "product-validator-group", handleMessage)
}
