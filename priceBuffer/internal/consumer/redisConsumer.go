// Package consumer package for all consumers
package consumer

import (
	"github.com/AndiVS/broker-api/priceBuffer/model"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"sync"
)

// RedisStream for grpc
type RedisStream struct {
	client         *redis.Client
	mu             *sync.Mutex // protects currencyMap
	subscribersMap map[string][]*chan *model.Currency
}

// NewRedisStream create redis stream object
func NewRedisStream(client *redis.Client, mu *sync.Mutex, subscribersMap map[string][]*chan *model.Currency) *RedisStream {
	return &RedisStream{client: client, mu: mu, subscribersMap: subscribersMap}
}

// RedisConsumer consume messages from redis
func (s *RedisStream) RedisConsumer() {
	for {
		streams, err := s.client.XRead(&redis.XReadArgs{
			Streams: []string{"PriceGenerator", "$"},
		}).Result()

		if err != nil {
			log.Printf("err on consume events: %+v\n", err)
		}

		stream := streams[0].Messages[0]

		cur := new(model.Currency)
		for _, v := range stream.Values {
			err = cur.UnmarshalBinary([]byte(v.(string)))
			if err != nil {
				log.Printf("err %v ", err)
			}
			s.mu.Lock()
			for _, v := range s.subscribersMap[cur.CurrencyName] {
				*v <- cur
			}
			s.mu.Unlock()
			log.Printf("Get new data CurrencyName: %v, CurrencyPrice: %v, Time: %v", cur.CurrencyName, cur.CurrencyPrice, cur.Time)

		}
	}
}
