package main

import (
	"context"
	"github.com/AndiVS/broker-api/priceBuffer/protocolPrice"
	"github.com/AndiVS/broker-api/transactionBroker/protocolBroker"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"time"
)

func main() {

	connectionBuffer := connectToBuffer()
	go getPrices(connectionBuffer)
	connectionBroker := connectToBroker()
	buyCurrency(connectionBroker, "BTC", 64)

}

func connectToBroker() protocolBroker.TransactionServiceClient {
	addressGRPC := os.Getenv("GRPC_BROKER_ADDRESS")
	con, err := grpc.Dial(addressGRPC, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	return protocolBroker.NewTransactionServiceClient(con)
}

func connectToBuffer() protocolPrice.CurrencyServiceClient {
	//addressGrcp := os.Getenv("GRPC_BUFFER_ADDRESS")
	addressGrcp := "localhost:8081"
	con, err := grpc.Dial(addressGrcp, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	return protocolPrice.NewCurrencyServiceClient(con)
}

func buyCurrency(client protocolBroker.TransactionServiceClient, currency string, amount int64) {
	search, err := client.BuyCurrency(context.Background(), &protocolBroker.BuyRequest{CurrencyName: currency, CurrencyAmount: amount})
	if err != nil {
		log.Panicf("Error while buying currency: %v", err)
	}
	log.Printf("Transaction completed: %s", search.GetTransactionID())
}

func getPrices(client protocolPrice.CurrencyServiceClient) {
	notes := []*protocolPrice.GetPriceRequest{
		{Name: "BTC"},
	}
	stream, err := client.GetPrice(context.Background())
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

			log.Printf("Got currency data Name: %v Price: %v at time %v",
				in.Currency.CurrencyName, in.Currency.CurrencyPrice, in.Currency.Time)
		}
	}()
	for {
		for _, note := range notes {
			if err := stream.Send(note); err != nil {
				log.Fatalf("Failed to send a note: %v", err)
			}
		}
		time.Sleep(5 * time.Second)
	}
	stream.CloseSend()
	<-waitc
}
