// Package model
package model

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	TransactionID uuid.UUID `param:"tid" query:"tid" header:"tid" form:"tid" bson:"tid" msg:"tid"`
	CurrencyName  string    `param:"cname" query:"cname" header:"cname" form:"cname" bson:"cname" msg:"cname"`
	Amount        int64     `param:"amount" query:"amount" header:"amount" form:"amount" bson:"amount" msg:"amount"`
	Price         float32   `param:"price" query:"price" header:"price" form:"price" bson:"price" msg:"price"`
	Time          time.Time `param:"time" query:"time" header:"time" form:"time" bson:"time" msg:"time"`
}
