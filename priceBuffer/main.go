package main

import (
	"fmt"
	"github.com/AndiVS/broker-api/priceBuffer/internal/consumer"
	"github.com/AndiVS/broker-api/priceBuffer/internal/server"
	"github.com/AndiVS/broker-api/priceBuffer/protocolPrice"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {

	clientRedis := conToRedis()
	currencyMap := new(map[string]*protocolPrice.Currency)
	go consumer.RedisConsumer(clientRedis, *currencyMap)
	//conToGrpc(*currencyMap)
}

func conToRedis() *redis.Client {
	adr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	client := redis.NewClient(&redis.Options{
		Addr:     adr,
		Password: "",
		DB:       0, // use default DB
	})
	return client
}

func conToGrpc(currencyMap map[string]*protocolPrice.Currency) {
	listener, err := net.Listen("tcp", os.Getenv("GRCP_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	grpcServer := grpc.NewServer()
	protocolPrice.RegisterCurrencyServiceServer(grpcServer, server.NewCurrencyServer(currencyMap))

	log.Println("start server")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
