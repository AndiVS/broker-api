// Package server package for grcp price buffer server
package server

import (
	"github.com/AndiVS/broker-api/priceServer/model"
	"github.com/AndiVS/broker-api/priceServer/priceProtocol"
	"github.com/google/uuid"
	"sync"
)

// GRCPServer for grpc
type GRCPServer struct {
	priceProtocol.UnimplementedCurrencyServiceServer
	mu             *sync.Mutex // protects currencyMap
	subscribersMap map[string]map[uuid.UUID]*chan *model.Currency
}

// NewCurrencyServer create object GRCPServer
func NewCurrencyServer(mu *sync.Mutex, subscribersMap map[string]map[uuid.UUID]*chan *model.Currency) *GRCPServer {
	return &GRCPServer{mu: mu, subscribersMap: subscribersMap}
}

// GetPrice method of price buffer server
func (s *GRCPServer) GetPrice(request *priceProtocol.GetPriceRequest, stream priceProtocol.CurrencyService_GetPriceServer) error {
	id := uuid.New()
	ac := make(chan *model.Currency)
	//	<-c
	for _, v := range request.Name {
		s.mu.Lock()
		s.subscribersMap[v][id] = &ac
		s.mu.Unlock()
	}

	for {
		cur := <-ac
		pcur := priceProtocol.Currency{CurrencyName: cur.CurrencyName, CurrencyPrice: cur.CurrencyPrice, Time: cur.Time}
		err := stream.Send(&priceProtocol.GetPriceResponse{Currency: &pcur})
		if err != nil {
			for _, v := range request.Name {
				//<-s.subscribersMap[v][id]
				<-ac
				delete(s.subscribersMap[v], id)
			}
			return err
		}
	}
}
