// Package model used in api
package model

import "encoding/json"

// Currency model for sending by redis
type Currency struct {
	CurrencyName  string  `protobuf:"bytes,1,opt,name=currencyName,proto3" json:"currencyName,omitempty"`
	CurrencyPrice float32 `protobuf:"fixed32,2,opt,name=currencyPrice,proto3" json:"currencyPrice,omitempty"`
	Time          string  `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
}

// MarshalBinary Marshal currency to byte
func (c *Currency) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

// UnmarshalBinary Marshal currency to byte
func (c *Currency) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}
