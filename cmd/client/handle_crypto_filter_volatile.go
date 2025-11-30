package main

import (
        "log"
	"strconv"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoFilterVolatile(cs *crypto.CryptoState, args []string) {
        controlFilterVolatile(cs, args)
	
        log.Println("Filtering the list by a swing range.")
        log.Println("")
        log.Println("swing_score(rate) = high_24h / low_24h")
	log.Println("")

        list := crypto.FindWildSwingCoins(cs.CurrentMinSwingScore, cs.CurrentMaxSwingScore, cs.CurrentList)
	fields := []string{"High24H", "Low24H"}
	crypto.PrintCryptoList(list, cs.CurrentListID, cs.ClientTimeframes, fields)
	log.Println("")
	log.Println("To update the current list with the result: mutate filter crypto volatile")
}

// max rate
// min rate
func controlFilterVolatile(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
	        log.Printf("No arguments are passed, using the client state values for min and max swing scores: %f, %f\n", cs.CurrentMinSwingScore, cs.CurrentMaxSwingScore)
	case 1:
	        log.Printf("1 argument is passed, using it as the min swing score. Using the client state max swing score. Max swing score: %f\n", cs.CurrentMaxSwingScore)
		minScore, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateCurrentSwingScore(minScore, cs.CurrentMaxSwingScore)
	case 2:
	        log.Println("All arguments are passed. Using the passed arguments as the min/max swing scores and updating the client state values.")
		minRate, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		maxRate, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateCurrentSwingScore(minRate, maxRate)
	default:
	        log.Println("Invalid usage of the command: filter crypto volatile <min_rate float64> <max_rate float64>")
	}
}