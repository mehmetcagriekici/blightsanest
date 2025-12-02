package pubsub

import(
        "fmt"

        "github.com/mehmetcagriekici/blightsanest/internal/routing"
	
        amqp "github.com/rabbitmq/amqp091-go"
)

// function to subscribe to crypto data from the server
func SubscribeCrypto(conn *amqp.Connection,
                     handler func(routing.CryptoExchangeBody)) (func() error, error) {
        // Subscribe to the crypto list
	bindingKey := fmt.Sprintf("%s.*", routing.BlightCrypto)
	
	return Subscribe(conn,
	                 routing.BlightDurable,
			 routing.CryptoGet,
			 bindingKey,
			 routing.CryptoExchange,
			 handler,
			 Decode)
}