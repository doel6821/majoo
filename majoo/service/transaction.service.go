package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"majoo/repo"
	transaction "majoo/response"
	"time"
)

type TransactionService interface {
	TransactionReport(dateFrom, dateTo time.Time, id int64, page, limit int) ([]transaction.TransactionResponse , error)
	
}

type transactionService struct {
	transactionRepo repo.TransactionRepository
}

func NewTransactionService(transactionRepo repo.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
	}
}


func (c *transactionService) TransactionReport(dateFrom, dateTo time.Time, id int64, page, limit int) ([]transaction.TransactionResponse, error) {
	var res []transaction.TransactionResponse 

	data , err := c.transactionRepo.TransactionReport(dateFrom, dateTo, id, page, limit)
	if err != nil {
		log.Println("failed get data db")
		return nil, errors.New("failed get data db")
	}
	dataByte, _ := json.Marshal(data)
	fmt.Println(string(dataByte))
	res = transaction.NewTransactionArrayResponse(data)
	return res, nil

}