package pubsub

import(
        "log"
        "context"

        "github.com/mehmetcagriekici/blightsanest/internal/routing"
	
        amqp "github.com/rabbitmq/amqp091-go"
)

// function to subscribe to crypto exchanges by declaring and binding queues
func SubscribeCrypto(ctx context.Context, conn *amqp.Connection, queueName, exchangeName, exchangeKey string, queueType routing.QueueType) {
        // create a new channel from the connection
	ch, err := conn.Channel()
	if err != nil {
	        return err
	}

	var isDurable bool
	var autoDelete bool
	var isExclusive bool
        if queueType == routing.BlightDurable {
	        isDurable = true
		autoDelete = false
		isExclusive = false
	} else if queueType == routing.BlightTransient {
	        isDurable = false
		autoDelete = true
		isExclusive = true
	} else {
	        log.Fatal("Invalid QueueType")
	}
	
	// durable = queueType == durable
	// autoDelete = queueType == transient
	// exclusive = queueType == transient
	q, err := ch.QueueDeclare(queueName, isDurable, autoDelete, isExclusive, false, nil);
	if err != nil {
	        return err
	}

        // bind the queue to the exchange
	if err := ch.QueueBind(queueName, exchangeKey, exchangeName, false, nil); err != nil {
	        return err
	}

        log.Println(q)
	return nil
}