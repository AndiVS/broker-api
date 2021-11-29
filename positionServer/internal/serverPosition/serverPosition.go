// Package serverPosition
package serverPosition

import (
	"context"
	"github.com/AndiVS/broker-api/positionServer/internal/service"
	"github.com/AndiVS/broker-api/positionServer/protocolPosition"
	"github.com/AndiVS/broker-api/priceServer/model"
	"github.com/google/uuid"
	"sync"
)

// PositionServer struct for grcp
type PositionServer struct {
	Service     service.Positions
	mu          *sync.Mutex
	currencyMap map[string]*model.Currency
	*protocolPosition.UnimplementedPositionServiceServer
}

// NewPositionServer constructor
func NewPositionServer(Service service.Positions, mu *sync.Mutex, currencyMap map[string]*model.Currency) *PositionServer {
	return &PositionServer{Service: Service, mu: mu, currencyMap: currencyMap}
}

// OpenPosition add transaction
func (t *PositionServer) OpenPosition(ctx context.Context, in *protocolPosition.OpenRequest) (*protocolPosition.OpenResponse, error) {
	if in.Price == t.currencyMap[in.CurrencyName].CurrencyPrice {
		id, err := t.Service.OpenPosition(ctx, in.CurrencyName, t.currencyMap[in.CurrencyName].Time, in.CurrencyAmount, t.currencyMap[in.CurrencyName].CurrencyPrice)
		if err != nil {
			return nil, err
		}
		return &protocolPosition.OpenResponse{PositionID: id.String()}, nil
	}
	return nil, nil //"price changed current price"
}

// ClosePosition add transaction
func (t *PositionServer) ClosePosition(ctx context.Context, in *protocolPosition.CloseRequest) (*protocolPosition.CloseResponse, error) {
	id, err := uuid.Parse(in.PositionID)
	if err != nil {
		return &protocolPosition.CloseResponse{Error: err.Error()}, nil
	}
	err = t.Service.ClosePosition(ctx, id, t.currencyMap[in.CurrencyName].CurrencyPrice)
	return &protocolPosition.CloseResponse{}, nil
}
