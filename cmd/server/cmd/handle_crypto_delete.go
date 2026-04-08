package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var deleteCryptoCmd = &cobra.Command{
	Use:   "crypto [args...]",
	Short: "Delete a crypto list from the database",
	Run:   handleCryptoDelete,
}

func handleCryptoDelete(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Println("To delete a crypto list from the database you need to provide the ID of the crypto list...")
		return
	}

	deleted, err := DbQueries.DeleteCryptoList(Ctx, args[0])
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range deleted {
		log.Printf("Crypto list %s is successfully deleted.\n", l.CryptoKey)
	}
}
