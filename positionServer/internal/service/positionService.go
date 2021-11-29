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
	OpenPosition(c context.Context, position *model.Position) (*uuid.UUID, error)
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
func (s *PositionsService) OpenPosition(c context.Context, position *model.Position) (*uuid.UUID, error) {
	return s.Rep.OpenPosition(c, position)
}

// ClosePosition add record about position
func (s *PositionsService) ClosePosition(c context.Context, id uuid.UUID, price float32) error {
	return s.Rep.ClosePosition(c, &id, &price)
}
