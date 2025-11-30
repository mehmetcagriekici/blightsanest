package main

import(
        "log"
	"strconv"
	"fmt"
	"strings"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoFilterPriceChangePercentage(cs *crypto.CryptoState, args []string) {
        controlFilterPriceChangePercentage(cs, args)
	
        log.Println("Filtering the crypto list with the client preferences min/max price change percentage and current timeframe")
	log.Println("")
	list := crypto.FilterCoinPriceChange(cs.CurrentMinPriceChangePercentage, cs.CurrentMaxPriceChangePercentage, cs.CurrentTimeframe, cs.CurrentList)
	
	t := fmt.Sprintf("%v", cs.CurrentTimeframe)
        frame := fmt.Sprintf("PriceChangePercentage%s", strings.ToUpper(t))
        fields := []string{frame}
        crypto.PrintCryptoList(list, cs.CurrentListID, cs.ClientTimeframes, fields)
	log.Println("")
	log.Println("To update the current client list with the result: mutate filter crypto price_change_percentage")
}

// min price change
// max price change
// timeframe
func controlFilterPriceChangePercentage(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
	        log.Println("No arguments are passed, using the client state values for min, max price change percentages and time frame.")
		log.Printf("min price change percentage: %f\n", cs.CurrentMinPriceChangePercentage)
		log.Printf("max price change percentage: %f\n", cs.CurrentMaxPriceChangePercentage)
		log.Printf("Current timeframe: %v\n", cs.CurrentTimeframe)
	case 1:
	        log.Println("Updating the min price change percentage preference.")
		log.Printf("Current max price change percentage: %f\n", cs.CurrentMaxPriceChangePercentage)
		log.Printf("Current timeframe: %v\n", cs.CurrentTimeframe)

                minChange, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdatePriceChangePercentage(minChange, cs.CurrentMaxPriceChangePercentage)
	case 2:
	        log.Println("Updating the min and max price change preferences.")
		log.Printf("Current timeframe: %v\n", cs.CurrentTimeframe)
		minChange, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		maxChange, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdatePriceChangePercentage(minChange, maxChange)
	case 3:
	        log.Println("Updating the min and max price change and timeframe preferences.")
		minChange, err := strconv.ParseFloat(args[0], 64)
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
		cs.UpdatePriceChangePercentage(minChange, maxChange)
		cs.UpdateCurrentTimeframe(frames[0])
	default:
	        log.Println("Unexpected usage of the command: filter crypto price_change_percentage <min_change float64> <max_change float64> <timeframe>")
	}
}