// Package service
package service

import (
	"context"
	"github.com/AndiVS/broker-api/transactionBroker/internal/model"
	"github.com/AndiVS/broker-api/transactionBroker/internal/repository"
	"github.com/google/uuid"
)

// Transactions interface for transactionService
type Transactions interface {
	BuyCurrency(c context.Context, name, time string, amount int64, price float32) (*uuid.UUID, error)
}

// TransactionService struct for service
type TransactionService struct {
	Rep repository.Transactions
}

// NewTransactionService constructor
func NewTransactionService(Rep interface{}) Transactions {
	return &TransactionService{Rep: Rep.(*repository.Postgres)}
}

// BuyCurrency add record about transaction
func (s *TransactionService) BuyCurrency(c context.Context, name, time string, amount int64, price float32) (*uuid.UUID, error) {
	transaction := model.Transaction{TransactionID: uuid.New(), CurrencyName: name, Amount: amount,
		Price: price, TransactionTime: time}
	return s.Rep.InsertTransaction(c, &transaction)
}
