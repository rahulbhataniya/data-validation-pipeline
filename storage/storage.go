package storage

import (
	"log"

	"github.com/rahulbhataniya/data-validation-pipeline/model"
)

func Save(product model.Product) {
	log.Printf("✅ Product stored: %+v\n", product)
}
