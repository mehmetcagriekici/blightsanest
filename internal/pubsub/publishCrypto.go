package pubsub

import(
        "log"
        "time"
        "context"
        "encoding/gob"
	"bytes"

        "github.com/mehmetcagriekici/blightsanest/internal/routing"
	
        amqp "github.com/rabbitmq/amqp091-go"
)

// function to create topic exchanges on the crypto channel and publish the crypto data to that exchange
func PublishCrypto(ctx context.Context, conn *amqp.Connection, name, key, exchangeType string, queueType routing.QueueType, body routing.CryptoExchangeBody) error {
        // create a channel from the connection
	ch, err := conn.Channel()
	if err != nil {
	        return err
	}

        var durability bool
	if queueType == routing.BlightDurable {
	        durability = true
	} else if queueType == routing.BlightTransient {
	        durability = false
	} else {
	        log.Fatal("Invalid Queue Type")
	}
	
        // declare the exchange on the channel - durable exchange
	if err := ch.ExchangeDeclare(name, exchangeType, durability, false, false, false, nil); err != nil {
	        return err
	}

        // encode the exchane message body
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	if err := enc.Encode(body); err != nil {
	        return err
	}

        // exchange message
	msg := amqp.Publishing{
	        DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "application/gob",
		Body:         network.Bytes(),
	}

        // publish the message to the exchange with the app context
	if err := ch.PublishWithContext(ctx, name, key, false, false, msg); err != nil {
	        return err
	}
	return nil
}
