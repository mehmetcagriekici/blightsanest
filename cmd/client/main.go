package main

import(
        "os"
	"log"
	"context"
	"strconv"
        "time"
	
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/clientlogic"
	"github.com/mehmetcagriekici/blightsanest/internal/logs"
	"github.com/mehmetcagriekici/blightsanest/internal/pubsub"
)

func main() {
        // load environment variables
	if err := godotenv.Load(); err != nil {
	        log.Fatal(err)
	}

        // rabbit url
	rabbitURL := os.Getenv("RABBIT_CONNECTION_STRING")
	// cache interval
	cacheInterval := os.Getenv("CACHE_INTERVAL")

        // create a context for the client
	ctx := context.Background()

        // create a connection to the rabbitmq for the client
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
	        log.Fatal(err)
	}
	defer conn.Close()

        // create a new crypto state
	cryptoState := crypto.CreateCryptoState()

        // create a new subscription manager for the crypto
	cryptoSubscriptionManager := pubsub.NewSubscriptionManager()

        // create a crypto cache for the client
	interval, err := strconv.ParseFloat(cacheInterval, 64)
	if err != nil {
	        log.Fatal(err)
	}
	cryptoCache := crypto.CreateCryptoCache(time.Duration(interval) * time.Hour)
        
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
		   words[0] != "quit"   &&
		   words[0] != "mutate" &&
		   words[0] != "switch" &&
		   words[0] != "save"   &&
		   words[0] != "list"   &&
		   words[0] != "get"    &&
		   words[0] != "rank"   &&
		   words[0] != "group"  &&
		   words[0] != "filter" &&
		   words[0] != "find"   &&
		   words[0] != "calc" {
		           log.Println("Invalid Command!")
			   clientlogic.PrintClientHelp()
			   log.Println("")
			   continue
		}

                // quit REPL
		if words[0] == "quit" {
		        log.Println("Ending client session...")
			cryptoSubscriptionManager.CloseAll()
			time.Sleep(200 * time.Millisecond)
			break
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
			log.Println("")
			log.Println("To see the use of commands for a specific asset: help <asset_name>")
			log.Println("* help crypto: will print the crypto part of the manual.")
			log.Println("")

                        if len(words) == 2 {
                                if words[1] == "crypto" {
			                clientlogic.PrintCryptoHelp()
			        }
			}
			continue
		}

                // feature commands requires at least one more argument
		if !controlFeatureCommands(words) {
		        continue
		}

                // mutate client state
		if words[0] == clientlogic.CLIENT_MUTATE {
		        // mutate <operation> <asset> <asset_feature>
			if len(words) < 4 {
			        log.Println("mutate command usage: mutate <operation> <asset> <feature> <operation_arguments - optional if not provided uses the client preferences ->")
				log.Println("  Example: mutate group crypto liquidity")
			        continue 
			}
			
		        // update the client state list with the result of the operation
			if words[2] == "crypto" {
			        if words[1] == "rank" {
				        handleCryptoMutate(cryptoState, cryptoCache, words[1], "", words[3:])
					continue
				}
			        handleCryptoMutate(cryptoState, cryptoCache, words[1], words[3], words[4:])
				continue
			}
			continue
		}

                // switch between cached data
		if words[0] == clientlogic.CLIENT_SWITCH {
		        if words[1] == "crypto" {
				handleCryptoSwitch(cryptoState,
				                   cryptoCache,
						   words[2:],
						   conn,
						   cryptoSubscriptionManager)
				continue
			}
		}

                // save the asset on the cache
		if words[0] == clientlogic.CLIENT_SAVE {
		        // to save the current list in the cache
			handleCryptoSave(cryptoState, cryptoCache, ctx, conn)
			continue
		}

                // list the existing lists in the cache
		if words[0] == clientlogic.CLIENT_LIST {
		        handleCryptoList(cryptoState, cryptoCache)
			continue
		}
		
	        // Get data from the serve
		if words[0] == clientlogic.CLIENT_GET {
		        if words[1] == "crypto" {
			        frames := words[2:]
				handleCryptoGet(cryptoCache, cryptoState, conn, frames, cryptoSubscriptionManager)
                                continue
			}
		}

                // ranking features
                if words[0] == clientlogic.CLIENT_RANK {
		        if words[1] == "crypto" {
			        handleCryptoRank(cryptoState, words[2:])
				continue
			}
		}

                // grouping features
		if words[0] == clientlogic.CLIENT_GROUP {
		        if !controlFeatureSub(words) {
			        continue
			}
			
		        if words[1] == "crypto" {
			        switch words[2] {
				case clientlogic.CRYPTO_GROUP_LIQUIDITY:
			               	handleCryptoGroupLiquidity(cryptoState, words[3:])
				case clientlogic.CRYPTO_GROUP_SCARCITY:
					handleCryptoGroupScarcity(cryptoState, words[3:])
				default:
				        log.Println("Invalid crypto grouping option. Available: <liquidity> <scarcity>")
				}
				continue
			}
		}

                // filtering features
		if words[0] == clientlogic.CLIENT_FILTER {
		        if !controlFeatureSub(words) {
			        continue
			}
			
                        if words[1] == "crypto" {
                                switch words[2] {
			        case "total_volume":
					handleCryptoFilterTotalVolume(cryptoState, words[3:])
			        case "market_cap":
					handleCryptoFilterMarketCap(cryptoState, words[3:])
				case "price_change_percentage":
					handleCryptoFilterPriceChangePercentage(cryptoState, words[3:])
				case "volatile":
					handleCryptoFilterVolatile(cryptoState, words[3:])
				case "high_risk":
					handleCryptoFilterHighRisk(cryptoState, words[3:])
				case "low_risk":
					handleCryptoFilterLowRisk(cryptoState, words[3:])
				default:
				        log.Println("Invalid crypto filtering option. Available: <total_volume> <market_cap> <price_change_percentage> <volatile> <high_risk> <low_risk>")
			        }
			continue
		        }
		}

                // searcing features
		if words[0] == clientlogic.CLIENT_FIND {
		        if !controlFeatureSub(words) {
			        continue
			}

                        if words[1] == "crypto" {
			        switch words[2] {
				case clientlogic.CRYPTO_FIND_NAME:
				        handleCryptoFindName(cryptoState, words[3])
				case clientlogic.CRYPTO_FIND_NEW_HIGH_PRICE:
				        handleCryptoNewHighPrice(cryptoState, cryptoCache)
				case clientlogic.CRYPTO_FIND_NEW_LOW_PRICE:
				        handleCryptoNewLowPrice(cryptoState, cryptoCache)
				case clientlogic.CRYPTO_FIND_HIGH_PRICE_SPIKE:
					handleCryptoNewPriceSpike(cryptoState, words[3:])
				case clientlogic.CRYPTO_FIND_POTENTIAL_RALLY:
					handleCryptoFindPotentialRally(cryptoState, words[3:])
				case clientlogic.CRYPTO_FIND_COIN_INFLATION:
				        handleCryptoFindCoinInflation(cryptoState, words[3:])
				default:
				        log.Println("Invalid crypto search command. Available: <name>, <new_high_price>, <high_price_spike>, <potential_rally>, <coin_inflation>")
				}
				continue
			}
		}

                // calculating features
		if words[0] == clientlogic.CLIENT_CALC {
		        if !controlFeatureSub(words) {
			        continue
			}

                        if words[1] == "crypto" {
			        switch words[2] {
				case clientlogic.CRYPTO_CALC_VOLATILITY:
				        handleCryptoCalcVolatility(cryptoState, words[3:])
				case clientlogic.CRYPTO_CALC_GROWTH_POTENTIAL:
					handleCryptoCalcGrowthPotential(cryptoState, words[3:])
				case clientlogic.CRYPTO_CALC_LIQUIDITY:
					handleCryptoCalcLiquidity(cryptoState, words[3:])
				case clientlogic.CRYPTO_CALC_TREND_STRENGTH:
					handleCryptoCalcTrendStrength(cryptoState, words[3:])
				default:
				        log.Println("Invalid crypto calculation command. Available: <volatility>, <growth_potential>, <liquidity>, <trend_strength>")
				}
				continue
			}
		}

	}
}
