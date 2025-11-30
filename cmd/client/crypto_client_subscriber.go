package main

import(
        "log"

	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/routing"
)

func subscriberClient(cc *crypto.CryptoCache) func(delivery routing.CryptoExchangeBody) {
        return func(delivery routing.CryptoExchangeBody) {
	        list := delivery.Payload
		id := delivery.ID
		if id == "" || len(list) == 0 {
		        log.Println("No lists from other clients")
			return
		}
		cc.Add(id, list)
		log.Printf("List %s is successfully added to the current client cache.\n", id)
	}
}
