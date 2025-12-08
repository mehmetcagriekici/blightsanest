package pubsub

import(
        "log"
        "errors"
	"os"
	"strconv"

        "github.com/mehmetcagriekici/blightsanest/internal/routing"

        "github.com/joho/godotenv"
        amqp "github.com/rabbitmq/amqp091-go"
)


// subscribe to a queue and consume it
func Subscribe[T any](conn *amqp.Connection,
                      queueType routing.QueueType,
		      queueName,
		      bindingKey,
		      exchangeName,
		      dlx string,
		      handler func(T) routing.AckType,
		      unmarshaller func([]byte) (T, error)) (func() error, error) {
        // fetch env variables
	if err := godotenv.Load(); err != nil {
	        log.Println(err)
	}

        prefetchCount := os.Getenv("SUBSCRIBER_PREFETCH")
	if prefetchCount == "" {
	        prefetchCount = "10"
	}
 	
	qosCount, err := strconv.Atoi(prefetchCount)
	if err != nil {
	        log.Fatal(err)
	}
	
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
		autoDelete = false
		exclusive  = false
	} else {
	        return nil, errors.New("Invalid queue type")
	}

        // declare a queue from the channel with the parameters
	qArgs := amqp.Table{
	        "x-dead-letter-exhange": dlx,
	}
	q, err := ch.QueueDeclare(queueName,
	                          durable,
				  autoDelete,
				  exclusive,
				  false,
				  qArgs)
	if err != nil {
	        return nil, err
	}
        
        // bind the queue to the exchange
	if err := ch.QueueBind(q.Name,
	                       bindingKey,
			       exchangeName,
			       false,
			       nil); err != nil {
	        return nil, err
	}

        // set the quality of service
	if err := ch.Qos(qosCount, 0, false); err != nil {
	        log.Fatal(err)
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
			ackType := handler(val)
			// acknowledgement
			switch ackType {
			case routing.ACK:
			        if err := dl.Ack(false); err != nil {
				        log.Fatal(err)
				}
			case routing.NACK_REQUEUE:
			        if err := dl.Nack(false, true); err != nil {
				        log.Fatal(err)
				}
			case routing.NACK_DISCARD:
			        if err := dl.Nack(false, false); err != nil {
				        log.Fatal(err)
				}
			default:
			        log.Fatal("Invalid Acknowledgement Type")
			}	
		}
	}()
	
        return func() error {
	        return ch.Cancel("", false)
	}, nil
}