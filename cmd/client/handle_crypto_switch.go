package main

import(
        "log"

        amqp "github.com/rabbitmq/amqp091-go"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/pubsub"
)

func handleCryptoSwitch(cs *crypto.CryptoState,
                        cc *crypto.CryptoCache,
			key string,
			conn *amqp.Connection) {
        if key == "" {
	        log.Println("Please provide an ID of an existing list to use this command.")
		return
	}

        // subscribe to the crypto lists from other clients and add them to the current client's cache
	if err := pubsub.SubscribeClientCrypto(conn, key, subscriberClient(cc)); err != nil {
	        log.Fatal(err)
	}

	list, ok := cc.Get(key)
	if !ok {
	        log.Println("There is no list in the cache with the provided ID...")
		return
	}
	
        cs.UpdateCurrentList(key, list)
}
