package pubsub

import(
        "fmt"

        "github.com/mehmetcagriekici/blightsanest/internal/routing"
	
        amqp "github.com/rabbitmq/amqp091-go"
)

// function to subscribe to fresh crypto data from the server
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

// function to subscribe to existing crypto data from clients
func SubscribeClientCrypto(conn *amqp.Connection,
                           handler func(routing.CryptoExchangeBody)) (func() error, error) {
        bindingKey := fmt.Sprintf("%s.*", routing.BlightClientCrypto)

        return Subscribe(conn,
	                 routing.BlightTransient,
			 routing.CryptoClientGet,
			 bindingKey,
			 routing.ClientExchange,
			 handler,
			 Decode)
}