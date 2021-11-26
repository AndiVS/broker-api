package main

import (
	"fmt"
	"github.com/AndiVS/broker-api/priceBuffer/internal/consumer"
	"github.com/AndiVS/broker-api/priceBuffer/internal/server"
	"github.com/AndiVS/broker-api/priceBuffer/model"
	"github.com/AndiVS/broker-api/priceBuffer/protocolPrice"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"sync"
)

func main() {
	currencyMap := map[string]*model.Currency{}
	mute := new(sync.Mutex)
	go conToGrpc(mute, currencyMap)
	conToRedis(mute, currencyMap)
}

func conToRedis(mu *sync.Mutex, currencyMap map[string]*model.Currency) {
	adr := fmt.Sprintf("%s%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	//adr := fmt.Sprintf("%s:%s", "172.28.1.1", "6379")
	client := redis.NewClient(&redis.Options{
		Addr:     adr,
		Password: "",
		DB:       0, // use default DB
	})
	redisStream := consumer.NewRedisStream(client, mu, currencyMap)
	redisStream.RedisConsumer()

}

func conToGrpc(mu *sync.Mutex, currencyMap map[string]*model.Currency) {
	listener, err := net.Listen("tcp", os.Getenv("GRPC_BUFFER_PORT"))
	//listener, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create grpc server
	grpcServer := grpc.NewServer()
	protocolPrice.RegisterCurrencyServiceServer(grpcServer, server.NewCurrencyServer(mu, currencyMap))

	log.Println("start server")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
