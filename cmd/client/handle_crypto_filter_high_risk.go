package main

import(
        "log"
	"strconv"
	"strings"
	"fmt"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoFilterHighRisk(cs *crypto.CryptoState, args []string) {
        defer log.Print("> ")
	
        controlFilterHighRisk(cs, args)
	
        list := crypto.FlagRiskCoins(cs.CurrentMaxATHChangePercentage, cs.CurrentMaxVolume, cs.CurrentList)

        baseID := strings.Split(cs.CurrentListID, "_")[0]
	newID := fmt.Sprintf("%s_filter_high_risk_%s_%s", baseID, cs.CurrentMaxATHChangePercentage, cs.CurrentMaxVolume)
	cs.UpdateCurrentList(newID, list)
	
	fields := []string{"TotalVolume", "ATH", "AthChangePercentage"}
	crypto.PrintCryptoList(cs.CurrentList, cs.CurrentListID, cs.ClientTimeframes, fields)
	log.Println("")
	
        return
}

// max ath change
// max volume
func controlFilterHighRisk(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
	        log.Printf("No arguments are passed, using the client state values for max ath change percentage and the max total volume: %f %f\n", cs.CurrentMaxATHChangePercentage, cs.CurrentMaxVolume)
	case 1:
	        log.Printf("Updating the max ath change percentage preference. Using the max volume value from the client state: %f\n", cs.CurrentMaxVolume)
		maxAth, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateAthChangePercentage(cs.CurrentMinATHChangePercentage, maxAth)
	case 2:
	        log.Println("Updating the max ath change and max volume preferences...")
		maxAth, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		maxVolume, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateAthChangePercentage(cs.CurrentMinATHChangePercentage, maxAth)
		cs.UpdateVolume(cs.CurrentMinVolume, maxVolume)
	default:
	        log.Println("Incorrect use of the command: filter crypto high_risk <max_ath_change_percentage float64> <max_total_volume float64>")
	}
}