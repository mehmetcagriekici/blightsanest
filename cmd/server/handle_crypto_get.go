package main

import(
        "log"
        "context"
	"encoding/json"
	
        amqp "github.com/rabbitmq/amqp091-go"

        "github.com/mehmetcagriekici/blightsanest/internal/database"
	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/routing"
	"github.com/mehmetcagriekici/blightsanest/internal/pubsub"
)

func handleCryptoGet(ctx context.Context,
	conn *amqp.Connection,
	cc *crypto.CryptoCache,
	args []string,
	queries *database.Queries) {
        if len(args) != 1 {
	        log.Printf("Please provide the ID of the crypto list you want to get from the database.")
		return
	}

        // check if requested list already exists in the cache
	if _, ok := cc.Get(args[0]); ok {
	        log.Println("Requested list already exists in the current cache. Quiting the process...")
		return
	}
	
        data, err := queries.GetCryptoList(ctx, args[0])
	if err != nil {
	        log.Fatal(err)
	}

        // decode the crypto list
	mb, err := data.CryptoList.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}
	var b []crypto.MarketData
	if err := json.Unmarshal(mb, &b); err != nil {
	        log.Fatal(err)
	}

        // add new list to the cache
	cc.Add(data.CryptoKey, b)

        // publish the new list from the server
	delivery := routing.CryptoExchangeBody{
	        ID: data.CryptoKey,
		CreatedAt: data.CreatedAt,
		Payload: b,
	}
	if err := pubsub.PublishCrypto(ctx, conn, delivery); err != nil {
	        log.Fatal(err)
	}
}
