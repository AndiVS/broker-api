package main

import (
	"context"
	"fmt"
	"github.com/AndiVS/broker-api/priceBuffer/protocolPrice"
	"github.com/AndiVS/broker-api/transactionBroker/internal/config"
	"github.com/AndiVS/broker-api/transactionBroker/internal/repository"
	"github.com/AndiVS/broker-api/transactionBroker/internal/serverBroker"
	"github.com/AndiVS/broker-api/transactionBroker/internal/service"
	"github.com/AndiVS/broker-api/transactionBroker/protocolBroker"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	setLog()

	cfg := config.Config{}
	config.New(&cfg)

	logLevel, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatalf("Unable to parse loglevel: %s", cfg.LogLevel)
	}
	log.SetLevel(logLevel)

	cfg.DBURL = getURL(&cfg)
	log.Infof("Using DB URL: %s", cfg.DBURL)

	pool := getPostgres(cfg.DBURL)

	recordRepository := repository.NewRepository(pool)

	log.Infof("Connected!")

	log.Infof("Starting HTTP server at %s...", cfg.Port)

	transactionService := service.NewTransactionService(recordRepository)

	transactionServer := serverBroker.NewTransactionServer(transactionService)

	err = runGRPCServer(transactionServer, &cfg)
	if err != nil {
		log.Printf("err in grpc run %v", err)
	}

	currencyMap := new(map[uuid.UUID]*protocolPrice.Currency)

	connectionBuffer := connectToBuffer()
	getPrices(connectionBuffer, *currencyMap)

}

func setLog() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
}

func getURL(cfg *config.Config) (URL string) {
	var str string
	str = fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		cfg.System,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
	return str
}

func getPostgres(url string) *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		log.Fatalf("Unable to connection to database: %v", err)
	}
	return pool
}

func runGRPCServer(recServer protocolBroker.TransactionServiceServer, cfg *config.Config) error {
	listener, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	protocolBroker.RegisterTransactionServiceServer(grpcServer, recServer)
	log.Printf("server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return grpcServer.Serve(listener)
}

func connectToBuffer() protocolPrice.CurrencyServiceClient {
	addressGrcp := os.Getenv("GRPC_BUFFER_ADDRESS")
	con, err := grpc.Dial(addressGrcp, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	return protocolPrice.NewCurrencyServiceClient(con)
}

func getPrices(client protocolPrice.CurrencyServiceClient, currencyMap map[uuid.UUID]*protocolPrice.Currency) {
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

			curr := protocolPrice.Currency{CurrencyID: in.Currency.CurrencyID,
				CurrencyName: in.Currency.CurrencyName, CurrencyPrice: in.Currency.CurrencyPrice, Time: in.Currency.Time}
			id, err := uuid.Parse(in.Currency.CurrencyID)
			if err != nil {
				log.Printf("err in parsing %v", err)
			}
			currencyMap[id] = &curr

			log.Printf("Got currency id: %v name: %v price: %v at time %v",
				in.Currency.CurrencyID, in.Currency.CurrencyName, in.Currency.CurrencyPrice, in.Currency.Time)
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
