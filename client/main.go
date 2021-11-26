package main

import (
	"context"
	"github.com/AndiVS/broker-api/priceServer/protocolPrice"
	"github.com/AndiVS/broker-api/transactionBroker/protocolBroker"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	connectionPriceServer := connectToPriceServer()

	subList := []string{"BTC", "ETH", "YFI"}
	subscribeToCurrency(subList, connectionPriceServer)
	for {

	}

	/*time.Sleep(10 * time.Second)
	//unsubscribeFromCurrency("ETH",subMap)
	connectionBroker := connectToBroker()
	buyCurrency(connectionBroker, "BTC", 64)*/

}

func connectToBroker() protocolBroker.TransactionServiceClient {
	//addressGRPC := os.Getenv("GRPC_BROKER_ADDRESS")
	addressGrcp := "localhost:8080"
	con, err := grpc.Dial(addressGrcp, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	return protocolBroker.NewTransactionServiceClient(con)
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

func buyCurrency(client protocolBroker.TransactionServiceClient, currency string, amount int64) {
	search, err := client.BuyCurrency(context.Background(), &protocolBroker.BuyRequest{CurrencyName: currency, CurrencyAmount: amount})
	if err != nil {
		log.Panicf("Error while buying currency: %v", err)
	}
	log.Printf("Transaction completed: %s", search.GetTransactionID())
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
