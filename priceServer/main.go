package main

import (
	"fmt"

	"github.com/AndiVS/broker-api/priceServer/internal/consumer"
	"github.com/AndiVS/broker-api/priceServer/internal/server"
	"github.com/AndiVS/broker-api/priceServer/model"
	"github.com/AndiVS/broker-api/priceServer/priceProtocol"
	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"net"
	"os"
	"sync"
)

func main() {
	subscribersMap := map[string]map[uuid.UUID]*chan *model.Currency{
		"BTC": {},
		"ETH": {},
		"YFI": {},
	}
	mute := new(sync.Mutex)
	go conToGrpc(mute, subscribersMap)
	conToRedis(mute, subscribersMap)
}

func conToRedis(mu *sync.Mutex, subscribersMap map[string]map[uuid.UUID]*chan *model.Currency) {
	adr := fmt.Sprintf("%s%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	client := redis.NewClient(&redis.Options{
		Addr:     adr,
		Password: "",
		DB:       0, // use default DB
	})
	redisStream := consumer.NewRedisStream(client, mu, subscribersMap)
	redisStream.RedisConsumer()
}

func conToGrpc(mu *sync.Mutex, subscribersMap map[string]map[uuid.UUID]*chan *model.Currency) {
	listener, err := net.Listen("tcp", os.Getenv("GRPC_BUFFER_PORT"))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	priceProtocol.RegisterCurrencyServiceServer(grpcServer, server.NewCurrencyServer(mu, subscribersMap))

	log.Println("start server")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
