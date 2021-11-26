package main

import (
	"fmt"
	"github.com/AndiVS/broker-api/priceServer/internal/consumer"
	"github.com/AndiVS/broker-api/priceServer/internal/server"
	"github.com/AndiVS/broker-api/priceServer/model"
	"github.com/AndiVS/broker-api/priceServer/protocolPrice"
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
		"BTC": map[uuid.UUID]*chan *model.Currency{},
		"ETH": map[uuid.UUID]*chan *model.Currency{},
		"YFI": map[uuid.UUID]*chan *model.Currency{},
	}
	mute := new(sync.Mutex)
	go conToGrpc(mute, subscribersMap)
	conToRedis(mute, subscribersMap)
}

func conToRedis(mu *sync.Mutex, subscribersMap map[string]map[uuid.UUID]*chan *model.Currency) {
	adr := fmt.Sprintf("%s%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	//adr := fmt.Sprintf("%s:%s", "172.28.1.1", "6379")
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
	//listener, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create grpc server
	grpcServer := grpc.NewServer()
	protocolPrice.RegisterCurrencyServiceServer(grpcServer, server.NewCurrencyServer(mu, subscribersMap))

	log.Println("start server")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
