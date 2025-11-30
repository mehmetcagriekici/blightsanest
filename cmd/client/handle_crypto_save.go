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
        // save current list to the cache
        log.Println("Saving the current list to the client cache and to the docker volume.")
	cc.Add(cs.CurrentListID, cs.CurrentList)

        // publish current list to the other clients
	data := routing.CryptoExchangeBody{
	        ID:        cs.CurrentListID,
		CreatedAt: time.Now(),
		Payload: cs.CurrentList,
	}
	if err := pubsub.PublishClientCrypto(ctx, conn, data); err != nil {
	        log.Fatal(err)
	}
}