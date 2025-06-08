package tests

import (
	"testing"

	"github.com/rahulbhataniya/data-validation-pipeline/model"
	"github.com/rahulbhataniya/data-validation-pipeline/validation"
)

func TestValidator_Validate(t *testing.T) {
	v := validation.Validator{}

	tests := []struct {
		name    string
		product model.Product
		wantErr bool
	}{
		{
			name: "valid product",
			product: model.Product{
				ID:                "123",
				Name:              "Product X",
				QuantityAvailable: 10,
				QuantityOrdered:   5,
				LeadTime:          30,
				SupplierID:        "SUP123",
				LastUpdated:       "2025-06-07",
			},
			wantErr: false,
		},
		{
			name: "missing product ID",
			product: model.Product{
				ID:                "",
				Name:              "Product X",
				QuantityAvailable: 10,
				QuantityOrdered:   5,
				LeadTime:          30,
				SupplierID:        "SUP123",
				LastUpdated:       "2025-06-07",
			},
			wantErr: true,
		},
		{
			name: "short product name",
			product: model.Product{
				ID:                "123",
				Name:              "AB",
				QuantityAvailable: 10,
				QuantityOrdered:   5,
				LeadTime:          30,
				SupplierID:        "SUP123",
				LastUpdated:       "2025-06-07",
			},
			wantErr: true,
		},
		{
			name: "negative quantity available",
			product: model.Product{
				ID:                "123",
				Name:              "Product X",
				QuantityAvailable: -1,
				QuantityOrdered:   5,
				LeadTime:          30,
				SupplierID:        "SUP123",
				LastUpdated:       "2025-06-07",
			},
			wantErr: true,
		},
		{
			name: "negative quantity ordered",
			product: model.Product{
				ID:                "123",
				Name:              "Product X",
				QuantityAvailable: 10,
				QuantityOrdered:   -5,
				LeadTime:          30,
				SupplierID:        "SUP123",
				LastUpdated:       "2025-06-07",
			},
			wantErr: true,
		},
		{
			name: "lead time too high",
			product: model.Product{
				ID:                "123",
				Name:              "Product X",
				QuantityAvailable: 10,
				QuantityOrdered:   5,
				LeadTime:          500,
				SupplierID:        "SUP123",
				LastUpdated:       "2025-06-07",
			},
			wantErr: true,
		},
		{
			name: "missing supplier ID",
			product: model.Product{
				ID:                "123",
				Name:              "Product X",
				QuantityAvailable: 10,
				QuantityOrdered:   5,
				LeadTime:          30,
				SupplierID:        "",
				LastUpdated:       "2025-06-07",
			},
			wantErr: true,
		},
		{
			name: "invalid date format",
			product: model.Product{
				ID:                "123",
				Name:              "Product X",
				QuantityAvailable: 10,
				QuantityOrdered:   5,
				LeadTime:          30,
				SupplierID:        "SUP123",
				LastUpdated:       "07-06-2025", // wrong format
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := v.Validate(tc.product)
			if (err != nil) != tc.wantErr {
				t.Errorf("Validate() error = %v, wantErr = %v", err, tc.wantErr)
			}
		})
	}
}
