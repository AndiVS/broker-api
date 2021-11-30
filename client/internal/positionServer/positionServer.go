package positionServer

import (
	"context"
	"github.com/AndiVS/broker-api/positionServer/positionProtocol"
	"github.com/AndiVS/broker-api/priceServer/model"
	"log"
	"sync"
)

// PositionServer struct for opening and closing a position
type PositionServer struct {
	connection  positionProtocol.PositionServiceClient
	currencyMap *map[string]*model.Currency
	mutex       *sync.Mutex
	positionMap map[string]map[string]bool
}

func NewPositionServer(connection positionProtocol.PositionServiceClient, subList []string, currencyMap *map[string]*model.Currency, mutex *sync.Mutex) *PositionServer {
	return &PositionServer{
		currencyMap: currencyMap,
		mutex:       mutex,
		connection:  connection,
		positionMap: createPositionMap(subList),
	}
}

func createPositionMap(sublist []string) map[string]map[string]bool {
	positionMap := make(map[string]map[string]bool)
	for _, v := range sublist {
		positionMap[v] = map[string]bool{}
	}
	return positionMap
}

func (s *PositionServer) OpenPosition(currency string, amount int64) string {
	open, err := s.connection.OpenPosition(context.Background(),
		&positionProtocol.OpenRequest{CurrencyName: currency, CurrencyAmount: amount, Price: (*s.currencyMap)[currency].CurrencyPrice})
	if err != nil {
		log.Printf("Error while opening position: %v", err)
	}
	s.positionMap[currency][open.GetPositionID()] = false
	log.Printf("Position open with id: %s", open.GetPositionID())
	return open.GetPositionID()
}

func (s *PositionServer) ClosePosition(id, currency string) {
	_, err := s.connection.ClosePosition(context.Background(), &positionProtocol.CloseRequest{PositionID: id, CurrencyName: currency})
	if err != nil {
		log.Printf("Error while closing position: %v", err)
	} else {
		s.positionMap[currency][id] = true
		log.Printf("Position with id: %s closed", id)
	}
}