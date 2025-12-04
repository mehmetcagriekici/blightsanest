package main

import(
        "log"
	"strconv"
	"strings"
	"fmt"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoFilterMarketCap(cs *crypto.CryptoState, args []string) {
        defer log.Print("> ")
	
        controlFilterMarketCap(cs, args)

        list := crypto.FilterCoinCap(cs.CurrentMinMarketCap, cs.CurrentMaxMarketCap, cs.CurrentList)

        baseID := strings.Split(cs.CurrentListID, "_")[0]
	newID := fmt.Sprintf("%s_filter_market_cap_%s_%s", baseID, cs.CurrentMinMarketCap, cs.CurrentMaxMarketCap)
	cs.UpdateCurrentList(newID, list)
	
	fields := []string{"MarketCapRank", "MarketCapChangePercentage"}
	crypto.PrintCryptoList(cs.CurrentList, cs.CurrentListID, cs.ClientTimeframes, fields)
	log.Println("")
	
	return
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