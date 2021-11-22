// Package service
package service

import (
	"context"
	"github.com/AndiVS/broker-api/priceBuffer/protocolPrice"
	"github.com/AndiVS/broker-api/transactionBroker/internal/model"
	"github.com/AndiVS/broker-api/transactionBroker/internal/repository"
	"github.com/google/uuid"
)

type Transactions interface {
	BuyCurrency(c context.Context, name *string, amount *int64) (*uuid.UUID, error)
}

type TransactionService struct {
	Rep         repository.Transactions
	CurrencyMap map[string]*protocolPrice.Currency
}

func NewTransactionService(Rep interface{}, currencyMap map[string]*protocolPrice.Currency) Transactions {
	return &TransactionService{Rep: Rep.(*repository.Postgres), CurrencyMap: currencyMap}
}

// BuyCurrency add record about transaction
func (s *TransactionService) BuyCurrency(c context.Context, name *string, amount *int64) (*uuid.UUID, error) {
	transaction := model.Transaction{TransactionID: uuid.New(), CurrencyName: *name, Amount: *amount,
		Price: s.CurrencyMap[*name].CurrencyPrice, Time: s.CurrencyMap[*name].Time}
	return s.Rep.InsertTransaction(c, &transaction)
}
