package main

import(
        "log"
	"strconv"
	"fmt"
	"strings"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoNewPriceSpike(cs *crypto.CryptoState, args []string) {
        controlHighPriceSpike(cs, args)
	
        log.Println("Finding the coins with a new high price spike with the preferred timeframe...")
	log.Println("")

        list := crypto.CoinsHighPriceSpike(cs.CurrentMinPriceChangePercentage, cs.CurrentTimeframe, cs.CurrentList)
	t := fmt.Sprintf("%v", cs.CurrentTimeframe)
	frame := fmt.Sprintf("PriceChangePercentage%s", strings.ToUpper(t))
	fields := []string{frame}
	crypto.PrintCryptoList(list, cs.CurrentListID, cs.ClientTimeframes, fields)
	log.Println("")
	log.Println("To update the current client list with the result: mutate find crypto high_price_spike")
}

// min price change percentage
// timeframe
func controlHighPriceSpike(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
	        log.Println("No arguments are passed for the command. Using the ones from the client state.")
		log.Printf("Current min price change percentage: %f\n", cs.CurrentMinPriceChangePercentage)
		log.Printf("Current timeframe: %v\n", cs.CurrentTimeframe)
	case 1:
	        log.Printf("Only one argument is passed as the min price change percentage. Using the client state current timeframe as the second argument: %v\n", cs.CurrentTimeframe)
		minChange, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdatePriceChangePercentage(minChange, cs.CurrentMaxPriceChangePercentage)
	case 2:
	        log.Println("Updating the client state min price change percentage and current timeframe preferences...")
		minChange, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		timeframes := crypto.GetInputTimeframes(args[1:])
		cs.UpdatePriceChangePercentage(minChange, cs.CurrentMaxPriceChangePercentage)
		cs.UpdateCurrentTimeframe(timeframes[0])
	default:
	        log.Println("Invalid use of command: find crypto high_price_spike <min_price_change_percentage float64> <timeframe>")
	}
}