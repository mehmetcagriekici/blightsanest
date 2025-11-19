package main

import(
        "os"
	"log"
	"context"
	"github.com/joho/godotenv"
	
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
        // load environment variables
	if err := godotenv.Load(); err != nil {
	        log.Fatal(err)
	}

        // rabbit url
	rabbitURL := os.Getenv("RABBIT_CONNECTION_STRING")

        // create a context for the client
	ctx := context.Background()

        // create a connection to the rabbitmq for the client
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
	        log.Fatal(err)
	}
	defer conn.Close()

        
}