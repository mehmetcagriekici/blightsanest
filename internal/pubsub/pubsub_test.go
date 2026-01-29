package pubsub

import(
        "testing"
	"context"

        amqp "github.com/rabbitmq/amqp091-go"
	
        "github.com/mehmetcagriekici/blightsanest/internal/routing"
	"github.com/mehmetcagriekici/blightsanest/internal/readwrite"
)

type TestMsg struct {
        Body string
}

// Start the RabbitMQ server
func TestPubsub(t *testing.T) {
	ctx := context.Background()
        conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
	        t.Errorf("Couldn't create the connection")
	}
	defer conn.Close()

        msg := TestMsg{
	        Body: "this is the test message.",
	}

        // publish the message
        if err := Publish(ctx,
	                  conn,
			  routing.BlightTransient,
			  routing.TestExchange,
			  routing.BlightDirect,
			  routing.BlightTesting,
			  msg); err != nil {
	        t.Errorf("Couldn't publish the message")
	}

        // subscribe to the message
	cancel, err := Subscribe(conn,
	                         routing.BlightTransient,
				 routing.TestQueue,
				 routing.BlightTesting,
				 routing.TestExchange,
				 routing.CryptoDLX,
				 func(delivery TestMsg) routing.AckType {
				         res := delivery.Body
					 if res != msg.Body {
					         t.Errorf("expected: %s, got: %s", msg.Body, res)
                                                 return routing.NACK_DISCARD
					 }
					 return routing.ACK
				 },
				 readwrite.Decode)
	if err != nil {
	        t.Errorf("Couldn't subscribe to the message")
	}

	defer cancel()
}
