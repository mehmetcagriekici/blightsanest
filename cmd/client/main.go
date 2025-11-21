package main

import(
        "os"
	"log"
	"context"
	
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/routing"
	"github.com/mehmetcagriekici/blightsanest/internal/clientlogic"
	"github.com/mehmetcagriekici/blightsanest/internal/logs"
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

        // create a new crypto state
	cryptoState := crypto.NewCryptoState()

        // Client REPL
	clientlogic.PrintClientIntroduction()
        for {
	        words := logs.GetInput()
		// no user input
		if len(words) == 0 {
		        log.Println("Please provide a command to continue")
			clientlogic.PrintClientHelp()
			continue
		}

                // invalid commands
		if words[0] != "manual" &&
		   words[0] != "help"   &&
		   words[0] != "mutate" &&
		   words[0] != "get"    &&
		   words[0] != "rank"   &&
		   words[0] != "group"  &&
		   words[0] != "filter" &&
		   words[0] != "find"   &&
		   words[0] != "calc" {
		           log.Println("Invalid Command!")
			   clientlogic.PrintClientHelp()
			   continue
		   }

                // print client manual
		if words[0] == "manual" {
		        log.Println("----------")
			log.Println("")
		        clientlogic.PrintClientManual()
			log.Println("")
			log.Println("----------")
			continue
		}

                // print help menu to the console
		if words[0] == "help" {
		        log.Println("----------")
			log.Println("")
			clientlogic.PrintClientHelp()
			log.Println("")
			log.Println("----------")
			continue
		}

                // mutate client state
		if words[0] == "mutate" {
		}
		
	        // Get data from the server
		if words[0] == "get" {
		        // get requires at least one more argument
			if len(words) < 2 {
			        log.Println("get command requires at least one additional argument...")
				log.Println("* get crypto")
				continue
			}
		}

                // ranking features
                if words[0] == "rank" {
		}

                // grouping features
		if words[0] == "group" {
		}

                // filtering features
		if words[0] == "filter" {
		}

                // searcing features
		if words[0] == "find" {
		}

                // calculating features
		if words[0] == "calc" {
		}
	}
}
