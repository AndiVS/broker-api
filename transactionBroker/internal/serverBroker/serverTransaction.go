package serverBroker

import (
	"context"
	"github.com/google/uuid"
	"serverBroker/internal/model"
	"serverBroker/internal/protocolBroker"
	"serverBroker/internal/service"
	"time"
)

const layout = "2006-01-02T15:04:05.000Z"

// TransactionServer struct for grcp
type TransactionServer struct {
	Service service.Transactions
	*protocolBroker.UnimplementedTransactionServiceServer
}

// NewTransactionServer as
func NewTransactionServer(Service service.Transactions) *TransactionServer {
	return &TransactionServer{Service: Service}
}

// BuyCurrency Cat about cat
func (t *TransactionServer) BuyCurrency(ctx context.Context, in *protocolBroker.BuyRequest) (*protocolBroker.BuyResponse, error) {
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
	return &protocolBroker.BuyResponse{TransactionID: id.String()}, nil
}
