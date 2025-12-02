package main

import(
        "log"

        amqp "github.com/rabbitmq/amqp091-go"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/pubsub"
	"github.com/mehmetcagriekici/blightsanest/internal/routing"
)

func handleCryptoSwitch(cs  *crypto.CryptoState,
                        cc  *crypto.CryptoCache,
			args []string,
			conn *amqp.Connection,
			sm   *pubsub.SubscriptionManager) {
	if len(args) != 1 {
	        log.Println("Please provide an ID of an existing list.")
		return
	}
	
         key := args[0]
        // subscribe to the crypto lists from other clients and add them to the current client's cache
	cancel, err := pubsub.SubscribeCrypto(conn, asyncCryptoSubscriptionHandler(func(delivery routing.CryptoExchangeBody) {
	        list := delivery.Payload
		id := delivery.ID

                if len(list) > 0 {
		        cc.Add(id, list)
			log.Println("New crypto list is successfully added to the client cache from the rabbitmq server.")
		}
		
                if len(cc.Market) == 0 {
		        log.Println("No other lists in the client cache...")
			return
		}
		
                newList, ok := cc.Get(key)
		if !ok {
		        log.Println("There is no list in the client cache with the passed ID.")
			return
		}
		
		cs.UpdateCurrentList(key, newList)
		log.Printf("Successfully switched to the list: %s\n", key)
	}))
	
	if err != nil {
	        log.Fatal(err)
	}
	
	sm.Add(cancel)
}
