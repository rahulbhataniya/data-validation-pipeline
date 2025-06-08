package model

type Product struct {
	ID                string `json:"id" validate:"required,uuid4"`
	Name              string `json:"name" validate:"required"`
	QuantityAvailable int    `json:"quantity_available" validate:"gte=0"`
	QuantityOrdered   int    `json:"quantity_ordered" validate:"gte=0"`
	LeadTime          int    `json:"lead_time" validate:"gte=0"` // days
	SupplierID        string `json:"supplier_id" validate:"required"`
	LastUpdated       string `json:"last_updated" validate:"required,datetime=2006-01-02"`
}
