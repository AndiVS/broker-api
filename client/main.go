package main

import (
	"context"
	"github.com/AndiVS/broker-api/positionServer/protocolPosition"
	"github.com/AndiVS/broker-api/priceServer/protocolPrice"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	connectionPriceServer := connectToPriceServer()

	subList := []string{"BTC", "ETH", "YFI"}
	go subscribeToCurrency(subList, connectionPriceServer)
	//unsubscribeFromCurrency("ETH",subMap)
	connectionPositionServer := connectToPositionServer()
	OpenPosition(connectionPositionServer, "BTC", 64)

}

func connectToPositionServer() protocolPosition.PositionServiceClient {
	//addressGRPC := os.Getenv("GRPC_BROKER_ADDRESS")
	addressGrcp := "localhost:8080"
	con, err := grpc.Dial(addressGrcp, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	return protocolPosition.NewPositionServiceClient(con)
}

func connectToPriceServer() protocolPrice.CurrencyServiceClient {
	//addressGrcp := os.Getenv("GRPC_BUFFER_ADDRESS")
	addressGrcp := "172.28.1.9:8081"
	//addressGrcp := "localhost:8081"
	con, err := grpc.Dial(addressGrcp, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	return protocolPrice.NewCurrencyServiceClient(con)
}

func OpenPosition(client protocolPosition.PositionServiceClient, currency string, amount int64) {
	search, err := client.OpenPosition(context.Background(), &protocolPosition.OpenRequest{CurrencyName: currency, CurrencyAmount: amount})
	if err != nil {
		log.Panicf("Error while buying currency: %v", err)
	}
	log.Printf("Transaction completed: %s", search.GetPositionID())
}

func subscribeToCurrency(subList []string, client protocolPrice.CurrencyServiceClient) {
	req := protocolPrice.GetPriceRequest{Name: subList}
	stream, err := client.GetPrice(context.Background(), &req)
	if err != nil {
		log.Fatalf("sub err  %v", err)
	}
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Failed to receive a note : %v", err)
		}

		log.Printf("Got currency data Name: %v Price: %v at time %v",
			in.Currency.CurrencyName, in.Currency.CurrencyPrice, in.Currency.Time)
	}

}
