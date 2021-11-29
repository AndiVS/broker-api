// Package positionServer
package positionServer

import (
	"context"
	modelLokal "github.com/AndiVS/broker-api/positionServer/internal/model"
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
	currencyMap *map[string]*model.Currency
	*protocolPosition.UnimplementedPositionServiceServer
}

// NewPositionServer constructor
func NewPositionServer(Service service.Positions, mu *sync.Mutex, currencyMap *map[string]*model.Currency) *PositionServer {
	return &PositionServer{Service: Service, mu: mu, currencyMap: currencyMap}
}

// OpenPosition add transaction
func (t *PositionServer) OpenPosition(ctx context.Context, in *protocolPosition.OpenRequest) (*protocolPosition.OpenResponse, error) {
	if in.Price == (*t.currencyMap)[in.CurrencyName].CurrencyPrice {
		id1 := uuid.New()
		position := modelLokal.Position{PositionID: id1, CurrencyName: in.CurrencyName, Amount: in.CurrencyAmount,
			OpenPrice: (*t.currencyMap)[in.CurrencyName].CurrencyPrice, OpenTime: (*t.currencyMap)[in.CurrencyName].Time}
		id, err := t.Service.OpenPosition(ctx, &position)
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
		return &protocolPosition.CloseResponse{}, err
	}
	err = t.Service.ClosePosition(ctx, id, (*t.currencyMap)[in.CurrencyName].CurrencyPrice)
	if err != nil {
		return &protocolPosition.CloseResponse{}, err
	}
	return &protocolPosition.CloseResponse{}, nil
}
