// Package server for work with grcp
package server

import (
	"context"
	"sync"

	modelLocal "github.com/AndiVS/broker-api/positionServer/internal/model"
	"github.com/AndiVS/broker-api/positionServer/internal/service"
	"github.com/AndiVS/broker-api/positionServer/positionProtocol"
	"github.com/AndiVS/broker-api/priceServer/model"
	"github.com/google/uuid"
)

// PositionServer struct for grcp
type PositionServer struct {
	Service     service.Positions
	mu          *sync.Mutex
	currencyMap map[string]*model.Currency
	profileMap  map[uuid.UUID]*modelLocal.Profile
	*positionProtocol.UnimplementedPositionServiceServer
}

// NewPositionServer constructor
func NewPositionServer(Service service.Positions, mu *sync.Mutex, currencyMap map[string]*model.Currency, profileMap map[uuid.UUID]*modelLocal.Profile) *PositionServer {
	return &PositionServer{Service: Service, mu: mu, currencyMap: currencyMap, profileMap: profileMap}
}

// OpenPosition add transaction
func (t *PositionServer) OpenPosition(ctx context.Context, in *positionProtocol.OpenRequest) (*positionProtocol.OpenResponse, error) {
	t.mu.Lock()
	if in.Price == t.currencyMap[in.CurrencyName].CurrencyPrice {
		if *t.profileMap[idU].Balance > in.Price*float32(in.CurrencyAmount) {
			*t.profileMap[idU].Balance -= in.Price * float32(in.CurrencyAmount)
			id1 := uuid.New()
			t.profileMap[idU].PositionList = append(t.profileMap[idU].PositionList, &id1)
			position := modelLocal.Position{PositionID: &id1, CurrencyName: in.CurrencyName, Amount: &in.CurrencyAmount,
				OpenPrice: &t.currencyMap[in.CurrencyName].CurrencyPrice, OpenTime: t.currencyMap[in.CurrencyName].Time, TakeProfit: &in.TakeProfit, StopLoss: &in.StopLoss}
			id, err := t.Service.OpenPosition(ctx, &position)
			t.mu.Unlock()
			if err != nil {
				return nil, err
			}
			return &positionProtocol.OpenResponse{PositionID: id.String()}, nil
		}
	}
	return nil, nil
}

// ClosePosition add transaction
func (t *PositionServer) ClosePosition(ctx context.Context, in *positionProtocol.CloseRequest) (*positionProtocol.CloseResponse, error) {
	id, err := uuid.Parse(in.PositionID)
	if err != nil {
		return &positionProtocol.CloseResponse{}, err
	}
	err = t.Service.ClosePosition(ctx, id, t.currencyMap[in.CurrencyName].CurrencyPrice)
	if err != nil {
		return &positionProtocol.CloseResponse{}, err
	}
	return &positionProtocol.CloseResponse{}, nil
}
