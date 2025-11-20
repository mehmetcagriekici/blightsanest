package pubsub

import(
        "log"

        "github.com/mehmetcagriekici/blightsanest/internal/routing"
	
        amqp "github.com/rabbitmq/amqp091-go"
)

// function to subscribe to crypto exchanges by declaring and binding queues
func SubscribeCrypto(conn *amqp.Connection,
                     handler func(routing.CryptoExchangeBody)) error {
        // Subscribe to the crypto list
	if err := Subscribe(conn,
	                    routing.BlightDurable,
			    routing.CryptoGet,
			    routing.BlightCrypto,
			    routing.CryptoExchange,
			    handler()); err != nil {
	        return err
	}
	return nil
}
