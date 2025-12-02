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
        if len(cs.CurrentList) == 0 {
	        log.Println("Current client list is empty...")
		return
	}
	
        // save current list to the cache
	cc.Add(cs.CurrentListID, cs.CurrentList)
	createdAt, ok := cc.GetCreatedAt(cs.CurrentListID)
	if !ok {
	        log.Println("Using the current time as the message created at.")
		createdAt = time.Now()
	}

        // publish current list to the other clients
	data := routing.CryptoExchangeBody{
	        ID:        cs.CurrentListID,
		CreatedAt: createdAt,
		Payload:   cs.CurrentList,
	}
	
	if err := pubsub.PublishCrypto(ctx, conn, data); err != nil {
	        log.Fatal(err)
	}
}