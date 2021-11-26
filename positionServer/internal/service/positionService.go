// Package service
package service

import (
	"context"
	"github.com/AndiVS/broker-api/positionServer/internal/model"
	"github.com/AndiVS/broker-api/positionServer/internal/repository"
	"github.com/google/uuid"
)

// Positions interface for transactionService
type Positions interface {
	OpenPosition(c context.Context, name, time string, amount int64, price float32) (*uuid.UUID, error)
	ClosePosition(c context.Context, id uuid.UUID, price float32) error
}

// PositionsService struct for service
type PositionsService struct {
	Rep repository.Positions
}

// NewPositionService constructor
func NewPositionService(Rep interface{}) Positions {
	return &PositionsService{Rep: Rep.(*repository.Postgres)}
}

// OpenPosition add record about position
func (s *PositionsService) OpenPosition(c context.Context, name, time string, amount int64, price float32) (*uuid.UUID, error) {
	transaction := model.Position{PositionID: uuid.New(), CurrencyName: name, Amount: amount,
		OpenPrice: price, OpenTime: time}
	return s.Rep.OpenPosition(c, &transaction)
}

// ClosePosition add record about position
func (s *PositionsService) ClosePosition(c context.Context, id uuid.UUID, price float32) error {
	return s.Rep.ClosePosition(c, &id, &price)
}
