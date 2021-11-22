package main

import (
	"context"
	"github.com/AndiVS/broker-api/priceBuffer/protocolPrice"
	"github.com/AndiVS/broker-api/transactionBroker/protocolBroker"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	connectionBroker := connectToBroker()
	connectionBuffer := connectToBuffer()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// todo balance check
	cur := protocolBroker.Currency{CurrencyID: uuid.New().String(), CurrencyName: "BTC", CurrencyPrice: 50000.000, Time: time.Now().String()}
	search, err := connectionBroker.BuyCurrency(ctx, &protocolBroker.BuyRequest{Currency: &cur, CurrencyAmount: 64})
	if err != nil {
		log.Panicf("Error while buying currency: %v", err)
	}
	log.Printf("Transaction completed: %s", search.GetTransactionID())

	getPrices(connectionBuffer)
}

func connectToBroker() protocolBroker.TransactionServiceClient {
	addressGrcp := os.Getenv("GRPC_BROKER_ADDRESS")
	con, err := grpc.Dial(addressGrcp, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	return protocolBroker.NewTransactionServiceClient(con)
}

func connectToBuffer() protocolPrice.CurrencyServiceClient {
	addressGrcp := os.Getenv("GRPC_BUFFER_ADDRESS")
	con, err := grpc.Dial(addressGrcp, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	return protocolPrice.NewCurrencyServiceClient(con)
}

func getPrices(client protocolPrice.CurrencyServiceClient) {
	notes := []*protocolPrice.GetPriceRequest{
		{Name: "BTC"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.GetPrice(ctx)
	if err != nil {
		log.Fatalf("%v.RouteChat(_) = _, %v", client, err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Got message %s at point(%d, %d)", in.Currency.Price, in.Currency.Name, in.Location.Longitude)
		}
	}()
	for _, note := range notes {
		if err := stream.Send(note); err != nil {
			log.Fatalf("Failed to send a note: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
}
