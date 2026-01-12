package main

import _ "github.com/lib/pq"

import(
	"log"
	"os"
	"time"
	"context"
	"strconv"
	"database/sql"
	
	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/serverlogic"
	"github.com/mehmetcagriekici/blightsanest/internal/logs"
	"github.com/mehmetcagriekici/blightsanest/internal/database"
	
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/joho/godotenv"
)

func main() {
        // load environment variables
	if err := godotenv.Load(); err != nil {
	        log.Fatal(err)
	}

        // environment variables
	cryptoAPIKey := os.Getenv("COIN_GECKO_KEY")
	rabbitURL := os.Getenv("RABBIT_CONNECTION_STRING")
	cacheInterval := os.Getenv("CACHE_INTERVAL")
	databaseURL := os.Getenv("DB_URL")

        // create a context for the server
	ctx := context.Background()

        // create a connection to the rabbitmq for the server
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
	        log.Fatal(err)
	}
	defer conn.Close()

        // create the server crypto cache with 24 hours reaping interval
	interval, err := strconv.ParseFloat(cacheInterval, 64)
	if err != nil {
	        log.Fatal(err)
	}
        cryptoCache := crypto.CreateCryptoCache(time.Duration(interval) * time.Hour)
	defer cryptoCache.Close()

	// open database
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	// get the gooese queries
	dbQueries := database.New(db)
	
        //REPL
	serverlogic.PrintServerIntroduction()
	for {
	        words := logs.GetInput()
		
		// no commands
		if len(words) == 0 {
		        log.Println("Please provide one of the commands to continue:")
			serverlogic.PrintServerHelp()
		        continue
		}
		
		// valid commands
		if words[0] != "quit" &&
		   words[0] != "fetch" &&
		   words[0] != "get" &&
		   words[0} != "save" &&
		   words[0] != "help" {
		        log.Println("Invalid server command! Please continue with one of these:")
			serverlogic.PrintServerHelp()
			continue
		}

                // print the server help
		if words[0] == "help" {
		        serverlogic.PrintServerHelp()
			printCryptoFetchHelp()
			continue
		}

                // quit
                if words[0] == "quit" {
		        log.Println("Exiting the BlightSanest server...")
			break
		}

                if !crypto.ControlFeatureCommands(words) {
		        continue
		}
		
                // fetch - from the api with cache
		if words[0] == "fetch" {			
			if words[1] == "crypto" {
			        handleCryptoFetch(ctx, conn, cryptoCache, cryptoAPIKey, words[2:])
                        }
		}

               // get - from the database
	       if words[0] == "get" {
	               if words[1] == "crypto" {
		               handleCryptoGet(ctx, conn, cryptoCache, words[:2], dbQueries)
		       }
	       }

		// save - to the database
		if words[0] == "save" {
			if words[1] == "crypto" {
				handleCryptoSave(ctx, cryptoCache, words[:2], dbQueries)
			}
		}
	}
}
