package main

import (
	"fmt"
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

	startPrice := 60000
	minRand := startPrice - 5000
	maxRand := startPrice + 5000

	rand.Seed(time.Now().UTC().UnixNano())
	price := rand.Intn(maxRand-minRand) + minRand

	_, err := client.XAdd(&redis.XAddArgs{
		Stream: "PriceGenerator",
		Values: map[string]interface{}{
			"Name":  "BTC",
			"price": price,
			"time":  time.Now(),
		},
	}).Result()
	if err != nil {
		log.Printf("err in add in stream %v", err)
	}
}
