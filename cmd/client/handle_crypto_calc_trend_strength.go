package main

import(
        "log"
	"fmt"
	
        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoCalcTrendStrength(cs *crypto.CryptoState, args []string) {
        controlCalcTrendStrength(cs, args)
        list, err := crypto.CheckRealTrend(cs.CurrentTimeframe, cs.CurrentList)
	if err != nil {
	        log.Fatal(err)
	}
	newID := fmt.Sprintf("calc_trend_strength_%s", cs.CurrentTimeframe)
        fields := []string{"PriceChangePercentage24h", "MarketCap", "MarketCapRank", "MarketCapRankPercentage"}
	commonCryptoHandler(cs, list, fields, newID)
}

func controlCalcTrendStrength(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
	        log.Println("No arguments are passed. Make sure the client preference for the current time frame is 24h. Using the client current timeframe value...")
		log.Printf("Current Timeframe: %v\n", cs.CurrentTimeframe)
	case 1:
	        log.Println("Updating the client current timeframe preference...")
		frames := crypto.GetInputTimeframes([]string{args[0]})
		if len(frames) != 1 || frames[0] != "24h"{
		        log.Fatal("You must use <24h> timeframe to use this feautre....")
		}
		cs.UpdateCurrentTimeframe(frames[0])
	default:
	        log.Println("Invalid use of command: calc crypto trend_strength <timeframe>")
	}
}