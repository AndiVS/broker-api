// Package positionserver for work with grcp
package positionserver

import (
	"context"
	modelLokal "github.com/AndiVS/broker-api/positionServer/internal/model"
	"github.com/AndiVS/broker-api/positionServer/internal/service"
	"github.com/AndiVS/broker-api/positionServer/positionProtocol"
	"github.com/AndiVS/broker-api/priceServer/model"
	"github.com/google/uuid"
	"sync"
)

// PositionServer struct for grcp
type PositionServer struct {
	Service     service.Positions
	mu          *sync.Mutex
	currencyMap *map[string]*model.Currency
	*positionProtocol.UnimplementedPositionServiceServer
}

// NewPositionServer constructor
func NewPositionServer(Service service.Positions, mu *sync.Mutex, currencyMap *map[string]*model.Currency) *PositionServer {
	return &PositionServer{Service: Service, mu: mu, currencyMap: currencyMap}
}

// OpenPosition add transaction
func (t *PositionServer) OpenPosition(ctx context.Context, in *positionProtocol.OpenRequest) (*positionProtocol.OpenResponse, error) {
	if in.Price == (*t.currencyMap)[in.CurrencyName].CurrencyPrice {
		id1 := uuid.New()
		position := modelLokal.Position{PositionID: id1, CurrencyName: in.CurrencyName, Amount: in.CurrencyAmount,
			OpenPrice: (*t.currencyMap)[in.CurrencyName].CurrencyPrice, OpenTime: (*t.currencyMap)[in.CurrencyName].Time}
		id, err := t.Service.OpenPosition(ctx, &position)
		if err != nil {
			return nil, err
		}
		return &positionProtocol.OpenResponse{PositionID: id.String()}, nil
	}
	return nil, nil //"price changed current price"
}

// ClosePosition add transaction
func (t *PositionServer) ClosePosition(ctx context.Context, in *positionProtocol.CloseRequest) (*positionProtocol.CloseResponse, error) {
	id, err := uuid.Parse(in.PositionID)
	if err != nil {
		return &positionProtocol.CloseResponse{}, err
	}
	err = t.Service.ClosePosition(ctx, id, (*t.currencyMap)[in.CurrencyName].CurrencyPrice)
	if err != nil {
		return &positionProtocol.CloseResponse{}, err
	}
	return &positionProtocol.CloseResponse{}, nil
}
