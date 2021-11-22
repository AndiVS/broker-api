package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
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

	price := generatePrice(55000)
	for {
		_, err := client.XAdd(&redis.XAddArgs{
			Stream: "PriceGenerator",
			Values: map[string]interface{}{
				"CurrID": uuid.New(),
				"Name":   "BTC",
				"Price":  price,
				"Time":   time.Now(),
			},
		}).Result()
		if err != nil {
			log.Printf("err in add in stream %v", err)
		}
		time.Sleep(5 * time.Second)
		price = generatePrice(price)
	}
}

func generatePrice(currentPrice int) int {
	minRand := currentPrice - 5000
	maxRand := currentPrice + 5000

	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(maxRand-minRand) + minRand
}
