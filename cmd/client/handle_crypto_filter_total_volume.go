package main

import (
        "log"
	"strconv"
	
        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoFilterTotalVolume(cs *crypto.CryptoState, args []string) {
        controlFilterTotalVolume(cs, args)
	
        list := crypto.FilterCoinVolume(cs.CurrentMinVolume, cs.CurrentMaxVolume, cs.CurrentList)
	log.Println("")
	log.Println("Successfully filtered the coins by their volumes with.")
	log.Println("")
	log.Println("")
	fields := []string{"TotalVolume"}
	crypto.PrintCryptoList(list, cs.CurrentListID, cs.ClientTimeframes, fields)
	log.Println("")
	log.Println("To update the list with the result: mutate filter crypto total_volume")
        log.Println("")
}

// min_volume
// max_volume
func controlFilterTotalVolume(cs *crypto.CryptoState, args []string) {
        switch len(args) {
        case 0:
	        log.Println("No arguments are passed, using the ones from the client state:")
		log.Printf("min volume: %f\n", cs.CurrentMinVolume)
		log.Printf("max volume: %f\n", cs.CurrentMaxVolume)
	case 1:
	        log.Println("min volume passed, using the max volume from the client state")
		log.Printf("max volume: %f\n", cs.CurrentMaxVolume)
		minVolume, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateVolume(minVolume, cs.CurrentMaxVolume)
	case 2:
	        log.Println("Using the input min and max volumes, and updating the client preferences.")
		min, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		max, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateVolume(min, max)
	default:
	        log.Println("Invalid arguments: filter crypto total_volume <min_volume float64> <max_volume float64>")
	}
}