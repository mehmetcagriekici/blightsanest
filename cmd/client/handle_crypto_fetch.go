package main

import(
	"log"
	
        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/pubsub"
	"github.com/mehmetcagriekici/blightsanest/internal/routing"
	
	amqp "github.com/rabbitmq/amqp091-go"
)

// gets data from the server
func handleCryptoFetch(cc     *crypto.CryptoCache,
                       cs     *crypto.CryptoState,
		       conn   *amqp.Connection,
		       args   []string,
		       sm     *pubsub.SubscriptionManager) {
	defer log.Print("> ")
        // control args
	if len(args) != 1 {
	        log.Println("<fetch crypto> command requires a key of a published crypto list as an argument.")
		log.Println("    fetch crypto <id_of_a_published_crypto_list_from_the_server>")
		return
	}
	
	key := args[0]
	
        // check if the requested list is the current list
	if key == cs.CurrentListID {
	        log.Println("Requested crypto list is already the list on the current client.")
		return
	}
	
        // check client cache if the crypto list exists
	_, ok := cc.Get(key)
	if ok {
	        log.Println("Requested crypto list already exists in the client cache.")
		log.Println("To make the requested list the current client list:")
		log.Printf("    switch crypto %s\n", key)
		return
	}

        log.Println("Subscribing to the server crypto channel to get the requested list...")
	cancel, err := pubsub.SubscribeCrypto(conn, func(delivery routing.CryptoExchangeBody) {
	        list := delivery.Payload
		id := delivery.ID

                if id == "" || len(list) == 0 {
		        log.Println("No crypto list is delivered to the client subscriber.")
			return
		}

                cc.Add(id, list)
		cs.UpdateCurrentList(id, list)
		log.Printf("New crypto list <%s> is successfully added to the client cache and the state.\n", id)
	})
	
	if err != nil {
	        log.Fatal(err)
	}
        
        sm.Add(cancel)	
}
