package main

import(
        "log"
	"strconv"
	"fmt"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoFindPotentialRally(cs *crypto.CryptoState, args []string) {
        controlFindPotentialRally(cs, args)
        list := crypto.CoinsGetCloseAthChange(cs.CurrentMaxATHChangePercentage, cs.CurrentList)
	newID := fmt.Sprintf("find_potential_rally_%f", cs.CurrentMaxATHChangePercentage)
	fields := []string{"ATH", "AthChangePercentage"}
	commonCryptoHandler(cs, list, fields, newID)
}

// max ath change
func controlFindPotentialRally(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
	        log.Printf("No arguments are passed. Using the client state max ath change percentage value. %f\n", cs.CurrentMaxATHChangePercentage)
	case 1:
	        log.Println("Updating the client max ath change percentage preference")
		maxChange, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateAthChangePercentage(cs.CurrentMinATHChangePercentage, maxChange)
	default:
	        log.Println("Invalid use of command: find crypto potential_rally <max_ath_change_percentage float64>")
	}
}