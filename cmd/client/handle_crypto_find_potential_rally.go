package main

import(
        "log"
	"strconv"
	"strings"
	"fmt"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoFindPotentialRally(cs *crypto.CryptoState, args []string) {
        defer log.Print("> ")
	
        controlFindPotentialRally(cs, args)

        list := crypto.CoinsGetCloseAthChange(cs.CurrentMaxATHChangePercentage, cs.CurrentList)

        baseID := strings.Split(cs.CurrentListID, "_")[0]
	newID := fmt.Sprintf("%s_find_potential_rally_%s", baseID, cs.CurrentMaxATHChangePercentage)
	cs.UpdateCurrentList(newID, list)
	
	fields := []string{"ATH", "AthChangePercentage"}
	crypto.PrintCryptoList(cs.CurrentList, cs.CurrentListID, cs.ClientTimeframes, fields)
	log.Println("")

        return
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