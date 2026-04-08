package cmd

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/spf13/cobra"

	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/database"
)

var saveCryptoCmd = &cobra.Command{
	Use:   "crypto [args...]",
	Short: "Save crypto data to the database",
	Run:   handleCryptoSave,
}

// args: mandotary crypto list cache id, arbitrary custom crypto list id -> saveed as key
func handleCryptoSave(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Println("To save a crypto list to the database, please provide the ID of the crypto list you want to save")
		return
	}

	// initial key -> cache id
	cryptoKey := args[0]
	if len(args) == 2 {
		cryptoKey = args[1]
	}

	// check if the passed list id exist on the cache
	cryptoEntry, ok := CryptoCache.Get(args[0])
	if !ok {
		log.Println("Provided crypto list id does not exist on the server cache!")
		return
	}

	// encode crypto list into json.RawMessage
	encoded, err := json.Marshal(cryptoEntry.Market)
	if err != nil {
		log.Fatal(err)
	}

	dbID := uuid.New()
	dbParams := database.CreateCryptoListParams{
		ID:         dbID,
		UpdatedAt:  cryptoEntry.CreatedAt,
		CryptoKey:  cryptoKey,
		CryptoList: json.RawMessage(encoded),
	}

	if _, err := DbQueries.CreateCryptoList(Ctx, dbParams); err != nil {
		log.Fatal(err)
	}
	log.Printf("List %s successfully saved to database as %s\n", args[0], cryptoKey)
}
