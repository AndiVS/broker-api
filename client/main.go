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
