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
	connectionBuffer := connectToBuffer()

	subMap := map[string]*protocolPrice.CurrencyService_GetPriceClient{}
	subscribeToCurrency("BTC", connectionBuffer, subMap)
	subscribeToCurrency("ETH", connectionBuffer, subMap)
	getPrices(subMap)
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

func connectToBuffer() protocolPrice.CurrencyServiceClient {
	//addressGrcp := os.Getenv("GRPC_BUFFER_ADDRESS")
	//addressGrcp := "172.28.1.9:8081"
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

func subscribeToCurrency(CurrencyName string, client protocolPrice.CurrencyServiceClient, subMap map[string]*protocolPrice.CurrencyService_GetPriceClient) {
	req := protocolPrice.GetPriceRequest{Name: CurrencyName}
	stream, err := client.GetPrice(context.Background(), &req)
	if err != nil {
		log.Fatalf("sub to %v err  %v", CurrencyName, err)
	}
	subMap[CurrencyName] = &stream
}

func unsubscribeFromCurrency(CurrencyName string, subMap map[string]*protocolPrice.CurrencyService_GetPriceClient) {
	str := *subMap[CurrencyName]
	str.CloseSend()
}

func getPrices(subMap map[string]*protocolPrice.CurrencyService_GetPriceClient) {
	for _, v := range subMap {
		go getPr(*v)
	}
}

func getPr(stream protocolPrice.CurrencyService_GetPriceClient) {
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
