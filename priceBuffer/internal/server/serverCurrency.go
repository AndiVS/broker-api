package server

import (
	"github.com/AndiVS/broker-api/priceBuffer/protocolPrice"
	"io"
	"sync"
)

// GRCPServer for grpc
type GRCPServer struct {
	protocolPrice.UnimplementedCurrencyServiceServer

	mu          *sync.Mutex // protects currencyMap
	currencyMap map[string]protocolPrice.Currency
}

func NewCurrencyServer(mu *sync.Mutex, currencyMap map[string]protocolPrice.Currency) *GRCPServer {
	return &GRCPServer{mu: mu, currencyMap: currencyMap}
}

func (s *GRCPServer) GetPrice(stream protocolPrice.CurrencyService_GetPriceServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		key := in.Name

		s.mu.Lock()
		resp := s.currencyMap[key]
		s.mu.Unlock()

		err = stream.Send(&protocolPrice.GetPriceResponse{Currency: &resp})
		if err != nil {
			return err
		}
	}
}
