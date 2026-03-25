package cmd

import (
        "log"

	"github.com/spf13/cobra"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

var switchCryptoCmd = &cobra.Command{
	Use:   "crypto [args...]",
	Short: "Switch current crypto list",
	Args:  cobra.MinimumNArgs(1),
	Run:   handleCryptoSwitch,
}

func handleCryptoSwitch(cmd *cobra.Command, args []string) {
	defer log.Print("> ")

	if len(args) != 1 {
	        log.Println("Please provide an ID of an existing list.")
		return
	}

	key := args[0]

        cryptoEntry, ok := CryptoCache.Get(key)
	if !ok {
	        log.Println("Requested list does not exist in the client cache.")
		log.Println("To make a get request to the server:")
		log.Printf("get crypto %s\n", key)
		return
	}

        log.Println("Updating the current list with the requested one...")
	CryptoState.UpdateCurrentList(key, cryptoEntry.Market)
	return
}
