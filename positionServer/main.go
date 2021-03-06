package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"sync"

	"github.com/AndiVS/broker-api/positionServer/internal/config"
	modelLocal "github.com/AndiVS/broker-api/positionServer/internal/model"
	"github.com/AndiVS/broker-api/positionServer/internal/repository"
	"github.com/AndiVS/broker-api/positionServer/internal/server"
	"github.com/AndiVS/broker-api/positionServer/internal/service"
	"github.com/AndiVS/broker-api/positionServer/positionProtocol"
	"github.com/AndiVS/broker-api/priceServer/model"
	"github.com/AndiVS/broker-api/priceServer/priceProtocol"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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

	currencyMap := map[string]*model.Currency{}
	positionMap := map[string]map[uuid.UUID]*chan *model.Currency{
		"BTC": {},
		"ETH": {},
		"YFI": {},
	}

	go listen(pool, positionMap)

	mute := new(sync.Mutex)
	connectionBuffer := connectToPriceServer()
	subList := []string{"BTC", "ETH"}
	go getPrices(subList, connectionBuffer, mute, currencyMap, positionMap)

	transactionService := service.NewPositionService(recordRepository)
	transactionServer := server.NewPositionServer(transactionService, mute, currencyMap)

	err = runGRPCServer(transactionServer)
	if err != nil {
		log.Printf("err in grpc run %v", err)
	}
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

func listen(pool *pgxpool.Pool, positionMap map[string]map[uuid.UUID]*chan *model.Currency) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Println("Error acquiring connection:", err)
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), "listen positions")
	if err != nil {
		log.Println("Error listening to transactions channel:", err)
	}

	for {
		notification, err := conn.Conn().WaitForNotification(context.Background())
		if err != nil {
			log.Println("Error waiting for notification:", err)
		}

		pos := modelLocal.Position{}
		err = pos.UnmarshalBinary([]byte(notification.Payload))
		if err != nil {
			log.Println("Error waiting for notification:", err)
		}
		mute := new(sync.Mutex)
		ch := make(chan *model.Currency)
		switch pos.Event {
		case "INSERT":
			log.Printf("Open position with id %v currency name %v open price %v open time %v", pos.PositionID, pos.CurrencyName, pos.OpenPrice, pos.OpenTime)
			mute.Lock()
			positionMap[pos.CurrencyName][*pos.PositionID] = &ch
			go evaluateProfit(&pos, ch)
		case "UPDATE":
			profit := (*pos.ClosePrice - *pos.OpenPrice) * float32(*pos.Amount)
			log.Printf("Position id %v close with profit %v", pos.PositionID.String(), profit)
			close(*positionMap[pos.CurrencyName][*pos.PositionID])
			delete(positionMap[pos.CurrencyName], *pos.PositionID)
		}
	}
}

func runGRPCServer(recServer positionProtocol.PositionServiceServer) error {
	port := os.Getenv("GRPC_BROKER_PORT")
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	positionProtocol.RegisterPositionServiceServer(grpcServer, recServer)
	log.Printf("server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return grpcServer.Serve(listener)
}

func connectToPriceServer() priceProtocol.CurrencyServiceClient {
	addressGrcp := fmt.Sprintf("%s%s", os.Getenv("GRPC_BUFFER_HOST"), os.Getenv("GRPC_BUFFER_PORT"))

	con, err := grpc.Dial(addressGrcp, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	return priceProtocol.NewCurrencyServiceClient(con)
}

func getPrices(CurrencyName []string, client priceProtocol.CurrencyServiceClient, mu *sync.Mutex,
	currencyMap map[string]*model.Currency, positionMap map[string]map[uuid.UUID]*chan *model.Currency) {
	req := priceProtocol.GetPriceRequest{Name: CurrencyName}
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
		currencyMap[cur.CurrencyName] = &cur
		for _, v := range positionMap[cur.CurrencyName] {
			*v <- &cur
		}
		mu.Unlock()
	}
}

func evaluateProfit(pos *modelLocal.Position, ch chan *model.Currency) {
	for {
		current, ok := <-ch
		if ok {
			profit := (current.CurrencyPrice - *pos.OpenPrice) * float32(*pos.Amount)

			log.Printf("For position %v, profit %v at time %v   (open price: %v current price %v amount %v)", pos.PositionID, profit,
				current.Time, pos.OpenPrice, current.CurrencyPrice, pos.Amount)
		} else {
			return
		}
	}
}
