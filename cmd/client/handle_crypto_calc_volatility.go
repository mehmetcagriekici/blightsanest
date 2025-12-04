package main

import(
        "log"
	"strconv"
	"strings"
	"fmt"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoCalcVolatility(cs *crypto.CryptoState, args []string) {
        defer log.Print("> ")
	
        controlCalcVolatility(cs, args)

        list := crypto.CalcCoinVolatility(cs.CurrentMinVolatility, cs.CurrentMaxVolatility, cs.CurrentList)

        baseID := strings.Split(cs.CurrentListID, "_")[0]
	newID := fmt.Sprintf("%s_calc_volatility_%s_%s", baseID, cs.CurrentMinVolatility, cs.CurrentMaxVolatility)
	cs.UpdateCurrentList(newID, list)

        fields := []string{"High24H", "Low24H"}
	crypto.PrintCryptoList(cs.CurrentList, cs.CurrentListID, cs.ClientTimeframes, fields)
	log.Println("")

        return
}

func controlCalcVolatility(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
	        log.Println("No arguments are passed. Using the current client preferences for min and max volatility values.")
		log.Println("volatility = (high_24h - lwo_24h) / current_price")
		log.Printf("Current Min Volatility: %f\n", cs.CurrentMinVolatility)
		log.Printf("Current Max Volatility: %f\n", cs.CurrentMaxVolatility)
	case 1:
	        log.Println("Only one argument is passed. Updating the client min volatility preference...")
 		log.Printf("Current Max Volatility: %f\n", cs.CurrentMaxVolatility)
 		minVolatility, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateCurrentVolatility(minVolatility, cs.CurrentMaxVolatility)
	case 2:
	        log.Println("Updating the client min and max volatility preferences...")
		minVolatility, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		maxVolatility, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateCurrentVolatility(minVolatility, maxVolatility)
	default:
	        log.Println("Invalid use of command: calc crypto volatility <min_volatility float64> <max_volatility float64>")
	}
}