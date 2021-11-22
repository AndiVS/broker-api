package main

import (
	"fmt"
	"github.com/AndiVS/broker-api/priceBuffer/protocolPrice"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"time"
)

func main() {
	adr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	client := redis.NewClient(&redis.Options{
		Addr:     adr,
		Password: "",
		DB:       0, // use default DB
	})
	currMap := generateCurrencyMap()
	for {
		_, err := client.XAdd(&redis.XAddArgs{
			Stream: "PriceGenerator",
			Values: map[string]interface{}{
				"CurrencyMap": currMap,
			},
		}).Result()
		if err != nil {
			log.Printf("err in add in stream %v", err)
		}
		time.Sleep(5 * time.Second)
		generatePrice(currMap)
	}
}

func generatePrice(currMap map[string]*protocolPrice.Currency) {
	for _, v := range currMap {
		rand.Seed(time.Now().UTC().UnixNano())
		a := rand.Float32() * 0.1
		b := float32(rand.Intn(2) - 1)
		v.CurrencyPrice *= a*b + 1
	}
}

func generateCurrencyMap() map[string]*protocolPrice.Currency {
	currMap := make(map[string]*protocolPrice.Currency)

	currMap["BTC"] = &protocolPrice.Currency{
		CurrencyName:  "BTC",
		CurrencyPrice: 55555.555,
		Time:          time.Now().String(),
	}

	return currMap
}
