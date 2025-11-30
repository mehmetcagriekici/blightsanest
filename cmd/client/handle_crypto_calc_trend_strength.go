package main

import(
        "log"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoCalcTrendStrength(cs *crypto.CryptoState, args []string) {
        controlCalcTrendStrength(cs, args)
	
        log.Println("Calculating a trend strength index for the coins...")
	log.Println("")

        list, err := crypto.CheckRealTrend(cs.CurrentTimeframe, cs.CurrentList)
	if err != nil {
	        log.Fatal(err)
	}
	fields := []string{"PriceChangePercentage24h", "MarketCap", "MarketCapRank", "MarketCapRankPercentage"}
	crypto.PrintCryptoList(list, cs.CurrentListID, cs.ClientTimeframes, fields)
	log.Println("")
	log.Println("To update the current client list with the result: mutate calc crypto trend_strength")
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