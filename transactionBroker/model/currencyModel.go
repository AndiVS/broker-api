package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

// Currency struct for currency
type Currency struct {
	ID    uuid.UUID `param:"_id" query:"_id" header:"_id" form:"_id" json:"_id" xml:"_id" bson:"_id" msg:"_id"`
	Name  string    `param:"currency" query:"currency" header:"currency" form:"currency" bson:"currency" msg:"currency"`
	Price float32   `param:"price" query:"price" header:"price" form:"price" bson:"price" msg:"price"`
	Time  time.Time `param:"time" query:"time" header:"time" form:"time" bson:"time" msg:"time"`
}

// MarshalBinary Marshal currency to byte
func (c *Currency) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

// UnmarshalBinary Marshal currency to byte
func (c *Currency) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}
