package main

import (
	"sync"
	"time"

	"github.com/AndiVS/broker-api/client/internal/positionserver"
	"github.com/AndiVS/broker-api/client/internal/priceserver"
	"github.com/AndiVS/broker-api/priceServer/model"
)

func main() {
	mute := new(sync.Mutex)
	subList := []string{"BTC", "ETH", "YFI"}
	curMap := make(map[string]*model.Currency)

	priceServ := priceserver.Newpositionserver(subList, &curMap, mute)

	posServ := positionserver.Newpositionserver(subList, &curMap, mute)

	go priceServ.SubscribeToCurrency()

	time.Sleep(5 * time.Second)

	posServ.OpenPosition("BTC", 64)
	id := posServ.OpenPosition("BTC", 64)

	time.Sleep(5 * time.Second)

	posServ.ClosePosition(id, "BTC")
}
