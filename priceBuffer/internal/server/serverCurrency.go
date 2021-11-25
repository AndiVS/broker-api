// Package server package for grcp price buffer server
package server

import (
	"github.com/AndiVS/broker-api/priceBuffer/model"
	"github.com/AndiVS/broker-api/priceBuffer/protocolPrice"
	"sync"
	"time"
)

// GRCPServer for grpc
type GRCPServer struct {
	protocolPrice.UnimplementedCurrencyServiceServer
	mu          *sync.Mutex // protects currencyMap
	currencyMap map[string]model.Currency
}

// NewCurrencyServer create object GRCPServer
func NewCurrencyServer(mu *sync.Mutex, currencyMap map[string]model.Currency) *GRCPServer {
	return &GRCPServer{mu: mu, currencyMap: currencyMap}
}

// GetPrice method of price buffer server
func (s *GRCPServer) GetPrice(request *protocolPrice.GetPriceRequest, stream protocolPrice.CurrencyService_GetPriceServer) error {
	key := request.Name
	for {
		time.Sleep(5 * time.Second)
		s.mu.Lock()
		resp := s.currencyMap[key]
		s.mu.Unlock()
		cur := protocolPrice.Currency{CurrencyName: resp.CurrencyName, CurrencyPrice: resp.CurrencyPrice, Time: resp.Time}
		err := stream.Send(&protocolPrice.GetPriceResponse{Currency: &cur})
		if err != nil {
			return err
		}
	}
}
