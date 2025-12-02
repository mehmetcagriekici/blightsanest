package main

import(
        "time"
	"log"
	
        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/pubsub"
	"github.com/mehmetcagriekici/blightsanest/internal/routing"
	
	amqp "github.com/rabbitmq/amqp091-go"
)

// consumes the crypto data in the queue
func handleCryptoGet(cc     *crypto.CryptoCache,
                     cs     *crypto.CryptoState,
		     conn   *amqp.Connection,
		     frames []string,
		     sm     *pubsub.SubscriptionManager) {
        // create a key from passed frames / args and interval cache
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

        // consume
	cancel, err := pubsub.SubscribeCrypto(conn, asyncCryptoSubscriptionHandler(func(delivery routing.CryptoExchangeBody) {
	        list := delivery.Payload
		id := delivery.ID

                if id == "" || len(list) == 0 {
		        log.Println("No crypto list is delivered to the client subscriber.")
			return
		}

                cc.Add(id, list)
		cs.UpdateCurrentList(id, list)
		log.Printf("New crypto list <%s> is successfully added to the client cache and the state.\n", id)
	}))
	
	if err != nil {
	        log.Fatal(err)
	}
        
        sm.Add(cancel)	
}