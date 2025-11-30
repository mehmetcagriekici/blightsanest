package main

import(
        "log"
	"strconv"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoFindPotentialRally(cs *crypto.CryptoState, args []string) {
        controlFindPotentialRally(cs, args)
	
        log.Println("Starting to find the coins with large potential upside remaining to their ATH with max ath change percentage preference.")
	log.Println("")

        list := crypto.CoinsGetCloseAthChange(cs.CurrentMaxATHChangePercentage, cs.CurrentList)
	fields := []string{"ATH", "AthChangePercentage"}
	crypto.PrintCryptoList(list, cs.CurrentListID, cs.ClientTimeframes, fields)
	log.Println("")
	log.Println("To update the current client list with the result: mutate find crypto potential_rally")
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