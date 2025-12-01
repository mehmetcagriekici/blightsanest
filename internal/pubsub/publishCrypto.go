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
		   routingKey := fmt.Sprintf("%s.*", routing.BlightCrypto)
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

// function to publish client modifed client lists - transient
func PublishClientCrypto(ctx context.Context,
                         conn *amqp.Connection,
			 data routing.CryptoExchangeBody) error {
	routingKey := fmt.Sprintf("%s*", routing.BlightClientCrypto)
        if err := Publish(ctx,
	                  conn,
			  routing.BlightTransient,
			  routing.CryptoExchange,
			  routing.BlightTopic,
			  routingKey,
			  data); err != nil {
	        return err
	}
	return nil

}
