package server

import (
	"github.com/AndiVS/broker-api/priceBuffer/protocolPrice"
	"github.com/google/uuid"
	"io"
	"sync"
)

// CurrencyServer for grpc
type CurrencyServer struct {
	protocolPrice.UnimplementedCurrencyServiceServer

	mu          sync.Mutex // protects currencyMap
	currencyMap map[uuid.UUID]*protocolPrice.Currency
}

func NewCurrencyServer(currencyMap map[uuid.UUID]*protocolPrice.Currency) *CurrencyServer {
	s := &CurrencyServer{currencyMap: currencyMap}
	return s
}

func (s *CurrencyServer) GetPrice(stream protocolPrice.CurrencyService_GetPriceServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		key := in.Name
		uKey, err := uuid.Parse(key)

		s.mu.Lock()
		resp := s.currencyMap[uKey]
		s.mu.Unlock()

		err = stream.Send(&protocolPrice.GetPriceResponse{Currency: resp})
		if err != nil {
			return err
		}
	}
}
