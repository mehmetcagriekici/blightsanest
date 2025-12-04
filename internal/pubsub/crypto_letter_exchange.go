package pubsub

import( 
        amqp "github.com/rabbitmq/amqp091-go"

        "github.com/mehmetcagriekici/blightsanest/internal/routing"
)

func CreateCryptoDLX(conn *amqp.Connection) error {
        // create a channel from the connection
	ch, err := conn.Channel()
	if err != nil {
	        return err
	}

        // declare an exhange from the channel
	if err := ch.ExchangeDeclare(routing.CryptoDLX,
	                            routing.BlightFanout,
				    true,
				    false,
				    false,
				    false,
				    nil); err != nil {
	        return err
	}

        // declare a queue from the channel
	if _, err := ch.QueueDeclare(routing.CryptoDLQ, true, false, false, false, nil); err != nil {
	        return err
	}

        // bind dlq to dlx
	if err := ch.QueueBind(routing.CryptoDLQ, "", routing.CryptoDLX, false, nil); err != nil {
	        return err
	}

        return nil
}