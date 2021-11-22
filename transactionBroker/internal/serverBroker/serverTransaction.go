// Package serverBroker
package serverBroker

import (
	"context"
	"github.com/AndiVS/broker-api/transactionBroker/internal/service"
	"github.com/AndiVS/broker-api/transactionBroker/protocolBroker"
)

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

	id, err := t.Service.BuyCurrency(ctx, &in.CurrencyName, &in.CurrencyAmount)
	if err != nil {
		return nil, err
	}
	return &protocolBroker.BuyResponse{TransactionID: id.String()}, nil
}
