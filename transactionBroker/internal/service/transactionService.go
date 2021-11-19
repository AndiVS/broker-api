// Package service
package service

import (
	"context"
	"github.com/AndiVS/broker-api/transactionBroker/internal/model"
	"github.com/AndiVS/broker-api/transactionBroker/internal/repository"
	"github.com/google/uuid"
)

type Transactions interface {
	BuyCurrency(c context.Context, curr *model.Currency, amount *int64) (*uuid.UUID, error)
}

type TransactionService struct {
	Rep repository.Transactions
}

func NewTransactionService(Rep interface{}) Transactions {
	return &TransactionService{Rep: Rep.(*repository.Postgres)}
}

// BuyCurrency add record about transaction
func (s *TransactionService) BuyCurrency(c context.Context, curr *model.Currency, amount *int64) (*uuid.UUID, error) {
	transaction := model.Transaction{TransactionID: uuid.New(), CurrencyID: curr.ID, Amount: *amount, Price: curr.Price, Time: curr.Time}
	/*	s.CatMap[cat.ID.String()] = cat
		s.Broker.ProduceEvent("cat", "Insert", *cat)*/
	return s.Rep.InsertTransaction(c, &transaction)
}
