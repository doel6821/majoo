package dto

type ReportDto struct {
	DateFrom string `json:"date_from" example:"2021-11-01"`
	DateTo string `json:"date_to" example:"2021-11-30"`
}