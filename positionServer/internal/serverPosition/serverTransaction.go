// Package serverPosition
package serverPosition

import (
	"context"
	"github.com/AndiVS/broker-api/positionServer/internal/service"
	"github.com/AndiVS/broker-api/positionServer/protocolPosition"
	"github.com/AndiVS/broker-api/priceServer/model"
	"sync"
)

// TransactionServer struct for grcp
type TransactionServer struct {
	Service     service.Transactions
	mu          *sync.Mutex
	currencyMap map[string]*model.Currency
	*protocolPosition.UnimplementedTransactionServiceServer
}

// NewTransactionServer constructor
func NewTransactionServer(Service service.Transactions, mu *sync.Mutex, currencyMap map[string]*model.Currency) *TransactionServer {
	return &TransactionServer{Service: Service, mu: mu, currencyMap: currencyMap}
}

// BuyCurrency add transaction
func (t *TransactionServer) BuyCurrency(ctx context.Context, in *protocolPosition.BuyRequest) (*protocolPosition.BuyResponse, error) {

	id, err := t.Service.BuyCurrency(ctx, in.CurrencyName, t.currencyMap[in.CurrencyName].Time, in.CurrencyAmount, t.currencyMap[in.CurrencyName].CurrencyPrice)
	if err != nil {
		return nil, err
	}
	return &protocolPosition.BuyResponse{TransactionID: id.String()}, nil
}
