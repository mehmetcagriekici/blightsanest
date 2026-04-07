package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

var searchCryptoPotentialRallyCmd = &cobra.Command{
	Use:   "crypto potential_rally [args...]",
	Short: "Find coins with a potential rally.",
	Run:   handleCryptoPotentialRally,
}

func handleCryptoFindPotentialRally(cmd *cobra.Command, args []string) {
	controlFindPotentialRally(CryptoState, args)
	list := crypto.CoinsGetCloseAthChange(CryptoState.CurrentMaxATHChangePercentage, CryptoState.CurrentList)
	newID := fmt.Sprintf("find_potential_rally_%f", CryptoState.CurrentMaxATHChangePercentage)
	fields := []string{"ATH", "AthChangePercentage"}
	commonCryptoHandler(CryptoState, list, fields, newID)
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
		cs.UpdateAthChangePercentage(maxChange)
	default:
		log.Println("Invalid use of command: find crypto potential_rally <max_ath_change_percentage float64>")
	}
}
