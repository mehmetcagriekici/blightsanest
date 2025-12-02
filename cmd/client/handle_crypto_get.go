package main

import(
        "time"
	"log"
	
        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/pubsub"
	"github.com/mehmetcagriekici/blightsanest/internal/routing"
	
	amqp "github.com/rabbitmq/amqp091-go"
)

func handleCryptoGet(cc     *crypto.CryptoCache,
                     cs     *crypto.CryptoState,
		     conn   *amqp.Connection,
		     frames []string,
		     sm     *pubsub.SubscriptionManager) {
        // create a key from passed frames
	key := crypto.CreateCryptoCacheKey(frames, time.Now().Unix())

        // check if the key exists in the cache
	list, ok := cc.Get(key)
	if ok {
	        // check if the list is already the crypto state list
		if key == cs.CurrentListID {
		        // skip
			log.Println("Fetched list is already the current list")
			return
		}
		// if not update the current list with the cached list
		cs.SetTimeframes(key)
		cs.UpdateCurrentList(key, list)
		return
	}
	
	cancel, err := pubsub.SubscribeCrypto(conn, subscriberServer(cc, cs, key))
	if err != nil {
	        log.Fatal(err)
	}
        
        sm.Add(cancel)
	
	log.Println("Successfully fetched the crypto list from the server.")
	
	return
}

func subscriberServer(cc *crypto.CryptoCache, cs *crypto.CryptoState, key string) func(routing.CryptoExchangeBody) {
        return func(delivery routing.CryptoExchangeBody) {
	        list := delivery.Payload
		id   := delivery.ID

                // if there is no id or no list or id is not a match
		if id == "" || len(list) == 0 {
		        log.Fatal("No Crypto List found on the server!")
		}
		
		if key != id {
		        log.Fatal("Requested Crypto Key does not match the server key")
		}

                // add list to the cache and to the state
		cc.Add(id, list)
		cs.UpdateCurrentList(id, list)
		log.Println("New Crypto List successfully added to the client")
	}
}
