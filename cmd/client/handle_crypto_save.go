package main

import(
        "log"
	"context"
	"time"

        amqp "github.com/rabbitmq/amqp091-go"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/routing"
	"github.com/mehmetcagriekici/blightsanest/internal/pubsub"
)

func handleCryptoSave(cs *crypto.CryptoState,
                      cc *crypto.CryptoCache,
		      ctx context.Context,
		      conn *amqp.Connection) {
        defer log.Print("> ")
	
        if len(cs.CurrentList) == 0 {
	        log.Println("Current client list is empty. Nothing to save...")
		return
	}
	
        // save current list to the cache
	cc.Add(cs.CurrentListID, cs.CurrentList)
	
        // publish current list to the other clients
	data := routing.CryptoExchangeBody{
	        ID:        cs.CurrentListID,
		CreatedAt: time.Now(),
		Payload:   cs.CurrentList,
	}

        log.Printf("Publishing the list: %s\n", cs.CurrentListID)
	if err := pubsub.PublishClientCrypto(ctx, conn, data); err != nil {
	        log.Fatal(err)
	}
	
	log.Println("List is successfully published.")
	return
}