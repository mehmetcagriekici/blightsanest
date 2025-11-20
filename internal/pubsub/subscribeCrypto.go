package pubsub

import(
        "fmt"

        "github.com/mehmetcagriekici/blightsanest/internal/routing"
	
        amqp "github.com/rabbitmq/amqp091-go"
)

// function to subscribe to crypto exchanges by declaring and binding queues
func SubscribeCrypto(conn *amqp.Connection,
                     msgID string,
                     handler func(routing.CryptoExchangeBody)) error {
        // Subscribe to the crypto list
	routingKey := fmt.Sprintf("%s-%s", routing.BlightCrypto, msgID)
	if err := Subscribe(conn,
	                    routing.BlightDurable,
			    routing.CryptoGet,
			    routingKey,
			    routing.CryptoExchange,
			    handler); err != nil {
	        return err
	}
	return nil
}
