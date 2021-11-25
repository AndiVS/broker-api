// Package model
package model

type Currency struct {
	CurrencyName  string  `protobuf:"bytes,1,opt,name=currencyName,proto3" json:"currencyName,omitempty"`
	CurrencyPrice float32 `protobuf:"fixed32,2,opt,name=currencyPrice,proto3" json:"currencyPrice,omitempty"`
	Time          string  `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
}
