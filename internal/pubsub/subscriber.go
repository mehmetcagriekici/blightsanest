package pubsub

import(
        "log"
        "errors"

        "github.com/mehmetcagriekici/blightsanest/internal/routing"
	
        amqp "github.com/rabbitmq/amqp091-go"
)


// subscribe to a queue and consume it
func Subscribe[T any](conn *amqp.Connection,
                      queueType routing.QueueType,
		      queueName,
		      routingKey,
		      exchangeName string,
		      handler func(T),
		      unmarshaller func([]byte) (T, error)) (func() error, error) {
        // create a channel from the connection
	ch, err := conn.Channel()
	if err != nil {
	        return nil, err
	}
	
        // queue parameters
	var durable    bool
	var autoDelete bool
	var exclusive  bool
	if queueType == routing.BlightDurable {
	        durable    = true
	        autoDelete = false
	        exclusive  = false
	} else if queueType == routing.BlightTransient {
	        durable    = false
		autoDelete = true
		exclusive  = true
	} else {
	        return nil, errors.New("Invalid queue type")
	}

        // declare a queue from the channel with the parameters
	q, err := ch.QueueDeclare(queueName,
	                          durable,
				  autoDelete,
				  exclusive,
				  false,
				  nil)
	if err != nil {
	        return nil, err
	}
        
        // bind the queue to the exchange
	if err := ch.QueueBind(q.Name,
	                       routingKey,
			       exchangeName,
			       false,
			       nil); err != nil {
	        return nil, err
	}

        // start delivering messages from the queue
	deliveries, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
	        ch.Close()
	        return nil, err
	}
	
        // range over the deliveries
	go func() {
	        defer ch.Close()
	        for dl := range deliveries {
		        // decode the delivery body
			val, err := unmarshaller(dl.Body)
			if err != nil {
			        log.Printf("Couldn't decode the delivery body: %v\n", err)
			}
			// use the decoded data on the handler function
			handler(val)
                        // remove the delivery from the queue
			if err := dl.Ack(false); err != nil {
			        log.Printf("ack error: %v\n", err)
			}
		}
	}()
	
        return func() error {
	        return ch.Cancel("", false)
	}, nil
}