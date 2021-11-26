// Package server package for grcp price buffer server
package server

import (
	"github.com/AndiVS/broker-api/priceBuffer/model"
	"github.com/AndiVS/broker-api/priceBuffer/protocolPrice"
	"sync"
)

// GRCPServer for grpc
type GRCPServer struct {
	protocolPrice.UnimplementedCurrencyServiceServer
	mu             *sync.Mutex // protects currencyMap
	subscribersMap map[string][]*chan *model.Currency
}

// NewCurrencyServer create object GRCPServer
func NewCurrencyServer(mu *sync.Mutex, subscribersMap map[string][]*chan *model.Currency) *GRCPServer {
	return &GRCPServer{mu: mu, subscribersMap: subscribersMap}
}

// GetPrice method of price buffer server
func (s *GRCPServer) GetPrice(request *protocolPrice.GetPriceRequest, stream protocolPrice.CurrencyService_GetPriceServer) error {
	key := request.Name
	c := make(chan *model.Currency)
	s.mu.Lock()
	s.subscribersMap[key] = append(s.subscribersMap[key], &c)
	s.mu.Unlock()
	for {
		cur := <-c
		pcur := protocolPrice.Currency{CurrencyName: cur.CurrencyName, CurrencyPrice: cur.CurrencyPrice, Time: cur.Time}
		err := stream.Send(&protocolPrice.GetPriceResponse{Currency: &pcur})
		if err != nil {
			return err
		}
	}
}
