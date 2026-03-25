package cmd

import (
        "log"

	"github.com/spf13/cobra"

	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

var listCryptoCmd = &cobra.Command{
	Use:   "crypto",
	Short: "Print existing crypto list ids."
	Run:   handleCryptoList,
}

func handleCryptoList(cmd *cobra.Command, args []string) {
       defer log.Print(">")

       if CryptoState.CurrentListID != "" {
               log.Printf("Current Crypto List ID: %s\n", CryptoState.CurrentListID)
       }

       if len(CryptoCache.Market) == 0 {
               log.Println("Client Cache is empty.")
	       return
       }

       for k := range CryptoCache.Market {
               log.Printf("List: %s\n", k)
       }

       return
}
