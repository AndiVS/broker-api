// Package serverBroker
package serverBroker

import (
	"context"
	"github.com/AndiVS/broker-api/priceBuffer/protocolPrice"
	"github.com/AndiVS/broker-api/transactionBroker/internal/service"
	"github.com/AndiVS/broker-api/transactionBroker/protocolBroker"
	"sync"
)

// TransactionServer struct for grcp
type TransactionServer struct {
	Service     service.Transactions
	mu          *sync.Mutex
	currencyMap map[string]protocolPrice.Currency
	*protocolBroker.UnimplementedTransactionServiceServer
}

// NewTransactionServer constructor
func NewTransactionServer(Service service.Transactions, mu *sync.Mutex, currencyMap map[string]protocolPrice.Currency) *TransactionServer {
	return &TransactionServer{Service: Service, mu: mu, currencyMap: currencyMap}
}

// BuyCurrency add transaction
func (t *TransactionServer) BuyCurrency(ctx context.Context, in *protocolBroker.BuyRequest) (*protocolBroker.BuyResponse, error) {

	id, err := t.Service.BuyCurrency(ctx, in.CurrencyName, t.currencyMap[in.CurrencyName].Time, in.CurrencyAmount, t.currencyMap[in.CurrencyName].CurrencyPrice)
	if err != nil {
		return nil, err
	}
	return &protocolBroker.BuyResponse{TransactionID: id.String()}, nil
}
