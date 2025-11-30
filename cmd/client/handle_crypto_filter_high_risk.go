package main

import(
        "log"
	"strconv"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoFilterHighRisk(cs *crypto.CryptoState, args []string) {
        controlFilterHighRisk(cs, args)
	
        log.Println("Filtering the coins by their volatility based on 24h high/low swing rate")
	log.Println("")

        list := crypto.FlagRiskCoins(cs.CurrentMaxATHChangePercentage, cs.CurrentMaxVolume, cs.CurrentList)
	fields := []string{"TotalVolume", "ATH", "AthChangePercentage"}
	crypto.PrintCryptoList(list, cs.CurrentListID, cs.ClientTimeframes, fields)
	log.Println("")
	log.Println("To update the current client list with the result: filter crypto high_risk")
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