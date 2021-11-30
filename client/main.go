package main

import (
	"sync"
	"time"

	"github.com/AndiVS/broker-api/client/internal/positionServer"
	"github.com/AndiVS/broker-api/client/internal/priceServer"
	"github.com/AndiVS/broker-api/priceServer/model"
)

func main() {
	mute := new(sync.Mutex)
	subList := []string{"BTC", "ETH", "YFI"}
	curMap := make(map[string]*model.Currency)

	priceServ := priceServer.NewPriceServer(subList, curMap, mute)

	posServ := positionServer.NewPositionServer(subList, curMap, mute)

	go priceServ.SubscribeToCurrency()

	time.Sleep(5 * time.Second)

	posServ.OpenPosition("BTC", 64)
	id := posServ.OpenPosition("BTC", 64)

	time.Sleep(5 * time.Second)

	posServ.ClosePosition(id, "BTC")
}
