package pubsub

import(
        "context"
	"fmt"
	
        amqp "github.com/rabbitmq/amqp091-go"
	
        "github.com/mehmetcagriekici/blightsanest/internal/routing"        
)

// function to publish fresh crypto data from the server - durable
func PublishCrypto(ctx context.Context,
                   conn *amqp.Connection,
		   data routing.CryptoExchangeBody) error {
        routingKey := fmt.Sprintf("%s.%s", routing.BlightCrypto, data.ID)
	
        if err := Publish(ctx,
	                  conn,
			  routing.BlightDurable,
			  routing.CryptoExchange,
			  routing.BlightTopic,
			  routingKey,
			  data); err != nil {
	        return err
	}
	
	return nil
}

// function to publish existing client data - transient
func PublishClientCrypto(ctx context.Context, conn *amqp.Connection, data routing.CryptoExchangeBody) error {
        routingKey := fmt.Sprintf("%s.%s", routing.BlightClientCrypto, data.ID)

        if err := Publish(ctx,
	                  conn,
			  routing.BlightTransient,
			  routing.ClientExchange,
			  routing.BlightTopic,
			  routingKey,
			  data); err != nil {
	        return err
	}

        return nil
}