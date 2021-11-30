// Package repository for working with postgres
package repository

import (
	"context"
	"errors"
	"time"

	"github.com/AndiVS/broker-api/positionServer/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

const timeFormat = "2006-01-02 15:04:05.000000000"

var (
	// ErrNotFound means entity is not found in repository
	ErrNotFound = errors.New("not found")
)

// Postgres struct for Pool
type Postgres struct {
	Pool *pgxpool.Pool
}

// NewRepository constructor
func NewRepository(pool *pgxpool.Pool) Positions {
	return &Postgres{Pool: pool}
}

// Positions used for structuring, function for working with records
type Positions interface {
	OpenPosition(c context.Context, position *model.Position) (*uuid.UUID, error)
	ClosePosition(c context.Context, id *uuid.UUID, closePrice *float32) error
}

// OpenPosition function for inserting item into table
func (repos *Postgres) OpenPosition(c context.Context, position *model.Position) (*uuid.UUID, error) {
	row := repos.Pool.QueryRow(c,
		"INSERT INTO positions (position_id, currency_name, amount, open_price, open_time, take_profit, stop_loss) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING position_id",
		position.PositionID, position.CurrencyName, position.Amount, position.OpenPrice, position.OpenTime, position.TakeProfit, position.StopLoss)

	err := row.Scan(&position.PositionID)
	if err != nil {
		log.Errorf("Unable to INSERT: %v", err)
		return position.PositionID, err
	}

	return position.PositionID, err
}

// ClosePosition function for updating item into table
func (repos *Postgres) ClosePosition(c context.Context, id *uuid.UUID, closePrice *float32) error {
	v, err := repos.Pool.Exec(c,
		"UPDATE positions SET close_price = $2, close_time = $3 WHERE position_id = $1",
		id, closePrice, time.Now().Format(timeFormat))

	if v.RowsAffected() == 0 {
		log.Errorf("Now sach row %v", ErrNotFound)
		return ErrNotFound
	}
	if err != nil {
		log.Errorf("Failed updating data in db: %s\n", err)
		return err
	}

	return nil
}
