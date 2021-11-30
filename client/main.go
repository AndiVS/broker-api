package main

import (
	"log"
	"sync"
	"time"

	"github.com/AndiVS/broker-api/client/internal/positionServer"
	"github.com/AndiVS/broker-api/client/internal/priceServer"
	"github.com/AndiVS/broker-api/positionServer/positionProtocol"
	"github.com/AndiVS/broker-api/priceServer/model"
	"github.com/AndiVS/broker-api/priceServer/priceProtocol"
	"google.golang.org/grpc"
)

func main() {
	mute := new(sync.Mutex)
	subList := []string{"BTC", "ETH", "YFI"}
	curMap := make(map[string]*model.Currency)

	conToPriceServer := connectToPriceServer()
	priceServ := priceServer.NewPriceServer(conToPriceServer, subList, &curMap, mute)

	conToPositionServer := connectToPositionServer()
	posServ := positionServer.NewPositionServer(conToPositionServer, subList, &curMap, mute)

	go priceServ.SubscribeToCurrency()

	time.Sleep(5 * time.Second)

	posServ.OpenPosition("BTC", 64)
	id := posServ.OpenPosition("BTC", 64)

	time.Sleep(5 * time.Second)

	posServ.ClosePosition(id, "BTC")
}

func connectToPriceServer() priceProtocol.CurrencyServiceClient {
	// addressGrcp := os.Getenv("GRPC_BUFFER_ADDRESS")
	addressGrcp := "172.28.1.9:8081"
	con, err := grpc.Dial(addressGrcp, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	return priceProtocol.NewCurrencyServiceClient(con)
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
