package pubsub

import(
        "context"
	"fmt"
	
        amqp "github.com/rabbitmq/amqp091-go"
	
        "github.com/mehmetcagriekici/blightsanest/internal/routing"        
)

// function to create topic exchanges on the crypto channel and publish the crypto data to that exchange
func PublishCrypto(ctx context.Context,
                   conn *amqp.Connection,
		   cryptoData routing.CryptoExchangeBody) error {
		   routingKey := fmt.Sprintf("%s*", routing.BlightCrypto)
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
