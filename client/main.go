package main

import (
	"context"
	"io"
	"log"
	"sync"
	"time"

	"github.com/AndiVS/broker-api/positionServer/positionProtocol"
	"github.com/AndiVS/broker-api/priceServer/model"
	"github.com/AndiVS/broker-api/priceServer/priceProtocol"
	"google.golang.org/grpc"
)

// PriceServer struct for renew price
type PriceServer struct {
	subList     []string
	connection  priceProtocol.CurrencyServiceClient
	currencyMap *map[string]*model.Currency
	mutex       *sync.Mutex
}

// PositionServer struct for opening and closing a position
type PositionServer struct {
	connection  positionProtocol.PositionServiceClient
	currencyMap *map[string]*model.Currency
	mutex       *sync.Mutex
	positionMap map[string]map[string]bool
}

func (s *PositionServer) createPositionMap(sublist []string) {
	s.positionMap = make(map[string]map[string]bool)
	for _, v := range sublist {
		s.positionMap[v] = map[string]bool{}
	}
}

func main() {
	mute := new(sync.Mutex)
	subList := []string{"BTC", "ETH", "YFI"}
	curMap := make(map[string]*model.Currency)

	priceServ := PriceServer{
		subList:     subList,
		currencyMap: &curMap,
		mutex:       mute,
	}
	priceServ.connectToPriceServer()

	posServ := PositionServer{
		currencyMap: &curMap,
		mutex:       mute,
	}
	posServ.connectToPositionServer()
	posServ.createPositionMap(subList)

	go priceServ.subscribeToCurrency()

	time.Sleep(5 * time.Second)

	posServ.openPosition("BTC", 64)
	id := posServ.openPosition("BTC", 64)

	time.Sleep(5 * time.Second)

	posServ.closePosition(id, "BTC")
}

func (s *PositionServer) connectToPositionServer() {
	// addressGRPC := os.Getenv("GRPC_BROKER_ADDRESS")
	addressGrcp := "172.28.1.8:8083"
	con, err := grpc.Dial(addressGrcp, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	s.connection = positionProtocol.NewPositionServiceClient(con)
}

func (s *PriceServer) connectToPriceServer() {
	// addressGrcp := os.Getenv("GRPC_BUFFER_ADDRESS")
	addressGrcp := "172.28.1.9:8081"
	con, err := grpc.Dial(addressGrcp, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	s.connection = priceProtocol.NewCurrencyServiceClient(con)
}

func (s *PositionServer) openPosition(currency string, amount int64) string {
	open, err := s.connection.OpenPosition(context.Background(),
		&positionProtocol.OpenRequest{CurrencyName: currency, CurrencyAmount: amount, Price: (*s.currencyMap)[currency].CurrencyPrice})
	if err != nil {
		log.Printf("Error while opening position: %v", err)
	}
	s.positionMap[currency][open.GetPositionID()] = false
	log.Printf("Position open with id: %s", open.GetPositionID())
	return open.GetPositionID()
}

func (s *PositionServer) closePosition(id, currency string) {
	_, err := s.connection.ClosePosition(context.Background(), &positionProtocol.CloseRequest{PositionID: id, CurrencyName: currency})
	if err != nil {
		log.Printf("Error while closing position: %v", err)
	} else {
		s.positionMap[currency][id] = true
		log.Printf("Position with id: %s closed", id)
	}
}

func (s *PriceServer) subscribeToCurrency() {
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
