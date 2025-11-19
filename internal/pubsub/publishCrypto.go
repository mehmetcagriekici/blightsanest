package pubsub

import(
        "context"
	
        amqp "github.com/rabbitmq/amqp091-go"
	
        "github.com/mehmetcagriekici/blightsanest/internal/routing"        
)

// function to create topic exchanges on the crypto channel and publish the crypto data to that exchange
func PublishCrypto(ctx context.Context,
                   conn *amqp.Connection,
		   cryptoData routing.CryptoExchangeBody) error {
        if err := Publish(ctx,
	                  conn,
			  routing.BlightDurable,
			  routing.CryptoExchange,
			  routing.BlightTopic,
			  routing.BlightCrypto,
			  cryptoData); err != nil {
	        return err
	}
	return nil
}
