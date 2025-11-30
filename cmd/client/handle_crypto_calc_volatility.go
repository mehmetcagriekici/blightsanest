package main

import(
        "log"
	"strconv"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoCalcVolatility(cs *crypto.CryptoState, args []string) {
        controlCalcVolatility(cs, args)
	
        log.Println("Calculating the daily volatility ranges of the coins...")
	log.Println("volatility = (high_24h - lwo_24h) / current_price")
	log.Println("")

        list := crypto.CalcCoinVolatility(cs.CurrentMinVolatility, cs.CurrentMaxVolatility, cs.CurrentList)
	fields := []string{"High24H", "Low24H"}
	crypto.PrintCryptoList(list, cs.CurrentListID, cs.ClientTimeframes, fields)
	log.Println("")
        log.Println("To update the current client list with the result: mutate calc crypto volatility")
	log.Println("")
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