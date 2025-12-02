package pubsub

import(
        "time"
        "context"
	"errors"

        "github.com/mehmetcagriekici/blightsanest/internal/routing"
	
        amqp "github.com/rabbitmq/amqp091-go"
)

func Publish[T any](ctx context.Context,
                    conn *amqp.Connection,
		    queueType routing.QueueType,
		    exchangeName,
		    exchangeType,
		    routingKey string,
		    val T) error {
        // create a channel from the connection
	ch, err := conn.Channel()
	if err != nil {
	        return err
	}

        // assign exchange declare durability depending on the queue type
	var ed bool
	if queueType == routing.BlightDurable {
	        ed = true
	} else if queueType == routing.BlightTransient {
	        ed = false
	} else {
	        return errors.New("Invalid queue type")
	}

        // declare the exchange
	if err := ch.ExchangeDeclare(exchangeName,
	                             exchangeType,
				     ed,
				     false,
				     false,
				     false,
				     nil); err != nil {
				             return err
				     }


        // encode the value into bytes
	data, err := Encode(val)
	if err != nil {
	        return err
	}

        // build the exchange message
	msg := amqp.Publishing{
	        DeliveryMode: amqp.Persistent,
		Timestamp: time.Now(),
		ContentType: "application/gob",
		Body: data,
	}

        // publish the message
	if err := ch.PublishWithContext(ctx,
	                                exchangeName,
					routingKey,
					false,
					false,
					msg); err != nil {
					        return err
					}
	return nil
}
