// Package model for working with position
package model

import (
	"encoding/json"

	"github.com/google/uuid"
)

// Position model of position
type Position struct {
	PositionID   *uuid.UUID `param:"position_id" query:"position_id" header:"position_id" form:"position_id" bson:"position_id" msg:"position_id" json:"position_id"`
	CurrencyName string     `param:"currency_name" query:"currency_name" header:"currency_name" form:"currency_name" bson:"currency_name" msg:"currency_name" json:"currency_name"`
	Amount       *int64     `param:"amount" query:"amount" header:"amount" form:"amount" bson:"amount" msg:"amount" json:"amount"`
	OpenPrice    *float32   `param:"open_price" query:"open_price" header:"open_price" form:"open_price" bson:"open_price" msg:"open_price" json:"open_price"`
	OpenTime     string     `param:"open_time" query:"open_time" header:"open_time" form:"open_time" bson:"open_time" msg:"open_time" json:"open_time"`
	ClosePrice   *float32   `param:"close_price" query:"close_price" header:"close_price" form:"close_price" bson:"close_price" msg:"close_price" json:"close_price"`
	CloseTime    string     `param:"close_time" query:"close_time" header:"close_time" form:"close_time" bson:"close_time" msg:"close_time" json:"close_time"`
	StopLoss     *float32   `param:"stop_loss" query:"stop_loss" header:"stop_loss" form:"stop_loss" bson:"stop_loss" msg:"stop_loss" json:"stop_loss"`
	TakeProfit   *float32   `param:"take_profit" query:"take_profit" header:"take_profit" form:"take_profit" bson:"take_profit" msg:"take_profit" json:"take_profit"`
	Event        string     `json:"event"`
}

// MarshalBinary Marshal currency to byte
func (p *Position) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

// UnmarshalBinary Marshal currency to byte
func (p *Position) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, p)
}
