package repo

import (
	"majoo/response"
	"time"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	TransactionReport(dateFrom, date_to time.Time, id int64, page, limit int) ([]response.Transaction, error)
	
}

type transactionRepo struct {
	connection *gorm.DB
}

func NewTransactionRepo(connection *gorm.DB) TransactionRepository {
	return &transactionRepo{
		connection: connection,
	}
}

func (c *transactionRepo) TransactionReport(dateFrom, dateTo time.Time, id int64, page, limit int) ([]response.Transaction, error) {
	var res []response.Transaction

	err := c.connection.Table("(SELECT i::date AS date FROM generate_series(?, ?, interval '1 day') i) s ", dateFrom, dateTo).
		Select("t.id, t.bill_total ,m.merchant_name, o.outlet_name, s.date ").
		Joins("LEFT JOIN transactions t on t.created_at::date = s.date").
		Joins("LEFT JOIN merchants m on t.merchant_id = m.id and m.user_id =?", id).
		Joins("LEFT JOIN outlets o on t.outlet_id = o.id").
		Where("s.date >= ? and s.date <= ? ", dateFrom, dateTo).
		Limit(limit).Offset((page - 1) * limit).Group("t.id,m.merchant_name, o.outlet_name, s.date").Order("s.date asc").
		Find(&res).Limit(-1).Offset(0).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

// select t.* ,m.merchant_name, o.outlet_name, s.date
// from (SELECT i::date AS date FROM generate_series('2021-11-01', '2021-11-30', interval '1 day') i) s 
// left join transactions t on t.created_at::date = s.date
// left join merchants m on t.merchant_id = m.id
// left join outlets o on t.outlet_id = o.id
// group by t.id,m.merchant_name, o.outlet_name, s.date
// order by s.date asc
