package entity

import "time"

// Merchant represent table merchants in database
type Merchant struct {
	ID           int64     `json:"id"`
	UserId       string    `json:"user_id"`
	MerchantName string    `json:"merchant_name"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    int64     `json:"created_by"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    int64     `json:"updated_by"`
}

func (t *Merchant) TableName() string {
	return "merchants"
}
