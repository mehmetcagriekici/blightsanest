package main

import(
        "log"
	"strconv"
	"fmt"
	"strings"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoFilterLowRisk(cs *crypto.CryptoState) {
        log.Println("Filtering the low risk coins by market cap rank and price change")
	log.Println("")
	
	list := crypto.FlagSafeCoins(cs.CurrentMaxRank, cs.CurrentMaxPriceChangePercentage, cs.CurrentTimeframe, cs.CurrentList)
	t := fmt.Sprintf("%v", cs.CurrentTimeframe)
	frame := fmt.Sprintf("PriceChangePercentage%s", strings.ToUpper(t))
	fields := []string{"MarketCapRank", "MarketCap", frame}
	crypto.PrintCryptoList(list, cs.CurrentListID, cs.ClientTimeframes, fields)
	log.Println("")
	log.Println("To update the client list with the result: mutate filter crypto low_risk")
}

// max market rank
// max price change
// timeframe
func controlFilterLowRisk(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
                log.Printf("No arguments are passed, using the client state values for max market rank and max price change percentage, and filter timeframe: %d %f %v\n", cs.CurrentMaxRank, cs.CurrentMaxPriceChangePercentage, cs.CurrentTimeframe)
	case 1:
	        log.Printf("Updating the max market rank preference. Using the client max price change percentage and the timeframe values. %f, %v\n", cs.CurrentMaxPriceChangePercentage, cs.CurrentTimeframe)
		maxRank, err := strconv.Atoi(args[0])
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateMarketRank(cs.CurrentMinRank, maxRank)
	case 2:
	        log.Printf("Updating the preferred max market rank and max price change percentage values. Using the client current timeframe: %v\n", cs.CurrentTimeframe)
		maxRank, err := strconv.Atoi(args[0])
		if err != nil {
		        log.Fatal(err)
		}
		maxChange, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateMarketRank(cs.CurrentMinRank, maxRank)
		cs.UpdatePriceChangePercentage(cs.CurrentMinPriceChangePercentage, maxChange)
	case 3:
	        log.Println("Updating the max market rank, max price change percentage, and current timeframe preferences")
		maxRank, err := strconv.Atoi(args[0])
		if err != nil {
		        log.Fatal(err)
		}
		maxChange, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
		        log.Fatal(err)
		}
		frames := crypto.GetInputTimeframes([]string{args[2]})
		if len(frames) != 1 {
		        log.Fatal("You must enter only one timeframe as the last argument")
		}
		cs.UpdateMarketRank(cs.CurrentMinRank, maxRank)
		cs.UpdatePriceChangePercentage(cs.CurrentMinPriceChangePercentage, maxChange)
		cs.UpdateCurrentTimeframe(frames[0])
	default:
	        log.Println("Invalid use of the command: filter crypto low_risk <max_market_cap_rank int> <max_price_change_percentage float64> <timeframe>")
       }
}