package main

import(
        "log"
	"fmt"
	
        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoRank(cs *crypto.CryptoState, order string, timeframe string) {
        var sortingOrder crypto.AvailableOrders
        if order == "asc" {
	        sortingOrder = crypto.CRYPTO_ASC
	} else if order == "desc" {
	        sortingOrder = crypto.CRYPTO_DESC
	} else {
	        log.Fatal("Invalid sorting order! <asc | desc>")
	}

        // get the timeframe using the user input
	frames := crypto.GetInputTimeframes([]string{timeframe})
	if len(frames) != 1 {
	        log.Fatal("To rank the coins bu price change percentage, an exsisting timeframe needed to be based on.")
	}

        cs.UpdateOrder(sortingOrder)
        cs.UpdateCurrentTimeframe(frames[0])
        sorted := crypto.RankCoins(frames[0], sortingOrder, cs.CurrentList)
	log.Printf("Sorting successfully completed in %s order by price change percentage %s\n", order, timeframe)
	log.Println("")
	log.Println("To update the list with the sorted one: mutate rank crypto")
	log.Println("")
	log.Println("")
	log.Println("")
	
	// print the sorted list
	sortedID := fmt.Sprintf("%s__sorted-%s", cs.CurrentListID, order)
	crypto.PrintCryptoList(sorted, sortedID, []string{timeframe}, []string{})
}