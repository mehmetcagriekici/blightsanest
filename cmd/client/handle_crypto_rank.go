package main

import(
        "log"
	
        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoRank(cs *crypto.CryptoState, args []string) {
        controlCryptoRank(cs, args)
	
        sorted := crypto.RankCoins(cs.CurrentTimeframe, cs.CurrentOrder, cs.CurrentList)
	
	log.Println("")
	log.Println("To update the list with the sorted one: mutate rank crypto")
	log.Println("")
	log.Println("")
	log.Println("")
	
	crypto.PrintCryptoList(sorted, cs.CurrentListID, cs.ClientTimeframes, []string{})
}

// order
// timeframe
func controlCryptoRank(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
	        log.Println("No arguments are provided. Using the existing client preferences for the order and the timeframe values.")
		log.Printf("Current Order: %v\n", cs.CurrentOrder)
		log.Printf("Current Timeframe: %v\n", cs.CurrentTimeframe)
	case 1:
	        log.Println("One argument is passed. Using the existing client preference for the current timeframe, and updating the client current order preference...")
		log.Printf("Current Timeframe: %v\n", cs.CurrentTimeframe)
		updateOrder(cs, args[0])
	case 2:
	        log.Println("Updating the current order and timeframe client preferences...")
		updateOrder(cs, args[0])
		timeframes := crypto.GetInputTimeframes([]string{args[1]})
		cs.UpdateCurrentTimeframe(timeframes[0])
	default:
	        log.Println("Invalid use of command. rank crypto <asc|desc> <timeframe>")
	}
}

func updateOrder(cs *crypto.CryptoState, order string) {
	if order == "asc" {
                cs.UpdateOrder(crypto.CRYPTO_ASC)
	} else if order == "desc" {
	        cs.UpdateOrder(crypto.CRYPTO_DESC)
	} else {
	        log.Println("Invalid sorting order. Available orders: <desc> <asc>")
	}
}