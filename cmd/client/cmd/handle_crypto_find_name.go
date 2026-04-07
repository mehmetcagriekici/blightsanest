package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

var searchCryptoNameCmd = &cobra.Command{
	Use:   "crypto name [args...]",
	Short: "Find a coin by its name.",
	Run:   handleCryptoFindName,
}

func handleCryptoFindName(cmd *cobra.Command, args []string) {
	name = args[0]
	if name == "" {
		log.Fatal("Invalid usage: find crypto name <coin_name string>")
	}

	coin, ok := crypto.SearchCoin(name, CryptoState.CurrentList)
	if !ok {
		log.Fatal("Couldn't find the coin.")
	}

	fields := []string{"ID", "MarketCap", "MarketCapChangePercentage", "TotalVolume", "ATH", "AthChangePercentage", "MaxSupply", "CirculatingSupply"}
	crypto.PrintCryptoList([]crypto.MarketData{coin},
		CryptoState.CurrentListID,
		CryptoState.ClientTimeframes,
		fields)
	log.Println("")
}
