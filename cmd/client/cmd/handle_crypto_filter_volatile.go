package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

var filterCryptoVolatileCmd = &cobra.Command{
	Use:   "crypto volatile [args...]",
	Short: "Filter coinst by min and max swing scores.",
	Run:   handleCryptoFilterVolatile,
}

func handleCryptoFilterVolatile(cmd *cobra.Command, args []string) {
	controlFilterVolatile(CryptoState, args)
	list := crypto.FindWildSwingCoins(CryptoState.CurrentMinSwingScore,
		CryptoState.CurrentMaxSwingScore,
		CryptoState.CurrentList)
	newID := fmt.Sprintf("filter_volatile_%f_%f",
		CryptoState.CurrentMinSwingScore,
		CryptoState.CurrentMaxSwingScore)
	fields := []string{"High24H", "Low24H"}
	commonCryptoHandler(CryptoState, list, fields, newID)
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
