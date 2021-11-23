package consumer

import (
	"github.com/AndiVS/broker-api/priceBuffer/protocolPrice"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"sync"
)

// RedisStream for grpc
type RedisStream struct {
	client      *redis.Client
	mu          *sync.Mutex // protects currencyMap
	currencyMap map[string]protocolPrice.Currency
}

func NewRedisStream(client *redis.Client, mu *sync.Mutex, currencyMap map[string]protocolPrice.Currency) *RedisStream {
	return &RedisStream{client: client, mu: mu, currencyMap: currencyMap}
}

func (s *RedisStream) RedisConsumer() {
	for {
		streams, err := s.client.XRead(&redis.XReadArgs{
			Streams: []string{"PriceGenerator", "$"},
		}).Result()

		if err != nil {
			log.Printf("err on consume events: %+v\n", err)
		}

		stream := streams[0].Messages[0]

		cur := new(protocolPrice.Currency)

		err = cur.UnmarshalBinary([]byte(stream.Values["Currency"].(string)))
		if err != nil {
			log.Printf("err %v ", err)
		}

		s.mu.Lock()
		s.currencyMap[cur.CurrencyName] = *cur
		s.mu.Unlock()
		log.Printf("Get new data CurrencyName: %v, CurrencyPrice: %v, Time: %v", cur.CurrencyName, cur.CurrencyPrice, cur.Time)
	}
}
