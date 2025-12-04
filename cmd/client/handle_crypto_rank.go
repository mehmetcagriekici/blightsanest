package main

import(
        "log"
	"fmt"
	"strings"
	
        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoRank(cs *crypto.CryptoState, args []string) {
        defer log.Print("> ")
        controlCryptoRank(cs, args)
	
        sorted := crypto.RankCoins(cs.CurrentSortingField, cs.CurrentOrder, cs.CurrentList)

        baseID := strings.Split(cs.CurrentListID, "_")[0]
        newID := fmt.Sprintf("%s_sorted_%s_%s", baseID, cs.CurrentOrder, cs.CurrentSortingField)
	cs.UpdateCurrentList(newID, sorted)
	
	crypto.PrintCryptoList(cs.CurrentList, cs.CurrentListID, cs.ClientTimeframes, []string{})
	log.Println("")
	
	return
}

// order
// timeframe
func controlCryptoRank(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
	        log.Println("No arguments are provided. Using the existing client preferences for the order and the sorting fieldname values.")
		log.Printf("Current Order:         %v\n", cs.CurrentOrder)
		log.Printf("Current Sorting Field: %v\n", cs.CurrentSortingField)
	case 1:
	        log.Println("One argument is passed. Using the existing client preference for the current sorting field, and updating the client current order preference...")
		log.Printf("Current Sorting Field: %v\n", cs.CurrentSortingField)
		updateOrder(cs, args[0])
	case 2:
	        log.Println("Updating the current order and sorting field client preferences...")
		updateOrder(cs, args[0])
		updateField(cs, args[1])
	default:
	        log.Println("Invalid use of command. rank crypto <asc|desc> <field_name>")
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

// Available Sorting Fiels
/*
<current_price>,
<market_cap>,
<market_cap_rank>,
<market_cap_change_percentage>,
<total_volume>,
<high_24h>,
<low_24h>,
<ath>,
<price_change_percentage (with an existing timeframe in the client state)>,
<ath_change_percentage>,
<max_supply>,
<circulating_supply>
*/

func updateField(cs *crypto.CryptoState, fieldName string) {
        switch fieldName {
	case "current_price":
	        cs.UpdateCurrentSortingField("CurrentPrice")
	case "market_cap":
	        cs.UpdateCurrentSortingField("MarketCap")
	case "market_cap_rank":
	        cs.UpdateCurrentSortingField("MarketCapRank")
	case "market_cap_change_percentage":
	        cs.UpdateCurrentSortingField("MarketCapChangePercentage")
	case "total_volume":
	        cs.UpdateCurrentSortingField("TotalVolume")
	case "high_24h":
	        cs.UpdateCurrentSortingField("High24H")
	case "low_24h":
	        cs.UpdateCurrentSortingField("Low24H")
	case "ath":
	        cs.UpdateCurrentSortingField("ATH")
	case "price_change_percentage":
	        cs.UpdateCurrentSortingField(fmt.Sprintf("PriceChangePercentage%s", cs.CurrentTimeframe))
	case "ath_change_percentage":
	        cs.UpdateCurrentSortingField("AthChangePercentage")
	case "max_supply":
	        cs.UpdateCurrentSortingField("MaxSupply")
	case "circulating_supply":
	        cs.UpdateCurrentSortingField("CirculatingSupply")
	default:
	        log.Println("Invalid sorting field.")
	}
}
