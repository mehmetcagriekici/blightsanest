package main

import(
        "log"
	"strconv"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoFilterMarketCap(cs *crypto.CryptoState) {
        log.Println("")
        log.Println("Filtering the client crypto list by the market cap using a min and max range.")

        list := crypto.FilterCoinCap(cs.CurrentMinMarketCap, cs.CurrentMaxMarketCap, cs.CurrentList)
	log.Println("")
	log.Println("Successfully filtered the crypto list by the preferred market cap range")
	log.Println("")
	
	fields := []string{"MarketCapRank", "MarketCapChangePercentage"}
	crypto.PrintCryptoList(list, cs.CurrentListID, cs.ClientTimeframes, fields)
	log.Println("")
	log.Println("To update the current client list with the result: mutate filter crypto market_cap")
}

// min market cap
// max market cap
func controlFilterMarketCap(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
	        log.Printf("No arguments are passed, using the client state values for min and max market caps: %f, %f\n", cs.CurrentMinMarketCap, cs.CurrentMaxMarketCap)
	case 1:
	        log.Printf("1 argument is passed. Using the user input as the min market cap value. Max market cap value is fetched from the client state. Max market cap: %f\n", cs.CurrentMaxMarketCap)
		minCap, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateMarketCap(minCap, cs.CurrentMaxMarketCap)
	case 2:
	        log.Println("All arguments are passed. Using the passed arguments as the min and max market caps...")
	        minCap, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		maxCap, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateMarketCap(minCap, maxCap)
	default:
	        log.Println("Invalid amount of arguments: filter crypto market_cap <min_market_cap float64> <max_market_cap float64>")
	}
}