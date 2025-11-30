package main

import(
        "log"
	"time"
	
        amqp "github.com/rabbitmq/amqp091-go"

	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/pubsub"
)

func handleCryptoList(cs *crypto.CryptoState,
                      cc *crypto.CryptoCache,
		      conn *amqp.Connection) {
        log.Println("Checking the rabbitmq server for other lists from other clients...")
	log.Println("Adding the lists from other clients to the current client's cache")
	
	key := crypto.CreateCryptoCacheKey(cs.ClientTimeframes, time.Now().Unix())
	if err := pubsub.SubscribeClientCrypto(conn, key, subscriberClient(cc)); err != nil {
	        log.Fatal(err)
	}
	
        log.Println("IDs of the existing lists on the current client cache:")
	for k := range cc.Market {
	        log.Printf("----> List: %s\n", k)
	}
}
