// Package model
package model

import (
	"github.com/google/uuid"
)

// Transaction model of transaction
type Transaction struct {
	//UserID 			uuid.UUID `param:"user_id" query:"user_id" header:"user_id" form:"user_id" bson:"user_id" msg:"user_id" json:"user_id"`
	TransactionID   uuid.UUID `param:"transaction_id" query:"transaction_id" header:"transaction_id" form:"transaction_id" bson:"transaction_id" msg:"transaction_id" json:"transaction_id"`
	CurrencyName    string    `param:"currency_name" query:"currency_name" header:"currency_name" form:"currency_name" bson:"currency_name" msg:"currency_name" json:"currency_name"`
	Amount          int64     `param:"amount" query:"amount" header:"amount" form:"amount" bson:"amount" msg:"amount" json:"amount"`
	Price           float32   `param:"price" query:"price" header:"price" form:"price" bson:"price" msg:"price" json:"price"`
	TransactionTime string    `param:"transaction_time" query:"transaction_time" header:"transaction_time" form:"transaction_time" bson:"transaction_time" msg:"transaction_time" json:"transaction_time"`
}
