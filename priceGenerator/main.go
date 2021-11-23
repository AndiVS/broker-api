package main

import (
	"fmt"
	"github.com/AndiVS/broker-api/priceBuffer/protocolPrice"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

const timeFormat = "2006-01-02 15:04:05.000000000"

func main() {
	//adr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	adr := fmt.Sprintf("%s:%s", "172.28.1.1", "6379")
	client := redis.NewClient(&redis.Options{
		Addr:     adr,
		Password: "",
		DB:       0, // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
	}
	currMap := generateCurrencyMap()
	for {
		for _, v := range currMap {
			_, err := client.XAdd(&redis.XAddArgs{
				Stream: "PriceGenerator",
				Values: map[string]interface{}{
					"Currency": v,
				},
			}).Result()
			if err != nil {
				log.Printf("err in add in stream %v", err)
			}
		}
		time.Sleep(5 * time.Second)
		generatePrice(currMap)
	}
}

func generatePrice(currMap map[string]*protocolPrice.Currency) {
	for _, v := range currMap {
		rand.Seed(time.Now().UTC().UnixNano())
		a := rand.Float32() * 0.01
		b := float32(rand.Intn(3) - 1)
		v.CurrencyPrice *= a*b + 1
		v.Time = time.Now().Format(timeFormat)
	}
}

func generateCurrencyMap() map[string]*protocolPrice.Currency {
	currMap := make(map[string]*protocolPrice.Currency)

	currMap["BTC"] = &protocolPrice.Currency{
		CurrencyName:  "BTC",
		CurrencyPrice: 55555.555,
		Time:          time.Now().Format(timeFormat),
	}
	/*	currMap["BTC1"] = &protocolPrice.Currency{
			CurrencyName:  "BTC1",
			CurrencyPrice: 55555.555,
			Time:          time.Now().Format(timeFormat),
		}
		currMap["BTC2"] = &protocolPrice.Currency{
			CurrencyName:  "BTC2",
			CurrencyPrice: 55555.555,
			Time:          time.Now().Format(timeFormat),
		}
		currMap["BTC3"] = &protocolPrice.Currency{
			CurrencyName:  "BTC3",
			CurrencyPrice: 55555.555,
			Time:          time.Now().Format(timeFormat),
		}
		currMap["BTC4"] = &protocolPrice.Currency{
			CurrencyName:  "BTC4",
			CurrencyPrice: 55555.555,
			Time:          time.Now().Format(timeFormat),
		}*/

	return currMap
}
