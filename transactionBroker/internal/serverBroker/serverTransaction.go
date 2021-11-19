// Package serverBroker
package serverBroker

import (
	"context"
	"github.com/AndiVS/broker-api/transactionBroker/internal/model"
	"github.com/AndiVS/broker-api/transactionBroker/internal/service"
	protocolBroker2 "github.com/AndiVS/broker-api/transactionBroker/protocolBroker"
	"github.com/google/uuid"
	"time"
)

const layout = "2006-01-02T15:04:05.000Z"

// TransactionServer struct for grcp
type TransactionServer struct {
	Service service.Transactions
	*protocolBroker2.UnimplementedTransactionServiceServer
}

// NewTransactionServer as
func NewTransactionServer(Service service.Transactions) *TransactionServer {
	return &TransactionServer{Service: Service}
}

// BuyCurrency Cat about cat
func (t *TransactionServer) BuyCurrency(ctx context.Context, in *protocolBroker2.BuyRequest) (*protocolBroker2.BuyResponse, error) {
	cid, err := uuid.Parse(in.Currency.CurrencyID)
	if err != nil {
		return nil, err
	}

	tim, err := time.Parse(layout, in.Currency.Time)
	if err != nil {
		return nil, err
	}

	curr := model.Currency{ID: cid, Name: in.Currency.CurrencyName, Price: in.Currency.CurrencyPrice, Time: tim}

	id, err := t.Service.BuyCurrency(ctx, &curr, &in.CurrencyAmount)
	if err != nil {
		return nil, err
	}
	return &protocolBroker2.BuyResponse{TransactionID: id.String()}, nil
}
