package entity

import "time"

// Outlet represent table outlets in database
type Outlet struct {
	ID         int64     `json:"id"`
	MerchantId string    `json:"merchant_id"`
	OutletName string    `json:"outlet_name"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  int64     `json:"created_by"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  int64     `json:"updated_by"`
}

func (t *Outlet) TableName() string {
	return "outlets"
}
