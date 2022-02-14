package response

import (
	"majoo/entity"
	"time"
)

type TransactionResponse struct {
	Date         time.Time `json:"Date"`
	MerchantName string    `json:"merchant_name"`
	OutletName   string    `json:"outlet_name"`
	Omzet        float64   `json:"omzet"`
}

type Transaction struct {
	ID           int64     `json:"id"`
	MerchantId   string    `json:"merchant_id"`
	MerchantName string    `json:"merchant_name"`
	OutletId     string    `json:"outlet_id"`
	OutletName   string    `json:"outlet_name"`
	BillTotal    float64   `json:"bill_total"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    int64     `json:"created_by"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    int64     `json:"updated_by"`
	Date         time.Time `json:"date"`
}

func (t *Transaction) TableName() string {
	return "transactions"
}

func NewTransactionResponse(transaction entity.Transaction) TransactionResponse {
	transactionRes := TransactionResponse{
		Date:         transaction.CreatedAt,
		MerchantName: transaction.Merchant.MerchantName,
		OutletName:   transaction.Outlet.OutletName,
		Omzet:        transaction.BillTotal,
	}

	return transactionRes
}

func NewTransactionArrayResponse(transaction []Transaction) []TransactionResponse {
	transactionRes := []TransactionResponse{}
	for _, v := range transaction {
		if (v.MerchantName == "" || v.OutletName == "") && v.BillTotal !=0 {
			continue
		}
		t := TransactionResponse{
			Date:         v.Date,
			MerchantName: v.MerchantName,
			OutletName:   v.OutletName,
			Omzet:        v.BillTotal,
		}
		transactionRes = append(transactionRes, t)
	}
	return transactionRes
}
