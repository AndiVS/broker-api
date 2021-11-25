package main

import (
	"context"
	"fmt"
	"github.com/AndiVS/broker-api/priceBuffer/model"
	"github.com/AndiVS/broker-api/priceBuffer/protocolPrice"
	"github.com/AndiVS/broker-api/transactionBroker/internal/config"
	"github.com/AndiVS/broker-api/transactionBroker/internal/repository"
	"github.com/AndiVS/broker-api/transactionBroker/internal/serverBroker"
	"github.com/AndiVS/broker-api/transactionBroker/internal/service"
	"github.com/AndiVS/broker-api/transactionBroker/protocolBroker"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"net"
	"sync"
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

	go listen(pool)

	recordRepository := repository.NewRepository(pool)

	log.Infof("Connected!")

	currencyMap := map[string]model.Currency{}

	mute := new(sync.Mutex)
	connectionBuffer := connectToBuffer()
	go getPrices("BTC", connectionBuffer, mute, currencyMap)
	go getPrices("ETH", connectionBuffer, mute, currencyMap)

	transactionService := service.NewTransactionService(recordRepository)
	transactionServer := serverBroker.NewTransactionServer(transactionService, mute, currencyMap)

	err = runGRPCServer(transactionServer)
	if err != nil {
		log.Printf("err in grpc run %v", err)
	}

	/*connectionBuffer := connectToBuffer()
	getPrices(connectionBuffer)
	*/
}

func setLog() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
}

func getURL(cfg *config.Config) (URL string) {
	str := fmt.Sprintf("%s://%s:%s@%s:%d/%s",
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

func listen(pool *pgxpool.Pool) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Println("Error acquiring connection:", err)
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), "listen transactions")
	if err != nil {
		log.Println("Error listening to transactions channel:", err)
	}

	for {
		notification, err := conn.Conn().WaitForNotification(context.Background())
		if err != nil {
			log.Println("Error waiting for notification:", err)
		}

		log.Println("PID:", notification.PID, "Channel:", notification.Channel, "Payload:", notification.Payload)
	}
}

func runGRPCServer(recServer protocolBroker.TransactionServiceServer) error {
	//port := os.Getenv("GRPC_BROKER_PORT")
	port := ":8080"
	listener, err := net.Listen("tcp", port)
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
	//addressGrcp := fmt.Sprintf("%s%s", os.Getenv("GRPC_BUFFER_HOST"), os.Getenv("GRPC_BUFFER_PORT"))
	addressGrcp := "172.28.1.9:8081"
	con, err := grpc.Dial(addressGrcp, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	return protocolPrice.NewCurrencyServiceClient(con)
}

func getPrices(CurrencyName string, client protocolPrice.CurrencyServiceClient, mu *sync.Mutex, currencyMap map[string]model.Currency) {
	req := protocolPrice.GetPriceRequest{Name: CurrencyName}
	stream, err := client.GetPrice(context.Background(), &req)
	if err != nil {
		log.Fatalf("%v.RouteChat(_) = _, %v", client, err)
	}
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Failed to receive a note : %v", err)
		}

		cur := model.Currency{CurrencyName: in.Currency.CurrencyName, CurrencyPrice: in.Currency.CurrencyPrice, Time: in.Currency.Time}
		mu.Lock()
		currencyMap[cur.CurrencyName] = cur
		mu.Unlock()

		log.Printf("Got currency data Name: %v Price: %v at time %v",
			in.Currency.CurrencyName, in.Currency.CurrencyPrice, in.Currency.Time)
	}
}
