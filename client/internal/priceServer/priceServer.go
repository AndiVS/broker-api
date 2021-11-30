package priceServer

import (
	"context"
	"github.com/AndiVS/broker-api/priceServer/model"
	"github.com/AndiVS/broker-api/priceServer/priceProtocol"
	"io"
	"log"
	"sync"
)

// PriceServer struct for renew price
type PriceServer struct {
	subList     []string
	connection  priceProtocol.CurrencyServiceClient
	currencyMap *map[string]*model.Currency
	mutex       *sync.Mutex
}

func NewPriceServer(connection priceProtocol.CurrencyServiceClient, subList []string, currencyMap *map[string]*model.Currency, mutex *sync.Mutex) *PriceServer {
	return &PriceServer{
		subList:     subList,
		currencyMap: currencyMap,
		mutex:       mutex,
		connection:  connection,
	}
}

func (s *PriceServer) SubscribeToCurrency() {
	req := priceProtocol.GetPriceRequest{Name: s.subList}
	stream, err := s.connection.GetPrice(context.Background(), &req)
	if err != nil {
		log.Fatalf("sub err  %v", err)
	}
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Failed to receive a note : %v", err)
		}
		log.Printf("Got currency data Name: %v Price: %v at time %v",
			in.Currency.CurrencyName, in.Currency.CurrencyPrice, in.Currency.Time)
		s.mutex.Lock()
		(*s.currencyMap)[in.Currency.CurrencyName] = &model.Currency{CurrencyName: in.Currency.CurrencyName, CurrencyPrice: in.Currency.CurrencyPrice, Time: in.Currency.Time}
		s.mutex.Unlock()
	}
}
