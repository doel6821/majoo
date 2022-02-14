package entity

import "time"

// Transaction represent table transactions in database
type Transaction struct {
	ID         int64     `json:"id"`
	MerchantId string    `json:"merchant_id"`
	Merchant   Merchant  `json:"merchant" gorm:"foreignKey:MerchantId;references:ID"`
	OutletId   string    `json:"outlet_id"`
	Outlet     Outlet    `json:"outlet" gorm:"foreignKey:OutletId;references:ID"`
	BillTotal  float64    `json:"bill_total"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  int64     `json:"created_by"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  int64     `json:"updated_by"`
}

func (t *Transaction) TableName() string {
	return "transactions"
}
