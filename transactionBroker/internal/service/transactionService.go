package service

import (
	"context"
	"github.com/google/uuid"
	"serverBroker/internal/model"
	"serverBroker/internal/repository"
)

type Transactions interface {
	BuyCurrency(c context.Context, curr *model.Currency, amount *int64) (*uuid.UUID, error)
}


type TransactionService struct {
	Rep repository.Transactions
}


// BuyCurrency add record about transaction
func (s *TransactionService) BuyCurrency(c context.Context, curr *model.Currency, amount *int64) (*uuid.UUID, error) {
	transaction := model.Transaction{TransactionID: uuid.New(), CurrencyID: curr.ID, Amount: *amount, Price: curr.Price, Time: curr.Time}
/*	s.CatMap[cat.ID.String()] = cat
	s.Broker.ProduceEvent("cat", "Insert", *cat)*/
	return s.Rep.InsertTransaction(c, &transaction)
}

