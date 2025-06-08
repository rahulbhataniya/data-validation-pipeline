package validation

import (
	"errors"
	"regexp"
	"time"

	"github.com/rahulbhataniya/data-validation-pipeline/model"
)

// Validator struct - you can add config fields if needed
type Validator struct{}

// Validate checks product fields for validity
func (v *Validator) Validate(p model.Product) error {
	if p.ID == "" {
		return errors.New("product ID is required")
	}

	if len(p.Name) < 3 {
		return errors.New("product name must be at least 3 characters")
	}

	if p.QuantityAvailable < 0 {
		return errors.New("quantity available cannot be negative")
	}

	if p.QuantityOrdered < 0 {
		return errors.New("quantity ordered cannot be negative")
	}

	if p.LeadTime < 0 || p.LeadTime > 365 {
		return errors.New("lead time must be between 0 and 365 days")
	}

	if p.SupplierID == "" {
		return errors.New("supplier ID is required")
	}

	if !isValidDate(p.LastUpdated) {
		return errors.New("last updated date is invalid or not in YYYY-MM-DD format")
	}

	return nil
}

// isValidDate validates ISO date strings (YYYY-MM-DD)
func isValidDate(dateStr string) bool {
	// quick regex check
	match, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, dateStr)
	if !match {
		return false
	}

	// parse to time.Time
	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}
