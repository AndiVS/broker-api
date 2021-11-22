package consumer

import (
	"github.com/AndiVS/broker-api/priceBuffer/protocolPrice"
	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func RedisConsumer(client *redis.Client, currencyMap map[uuid.UUID]*protocolPrice.Currency) {
	for {
		streams, err := client.XRead(&redis.XReadArgs{
			Streams: []string{"PriceGenerator", "$"},
		}).Result()

		if err != nil {
			log.Printf("err on consume events: %+v\n", err)
		}

		stream := streams[0].Messages[0]
		currencyMap = stream.Values["CurrencyMap"].(map[uuid.UUID]*protocolPrice.Currency)
		//processRedisStream(stream, currencyMap)
	}
}

/*
func processRedisStream(message redis.XMessage, currencyMap map[uuid.UUID]*protocolPrice.Currency) {
	currencyID := message.Values["CurrID"].(uuid.UUID)
	currencyName := message.Values["Name"].(string)
	currencyPrice := message.Values["Price"].(float32)
	currencyTime := message.Values["Time"].(time.Time)

	curr := protocolPrice.Currency{CurrencyID: currencyID.String(),
		CurrencyName: currencyName, CurrencyPrice: currencyPrice, Time: currencyTime.String()}

	currencyMap[currencyID] = &curr
}
*/
