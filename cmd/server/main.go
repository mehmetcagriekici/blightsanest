package main

import(
	"log"
	"os"
	"time"
	"context"
	"strconv"
	
	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/serverlogic"
	"github.com/mehmetcagriekici/blightsanest/internal/logs"
	"github.com/mehmetcagriekici/blightsanest/internal/pubsub"
	"github.com/mehmetcagriekici/blightsanest/internal/routing"
	
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
		   words[0] != "help" {
		        log.Println("Invalid server command! Please continue with one of these:")
			serverlogic.PrintServerHelp()
			continue
		}

                // print the server help
		if words[0] == "help" {
		        serverlogic.PrintServerHelp()
			continue
		}

                // quit
                if words[0] == "quit" {
		        log.Println("Exiting the BlightSanest server...")
			break
		}
		
                // fetch
		if words[0] == "fetch" {
		        if len(words) < 2 {
			        log.Println("fetch command requires at least one additional argument: <crypto>")
				continue
			}
			
			if words[1] == "crypto" {
			        if len(words) < 3 {
				        log.Println("While fetching crpyto you need at least one more arguments <timeframe: 1h/24h/7d/30d/200d/1y - one or multiple>")
					log.Println("ex: fetch crypto 1h 24h 7d")
					continue
				}

                                // get the timeframe queries for price_change_percentage
				frames := words[2:]
				inputTimeframes := crypto.GetInputTimeframes(frames)
				cck := crypto.CreateCryptoCacheKey(frames, time.Now().Unix()) 
				
                                // check if exists in the cache
                                if data, ok := cryptoCache.Get(cck); ok {
		                        // fetched data already exists and published on the crypto channel
					log.Println("Crypto data already exists on the server cache...")
					// publish the existing data
					createdAt, ok := cryptoCache.GetCreatedAt(cck)
					if !ok {
					        log.Println("Using the zero value for msg CreatedAt...")
					}
					msg := routing.CryptoExchangeBody{
					        ID: cck,
						CreatedAt: createdAt,
						Payload: data,
					}
					if err := pubsub.PublishCrypto(ctx, conn, msg); err != nil {
					        log.Fatal(err)
					}
					continue
				}
				
				// fetch new data from the api with the new credentials
				log.Println("Fetching new crypto data...")
				newCryptoList, err := crypto.CryptoFetchMarket(inputTimeframes, cryptoAPIKey)
				if err != nil {
				        log.Fatal(err)
				}
				log.Println("New crypto list is successfully fetched, publishing a new crypto exchange on the cryto channel")

				// add new list to the cache
				cryptoCache.Add(cck, newCryptoList)
				cryptoExchangeMessage := routing.CryptoExchangeBody{
				        ID:        cck,
					CreatedAt: time.Now(),
					Payload:   newCryptoList,
				}

                                // publish the new list on the crypto channel
				if err := pubsub.PublishCrypto(ctx,
				                               conn,
							       cryptoExchangeMessage); err != nil {
				        log.Fatal("Couldn't publish the new crypto list to the crypto channel" ,err)
				}
			}
		}
	}
}