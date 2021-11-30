// Package positionserver for working with position server
package positionserver

import (
	"context"
	"log"
	"sync"

	"github.com/AndiVS/broker-api/positionServer/positionProtocol"
	"github.com/AndiVS/broker-api/priceServer/model"
	"google.golang.org/grpc"
)

// PositionServer struct for opening and closing a position
type PositionServer struct {
	connection  positionProtocol.PositionServiceClient
	currencyMap map[string]*model.Currency
	mutex       *sync.Mutex
	positionMap map[string]map[string]bool
}

// NewPositionServer constructor
func NewPositionServer(subList []string, currencyMap map[string]*model.Currency, mutex *sync.Mutex) *PositionServer {
	return &PositionServer{
		currencyMap: currencyMap,
		mutex:       mutex,
		connection:  connectToPositionServer(),
		positionMap: createPositionMap(subList),
	}
}

func connectToPositionServer() positionProtocol.PositionServiceClient {
	// addressGRPC := os.Getenv("GRPC_BROKER_ADDRESS")
	addressGrcp := "172.28.1.8:8083"
	con, err := grpc.Dial(addressGrcp, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	return positionProtocol.NewPositionServiceClient(con)
}

func createPositionMap(sublist []string) map[string]map[string]bool {
	positionMap := make(map[string]map[string]bool)
	for _, v := range sublist {
		positionMap[v] = map[string]bool{}
	}
	return positionMap
}

// OpenPosition method that allow open positions
func (s *PositionServer) OpenPosition(currency string, amount int64) string {
	open, err := s.connection.OpenPosition(context.Background(),
		&positionProtocol.OpenRequest{CurrencyName: currency, CurrencyAmount: amount, Price: s.currencyMap[currency].CurrencyPrice})
	if err != nil {
		log.Printf("Error while opening position: %v", err)
	}
	s.positionMap[currency][open.GetPositionID()] = false
	log.Printf("Position open with id: %s", open.GetPositionID())
	return open.GetPositionID()
}

// ClosePosition method that allow close positions
func (s *PositionServer) ClosePosition(id, currency string) {
	_, err := s.connection.ClosePosition(context.Background(), &positionProtocol.CloseRequest{PositionID: id, CurrencyName: currency})
	if err != nil {
		log.Printf("Error while closing position: %v", err)
	} else {
		s.positionMap[currency][id] = true
		log.Printf("Position with id: %s closed", id)
	}
}
