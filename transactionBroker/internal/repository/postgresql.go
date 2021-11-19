// Package repository
package repository

import (
	"context"
	"errors"
	"github.com/AndiVS/broker-api/transactionBroker/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

var (
	// ErrNotFound means entity is not found in repository
	ErrNotFound = errors.New("not found")
)

// Postgres struct for Pool
type Postgres struct {
	Pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) Transactions {
	return &Postgres{Pool: pool}
}

// Transactions used for structuring, function for working with records
type Transactions interface {
	InsertTransaction(c context.Context, transaction *model.Transaction) (*uuid.UUID, error)
	SelectTransaction(c context.Context, id *uuid.UUID) (*model.Transaction, error)
	SelectAllTransactions(c context.Context) ([]*model.Transaction, error)
	SelectAllTransactionWithCurrency(c context.Context, id *uuid.UUID) ([]*model.Transaction, error)
	UpdateTransaction(c context.Context, transaction *model.Transaction) error
	DeleteTransaction(c context.Context, id *uuid.UUID) error
	DeleteALLTransactions(c context.Context, id *uuid.UUID) error
}

// InsertTransaction function for inserting item from a table
func (repos *Postgres) InsertTransaction(c context.Context, transaction *model.Transaction) (*uuid.UUID, error) {
	row := repos.Pool.QueryRow(c,
		"INSERT INTO transactions (tid, cid, amount, price, time) VALUES ($1, $2, $3, $4, $5) RETURNING tid",
		transaction.TransactionID, transaction.CurrencyID, transaction.Amount, transaction.Price, transaction.Time)

	err := row.Scan(&transaction.TransactionID)
	if err != nil {
		log.Errorf("Unable to INSERT: %v", err)
		return &transaction.TransactionID, err
	}

	return &transaction.TransactionID, err
}

// SelectTransaction function for selecting item from a table
func (repos *Postgres) SelectTransaction(c context.Context, id *uuid.UUID) (*model.Transaction, error) {
	var transaction model.Transaction
	row := repos.Pool.QueryRow(c,
		"SELECT tid, cid, amount, price, time FROM transactions WHERE tid = $1", id)

	err := row.Scan(&transaction.TransactionID, &transaction.CurrencyID, &transaction.Amount, &transaction.Price, &transaction.Time)
	if errors.Is(err, pgx.ErrNoRows) {
		log.Errorf("Not found : %s\n", err)
		return &transaction, ErrNotFound
	} else if err != nil {
		return &transaction, err
	}

	log.Printf("sec")

	return &transaction, err
}

// SelectAllTransactions function for selecting items from a table
func (repos *Postgres) SelectAllTransactions(c context.Context) ([]*model.Transaction, error) {
	var transactions []*model.Transaction

	row, err := repos.Pool.Query(c,
		"SELECT tid, cid, amount, price, time FROM transactions")

	for row.Next() {
		var rc model.Transaction
		err := row.Scan(&rc.TransactionID, &rc.CurrencyID, &rc.Amount, &rc.Price, &rc.Time)
		if err == pgx.ErrNoRows {
			return transactions, err
		}
		transactions = append(transactions, &rc)
	}

	return transactions, err
}

// SelectAllTransactionWithCurrency function for selecting items from a table
func (repos *Postgres) SelectAllTransactionWithCurrency(c context.Context, id *uuid.UUID) ([]*model.Transaction, error) {
	var transactions []*model.Transaction

	row, err := repos.Pool.Query(c,
		"SELECT tid, cid, amount, price, time FROM transactions WHERE cid = $1", id)

	for row.Next() {
		var rc model.Transaction
		err := row.Scan(&rc.TransactionID, &rc.CurrencyID, &rc.Amount, &rc.Price, &rc.Time)
		if err == pgx.ErrNoRows {
			return transactions, err
		}
		transactions = append(transactions, &rc)
	}

	return transactions, err
}

// UpdateTransaction function for updating item from a table
func (repos *Postgres) UpdateTransaction(c context.Context, transaction *model.Transaction) error {
	_, err := repos.Pool.Exec(c,
		"UPDATE transactions SET cid = $2, amount = $3, price = $4, time = $5 WHERE tid = $1",
		transaction.TransactionID, transaction.CurrencyID, transaction.Amount, transaction.Price, transaction.Time)

	if err != nil {
		log.Errorf("Failed updating data in db: %s\n", err)
		return err
	}

	return nil
}

// DeleteTransaction function for deleting item from a table
func (repos *Postgres) DeleteTransaction(c context.Context, id *uuid.UUID) error {
	_, err := repos.Pool.Exec(c, "DELETE FROM transactions WHERE tid = $1", id)

	if err != nil {
		return err
	}

	return nil
}

// DeleteALLTransactions function for deleting item from a table
func (repos *Postgres) DeleteALLTransactions(c context.Context, id *uuid.UUID) error {
	_, err := repos.Pool.Exec(c, "DELETE FROM transactions WHERE cid = $1", id)

	if err != nil {
		return err
	}

	return nil
}
