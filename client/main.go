package main

import (
	"transactionBroker/serverBroker/internal/protocolBroker"
	"context"
	"log"
	"time"
)

func main()  {
	con, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	c := protocolBroker.NewUserServiceClient(cc2)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	serch, err := c.SearchUser(ctx, &protocol.SearchUserRequest{Username: "admin"})
	if err != nil {
		log.Panicf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", serch.GetUser())
}