package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/AndiVS/broker-api/priceServer/model"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
)

const timeFormat = "2006-01-02 15:04:05.000000000"

func main() {
	client := connectToRedis()
	currMap := generateCurrencyMap()

	for {
		_, err := client.XAdd(&redis.XAddArgs{
			Stream: "PriceGenerator",
			Values: currMap,
		}).Result()
		if err != nil {
			log.Printf("err in add in stream %v", err)
		}
		time.Sleep(5 * time.Second)
		generatePrice(currMap)
	}
}

func connectToRedis() *redis.Client {
	adr := fmt.Sprintf("%s%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	client := redis.NewClient(&redis.Options{
		Addr:     adr,
		Password: "",
		DB:       0, // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Printf("err in ping redis conn %v", err)
	}
	return client
}

func generatePrice(currMap map[string]interface{}) {
	for _, v := range currMap {
		rand.Seed(time.Now().UTC().UnixNano())
		a := rand.Float32() * 0.01
		b := float32(rand.Intn(3) - 1)
		v.(*model.Currency).CurrencyPrice *= a*b + 1
		v.(*model.Currency).Time = time.Now().Format(timeFormat)
	}
}

func generateCurrencyMap() map[string]interface{} {
	currMap := make(map[string]interface{})

	currMap["BTC"] = &model.Currency{
		CurrencyName:  "BTC",
		CurrencyPrice: 56000.555,
		Time:          time.Now().Format(timeFormat),
	}
	currMap["ETH"] = &model.Currency{
		CurrencyName:  "ETH",
		CurrencyPrice: 4300.555,
		Time:          time.Now().Format(timeFormat),
	}
	currMap["YFI"] = &model.Currency{
		CurrencyName:  "YFI",
		CurrencyPrice: 30000.000,
		Time:          time.Now().Format(timeFormat),
	}
	return currMap
}
