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
	// "github.com/mehmetcagriekici/blightsanest/internal/routing"
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
	// cache interval
	cacheInterval := os.Getenv("CACHE_INTERVAL")

        // create a context for the client
	ctx := context.Background()
	ctx = ctx

        // create a connection to the rabbitmq for the client
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
	        log.Fatal(err)
	}
	defer conn.Close()

        // create a new crypto state
	cryptoState := crypto.CreateCryptoState()

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

                        if words[1] == "crypto" {
			        clientlogic.PrintCryptoHelp()
			}
			continue
		}

                // feature commands requires at least one more argument
		if !controlFeatureCommands(words) {
		        continue
		}

                // mutate client state
		if words[0] == "mutate" {
		        // update the client state list with the result of the operation
		}

                // switch between cached data
		if words[0] == "switch" {
		}

                // save the asset on the cache
		if words[0] == "save" {
		        // to save the current list on the state in the cache and to the docker volume
		}
		
	        // Get data from the server
		if words[0] == "get" {
		        if words[1] == "crypto" {
			        frames := words[2:]
				handleCryptoGet(cryptoCache, cryptoState, conn, frames)
				crypto.PrintCryptoList(cryptoState.CurrentList,
				                       cryptoState.CurrentListID,
						       cryptoState.ClientTimeframes,
						       []string{})
                                continue
			}
		}

                // ranking features
                if words[0] == "rank" {
		        if words[1] == "crypto" {
			        handleCryptoRank(cryptoState, words[2], words[3])
				continue
			}
		}

                // grouping features
		if words[0] == "group" {
		        if !controlFeatureSub(words) {
			        continue
			}
			
		        if words[1] == "crypto" {
			        switch words[2] {
				case "liquidity":
			                controlLiquidityArguments(cryptoState, words[3:])
					handleCryptoLiquidity(cryptoState)
				case "scarcity":
				        controlScarcityArguments(cryptoState, words[3:])
					handleCryptoScarcity(cryptoState)
				default:
				        log.Println("Invalid crypto grouping option. Available: <liquidity> <scarcity>")
				}
				continue
			}
		}

                // filtering features
		if words[0] == "filter" {
		        if !controlFeatureSub(words) {
			        continue
			}
			
                        if words[1] == "crypto" {
                                switch words[2] {
			        case "total_volume":
				        controlFilterTotalVolume(cryptoState, words[3:])
					handleCryptoFilterTotalVolume(cryptoState)
			        case "market_cap":
				        controlFilterMarketCap(cryptoState, words[3:])
					handleCryptoFilterMarketCap(cryptoState)
				case "price_change_percentage":
				        controlFilterPriceChangePercentage(cryptoState, words[3:])
					handleCryptoFilterPriceChangePercentage(cryptoState)
				case "volatile":
				        controlFilterVolatile(cryptoState, words[3:])
					handleCryptoFilterVolatile(cryptoState)
				case "high_risk":
				        controlFilterHighRisk(cryptoState, words[3:])
					handleCryptoFilterHighRisk(cryptoState)
				case "low_risk":
				        controlFilterLowRisk(cryptoState, words[3:])
					handleCryptoFilterLowRisk(cryptoState)
				default:
				        log.Println("Invalid crypto filtering option. Available: <total_volume> <market_cap> <price_change_percentage> <volatile> <high_risk> <low_risk>")
			        }
			continue
		        }
		}

                // searcing features
		if words[0] == "find" {
		        if !controlFeatureSub(words) {
			        continue
			}

                        if words[1] == "crypto" {
			        switch words[2] {
				case "name":
				case "new_high_price":
				case "new_low_price":
				case "high_price_spike":
				case "potential_rally":
				case "coin_inflation":
				default:
				        log.Println("Invalid crypto search command. Available: <name>, <new_high_price>, <high_price_spike>, <potential_rally>, <coin_inflation>")
				}
				continue
			}
		}

                // calculating features
		if words[0] == "calc" {
		        if !controlFeatureSub(words) {
			        continue
			}

                        if words[1] == "crypto" {
			        switch words[2] {
				case "volatility":
				case "growth_potential":
				case "liquidity":
				case "trend_strength":
				default:
				        log.Println("Invalid crypto calculation command. Available: <volatility>, <growth_potential>, <liquidity>, <trend_strength>")
				}
				continue
			}
		}

	}
}
