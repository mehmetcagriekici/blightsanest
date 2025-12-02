package pubsub

import(
        "context"
	"fmt"
	
        amqp "github.com/rabbitmq/amqp091-go"
	
        "github.com/mehmetcagriekici/blightsanest/internal/routing"        
)

// function to publish raw crypto data from the server - durable
func PublishCrypto(ctx context.Context,
                   conn *amqp.Connection,
		   cryptoData routing.CryptoExchangeBody) error {
        routingKey := fmt.Sprintf("%s.%s", routing.BlightCrypto, cryptoData.ID)
	
        if err := Publish(ctx,
	                  conn,
			  routing.BlightDurable,
			  routing.CryptoExchange,
			  routing.BlightTopic,
			  routingKey,
			  cryptoData); err != nil {
	        return err
	}
	
	return nil
}
