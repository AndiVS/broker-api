package main

import (
	"context"
	"fmt"
	"github.com/AndiVS/broker-api/priceBuffer/internal/model"
	"github.com/AndiVS/broker-api/priceBuffer/internal/server"
	"github.com/AndiVS/broker-api/priceBuffer/protocolPrice"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"net"
	"os"
)

func main() {

	clientRedis := conToRedis()
	connGrpc := conToGrpc()
	go redisConsumer(clientRedis, connGrpc)

	lis, err := net.Listen("tcp", ":50005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	grpcServer := grpc.NewServer()
	protocol.RegisterCurrencyServiceServer(grpcServer, server.CurrencyServer)

	log.Println("start server")
	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

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

func conToGrpc() *grpc.ClientConn {
	conn, err := grpc.Dial(":50005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	return conn
}

func redisConsumer(client *redis.Client, connGrpc *grpc.ClientConn) {
	for {
		streams, err := client.XRead(&redis.XReadArgs{
			Streams: []string{"PriceGenerator", "$"},
		}).Result()

		if err != nil {
			log.Printf("err on consume events: %+v\n", err)
		}

		stream := streams[0].Messages[0]
		processRedisStream(stream, connGrpc)
	}
}

func processRedisStream(message redis.XMessage, connGrpc *grpc.ClientConn) {

	// create stream
	client := protocol.NewCurrencyServiceClient(connGrpc)
	in := &protocol.GetPriceRequest{Name: "sad"}
	stream, err := client.GetPrice(context.Background(), in)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("Resp received: %s", resp.Result)
		}
	}()

	<-done //we will wait until all response is received
	log.Printf("finished")

}
