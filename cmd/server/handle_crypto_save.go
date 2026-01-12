package main

import(
	"log"
	"time"
	"context"
	"encoding/json"
	
	"github.com/google/uuid"

	"github.com/mehmetcagriekici/blightsanest/internal/database"
	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

// args: mandotary crypto list cache id, arbitrary custom crypto list id -> saveed as key
func handleCryptoSave(ctx context.Context, cc *crypto.CryptoCache, args []string, queries *database.Queries) {
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
	cryptoEntry, ok := cc.Get(args[0])
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
		ID: dbID,
		UpdatedAt: cryptoEntry.CreatedAt,
		CryptoKey: cryptoKey,
		CryptoList: json.RawMessage(encoded),
	}

	if _, err := queries.CreateCryptoList(ctx, dbParams); err != nil {
		log.Fatal(err)
	}
	log.Printf("List %s successfully saved to database as %s\n", args[0], cryptoKey)
}
